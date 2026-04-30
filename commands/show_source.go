package commands

import (
	"bytes"
	"cmp"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
	"benchspotter/repository"
)

// sourceViewBase holds the source-reading infrastructure shared by show escape
// and show inline. Embed it in a view model to get caching, git-aware source
// lookup, and the shared rendering primitives.
type sourceViewBase struct {
	env         *execenv.Env
	sourcesRoot string
	gitCommit   string
	gitDiff     []byte
	git         repository.GitSource
	rawCache    map[string][]byte
	fileCache   map[string][]string
	funcCache   map[string][]engine.FuncBoundary

	// navigation state
	yOffset        int
	viewportHeight int
	headers        []headerEntry
	pendingJump    int
	hasPendingJump bool

	// search state
	searchMode  bool
	searchQuery string
	searchIdx   int
}

type headerEntry struct {
	line  int    // 0-based line index in rendered content (set by updateHeaderLines)
	file  string // display path (relPath)
	fn    string // function name, or "" for file-level groups
	count int    // visible site count for this group
}

// SetCurrentYOffset implements YOffsetSetter. Called by the viewport before
// each Status() render so the breadcrumb reflects the live scroll position.
func (b *sourceViewBase) SetCurrentYOffset(yOffset, viewportHeight int) {
	b.yOffset = yOffset
	b.viewportHeight = viewportHeight
}

// ConsumeJumpOffset implements JumpRequester.
func (b *sourceViewBase) ConsumeJumpOffset() (int, bool) {
	if b.hasPendingJump {
		b.hasPendingJump = false
		return b.pendingJump, true
	}
	return 0, false
}

func (b *sourceViewBase) jumpNext() {
	for _, h := range b.headers {
		if h.line > b.yOffset {
			b.pendingJump = h.line
			b.hasPendingJump = true
			return
		}
	}
}

func (b *sourceViewBase) jumpPrev() {
	for i := len(b.headers) - 1; i >= 0; i-- {
		if b.headers[i].line < b.yOffset {
			b.pendingJump = b.headers[i].line
			b.hasPendingJump = true
			return
		}
	}
}

// updateHeaderLines scans rendered content to fill in the line-number field
// of each entry in b.headers. Call after Render() has populated b.headers
// with file/fn/count data and returned the content string.
func (b *sourceViewBase) updateHeaderLines(content string) {
	idx := 0
	for i, line := range strings.Split(content, "\n") {
		if idx >= len(b.headers) {
			break
		}
		if isGroupHeader(line) {
			b.headers[idx].line = i
			idx++
		}
	}
}

// currentBreadcrumb returns the label of the group whose header is at or
// above the current scroll position, or "" if none has been scrolled past.
func (b *sourceViewBase) currentBreadcrumb() string {
	var h *headerEntry
	for i := range b.headers {
		if b.headers[i].line <= b.yOffset {
			h = &b.headers[i]
		} else {
			break
		}
	}
	if h == nil {
		return ""
	}
	if h.fn != "" {
		return h.file + " · func " + h.fn
	}
	return h.file
}

// AdjustYOffset implements execenv.YOffsetAdjuster. It anchors on the source
// line nearest the middle of the viewport so that the middle of the screen
// stays stable when content length changes (e.g. toggling flow chains).
func (b *sourceViewBase) AdjustYOffset(oldContent, newContent string, oldOffset, viewportHeight int) int {
	mid := oldOffset + viewportHeight/2
	oldLines := strings.Split(oldContent, "\n")

	anchor := ""
	for delta := 0; delta < len(oldLines); delta++ {
		if i := mid + delta; i < len(oldLines) && isSourceLine(oldLines[i]) {
			anchor = oldLines[i]
			break
		}
		if i := mid - delta; i >= oldOffset && isSourceLine(oldLines[i]) {
			anchor = oldLines[i]
			break
		}
	}
	if anchor == "" {
		return oldOffset
	}
	for i, l := range strings.Split(newContent, "\n") {
		if l == anchor {
			return i - viewportHeight/2
		}
	}
	return oldOffset
}

