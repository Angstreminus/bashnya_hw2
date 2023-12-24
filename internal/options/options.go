package options

import (
	"errors"
	"flag"
)

type Options struct {
	C bool
	U bool
	I bool
	D bool
	F int
	S int
}

func New() *Options {
	return &Options{}
}
func ParseFilePath() (inputFileName string, outputFileName string) {
	args := flag.Args()
	switch len(args) {
	case 0:
	case 1:
		inputFileName = args[0]
	case 2:
		inputFileName = args[0]
		outputFileName = args[1]
	}
	return inputFileName, outputFileName
}
func (opt *Options) Init() {
	flag.IntVar(&opt.F, "f", 0, "Provide -f key")
	flag.IntVar(&opt.S, "s", 0, "Provide -s key")
	flag.BoolVar(&opt.C, "c", false, "Provide -c key")
	flag.BoolVar(&opt.U, "u", false, "Provide -u key")
	flag.BoolVar(&opt.I, "i", false, "Provide -i key")
	flag.BoolVar(&opt.D, "d", false, "Provide -d key")
	flag.Parse()
}

func ValidateKeys(keys ...bool) error {
	var cnt int
	for _, key := range keys {
		if !key {
			continue
		}
		cnt++
	}
	if cnt > 1 {
		return errors.New("Flags -c -d -u collision detected. Use only one of them")
	}
	return nil
}
