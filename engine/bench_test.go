package engine

import (
	"context"
	"testing"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/stretchr/testify/require"
)

func TestLocateBenchmarks(t *testing.T) {
	fs := osfs.New(".")

	benchs, err := LocateBenchmarks(context.Background(), fs)
	require.NoError(t, err)

	require.Equal(t, []BenchInfo{
		{"BenchmarkBaz", "internal/anotherpackage"},
		{"BenchmarkBoz", "internal/anotherpackage"},
		{"BenchmarkFoo", "internal"},
		{"BenchmarkBar", "internal"},
		{"BenchmarkBaz/foo/bar", "internal"},
	}, benchs)
}