// renderFunctionBody renders the body of bound with source annotations.
// renderLine is called after each source line with its 1-based line number.
// fallback is called when source is unavailable.
func (b *sourceViewBase) renderFunctionBody(
	sb *strings.Builder,
	absFile string,
	bound *engine.FuncBoundary,
	fallback func(*strings.Builder),
	renderLine func(*strings.Builder, int),
) {
	lines := b.fileLines(absFile)
	if len(lines) == 0 {
		fallback(sb)
		return
	}

	start, end := bound.StartLine, bound.EndLine
	if end > len(lines) {
		end = len(lines)
	}
	if start < 1 {
		start = 1
	}
	if start > end {
		start = end
	}

	raw := make([]string, end-start+1)
	for i, l := range lines[start-1 : end] {
		raw[i] = expandTabs(l, 4)
	}
	for i, hl := range highlightGoLines(raw) {
		lineNo := start + i
		fmt.Fprintf(sb, "  %4d  %s\n", lineNo, hl)
		renderLine(sb, lineNo)
	}
}

// renderContextLines renders ±ctx source lines around each line number in
// lineNums, merging overlapping windows.
// renderLine is called after each source line with its 1-based line number.
// fallback is called when source is unavailable.
func (b *sourceViewBase) renderContextLines(
	sb *strings.Builder,
	absFile string,
	lineNums []int,
	fallback func(*strings.Builder),
	renderLine func(*strings.Builder, int),
) {
	const ctx = 4
	lines := b.fileLines(absFile)
	if len(lines) == 0 {
		fallback(sb)
		return
	}

	type span struct{ start, end int }
	var spans []span
	for _, ln := range lineNums {
		lo, hi := ln-ctx, ln+ctx
		if lo < 1 {
			lo = 1
		}
		if hi > len(lines) {
			hi = len(lines)
		}
		if lo > hi {
			lo = hi
		}
		if len(spans) > 0 && lo <= spans[len(spans)-1].end+1 {
			if hi > spans[len(spans)-1].end {
				spans[len(spans)-1].end = hi
			}
		} else {
			spans = append(spans, span{lo, hi})
		}
	}

	for i, sp := range spans {
		if i > 0 {
			fmt.Fprintf(sb, "  %4s  …\n", "")
		}
		raw := make([]string, sp.end-sp.start+1)
		for j, l := range lines[sp.start-1 : sp.end] {
			raw[j] = expandTabs(l, 4)
		}
		for j, hl := range highlightGoLines(raw) {
			lineNo := sp.start + j
			fmt.Fprintf(sb, "  %4d  %s\n", lineNo, hl)
			renderLine(sb, lineNo)
		}
	}
}

func (b *sourceViewBase) rawContent(absFile string) []byte {
	if data, ok := b.rawCache[absFile]; ok {
		return data
	}
	var data []byte
	if b.git != nil && b.gitCommit != "" && isProjectFile(b.sourcesRoot, absFile) {
		if rel := b.gitRelPath(absFile); rel != "" {
			if raw, err := b.git.FileAtCommit(b.env.Ctx, b.gitCommit, rel); err == nil {
				if len(b.gitDiff) > 0 {
					ls := splitFileContent(raw)
					ls = engine.ApplyUnifiedDiff(ls, b.gitDiff, rel)
					data = []byte(strings.Join(ls, "\n"))
				} else {
					data = raw
				}
			}
		}
	}
	if data == nil {
		rel, err := filepath.Rel(b.sourcesRoot, absFile)
		if err != nil || strings.HasPrefix(rel, "..") {
			rel = absFile
		}
		data, err = os.ReadFile(filepath.Join(b.sourcesRoot, rel))
		if err != nil {
			data, _ = os.ReadFile(absFile)
		}
	}
	b.rawCache[absFile] = data
	return data
}

func (b *sourceViewBase) fileLines(absFile string) []string {
	if lines, ok := b.fileCache[absFile]; ok {
		return lines
	}
	lines := splitFileContent(b.rawContent(absFile))
	b.fileCache[absFile] = lines
	return lines
}

func (b *sourceViewBase) funcBoundaries(absFile string) []engine.FuncBoundary {
	if bounds, ok := b.funcCache[absFile]; ok {
		return bounds
	}
	bounds, _ := engine.ParseFuncBoundaries(b.resolveFullPath(absFile), b.rawContent(absFile))
	b.funcCache[absFile] = bounds
	return bounds
}

func (b *sourceViewBase) gitRelPath(absFile string) string {
	if filepath.IsAbs(absFile) {
		rel, err := filepath.Rel(b.sourcesRoot, absFile)
		if err != nil || strings.HasPrefix(rel, "..") {
			return ""
		}
		return rel
	}
	clean := filepath.Clean(absFile)
	if strings.HasPrefix(clean, "..") {
		return ""
	}
	return clean
}

