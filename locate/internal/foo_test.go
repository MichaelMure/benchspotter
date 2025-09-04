package internal

import (
	"testing"

	"benchspotter/benchinput"
)

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[string]int)
	}
}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[string]int)
	}
}

var input1 = benchinput.Bool("name1", false)
var input2 = benchinput.Int("name2", 12, 0, 24)
