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
)

const EscapeFilename = "escape.txt"

// EscapeSite is one line from the escape analysis output.
type EscapeSite struct {
	File      string // absolute path as emitted by the compiler
	Line      int
	Col       int
	Message   string   // everything after "file:line:col: "
	FlowChain []string // data-flow explanation lines from -m=2 (nil for -m output)
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
// If src is non-nil it is used as the file content instead of reading filename.
func ParseFuncBoundaries(filename string, src []byte) ([]FuncBoundary, error) {
	if !strings.HasSuffix(filename, ".go") {
		return nil, nil
	}
	fset := token.NewFileSet()
	var srcArg any
	if src != nil {
		srcArg = src
	}
	f, err := parser.ParseFile(fset, filename, srcArg, parser.SkipObjectResolution)
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

// RecordEscape captures escape analysis for a session.
// It delegates to RecordCompilerAnalysis with wantEscape=true.
func RecordEscape(ctx context.Context, sourcesRoot string, storage billy.Filesystem, sessionID string) error {
	return RecordCompilerAnalysis(ctx, sourcesRoot, storage, sessionID, true, false)
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
//
// With -m=2, the compiler emits verbose flow-chain blocks before each summary
// line. These are attached to the corresponding EscapeSite as FlowChain.
// -m output (no flow chains) is handled transparently.
func ParseEscapeOutput(data []byte) []EscapeSite {
	var sites []EscapeSite
	seen := make(map[string]struct{})
	// pendingFlow accumulates flow-chain lines keyed by "file:line:col".
	// Verbose blocks may appear in a different order than their summary lines,
	// so we collect all blocks first and attach on the matching summary.
	pendingFlow := make(map[string][]string)
	var currentFlowKey string

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		raw := scanner.Text()
		m := escapeLineRe.FindStringSubmatch(raw)
		if m == nil {
			continue
		}
		file, lineStr, colStr, msg := m[1], m[2], m[3], m[4]

		switch {
		case strings.HasPrefix(msg, " "):
			// Continuation line: flow step or from-clause, indented by the compiler.
			if currentFlowKey != "" {
				pendingFlow[currentFlowKey] = append(pendingFlow[currentFlowKey], strings.TrimSpace(msg))
			}
		case strings.HasSuffix(msg, ":"):
			// Verbose block header: e.g. '"x" escapes to heap in Foo:'
			// Ends with ":" because the compiler appends " in FuncName:" or
			// " for FuncName with derefs=N:". Start accumulating for this site.
			currentFlowKey = file + ":" + lineStr + ":" + colStr
		default:
			// Normal summary line — may follow a verbose block for the same location.
			currentFlowKey = ""
			if _, dup := seen[raw]; dup {
				continue
			}
			seen[raw] = struct{}{}
			var lineNum, colNum int
			fmt.Sscanf(lineStr, "%d", &lineNum)
			fmt.Sscanf(colStr, "%d", &colNum)
			key := file + ":" + lineStr + ":" + colStr
			site := EscapeSite{
				File:    file,
				Line:    lineNum,
				Col:     colNum,
				Message: msg,
			}
			if flows := pendingFlow[key]; len(flows) > 0 {
				site.FlowChain = flows
				delete(pendingFlow, key)
			}
			sites = append(sites, site)
		}
	}
	return sites
}