func (b *sourceViewBase) resolveFullPath(absFile string) string {
	if filepath.IsAbs(absFile) {
		return absFile
	}
	return filepath.Join(b.sourcesRoot, filepath.Clean(absFile))
}

func (b *sourceViewBase) relPath(file string) string {
	if !filepath.IsAbs(file) {
		return filepath.Clean(file)
	}
	rel, err := filepath.Rel(b.sourcesRoot, file)
	if err != nil || strings.HasPrefix(rel, "..") {
		return file
	}
	return rel
}

// ── Shared utilities ──────────────────────────────────────────────────────────

func splitFileContent(data []byte) []string {
	if len(data) == 0 {
		return nil
	}
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func expandTabs(s string, tabWidth int) string {
	var b strings.Builder
	col := 0
	for _, ch := range s {
		if ch == '\t' {
			spaces := tabWidth - col%tabWidth
			b.WriteString(strings.Repeat(" ", spaces))
			col += spaces
		} else {
			b.WriteRune(ch)
			col++
		}
	}
	return b.String()
}

func highlightGoLines(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}
	lexer := cmp.Or(lexers.Get("go"), lexers.Fallback)
	style := cmp.Or(styles.Get("monokai"), styles.Fallback)
	var buf bytes.Buffer
	it, err := lexer.Tokenise(nil, strings.Join(lines, "\n"))
	if err != nil {
		return lines
	}
	if err := formatters.TTY16m.Format(&buf, style, it); err != nil {
		return lines
	}
	out := strings.Split(buf.String(), "\n")
	if len(out) > len(lines) {
		out = out[:len(lines)]
	}
	const reset = "\x1b[0m"
	for i, l := range out {
		if !strings.HasSuffix(l, reset) {
			out[i] = l + reset
		}
	}
	return out
}

// ── Search ────────────────────────────────────────────────────────────────────

// IsModal implements execenv.ModalModel so Esc is forwarded to the model
// rather than quitting the application while the search prompt is open.
func (b *sourceViewBase) IsModal() bool { return b.searchMode }

// handleSearchKey processes a key press when in search mode, or intercepts
// ctrl+s to enter search mode. Returns true when the key was consumed.
// It always sets a pending jump so the viewport does not also process the key.
func (b *sourceViewBase) handleSearchKey(key string) bool {
	if !b.searchMode {
		if key == "ctrl+s" {
			b.searchMode = true
			b.searchQuery = ""
			b.searchIdx = 0
			b.pendingJump = b.yOffset
			b.hasPendingJump = true
			return true
		}
		return false
	}

	// Inside search mode every key is consumed; set a no-op jump by default so
	// the key doesn't fall through to the underlying viewport.
	defer func() {
		if !b.hasPendingJump {
			b.pendingJump = b.yOffset
			b.hasPendingJump = true
		}
	}()

	switch key {
	case "esc", "ctrl+s":
		b.searchMode = false
		b.searchQuery = ""
	case "enter":
		b.jumpToSearchMatch()
		b.searchMode = false
		b.searchQuery = ""
	case "backspace", "ctrl+h":
		if r := []rune(b.searchQuery); len(r) > 0 {
			b.searchQuery = string(r[:len(r)-1])
			b.searchIdx = 0
		}
	case "tab":
		if n := b.searchMatchCount(); n > 0 && b.searchIdx < n-1 {
			b.searchIdx++
			b.jumpToSearchMatch()
		}
	case "shift+tab":
		if b.searchIdx > 0 {
			b.searchIdx--
			b.jumpToSearchMatch()
		}
	default:
		if r := []rune(key); len(r) == 1 && r[0] >= 32 && r[0] != 127 {
			b.searchQuery += key
			b.searchIdx = 0
		}
	}
	return true
}

func (b *sourceViewBase) searchMatchCount() int {
	if b.searchQuery == "" {
		return len(b.headers)
	}
	n := 0
	for i := range b.headers {
		if b.searchMatch(i) {
			n++
		}
	}
	return n
}

func (b *sourceViewBase) searchMatch(i int) bool {
	h := &b.headers[i]
	return fuzzyMatch(b.searchQuery, h.file) || fuzzyMatch(b.searchQuery, h.fn)
}

