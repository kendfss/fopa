package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/kendfss/pipe"
	"github.com/unilibs/uniwidth"

	"github.com/kendfss/fopa"
)

var (
	flags = struct {
		Split       bool
		InPlace     bool
		Keep        bool
		NoRollback  bool
		NoTrim      bool
		Whole       bool
		FollowLinks bool
		Overwrite   bool
		Filler      *string
	}{
		Filler: &fopa.Filler,
	}
	version string
)

func init() {
	flag.BoolVar(&flags.Overwrite, "o", false, "overwrite any file that exists at the new path; overwritten files will be moved to the temp directory under /fopa")
	flag.BoolVar(&flags.Split, "s", false, "split path(s) on OS path-separator before sanitizing?")
	flag.BoolVar(&flags.InPlace, "i", false, "in place; rename files")
	flag.BoolVar(&flags.Keep, "K", false, "keep old files instead of removing them")
	flag.BoolVar(&flags.NoRollback, "R", false, "no rollback; do not rollback files from collisions while moving files if a fatal error is incurred")
	flag.BoolVar(&flags.NoTrim, "t", false, "no trim; do not trim leading and trailing whitespace from file paths. blobs from stdin will always be split on new lines")
	flag.BoolVar(&flags.Whole, "w", false, "whole path; clean whole path, not just base name")
	flag.BoolVar(&flags.FollowLinks, "l", false, "follow symbolic links when calling stat")
	flag.StringVar(flags.Filler, "f", fopa.Filler, "fill character to remove consecutive occurences of")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s %s:\n", os.Args[0], version)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	mode := modeBits()
	err := mode.validate()
	if err != nil {
		panic(impossible(err))
	}
	args := flag.Args()
	if len(args) == 0 {
		buf := pipe.Get()
		if len(buf) == 0 {
			flag.Usage()
			fatal("no file paths received")
		}
		args = regexp.MustCompile(`(\r?\n)+`).Split(string(buf), -1)
	}
	ops, widest, err := mkOps(args, mode)
	if err != nil {
		fatal(err.Error())
	}
	err = moveAllAndPrint(ops, widest, mode)
	if err != nil {
		fatal(err.Error())
	}
}

type mode uint8

const (
	followLinks mode = 1 << iota
	inPlace
	keep
	noRollback
	noTrim
	overwrite
	split
	whole
)

// validate checks that only valid combinations of mode bits are used
func (m mode) validate() error {
	const all = followLinks | inPlace | keep | noRollback | noTrim | split | whole
	if m&^all != 0 {
		return fmt.Errorf("invalid mode bits set: have %08b, want combination of followLinks (%08b) inPlace (%08b), noRollback (%08b), keep (%08b), noTrim (%08b), overwrite (%08b), split (%08b), and whole (%08b) - or zero", m, followLinks, inPlace, keep, noRollback, noTrim, overwrite, split, whole)
	}
	return nil
}

// modeBits computes the relevant mode bitfield given the boolean flags
func modeBits() (m mode) {
	if flags.FollowLinks {
		m |= followLinks
	}
	if flags.InPlace {
		m |= inPlace
	}
	if flags.Keep {
		m |= keep
	}
	if flags.NoRollback {
		m |= noRollback
	}
	if flags.NoTrim {
		m |= noTrim
	}
	if flags.Overwrite {
		m |= overwrite
	}
	if flags.Split {
		m |= split
	}
	if flags.Whole {
		m |= whole
	}
	return
}

type status uint8

const (
	pending status = iota
	done
)

type renameOp struct {
	old, new    string
	width       int
	status      status
	obsolescent string
}

// logf prints a line-feed-terminated string to stderr
func logf(msg string, args ...any) {
	if len(msg) > 0 && msg[len(msg)-1] != '\n' {
		msg += "\n"
	}
	fmt.Fprintf(os.Stderr, msg, args...)
}

// fatal prints a message and exits the program with status 1
func fatal(msg string, args ...any) {
	logf(msg, args...)
	os.Exit(1)
}

// max determines the larger of two ints
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// clone copies a file/directory from one path to another
// files and directories will have the same mode/permissions as their source versions
func clone(op *renameOp, mode mode) (err error) {
	var ostat os.FileInfo
	ostat, err = stat(op.old, mode)
	if err != nil {
		return
	}
	_, err = stat(op.new, mode)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if mode&overwrite == 0 {
			return fmt.Errorf("%w: new file %q", os.ErrExist, op.new)
		}
		err = clone(&renameOp{old: op.new, new: op.obsolescent}, mode&^overwrite)
		if err != nil {
			return
		}
	}
	newDir := filepath.Dir(op.new)
	var dstat os.FileInfo
	dstat, err = stat(newDir, mode)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !dstat.IsDir() {
			return fmt.Errorf("%w: %q should be a directory or non-existant", os.ErrExist, newDir)
		}
	}
	err = os.MkdirAll(newDir, os.ModePerm)
	if err != nil {
		return
	}
	if ostat.Mode()&os.ModeSymlink != 0 {
		var target string
		target, err = os.Readlink(op.old)
		if err != nil {
			return
		}
		return os.Symlink(target, op.new)
	}
	var of, nf *os.File
	if ostat.IsDir() {
		err = os.CopyFS(newDir, os.DirFS(op.old))
		if err != nil {
			return
		}
		err = os.Rename(filepath.Join(newDir, filepath.Base(op.old)), op.new)
		if err != nil {
			return
		}
		nf, err = os.Open(op.new)

	} else {
		nf, err = os.Create(op.new)
	}
	if err != nil {
		return
	}
	defer nf.Close()
	of, err = os.Open(op.old)
	if err != nil {
		return
	}
	defer of.Close()
	if !ostat.IsDir() {
		_, err = nf.ReadFrom(of)
		if err != nil {
			return
		}
	}
	return nf.Chmod(ostat.Mode())
}

