package engine

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

// ApplyUnifiedDiff applies the hunks in diffData that affect relPath to lines
// (each element is one source line, no trailing newline) and returns the result.
// relPath is matched against the "+++ b/<path>" diff header.
// If no matching hunks are found the original lines are returned unchanged.
func ApplyUnifiedDiff(lines []string, diffData []byte, relPath string) []string {
	hunks := parseHunksForFile(diffData, relPath)
	if len(hunks) == 0 {
		return lines
	}
	return applyHunks(lines, hunks)
}

type diffHunk struct {
	oldStart int
	lines    []diffHunkLine
}

type diffHunkLine struct {
	op   byte // ' ' context, '-' removed, '+' added
	text string
}

func parseHunksForFile(diffData []byte, relPath string) []diffHunk {
	var result []diffHunk
	inFile := false
	var cur *diffHunk

	scanner := bufio.NewScanner(bytes.NewReader(diffData))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "diff ") {
			inFile = false
			cur = nil
			continue
		}
		if strings.HasPrefix(line, "+++ ") {
			path := strings.TrimPrefix(line[4:], "b/")
			inFile = path == relPath
			cur = nil
			continue
		}
		if strings.HasPrefix(line, "--- ") || !inFile {
			continue
		}
		if strings.HasPrefix(line, "@@ ") {
			h := parseHunkHeader(line)
			if h != nil {
				result = append(result, *h)
				cur = &result[len(result)-1]
			}
			continue
		}
		if cur == nil {
			continue
		}
		if len(line) == 0 {
			cur.lines = append(cur.lines, diffHunkLine{' ', ""})
			continue
		}
		op := line[0]
		switch op {
		case ' ', '-', '+':
			cur.lines = append(cur.lines, diffHunkLine{op, line[1:]})
		case '\\':
			// "\ No newline at end of file" — ignore
		}
	}
	return result
}

func parseHunkHeader(line string) *diffHunk {
	// "@@ -oldStart[,oldCount] +newStart[,newCount] @@"
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return nil
	}
	old := strings.TrimPrefix(parts[1], "-")
	h := &diffHunk{}
	if comma := strings.Index(old, ","); comma >= 0 {
		h.oldStart, _ = strconv.Atoi(old[:comma])
	} else {
		h.oldStart, _ = strconv.Atoi(old)
	}
	if h.oldStart < 1 {
		h.oldStart = 1
	}
	return h
}

func applyHunks(lines []string, hunks []diffHunk) []string {
	var result []string
	origPos := 1 // 1-based position in original lines

	for _, h := range hunks {
		for origPos < h.oldStart && origPos <= len(lines) {
			result = append(result, lines[origPos-1])
			origPos++
		}
		for _, hl := range h.lines {
			switch hl.op {
			case ' ':
				if origPos <= len(lines) {
					result = append(result, lines[origPos-1])
				}
				origPos++
			case '-':
				origPos++
			case '+':
				result = append(result, hl.text)
			}
		}
	}
	for origPos <= len(lines) {
		result = append(result, lines[origPos-1])
		origPos++
	}
	return result
}
