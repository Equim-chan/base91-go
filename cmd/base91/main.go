// base91: a cli utility for base91 encode/decode. Run with --help for details.
package main // import "ekyu.moe/base91/cmd/base91"

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ekyu.moe/base91"
	"ekyu.moe/util/cli"
)

const (
	title = "base91"
	usage = "Usage: " + title + ` [OPTION]... <FILE>

Options:
   -d, --decode      Decode data.
   -o, --output      Use as output file. By default, it appends ".asc" to the input
                     filename under dump mode, and ".plain" under load mode. If the
                     input file is suffixed with ".asc", it will strip it. If "-"
                     is specified, write to stdout in raw. Default "-".
   -w, --wrap=COLS   Wrap encoded lines after COLS character (default 76). Use 0 to
                     disable line wrapping.

Example:
   ` + title + ` -o imgs.zip.asc dump imgs.zip
   ` + title + ` -o imgs.zip load imgs.zip.asc
`
)

func main() {
	var decodeMode bool
	var outFilename string
	var wrap int

	flag.BoolVar(&decodeMode, "d", false, "")
	flag.BoolVar(&decodeMode, "decode", false, "")
	flag.StringVar(&outFilename, "o", "-", "")
	flag.StringVar(&outFilename, "ouput", "-", "")
	flag.IntVar(&wrap, "w", base91.EmailLineWrap, "")
	flag.IntVar(&wrap, "wrap", base91.EmailLineWrap, "")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}

	flag.Parse()

	inFilename := flag.Arg(0)
	if len(inFilename) == 0 {
		inFilename = "-"
	}

	var err error
	if decodeMode {
		if len(outFilename) == 0 && inFilename != "-" {
			if strings.HasSuffix(strings.ToLower(inFilename), ".asc") {
				outFilename = strings.TrimSuffix(inFilename, ".asc")
			} else {
				outFilename = inFilename + ".plain"
			}
			outFilename = inFilename + ".asc"
		}
		err = load(outFilename, inFilename)
	} else {
		if len(outFilename) == 0 && inFilename != "-" {
			outFilename = inFilename + ".asc"
		}
		err = dump(outFilename, inFilename, wrap)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if err != cli.ErrAbortedByUser && outFilename != "-" {
			os.Remove(outFilename)
		}
		os.Exit(1)
	}
}
