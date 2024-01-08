package flow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUniqStrings(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name: "Default uniq test",
			input: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			want: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := GetUniqStrings(test.input)
			assert.Equal(t, test.want, res)
		})
	}
}

func TestUnifyCUD(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
		name  string
		c     bool
		u     bool
		d     bool
	}{
		{
			name: "C param test",
			input: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			c: true,
			d: false,
			u: false,
			want: []string{
				"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks.",
			},
		},
		{
			name: "D param test",
			input: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			c: false,
			d: true,
			u: false,
			want: []string{
				"I love music.",
				"I love music of Kartik.",
			},
		},
		{
			name: "U param test",
			input: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
			},
			c: false,
			d: false,
			u: true,
			want: []string{
				"",
				"Thanks.",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := UnifyCUD(test.c, test.u, test.d, test.input, test.input)
			assert.Equal(t, test.want, res)
		})
	}
}

func TestParamI(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name: "I Param test",
			input: []string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
			},
			want: []string{
				"i love music.",
				"i love music.",
				"i love music.",
				"",
				"i love music of kartik.",
				"i love music of kartik.",
				"thanks.",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ParamI(&test.input)
			assert.Equal(t, test.want, res)
		})
	}
}

func TestParamS(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name: "S Param test",
			input: []string{
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			want: []string{
				" love music.",
				" love music.",
				" love music.",
				"",
				" love music of Kartik.",
				"e love music of Kartik.",
				"hanks.",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ParamS(&test.input, 1)
			assert.Equal(t, test.want, res)
		})
	}
}

func TestParamF(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name: "F Param test",
			input: []string{
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			want: []string{
				"love music.",
				"love music.",
				"love music.",
				"",
				"love music of Kartik.",
				"love music of Kartik.",
				"Thanks.",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ParamF(&test.input, 1)
			assert.Equal(t, test.want, res)
		})
	}
}
