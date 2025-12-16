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

func BenchmarkBaz(b *testing.B) {
	b.Run("foo", func(foo *testing.B) {
		// this will get ignored, it's not a true benchmark as there is a .Run() call.
		for i := 0; i < foo.N; i++ {
			_ = make(map[string]int, input2)
		}

		foo.Run("bar", func(bar *testing.B) {
			for i := 0; i < bar.N; i++ {
				_ = make(map[string]int, input2)
			}
		})
	})
}

var input1 = benchinput.Bool("name1", false)
var input2 = benchinput.Int("name2", 12, 0, 24)
