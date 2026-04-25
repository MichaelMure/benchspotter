package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseInlineOutput(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		input := []byte(`# benchspotter/engine
engine/foo.go:10:5: cannot inline Foo: too complex
engine/foo.go:20:3: inlining call to bar.Baz
engine/foo.go:30:1: can inline Qux
`)
		sites := ParseInlineOutput(input)
		require.Len(t, sites, 3)

		assert.Equal(t, "engine/foo.go", sites[0].File)
		assert.Equal(t, 10, sites[0].Line)
		assert.Equal(t, 5, sites[0].Col)
		assert.Equal(t, "cannot inline Foo: too complex", sites[0].Message)

		assert.Equal(t, 20, sites[1].Line)
		assert.Equal(t, "inlining call to bar.Baz", sites[1].Message)

		assert.Equal(t, 30, sites[2].Line)
		assert.Equal(t, "can inline Qux", sites[2].Message)
	})

	t.Run("deduplication", func(t *testing.T) {
		// compiler emits the same line multiple times when building with ./...
		input := []byte(`engine/foo.go:10:5: cannot inline Foo: too complex
engine/foo.go:10:5: cannot inline Foo: too complex
engine/foo.go:20:3: inlining call to bar.Baz
`)
		sites := ParseInlineOutput(input)
		assert.Len(t, sites, 2)
	})

	t.Run("empty", func(t *testing.T) {
		assert.Empty(t, ParseInlineOutput(nil))
		assert.Empty(t, ParseInlineOutput([]byte("# only headers\n# no sites\n")))
	})

	t.Run("filters_non_inline", func(t *testing.T) {
		// escape analysis lines must be dropped
		input := []byte(`engine/foo.go:10:5: bar escapes to heap
engine/foo.go:20:3: leaking param: buf
engine/foo.go:30:1: cannot inline Foo: too complex
`)
		sites := ParseInlineOutput(input)
		require.Len(t, sites, 1)
		assert.Equal(t, "cannot inline Foo: too complex", sites[0].Message)
	})
}

func TestInlineSiteKind(t *testing.T) {
	cases := []struct {
		msg  string
		kind InlineKind
	}{
		{"cannot inline Foo: too complex", InlineCannotInline},
		{"cannot inline Bar", InlineCannotInline},
		{"inlining call to pkg.Baz", InlineInliningCall},
		{"can inline Qux", InlineCanInline},
		{"can inline Qux with cost 64 as:", InlineCanInline},
	}
	for _, tc := range cases {
		s := InlineSite{Message: tc.msg}
		assert.Equal(t, tc.kind, s.Kind(), "Kind() for %q", tc.msg)
	}
}

func TestSplitCompilerOutput(t *testing.T) {
	t.Run("separates_escape_and_inline", func(t *testing.T) {
		input := []byte(`engine/foo.go:10:5: bar escapes to heap
engine/foo.go:20:3: leaking param: buf
engine/foo.go:30:1: cannot inline Foo: too complex
engine/foo.go:40:2: inlining call to bar.Baz
engine/foo.go:50:1: can inline Qux
`)
		escapeData, inlineData := splitCompilerOutput(input)

		escapeSites := ParseEscapeOutput(escapeData)
		require.Len(t, escapeSites, 2)
		assert.Equal(t, "bar escapes to heap", escapeSites[0].Message)
		assert.Equal(t, "leaking param: buf", escapeSites[1].Message)

		inlineSites := ParseInlineOutput(inlineData)
		require.Len(t, inlineSites, 3)
		assert.Equal(t, "cannot inline Foo: too complex", inlineSites[0].Message)
		assert.Equal(t, "inlining call to bar.Baz", inlineSites[1].Message)
		assert.Equal(t, "can inline Qux", inlineSites[2].Message)
	})

	t.Run("continuation_lines_follow_escape", func(t *testing.T) {
		// Flow-chain continuation lines (leading space in message) must stay
		// with the escape block they belong to, not bleed into inline output.
		input := []byte(`mypkg/foo.go:10:9: "x" escapes to heap in Foo:
mypkg/foo.go:10:9:   flow: {heap} <- &x:
mypkg/foo.go:10:9:     from &x (address-of) at mypkg/foo.go:10:9
mypkg/foo.go:10:9: "x" escapes to heap
mypkg/foo.go:5:3: cannot inline Bar: too complex
`)
		escapeData, inlineData := splitCompilerOutput(input)

		escapeSites := ParseEscapeOutput(escapeData)
		require.Len(t, escapeSites, 1)
		assert.Equal(t, `"x" escapes to heap`, escapeSites[0].Message)
		require.Len(t, escapeSites[0].FlowChain, 2)

		inlineSites := ParseInlineOutput(inlineData)
		require.Len(t, inlineSites, 1)
		assert.Equal(t, "cannot inline Bar: too complex", inlineSites[0].Message)
	})

	t.Run("package_headers_to_escape", func(t *testing.T) {
		// Lines that don't match the file:line:col pattern (e.g. "# pkg/name")
		// should be included in the escape output and not lost.
		input := []byte(`# benchspotter/engine
engine/foo.go:10:5: bar escapes to heap
`)
		escapeData, _ := splitCompilerOutput(input)
		assert.Contains(t, string(escapeData), "# benchspotter/engine")
	})
}
