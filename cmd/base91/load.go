package main

import (
	"io"

	"ekyu.moe/base91"
	"ekyu.moe/util/cli"
)

func load(outFilename, inFilename string) error {
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

	d := base91.NewDecoder(inFile)

	if _, err := io.Copy(outFile, d); err != nil {
		return err
	}

	return nil
}
