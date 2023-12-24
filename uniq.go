package main

import (
	"uniq/internal/errorcheck"
	"uniq/internal/flow"
	"uniq/internal/operations"
	"uniq/internal/options"
)

func main() {
	opt := options.New()
	opt.Init()
	reader, err := operations.NewReader(opt.InputFile)
	errorcheck.Check(err)
	data, err := reader.ReadLines()
	reader.Close()
	errorcheck.Check(err)
	prepStr := flow.UnifyFSI(opt.F, opt.S, opt.I, data)
	if !opt.C && !opt.U && !opt.D {
		data = flow.Unify(data, prepStr)
	} else {
		data = flow.UnifyCUD(opt.C, opt.U, opt.D, prepStr, data)
	}
	writer, err := operations.NewWriter(opt.OutputFile)
	errorcheck.Check(err)
	err = writer.WriteLines(data)
	writer.Close()
	errorcheck.Check(err)
}
