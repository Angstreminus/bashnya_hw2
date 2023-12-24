package flow

import "testing"

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
