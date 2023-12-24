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
	inputFile, outputFile := options.ParseFilePath()
	reader, err := utils.NewReader(inputFile)
	errorcheck.Check(err)
	defer reader.Close()
	data, err := reader.ReadLines()
	errorcheck.Check(err)
	writer, err := utils.NewWriter(outputFile)
	errorcheck.Check(err)
	defer writer.Close()
	prepStr := flow.UnifyFSI(opt.F, opt.S, opt.I, data)
	if !opt.C && !opt.U && !opt.D {
		data = flow.Unify(data, prepStr)
	} else {
		data = flow.UnifyCUD(opt.C, opt.U, opt.D, prepStr, data)
	}
	err = writer.WriteLines(data)
	errorcheck.Check(err)
}
