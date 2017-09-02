package main

import (
	"io"

	"ekyu.moe/base91"
	"ekyu.moe/util/cli"
)

func dump(outFilename, inFilename string, wrap int) error {
	// validate and read in file
	inFile, _, err := cli.AccessOpenFile(inFilename)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// validate and create out file
	outFile, err := cli.PromptOverrideCreate(outFilename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var e io.WriteCloser
	if wrap <= 0 {
		e = base91.NewEncoder(outFile)
	} else {
		e = base91.NewLineWrapper(outFile, wrap)
	}
	defer e.Close()

	if _, err := io.Copy(e, inFile); err != nil {
		return err
	}

	return nil
}