func (b *sourceViewBase) jumpToSearchMatch() {
	count := 0
	for i := range b.headers {
		if b.searchQuery == "" || b.searchMatch(i) {
			if count == b.searchIdx {
				b.pendingJump = b.headers[i].line
				b.hasPendingJump = true
				return
			}
			count++
		}
	}
}

// searchStatusLine returns the footer string to show while in search mode, or
// "" when search is inactive.
func (b *sourceViewBase) searchStatusLine() string {
	if !b.searchMode {
		return ""
	}
	n := b.searchMatchCount()
	var info string
	switch {
	case b.searchQuery == "":
		info = ""
	case n == 0:
		info = "  (no match)"
	default:
		idx := b.searchIdx
		if idx >= n {
			idx = n - 1
		}
		info = fmt.Sprintf("  (%d/%d)", idx+1, n)
	}
	return fmt.Sprintf("search: %s█%s    [⇥] next  [⇤] prev  [↵] jump  [esc] cancel", b.searchQuery, info)
}

func fuzzyMatch(query, target string) bool {
	if query == "" {
		return true
	}
	qi, q, t := 0, []rune(strings.ToLower(query)), []rune(strings.ToLower(target))
	for _, r := range t {
		if qi < len(q) && r == q[qi] {
			qi++
		}
	}
	return qi == len(q)
}

// ── Group helpers ─────────────────────────────────────────────────────────────

// funcKey identifies a group of diagnostic sites sharing the same context:
// the absolute file path and the name of the enclosing function. fn is ""
// for sites that fall outside any parsed function boundary (e.g. package-level
// variable initialisers).
type funcKey struct{ file, fn string }

// siteGroup collects all visible diagnostic sites that share the same
// (file, enclosing-function) context. bound is the parsed function boundary
// used for full-body source rendering; it is nil when fn is "" (no enclosing
// function found), in which case context-window rendering is used instead.
type siteGroup[S any] struct {
	key   funcKey
	bound *engine.FuncBoundary
	sites []S
}

// buildGroups groups visible sites by (file, enclosing-func) and sorts them by
// file then by function start line.
func buildGroups[S any](b *sourceViewBase, visible []S, getFile func(S) string, getLine func(S) int) []siteGroup[S] {
	var groups []siteGroup[S]
	groupIdx := make(map[funcKey]int)
	for _, s := range visible {
		file, line := getFile(s), getLine(s)
		fn := engine.FindFunc(b.funcBoundaries(file), line)
		key := funcKey{file: file}
		if fn != nil {
			key.fn = fn.Name
		}
		i, ok := groupIdx[key]
		if !ok {
			i = len(groups)
			groups = append(groups, siteGroup[S]{key: key, bound: fn})
			groupIdx[key] = i
		}
		groups[i].sites = append(groups[i].sites, s)
	}
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].key.file != groups[j].key.file {
			return groups[i].key.file < groups[j].key.file
		}
		if groups[i].bound != nil && groups[j].bound != nil {
			return groups[i].bound.StartLine < groups[j].bound.StartLine
		}
		return groups[i].key.fn < groups[j].key.fn
	})
	return groups
}

// writeGroupHeader writes the  ──  file · func (N) ──  separator line.
func (b *sourceViewBase) writeGroupHeader(sb *strings.Builder, dispFile, fn string, count int) {
	var title string
	badge := fmt.Sprintf(" (%d)", count)
	if fn != "" {
		title = fmt.Sprintf(" %s · func %s%s ", dispFile, fn, badge)
	} else {
		title = fmt.Sprintf(" %s%s ", dispFile, badge)
	}
	const lineWidth = 72
	left := 2
	right := lineWidth - left - len(title)
	if right < 2 {
		right = 2
	}
	fmt.Fprintf(sb, "%s%s%s\n", strings.Repeat("─", left), title, strings.Repeat("─", right))
}

// isGroupHeader reports whether a rendered line is a group separator written
// by writeGroupHeader (starts with two U+2500 BOX DRAWINGS LIGHT HORIZONTAL).
func isGroupHeader(line string) bool {
	return strings.HasPrefix(line, "──")
}

// truncateLeft truncates s to maxLen runes, replacing the removed prefix with
// "…". Useful for file paths where the filename at the end matters most.
func truncateLeft(s string, maxLen int) string {
	r := []rune(s)
	if len(r) <= maxLen {
		return s
	}
	return "…" + string(r[len(r)-(maxLen-1):])
}

// ── Sidebar ───────────────────────────────────────────────────────────────────

