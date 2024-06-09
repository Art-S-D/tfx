package cmd

import (
	"flag"
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

type Args struct {
	Src io.Reader
}

func ParseArgs() *Args {
	var out Args

	fs := flag.NewFlagSet("tfx", flag.ExitOnError)
	input := fs.String("in", "", "input file")
	fs.Parse(os.Args[1:])

	hasInputFile := input != nil && len(*input) > 0
	stdinIsATty := isatty.IsTerminal(os.Stdin.Fd())
	if stdinIsATty {
		if hasInputFile {
			f, err := os.Open(*input)
			if err != nil {
				panic(err.Error())
			}
			out.Src = f
		} else {
			fs.Usage()
			os.Exit(0)
		}
	} else {
		if hasInputFile {
			fs.Usage()
			os.Exit(0)
		} else {
			out.Src = os.Stdin
		}
	}

	return &out
}
