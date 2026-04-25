package engine

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/util"
	"golang.org/x/sys/execabs"
)

const InlineFilename = "inline.txt"

// InlineKind classifies an inline analysis message.
type InlineKind int

const (
	InlineCannotInline InlineKind = iota // cannot inline Foo: reason  (default filter)
	InlineInliningCall                   // inlining call to pkg.Foo   (call was inlined)
	InlineCanInline                      // can inline Foo             (eligible but not called here)
)

// InlineSite is one parsed inline-analysis message.
type InlineSite struct {
	File    string
	Line    int
	Col     int
	Message string
}

// Kind classifies the site by its message prefix.
func (s InlineSite) Kind() InlineKind {
	switch {
	case strings.HasPrefix(s.Message, "cannot inline"):
		return InlineCannotInline
	case strings.HasPrefix(s.Message, "inlining call to"):
		return InlineInliningCall
	default:
		return InlineCanInline
	}
}

// isInlineMessage reports whether a compiler message belongs to the inline
// analysis category (as opposed to escape analysis).
func isInlineMessage(msg string) bool {
	return strings.HasPrefix(msg, "can inline") ||
		strings.HasPrefix(msg, "cannot inline") ||
		strings.HasPrefix(msg, "inlining call to")
}

// RecordCompilerAnalysis runs go build -gcflags=./...=-m=2 once and writes
// the filtered output to escape.txt and/or inline.txt depending on the flags.
// When both are requested the single build pass is split by message type.
func RecordCompilerAnalysis(ctx context.Context, sourcesRoot string, storage billy.Filesystem, sessionID string, wantEscape, wantInline bool) error {
	if !wantEscape && !wantInline {
		return nil
	}

	dir := filepath.Join(sessionDir, sessionID)
	if err := storage.MkdirAll(dir, 0755); err != nil {
		return err
	}

	var buf bytes.Buffer
	cmd := execabs.CommandContext(ctx, "go", "build", "-gcflags=./...=-m=2", "./...")
	cmd.Dir = sourcesRoot
	cmd.Stderr = &buf
	_ = cmd.Run()

	if buf.Len() == 0 {
		return nil
	}

	escapeData, inlineData := splitCompilerOutput(buf.Bytes())

	if wantEscape && len(escapeData) > 0 {
		if err := util.WriteFile(storage, filepath.Join(dir, EscapeFilename), escapeData, 0644); err != nil {
			return err
		}
	}
	if wantInline && len(inlineData) > 0 {
		if err := util.WriteFile(storage, filepath.Join(dir, InlineFilename), inlineData, 0644); err != nil {
			return err
		}
	}
	return nil
}

// splitCompilerOutput partitions raw -m=2 compiler output into escape and
// inline lines. Continuation lines (indented) follow the same destination as
// the preceding non-continuation line.
func splitCompilerOutput(data []byte) (escapeData, inlineData []byte) {
	var escapeBuf, inlineBuf bytes.Buffer
	var currentDest *bytes.Buffer

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		m := escapeLineRe.FindStringSubmatch(line)
		if m == nil {
			// Package header or build error — keep in escape output.
			escapeBuf.WriteString(line + "\n")
			continue
		}
		msg := m[4]
		switch {
		case strings.HasPrefix(msg, " "):
			// Continuation/flow line — same destination as enclosing block.
			if currentDest != nil {
				currentDest.WriteString(line + "\n")
			}
		case isInlineMessage(msg):
			inlineBuf.WriteString(line + "\n")
			currentDest = &inlineBuf
		default:
			escapeBuf.WriteString(line + "\n")
			currentDest = &escapeBuf
		}
	}
	return escapeBuf.Bytes(), inlineBuf.Bytes()
}

// ReadInlineAnalysis reads and parses the saved inline analysis for a session.
func ReadInlineAnalysis(storage billy.Filesystem, sessionPath string) ([]InlineSite, error) {
	data, err := util.ReadFile(storage, filepath.Join(sessionPath, InlineFilename))
	if err != nil {
		return nil, fmt.Errorf("no inline analysis recorded for this session")
	}
	return ParseInlineOutput(data), nil
}

// ParseInlineOutput parses raw inline analysis output into InlineSite entries.
// Duplicate lines are silently dropped.
func ParseInlineOutput(data []byte) []InlineSite {
	var sites []InlineSite
	seen := make(map[string]struct{})

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		raw := scanner.Text()
		m := escapeLineRe.FindStringSubmatch(raw)
		if m == nil {
			continue
		}
		msg := m[4]
		if !isInlineMessage(msg) {
			continue
		}
		if _, dup := seen[raw]; dup {
			continue
		}
		seen[raw] = struct{}{}
		var lineNum, colNum int
		fmt.Sscanf(m[2], "%d", &lineNum)
		fmt.Sscanf(m[3], "%d", &colNum)
		sites = append(sites, InlineSite{
			File:    m[1],
			Line:    lineNum,
			Col:     colNum,
			Message: msg,
		})
	}
	return sites
}
