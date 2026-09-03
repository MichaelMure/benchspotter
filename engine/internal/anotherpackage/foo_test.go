package internal

import (
	"testing"

	"github.com/MichaelMure/benchspotter/benchinput"
)

func BenchmarkBaz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[string]int)
	}
}

func BenchmarkBoz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[string]int)
	}
}

var input3 = benchinput.Bool("name3", true)
var input4 = benchinput.Int("name4", 270, 180, 360)
