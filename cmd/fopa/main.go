package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kendfss/pipe"

	"github.com/kendfss/fopa"
)

var (
	splitFlag, inPlaceFlag, noTrimFlag bool
	version                            string
)

func init() {
	flag.BoolVar(&splitFlag, "s", false, "split path(s) on OS path-separator before sanitizing?")
	flag.BoolVar(&inPlaceFlag, "i", false, "in place; rename files")
	flag.BoolVar(&noTrimFlag, "n", false, "no trim; do not trim leading and trailing whitespace from file paths")
	flag.StringVar(&fopa.Filler, "f", fopa.Filler, "fill character to remove consecutive occurences of")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s %s:\n", os.Args[0], version)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		buf := pipe.Get()
		if len(buf) == 0 {
			flag.Usage()
			fatal("no file paths received")
		}
		args = regexp.MustCompile(`(\r?\n)+`).Split(string(buf), -1)
	}
	oldNames := make([]string, 0, len(args))
	newNames := make([]string, 0, len(args))
	oldLens, newLens := make([]int, 0, len(args)), make([]int, 0, len(args))
	longestOld, longestNew := 0, 0
	for _, old := range args {
		new := old
		if !noTrimFlag {
			new = strings.TrimSpace(old)
		}
		if inPlaceFlag && !exists(new) {
			fatal("file not found: %q", old)
		}
		if splitFlag {
			parts := filepath.SplitList(new)
			for _, part := range parts {
				oldNames = append(oldNames, part)
				if !noTrimFlag {
					part = strings.TrimSpace(part)
				}
				newNames = append(newNames, fopa.Clean(part))
				oldLens = append(oldLens, runeLen(old))
				newLens = append(newLens, runeLen(new))
				longestOld, longestNew = max(longestOld, oldLens[len(oldLens)-1]), max(longestNew, newLens[len(newLens)-1])
			}
			continue
		}
		oldNames = append(oldNames, old)
		newNames = append(newNames, fopa.Clean(new))
		oldLens = append(oldLens, runeLen(old))
		newLens = append(newLens, runeLen(new))
		longestOld, longestNew = max(longestOld, oldLens[len(oldLens)-1]), max(longestNew, newLens[len(newLens)-1])
	}
	for i, old := range oldNames {
		new := newNames[i]
		if inPlaceFlag {
			err := os.Rename(old, new)
			if err != nil {
				fatal("%s: %q -> %q", err, old, newNames[i])
			}
		}
		old += strings.Repeat(" ", longestOld-oldLens[i])
		new += strings.Repeat(" ", longestNew-newLens[i])
		fmt.Println(old, "->", new)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func logf(msg string, args ...any) {
	if len(msg) > 0 && msg[len(msg)-1] != '\n' {
		msg += "\n"
	}
	fmt.Fprintf(os.Stderr, msg, args...)
}

func fatal(msg string, args ...any) {
	logf(msg, args...)
	os.Exit(1)
}

func runeLen(s string) int {
	return len([]rune(s))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
