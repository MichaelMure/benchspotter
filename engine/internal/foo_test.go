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

// Sibling sub-benchmarks with different parameter names — exercises the
// testingVarName mutation bug where the second sibling was silently missed.
func BenchmarkSiblings(b *testing.B) {
	b.Run("first", func(c *testing.B) {
		for i := 0; i < c.N; i++ {
			_ = make(map[string]int)
		}
	})
	b.Run("second", func(d *testing.B) {
		for i := 0; i < d.N; i++ {
			_ = make(map[string]int)
		}
	})
}

var input1 = benchinput.Bool("name1", false)
var input2 = benchinput.Int("name2", 12, 0, 24)

// Non-literal max argument: static analysis cannot extract it, so this input
// must be omitted from results rather than appended with max=0.
var nonLiteralMax = 50
var inputBad = benchinput.Int("name_bad", 0, 0, nonLiteralMax)