const sidebarMinTermWidth = 80

func (b *sourceViewBase) SidebarWidth(totalWidth int) int {
	if totalWidth < sidebarMinTermWidth {
		return 0
	}
	w := totalWidth / 4
	if w < 20 {
		w = 20
	}
	if w > 48 {
		w = 48
	}
	return w
}

func (b *sourceViewBase) RenderSidebar(width, height int) string {
	inner := width - 1 // last column is the │ divider
	divider := b.env.Style.SidebarDivider("│")

	// In search mode, find which header index is the selected match.
	selectedIdx := -1
	if b.searchMode && b.searchQuery != "" {
		count := 0
		for i := range b.headers {
			if b.searchMatch(i) {
				if count == b.searchIdx {
					selectedIdx = i
					break
				}
				count++
			}
		}
	}

	// When not in search mode, find the scroll-based active header.
	var activeHeader *headerEntry
	if !b.searchMode {
		for i := range b.headers {
			if b.headers[i].line <= b.yOffset {
				activeHeader = &b.headers[i]
			} else {
				break
			}
		}
	}

	type entry struct {
		text      string
		prominent bool // auto-scroll to keep this visible
	}
	var entries []entry
	prominentIdx := -1

	applyStyle := func(text string, isActive, isSelected, isMatch, isFn bool) string {
		switch {
		case isSelected:
			return b.env.Style.SidebarSelected(text)
		case isActive:
			return b.env.Style.SidebarActive(text)
		case isMatch:
			if isFn {
				return b.env.Style.SidebarFn(text)
			}
			return b.env.Style.SidebarFile(text)
		default:
			return b.env.Style.SidebarDim(text)
		}
	}

	prevFile := ""
	for i := range b.headers {
		h := &b.headers[i]
		isActive := activeHeader != nil && &b.headers[i] == activeHeader
		isSelected := selectedIdx == i
		isMatch := !b.searchMode || b.searchQuery == "" || b.searchMatch(i)

		if h.file != prevFile {
			prevFile = h.file
			fileHasMatch := isMatch
			if !fileHasMatch {
				for j := i + 1; j < len(b.headers) && b.headers[j].file == h.file; j++ {
					if b.searchMatch(j) {
						fileHasMatch = true
						break
					}
				}
			}
			prominent := (isActive || isSelected) && h.fn == ""
			text := applyStyle(" "+truncateLeft(h.file, inner-1), isActive && h.fn == "", isSelected && h.fn == "", fileHasMatch, false)
			entries = append(entries, entry{text: text, prominent: prominent})
			if prominent {
				prominentIdx = len(entries) - 1
			}
		}

		if h.fn != "" {
			badge := fmt.Sprintf(" (%d)", h.count)
			name := "· " + h.fn
			maxName := inner - 2 - len(badge)
			if maxName < 3 {
				maxName = 3
			}
			if runes := []rune(name); len(runes) > maxName {
				name = string(runes[:maxName-1]) + "…"
			}
			prominent := isActive || isSelected
			text := applyStyle("  "+name+badge, isActive, isSelected, isMatch, true)
			entries = append(entries, entry{text: text, prominent: prominent})
			if prominent {
				prominentIdx = len(entries) - 1
			}
		}
	}

	// Auto-scroll to keep the prominent entry near the middle.
	scrollStart := 0
	if prominentIdx >= 0 {
		scrollStart = prominentIdx - height/2
		if scrollStart < 0 {
			scrollStart = 0
		}
	}

	var sb strings.Builder
	for row := 0; row < height; row++ {
		if row > 0 {
			sb.WriteByte('\n')
		}
		idx := scrollStart + row
		var line string
		if idx < len(entries) {
			line = lipgloss.NewStyle().Width(inner).MaxWidth(inner).Render(entries[idx].text)
		} else {
			line = lipgloss.NewStyle().Width(inner).Render("")
		}
		sb.WriteString(line + divider)
	}
	return sb.String()
}

// isSourceLine reports whether a rendered line has the "  %4d  …" source format.
func isSourceLine(line string) bool {
	if len(line) < 8 || line[0] != ' ' || line[1] != ' ' {
		return false
	}
	numPart := strings.TrimLeft(line[2:6], " ")
	if numPart == "" {
		return false
	}
	for _, c := range numPart {
		if c < '0' || c > '9' {
			return false
		}
	}
	return line[6] == ' ' && line[7] == ' '
}
