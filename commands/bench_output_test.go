package commands

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/perf/benchfmt"

	"benchspotter/commands/execenv"
	"benchspotter/repository"
)

func newBenchOutputTestEnv(t *testing.T) *execenv.Env {
	t.Helper()
	return execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), memfs.New(), nil))
}

func makeResult(name string, value float64, unit string) *benchfmt.Result {
	return &benchfmt.Result{
		Name:   benchfmt.Name(name),
		Iters:  1,
		Values: []benchfmt.Value{{Value: value, Unit: unit, OrigValue: value, OrigUnit: unit}},
	}
}

func TestBenchPrinterOutputRouting(t *testing.T) {
	t.Run("progress goes to stderr not stdout", func(t *testing.T) {
		env := newBenchOutputTestEnv(t)
		bp := newBenchPrinter(env, 3, 20, 5)

		bp.Add(makeResult("BenchmarkFoo-8", 1000, "ns/op"))
		bp.Add(makeResult("BenchmarkFoo-8", 1100, "ns/op"))
		bp.Add(makeResult("BenchmarkFoo-8", 900, "ns/op"))
		bp.Flush()

		assert.Empty(t, env.Out.String())
		assert.NotEmpty(t, env.Err.String())
	})

	t.Run("final line contains benchmark name and values", func(t *testing.T) {
		env := newBenchOutputTestEnv(t)
		bp := newBenchPrinter(env, 1, 20, 5)

		bp.Add(makeResult("BenchmarkFoo-8", 1234, "ns/op"))
		bp.Flush()

		out := env.Err.String()
		assert.Contains(t, out, "Foo")
		assert.Contains(t, out, "ns/op")
	})

	t.Run("multiple benchmarks each get a line", func(t *testing.T) {
		env := newBenchOutputTestEnv(t)
		bp := newBenchPrinter(env, 1, 20, 5)

		bp.Add(makeResult("BenchmarkFoo-8", 1000, "ns/op"))
		bp.Add(makeResult("BenchmarkBar-8", 2000, "ns/op"))
		bp.Flush()

		out := env.Err.String()
		assert.Contains(t, out, "Foo")
		assert.Contains(t, out, "Bar")
		assert.Empty(t, env.Out.String())
	})

	t.Run("count>1 prints spread on final line", func(t *testing.T) {
		env := newBenchOutputTestEnv(t)
		bp := newBenchPrinter(env, 3, 20, 5)

		bp.Add(makeResult("BenchmarkFoo-8", 1000, "ns/op"))
		bp.Add(makeResult("BenchmarkFoo-8", 2000, "ns/op"))
		bp.Add(makeResult("BenchmarkFoo-8", 1500, "ns/op"))
		bp.Flush()

		out := env.Err.String()
		assert.Contains(t, out, "±")
		assert.Empty(t, env.Out.String())
	})
}

func TestProfilePrinterOutputRouting(t *testing.T) {
	t.Run("progress goes to stderr not stdout", func(t *testing.T) {
		env := newBenchOutputTestEnv(t)
		pp := newProfilePrinter(env, "cpu", 2, 20, 5)

		pp.Done("BenchmarkFoo", 0)
		pp.Done("BenchmarkBar", 0)

		assert.Empty(t, env.Out.String())
		assert.NotEmpty(t, env.Err.String())
	})

	t.Run("output contains benchmark name and kind", func(t *testing.T) {
		env := newBenchOutputTestEnv(t)
		pp := newProfilePrinter(env, "cpu", 1, 20, 5)

		pp.Done("BenchmarkFoo", 0)

		out := env.Err.String()
		assert.Contains(t, out, "Foo")
		assert.Contains(t, out, "cpu")
	})
}

func TestBenchPrinterFormatJSON(t *testing.T) {
	t.Run("progress still goes to stderr in json mode", func(t *testing.T) {
		env := newBenchOutputTestEnv(t)
		env.Format = execenv.FormatJSON
		bp := newBenchPrinter(env, 1, 20, 5)

		bp.Add(makeResult("BenchmarkFoo-8", 1000, "ns/op"))
		bp.Flush()

		// stdout is reserved for JSON data; progress must not pollute it
		require.Empty(t, env.Out.String())
	})
}
