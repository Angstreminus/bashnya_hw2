package options

import (
	"errors"
	"flag"
)

type Options struct {
	C          bool
	U          bool
	I          bool
	D          bool
	InputFile  string
	OutputFile string
	F          int
	S          int
}

func New() *Options {
	return &Options{}
}

func (opt *Options) Init() {
	flag.IntVar(&opt.F, "f", 0, "Provide -f key")
	flag.IntVar(&opt.S, "s", 0, "Provide -s key")
	flag.BoolVar(&opt.C, "c", false, "Provide -c key")
	flag.BoolVar(&opt.U, "u", false, "Provide -u key")
	flag.BoolVar(&opt.I, "i", false, "Provide -i key")
	flag.BoolVar(&opt.D, "d", false, "Provide -d key")
	flag.StringVar(&opt.InputFile, "input_file", "", "Provide -input_file key")
	flag.StringVar(&opt.OutputFile, "output_file", "", "Provide -output_file key")
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