// stat performs a status syscall on the given path. if the followlinks bit is set, it will follow symbolic links
func stat(path string, mode mode) (os.FileInfo, error) {
	if mode&followLinks != 0 {
		return os.Stat(path)
	}
	return os.Lstat(path)
}

// mkOps handles the generation of rename ops and determines the widest old file name
func mkOps(args []string, mode mode) ([]*renameOp, int, error) {
	ops := make([]*renameOp, 0, len(args))
	widest := 0
	for _, old := range args {
		if mode&noTrim == 0 {
			old = strings.TrimSpace(old)
		}
		if mode&inPlace != 0 {
			_, err := stat(old, mode)
			if err != nil {
				return nil, -1, err
			}
		}
		if mode&split != 0 {
			for _, old := range filepath.SplitList(old) {
				if mode&noTrim == 0 {
					old = strings.TrimSpace(old)
				}
				op, err := newOp(old, mode)
				if err != nil {
					return nil, -1, err
				}
				ops = append(ops, op)
				widest = max(widest, op.width)
			}
			continue
		}
		op, err := newOp(old, mode)
		if err != nil {
			return nil, -1, err
		}
		ops = append(ops, op)
		widest = max(widest, op.width)
	}
	return ops, widest, nil
}

// impossible creates a prompt for the user to file a bug report
func impossible(err error) error {
	format := "%w\n\tthis should NEVER happen! please file a bug report"
	info, ok := debug.ReadBuildInfo()
	if ok {
		format += fmt.Sprintf("at https://%s/issues", info.Path)
	}
	return fmt.Errorf(format, err)
}

// impossiblef creates a prompt for the user to file a bug report using a format string
func impossiblef(msg string, args ...any) error {
	return impossible(fmt.Errorf(msg, args...))
}

// newOp returns a new renameOp for the given path
func newOp(old string, mode mode) (*renameOp, error) {
	op := renameOp{old: old, width: uniwidth.StringWidth(old)}
	if mode&whole != 0 {
		op.new = fopa.Clean(op.old)
	} else {
		dir, base := filepath.Split(op.old)
		op.new = filepath.Join(dir, fopa.Clean(base))
	}
	err := checkNew(&op, mode)
	if err != nil {
		return nil, err
	}
	return &op, nil
}

// checkNew assures the new file is writable
func checkNew(op *renameOp, mode mode) error {
	_, err := stat(op.new, mode)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if mode&overwrite == 0 {
			return os.ErrExist
		}
		dir := filepath.Join(os.TempDir(), "fopa")
		err := touchIn(dir)
		if err != nil {
			return err
		}
		if mode&noRollback == 0 {
			op.obsolescent = filepath.Join(dir, strings.ReplaceAll(stripVolume(op.new), string(os.PathSeparator), "-"))
		}
	}
	// File doesn't exist, check if parent directory is writable
	dir := filepath.Dir(op.new)
	return touchIn(dir)
}

// touchIn checks if a directory is writable by touching a file and deleting it
func touchIn(dir string) error {
	tmp, err := os.CreateTemp(dir, ".write-check-*")
	if err != nil {
		return err
	}
	defer tmp.Close()
	return os.Remove(tmp.Name())
}

// stripVolume removes the volume name from a file's path
func stripVolume(path string) string {
	volume := filepath.VolumeName(path)
	return strings.Replace(path, volume, "", 1)
}

// moveAllAndPrint executes all renameOp instructions and prints old and new names
// in the event of failure, ie non-nil err, it indoes all of its work by removing new files
// it first tries to copy all files, so this cleanup operation is safe!
func moveAllAndPrint(ops []*renameOp, widest int, mode mode) (err error) {
	defer func() {
		for _, op := range ops {
			err := finalize(op, mode, err)
			if err != nil {
				logf(err.Error())
			}
		}
	}()
	for i, op := range ops {
		if mode&inPlace != 0 {
			err = clone(op, mode)
			if err != nil {
				return
			}
			op.status = done
		}
		op.old += strings.Repeat(" ", widest-ops[i].width)
		fmt.Println(op.old, "->", op.new)
	}
	return nil
}

// finalize handles the removal of old files/dirs and the restoration of pre-existing files/dirs in case of error
func finalize(op *renameOp, mode mode, err error) error {
	if op.status == done {
		if err != nil {
			if mode&noRollback == 0 {
				return rollback(op, mode)
			}
		} else {
			if mode&keep == 0 {
				err := os.RemoveAll(op.old)
				if err != nil {
					if !os.IsNotExist(err) {
						return err
					}
				}
			}
		}
	}
	// if status is not done then there should be nothing to finalize
	return nil
}

// rollback handles the restoration of pre-exsisting files in the event of an error
func rollback(op *renameOp, mode mode) error {
	err := os.RemoveAll(op.new)
	if err != nil {
		if !os.IsNotExist(err) { // user may specify nested paths
			return err
		}
	}
	if op.obsolescent == "" {
		panic(impossiblef("rollback operation for %q has no obsolescent", op.old))
	}
	err = clone(&renameOp{old: op.obsolescent, new: op.new}, mode&^overwrite)
	if err != nil {
		return err
	}
	return os.RemoveAll(op.obsolescent)
}
