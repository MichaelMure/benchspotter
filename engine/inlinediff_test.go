package engine_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MichaelMure/benchspotter/engine"
)

func TestDiffInlineAnalysis(t *testing.T) {
	base := []engine.InlineSite{
		{File: "foo.go", Line: 10, Message: "cannot inline Foo: too complex"},
		{File: "foo.go", Line: 20, Message: "can inline Bar"},
	}
	new := []engine.InlineSite{
		{File: "foo.go", Line: 11, Message: "cannot inline Foo: too complex"},        // same, line shifted
		{File: "foo.go", Line: 30, Message: "cannot inline Baz: function too large"}, // added
	}

	diffs := engine.DiffInlineAnalysis(base, new)

	statusOf := func(msg string) engine.DiffStatus {
		for _, d := range diffs {
			if d.Site().Message == msg {
				return d.Status()
			}
		}
		t.Fatalf("site %q not found in diffs", msg)
		return -1
	}

	require.Equal(t, engine.DiffSame, statusOf("cannot inline Foo: too complex"))
	require.Equal(t, engine.DiffRemoved, statusOf("can inline Bar"))
	require.Equal(t, engine.DiffAdded, statusOf("cannot inline Baz: function too large"))

	// Same site: Base.Line=10, New.Line=11
	for _, d := range diffs {
		if d.Site().Message == "cannot inline Foo: too complex" {
			require.NotNil(t, d.Base)
			require.NotNil(t, d.New)
			require.Equal(t, 10, d.Base.Line)
			require.Equal(t, 11, d.New.Line)
		}
	}
}

func TestDiffInlineAnalysis_MultipleMatchesByKey(t *testing.T) {
	// Two calls to the same function in the same file — both have the same message.
	base := []engine.InlineSite{
		{File: "foo.go", Line: 5, Message: "inlining call to helper"},
		{File: "foo.go", Line: 8, Message: "inlining call to helper"},
	}
	new := []engine.InlineSite{
		{File: "foo.go", Line: 5, Message: "inlining call to helper"},
		// second call removed
	}

	// base[0] matches new[0] → Same; base[1] unmatched → Removed
	diffs := engine.DiffInlineAnalysis(base, new)
	require.Len(t, diffs, 2)
	same, removed := 0, 0
	for _, d := range diffs {
		switch d.Status() {
		case engine.DiffSame:
			same++
		case engine.DiffRemoved:
			removed++
		}
	}
	require.Equal(t, 1, same)
	require.Equal(t, 1, removed)
}

func TestDiffInlineAnalysis_CostNumberNormalization(t *testing.T) {
	// Same decision (cannot inline), different cost values — should be DiffSame.
	base := []engine.InlineSite{
		{File: "foo.go", Line: 10, Message: "cannot inline (*Table).Format: function too complex: cost 1088 exceeds budget 80"},
	}
	new := []engine.InlineSite{
		{File: "foo.go", Line: 10, Message: "cannot inline (*Table).Format: function too complex: cost 1253 exceeds budget 80"},
	}

	diffs := engine.DiffInlineAnalysis(base, new)
	require.Len(t, diffs, 1)
	require.Equal(t, engine.DiffSame, diffs[0].Status())
	require.NotNil(t, diffs[0].Base)
	require.NotNil(t, diffs[0].New)
}
