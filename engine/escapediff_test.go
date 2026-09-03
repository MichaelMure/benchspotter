package engine_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MichaelMure/benchspotter/engine"
)

func TestDiffEscapeAnalysis(t *testing.T) {
	base := []engine.EscapeSite{
		{File: "foo.go", Line: 5, Message: "x escapes to heap"},
		{File: "foo.go", Line: 10, Message: "y escapes to heap"},
	}
	new := []engine.EscapeSite{
		{File: "foo.go", Line: 5, Message: "x escapes to heap"},  // same
		{File: "foo.go", Line: 15, Message: "z escapes to heap"}, // added
	}

	diffs := engine.DiffEscapeAnalysis(base, new)

	statusOf := func(msg string) engine.DiffStatus {
		for _, d := range diffs {
			if d.Site().Message == msg {
				return d.Status()
			}
		}
		t.Fatalf("site %q not found in diffs", msg)
		return -1
	}

	require.Equal(t, engine.DiffSame, statusOf("x escapes to heap"))
	require.Equal(t, engine.DiffRemoved, statusOf("y escapes to heap"))
	require.Equal(t, engine.DiffAdded, statusOf("z escapes to heap"))
}
