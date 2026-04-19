package engine

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/util"
	"golang.org/x/sys/execabs"
)

const EscapeFilename = "escape.txt"

// EscapeSite is one line from the escape analysis output.
type EscapeSite struct {
	File    string // absolute path as emitted by the compiler
	Line    int
	Col     int
	Message string // everything after "file:line:col: "
}

// IsHeapEscape reports whether the message describes a value escaping to the heap.
func (e EscapeSite) IsHeapEscape() bool {
	return strings.Contains(e.Message, "escapes to heap") ||
		strings.Contains(e.Message, "moved to heap")
}

// IsLeakingParam reports whether the message describes a parameter that causes
// the caller's value to escape to the heap.
func (e EscapeSite) IsLeakingParam() bool {
	return strings.HasPrefix(e.Message, "leaking param")
}

// FuncBoundary records the name and line extent of a Go function declaration.
type FuncBoundary struct {
	Name      string
	StartLine int
	EndLine   int
}

// ParseFuncBoundaries parses a Go source file and returns the line boundaries
// of every function declaration. Returns nil (not an error) for non-Go files.
func ParseFuncBoundaries(filename string) ([]FuncBoundary, error) {
	if !strings.HasSuffix(filename, ".go") {
		return nil, nil
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	var funcs []FuncBoundary
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		start := fset.Position(fd.Pos())
		end := fset.Position(fd.End())
		funcs = append(funcs, FuncBoundary{
			Name:      fd.Name.Name,
			StartLine: start.Line,
			EndLine:   end.Line,
		})
	}
	return funcs, nil
}

// FindFunc returns the FuncBoundary that contains line, or nil if none does.
func FindFunc(funcs []FuncBoundary, line int) *FuncBoundary {
	for i := range funcs {
		if line >= funcs[i].StartLine && line <= funcs[i].EndLine {
			return &funcs[i]
		}
	}
	return nil
}

// RecordEscape runs go build with escape-analysis flags in the sources root,
// captures stderr, and writes the result to the session directory.
func RecordEscape(ctx context.Context, sourcesRoot string, storage billy.Filesystem, sessionID string) error {
	dir := filepath.Join(sessionDir, sessionID)
	if err := storage.MkdirAll(dir, 0755); err != nil {
		return err
	}

	var buf bytes.Buffer
	cmd := execabs.CommandContext(ctx, "go", "build", "-gcflags=./...=-m", "./...")
	cmd.Dir = sourcesRoot
	cmd.Stderr = &buf

	_ = cmd.Run() // non-zero exit is fine (build errors go to stderr too)

	if buf.Len() == 0 {
		return nil
	}
	return util.WriteFile(storage, filepath.Join(dir, EscapeFilename), buf.Bytes(), 0644)
}

// ReadEscapeAnalysis reads and parses the saved escape analysis for a session.
func ReadEscapeAnalysis(storage billy.Filesystem, sessionPath string) ([]EscapeSite, error) {
	data, err := util.ReadFile(storage, filepath.Join(sessionPath, EscapeFilename))
	if err != nil {
		return nil, fmt.Errorf("no escape analysis recorded for this session")
	}
	return ParseEscapeOutput(data), nil
}

var escapeLineRe = regexp.MustCompile(`^(.+):(\d+):(\d+): (.+)$`)

// ParseEscapeOutput parses the raw compiler output into EscapeSite entries.
// Duplicate lines (same file:line:col:message) are silently dropped — the
// compiler can emit the same diagnostic twice when multiple build targets
// share the same package.
func ParseEscapeOutput(data []byte) []EscapeSite {
	var sites []EscapeSite
	seen := make(map[string]struct{})
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		raw := scanner.Text()
		m := escapeLineRe.FindStringSubmatch(raw)
		if m == nil {
			continue
		}
		if _, dup := seen[raw]; dup {
			continue
		}
		seen[raw] = struct{}{}
		var lineNum, colNum int
		fmt.Sscanf(m[2], "%d", &lineNum)
		fmt.Sscanf(m[3], "%d", &colNum)
		sites = append(sites, EscapeSite{
			File:    m[1],
			Line:    lineNum,
			Col:     colNum,
			Message: m[4],
		})
	}
	return sites
}
