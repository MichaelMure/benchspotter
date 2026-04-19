package engine

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseEscapeOutput_basic(t *testing.T) {
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
}

func TestParseEscapeOutput_deduplication(t *testing.T) {
	// compiler emits the same line multiple times when building with ./...
	input := []byte(`engine/foo.go:10:5: bar escapes to heap
engine/foo.go:10:5: bar escapes to heap
engine/foo.go:20:3: leaking param: buf
`)
	sites := ParseEscapeOutput(input)
	assert.Len(t, sites, 2)
}

func TestParseEscapeOutput_empty(t *testing.T) {
	assert.Empty(t, ParseEscapeOutput(nil))
	assert.Empty(t, ParseEscapeOutput([]byte("# only headers\n# no sites\n")))
}

func TestEscapeSite_classification(t *testing.T) {
	cases := []struct {
		msg          string
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
}

func TestParseFuncBoundaries(t *testing.T) {
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

	funcs, err := ParseFuncBoundaries(path)
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
}

func TestParseFuncBoundaries_nonGoFile(t *testing.T) {
	funcs, err := ParseFuncBoundaries("somefile.txt")
	assert.NoError(t, err)
	assert.Nil(t, funcs)
}

func TestParseFuncBoundaries_notFound(t *testing.T) {
	_, err := ParseFuncBoundaries("/nonexistent/path/file.go")
	assert.Error(t, err)
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
