package locate

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocateBenchmarks(t *testing.T) {
	benchs, err := Benchmarks(context.Background(), ".")
	require.NoError(t, err)

	require.Equal(t, []BenchInfo{
		{"BenchmarkFoo", "internal"},
		{"BenchmarkBar", "internal"},
		{"BenchmarkBaz", "internal/anotherpackage"},
		{"BenchmarkBoz", "internal/anotherpackage"},
	}, benchs)
}
