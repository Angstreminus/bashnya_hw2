package flow

import (
	"testing"
	"uniq/internal/options"
)

func Benchmark1(b *testing.B) {
	lineInput := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	for i := 0; i < b.N; i++ {
		Unify(lineInput, lineInput)
		Unify(lineInput, lineInput)
	}
}

func Benchmark2(b *testing.B) {
	opt := options.New()
	lineInput := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	for i := 0; i < b.N; i++ {
		UnifyString(lineInput, *opt)
	}
}
