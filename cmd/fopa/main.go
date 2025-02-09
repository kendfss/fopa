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
	splitFlag bool
	version   string
)

func init() {
	flag.BoolVar(&splitFlag, "s", false, "split path(s) before sanitizing?")
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
			os.Exit(1)
		}
		args = regexp.MustCompile(`(\r?\n)+`).Split(string(buf), -1)
	}
	for i, arg := range args {
		if splitFlag {
			parts := filepath.SplitList(arg)
			for j, part := range parts {
				parts[j] = fopa.Clean(part)
			}
			args[i] = filepath.Join(parts...)
			continue
		}
		args[i] = fopa.Clean(arg)
	}
	fmt.Println(strings.Join(args, "\n"))
}
