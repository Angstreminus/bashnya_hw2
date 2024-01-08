package main

import (
	"uniq/errorcheck"
	"uniq/internal/flow"
	"uniq/internal/options"
	"uniq/internal/utils"
)

func main() {
	opt := options.New()
	opt.Init()

	err := options.ValidateKeys(opt.C, opt.U, opt.D)
	errorcheck.Check(err)

	inputFile, outputFile := options.ParseFilePath()

	reader, err := utils.NewReader(inputFile)
	errorcheck.Check(err)
	defer reader.Close()

	data, err := reader.ReadLines()
	errorcheck.Check(err)

	writer, err := utils.NewWriter(outputFile)
	errorcheck.Check(err)

	data = flow.UnifyString(data, *opt)

	err = writer.WriteLines(data)
	errorcheck.Check(err)
}
