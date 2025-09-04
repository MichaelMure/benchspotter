package locate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocateBenchmarks(t *testing.T) {
	benchs, err := Benchmarks(".")
	require.NoError(t, err)

	require.Equal(t, []BenchInfo{
		{"BenchmarkFoo", "internal"},
		{"BenchmarkBar", "internal"},
		{"BenchmarkBaz", "internal/anotherpackage"},
		{"BenchmarkBoz", "internal/anotherpackage"},
	}, benchs)
}
