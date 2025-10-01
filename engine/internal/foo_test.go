package internal

import (
	"testing"

	"benchspotter/benchinput"
)

// normal benchmark signature
func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[string]int, input2)
	}
}

// parameter with a different name
func BenchmarkBar(foo *testing.B) {
	for i := 0; i < foo.N; i++ {
		_ = make(map[string]int, input2)
	}
}

var input1 = benchinput.Bool("name1", false)
var input2 = benchinput.Int("name2", 12, 0, 24)
