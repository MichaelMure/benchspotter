package engine

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseEscapeOutput(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		input := []byte(`# benchspotter/engine
engine/foo.go:10:5: bar escapes to heap
engine/foo.go:20:3: leaking param: buf
engine/foo.go:30:1: can inline Baz
`)
		sites := ParseEscapeOutput(input)
		require.Len(t, sites, 3)

		assert.Equal(t, "engine/foo.go", sites[0].File)
		assert.Equal(t, 10, sites[0].Line)
		assert.Equal(t, 5, sites[0].Col)
		assert.Equal(t, "bar escapes to heap", sites[0].Message)

		assert.Equal(t, 20, sites[1].Line)
		assert.Equal(t, "leaking param: buf", sites[1].Message)

		assert.Equal(t, 30, sites[2].Line)
		assert.Equal(t, "can inline Baz", sites[2].Message)
	})

	t.Run("deduplication", func(t *testing.T) {
		// compiler emits the same line multiple times when building with ./...
		input := []byte(`engine/foo.go:10:5: bar escapes to heap
engine/foo.go:10:5: bar escapes to heap
engine/foo.go:20:3: leaking param: buf
`)
		sites := ParseEscapeOutput(input)
		assert.Len(t, sites, 2)
	})

	t.Run("empty", func(t *testing.T) {
		assert.Empty(t, ParseEscapeOutput(nil))
		assert.Empty(t, ParseEscapeOutput([]byte("# only headers\n# no sites\n")))
	})

	t.Run("flow_chain", func(t *testing.T) {
		// -m=2 emits a verbose block before the summary line.
		// Blocks and summaries may appear in different orders.
		input := []byte(`# mypkg
mypkg/foo.go:10:9: "x" escapes to heap in Foo:
mypkg/foo.go:10:9:   flow: {heap} <- &{storage for "x"}:
mypkg/foo.go:10:9:     from "x" (spill) at mypkg/foo.go:10:9
mypkg/foo.go:10:9:     from panic("x") (call parameter) at mypkg/foo.go:10:8
mypkg/foo.go:8:6: leaking param: name
mypkg/foo.go:10:9: "x" escapes to heap
`)
		sites := ParseEscapeOutput(input)
		require.Len(t, sites, 2)

		// leaking param appears first (no flow chain)
		assert.Equal(t, 8, sites[0].Line)
		assert.Equal(t, "leaking param: name", sites[0].Message)
		assert.Empty(t, sites[0].FlowChain)

		// heap escape with attached flow chain
		assert.Equal(t, 10, sites[1].Line)
		assert.Equal(t, `"x" escapes to heap`, sites[1].Message)
		require.Len(t, sites[1].FlowChain, 3)
		assert.Equal(t, `flow: {heap} <- &{storage for "x"}:`, sites[1].FlowChain[0])
		assert.Equal(t, `from "x" (spill) at mypkg/foo.go:10:9`, sites[1].FlowChain[1])
		assert.Equal(t, `from panic("x") (call parameter) at mypkg/foo.go:10:8`, sites[1].FlowChain[2])
	})

	t.Run("multiple_flow_blocks", func(t *testing.T) {
		// Two verbose blocks; summaries appear in reverse order (as the compiler emits them).
		input := []byte(`mypkg/foo.go:5:9: "a" escapes to heap in Bar:
mypkg/foo.go:5:9:   flow: {heap} <- a:
mypkg/foo.go:5:9:     from return a (return) at mypkg/foo.go:5:9
mypkg/foo.go:3:6: leaking param: p in Bar:
mypkg/foo.go:3:6:   flow: {heap} <- p:
mypkg/foo.go:3:6:     from Bar(p) (call parameter) at mypkg/foo.go:3:1
mypkg/foo.go:3:6: leaking param: p
mypkg/foo.go:5:9: "a" escapes to heap
`)
		sites := ParseEscapeOutput(input)
		require.Len(t, sites, 2)

		assert.Equal(t, 3, sites[0].Line)
		assert.Equal(t, "leaking param: p", sites[0].Message)
		require.Len(t, sites[0].FlowChain, 2)

		assert.Equal(t, 5, sites[1].Line)
		assert.Equal(t, `"a" escapes to heap`, sites[1].Message)
		require.Len(t, sites[1].FlowChain, 2)
	})
}

func TestEscapeSite(t *testing.T) {
	t.Run("classification", func(t *testing.T) {
		cases := []struct {
			msg           string
			heap, leaking bool
		}{
			{"bar escapes to heap", true, false},
			{"moved to heap: result", true, false},
			{"leaking param: buf", false, true},
			{"leaking param content: r", false, true},
			{"can inline Foo", false, false},
			{"inlining call to Foo", false, false},
			{"... argument does not escape", false, false},
		}
		for _, tc := range cases {
			s := EscapeSite{Message: tc.msg}
			assert.Equal(t, tc.heap, s.IsHeapEscape(), "IsHeapEscape(%q)", tc.msg)
			assert.Equal(t, tc.leaking, s.IsLeakingParam(), "IsLeakingParam(%q)", tc.msg)
		}
	})
}

func TestParseFuncBoundaries(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		src := `package p

func Foo() {}

func Bar(x int) int {
	return x + 1
}

type T struct{}

func (t T) Method() {}
`
		dir := t.TempDir()
		path := filepath.Join(dir, "p.go")
		require.NoError(t, os.WriteFile(path, []byte(src), 0644))

		funcs, err := ParseFuncBoundaries(path, nil)
		require.NoError(t, err)
		require.Len(t, funcs, 3)

		byName := make(map[string]FuncBoundary)
		for _, f := range funcs {
			byName[f.Name] = f
		}

		assert.Equal(t, 3, byName["Foo"].StartLine)
		assert.Equal(t, 3, byName["Foo"].EndLine)

		assert.Equal(t, 5, byName["Bar"].StartLine)
		assert.Equal(t, 7, byName["Bar"].EndLine)

		assert.Equal(t, 11, byName["Method"].StartLine)
		assert.Equal(t, 11, byName["Method"].EndLine)
	})

	t.Run("non_go_file", func(t *testing.T) {
		funcs, err := ParseFuncBoundaries("somefile.txt", nil)
		assert.NoError(t, err)
		assert.Nil(t, funcs)
	})

	t.Run("not_found", func(t *testing.T) {
		_, err := ParseFuncBoundaries("/nonexistent/path/file.go", nil)
		assert.Error(t, err)
	})

	t.Run("from_src", func(t *testing.T) {
		src := []byte(`package p

func Foo() {}

func Bar(x int) int {
	return x + 1
}
`)
		// filename need not exist on disk when src is provided
		funcs, err := ParseFuncBoundaries("virtual.go", src)
		require.NoError(t, err)
		require.Len(t, funcs, 2)
		assert.Equal(t, "Foo", funcs[0].Name)
		assert.Equal(t, "Bar", funcs[1].Name)
	})
}

func TestFindFunc(t *testing.T) {
	funcs := []FuncBoundary{
		{Name: "Foo", StartLine: 3, EndLine: 5},
		{Name: "Bar", StartLine: 8, EndLine: 15},
	}

	assert.Equal(t, "Foo", FindFunc(funcs, 3).Name)
	assert.Equal(t, "Foo", FindFunc(funcs, 5).Name)
	assert.Equal(t, "Bar", FindFunc(funcs, 10).Name)
	assert.Nil(t, FindFunc(funcs, 6))  // gap between functions
	assert.Nil(t, FindFunc(funcs, 1))  // before first function
	assert.Nil(t, FindFunc(funcs, 20)) // after last function
}
