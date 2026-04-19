package commands

import (
	"bytes"
	"cmp"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

type showEscapeOptions struct {
	session     string
	all         bool
	includeDeps bool
}

func newShowEscapeCommand(env *execenv.Env) *cobra.Command {
	options := showEscapeOptions{}

	cmd := &cobra.Command{
		Use:     "escape",
		Short:   "Show compiler escape analysis recorded for a session",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShowEscape(cmd.Context(), env, options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.session, "session", "", "Session ID to inspect")
	flags.BoolVar(&options.all, "all", false, "Show all compiler notes, not just heap escapes and leaking params")
	flags.BoolVar(&options.includeDeps, "deps", false, "Include stdlib and dependencies (default: project code only)")

	return cmd
}

func runShowEscape(ctx context.Context, env *execenv.Env, options showEscapeOptions) error {
	var selection *engine.SessionInfo

	if options.session == "" {
		const recallKey = "show_escape_session"
		preSelected := env.Repo.GetRecall(recallKey)
		var err error
		selection, err = inputs.SelectSession(ctx, env, preSelected,
			func(info *engine.SessionInfo) bool { return info.HasProfile(engine.ProfileEscape) })
		if err != nil {
			return err
		}
		if err = env.Repo.SetRecall(recallKey, selection.Id); err != nil {
			return err
		}
	} else {
		sessions, err := engine.LocateSessions(env.Repo.Storage())
		if err != nil {
			return err
		}
		for _, s := range sessions {
			if s.Id == options.session {
				selection = s
				break
			}
		}
		if selection == nil {
			return fmt.Errorf("session %q not found", options.session)
		}
	}

	if !selection.HasProfile(engine.ProfileEscape) {
		return fmt.Errorf("session %q has no escape analysis recorded", selection.HumanName)
	}

	sites, err := engine.ReadEscapeAnalysis(env.Repo.Storage(), selection.Path)
	if err != nil {
		return err
	}

	sourcesRoot := env.Repo.Sources().Root()

	switch env.Format {
	case execenv.FormatJSON:
		type jsonSite struct {
			File         string `json:"file"`
			Line         int    `json:"line"`
			Col          int    `json:"col"`
			Message      string `json:"message"`
			HeapEscape   bool   `json:"heap_escape"`
			LeakingParam bool   `json:"leaking_param"`
		}
		out := make([]jsonSite, 0, len(sites))
		for _, s := range sites {
			if !options.all && !s.IsHeapEscape() && !s.IsLeakingParam() {
				continue
			}
			if !options.includeDeps && !isProjectFile(sourcesRoot, s.File) {
				continue
			}
			out = append(out, jsonSite{
				File:         s.File,
				Line:         s.Line,
				Col:          s.Col,
				Message:      s.Message,
				HeapEscape:   s.IsHeapEscape(),
				LeakingParam: s.IsLeakingParam(),
			})
		}
		return env.Out.PrintJSON(out)

	case execenv.FormatText:
		model := &escapeViewModel{
			sites:       sites,
			sourcesRoot: sourcesRoot,
			showAll:     options.all,
			projectOnly: !options.includeDeps,
			fileCache:   make(map[string][]string),
			funcCache:   make(map[string][]engine.FuncBoundary),
		}
		return env.ViewportWithKeys(ctx, model)()

	default:
		return fmt.Errorf("unsupported format %v for show escape (text, json)", env.Format)
	}
}

type escapeViewModel struct {
	sites       []engine.EscapeSite
	sourcesRoot string
	showAll     bool
	projectOnly bool
	fileCache   map[string][]string              // abs path → lines
	funcCache   map[string][]engine.FuncBoundary // abs path → func boundaries
}

func (m *escapeViewModel) HandleKey(key string) bool {
	switch key {
	case "a":
		m.showAll = !m.showAll
		return true
	case "p":
		m.projectOnly = !m.projectOnly
		return true
	}
	return false
}

func (m *escapeViewModel) Status() string {
	filter := "heap + leaking params"
	if m.showAll {
		filter = "all notes"
	}
	scope := "project"
	if !m.projectOnly {
		scope = "all (incl. deps)"
	}
	return fmt.Sprintf("[a] filter: %s    [p] scope: %s    [q] quit", filter, scope)
}

func (m *escapeViewModel) Render() string {
	visible := m.filteredSites()
	if len(visible) == 0 {
		if m.showAll && !m.projectOnly {
			return "(no escape analysis output)\n"
		}
		hints := []string{}
		if !m.showAll {
			hints = append(hints, "[a] to show all compiler notes")
		}
		if m.projectOnly {
			hints = append(hints, "[p] to include deps and stdlib")
		}
		msg := "(no results"
		if len(hints) > 0 {
			msg += " — press " + strings.Join(hints, " or ")
		}
		return msg + ")\n"
	}

	// Group sites by (file, func). Preserve order of first appearance.
	type funcKey struct{ file, fn string }
	type funcGroup struct {
		key   funcKey
		bound *engine.FuncBoundary // nil = not in a known function
		sites []engine.EscapeSite
	}
	var groups []funcGroup
	groupIdx := make(map[funcKey]int)

	for _, s := range visible {
		bounds := m.funcBoundaries(s.File)
		fn := engine.FindFunc(bounds, s.Line)
		var key funcKey
		if fn != nil {
			key = funcKey{s.File, fn.Name}
		} else {
			key = funcKey{s.File, ""}
		}
		i, ok := groupIdx[key]
		if !ok {
			i = len(groups)
			groups = append(groups, funcGroup{key: key, bound: fn})
			groupIdx[key] = i
		}
		groups[i].sites = append(groups[i].sites, s)
	}

	// Sort groups: by file then function start line (or name for unknown).
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].key.file != groups[j].key.file {
			return groups[i].key.file < groups[j].key.file
		}
		if groups[i].bound != nil && groups[j].bound != nil {
			return groups[i].bound.StartLine < groups[j].bound.StartLine
		}
		return groups[i].key.fn < groups[j].key.fn
	})

	// Summary.
	heapCount := 0
	for _, s := range visible {
		if s.IsHeapEscape() {
			heapCount++
		}
	}
	var sb strings.Builder
	if heapCount > 0 {
		fmt.Fprintf(&sb, "%d heap escape(s) across %d function(s)\n", heapCount, len(groups))
	} else {
		fmt.Fprintf(&sb, "%d note(s) across %d function(s)\n", len(visible), len(groups))
	}

	for _, g := range groups {
		// Section header: "── file · func Name ───────..."
		dispFile := m.relPath(g.key.file)
		sb.WriteByte('\n')
		var title string
		if g.key.fn != "" {
			title = fmt.Sprintf(" %s · func %s ", dispFile, g.key.fn)
		} else {
			title = fmt.Sprintf(" %s ", dispFile)
		}
		const lineWidth = 72
		left := 2
		right := lineWidth - left - len(title)
		if right < 2 {
			right = 2
		}
		fmt.Fprintf(&sb, "%s%s%s\n", strings.Repeat("─", left), title, strings.Repeat("─", right))
		sb.WriteByte('\n')

		if g.bound != nil {
			m.renderFunction(&sb, g.key.file, g.bound, g.sites)
		} else {
			m.renderContext(&sb, g.key.file, g.sites)
		}
	}

	return sb.String()
}

// renderFunction writes the full body of bound with escape annotations inline.
func (m *escapeViewModel) renderFunction(sb *strings.Builder, absFile string, bound *engine.FuncBoundary, sites []engine.EscapeSite) {
	lines := m.fileLines(absFile)
	if len(lines) == 0 {
		for _, s := range sites {
			fmt.Fprintf(sb, "  %4d  %s\n", s.Line, s.Message)
		}
		return
	}

	// Build per-line annotation map.
	annotations := make(map[int][]engine.EscapeSite) // 1-based line → sites
	for _, s := range sites {
		annotations[s.Line] = append(annotations[s.Line], s)
	}

	start := bound.StartLine
	end := bound.EndLine
	if end > len(lines) {
		end = len(lines)
	}
	if start > end {
		start = end
	}
	if start < 1 {
		start = 1
	}

	// Expand tabs before highlighting so visual columns are predictable.
	raw := make([]string, end-start+1)
	for i, l := range lines[start-1 : end] {
		raw[i] = expandTabs(l, 4)
	}
	highlighted := highlightGoLines(raw)

	annotIndentStr := strings.Repeat(" ", annotIndent)
	for i, hl := range highlighted {
		lineNo := start + i
		fmt.Fprintf(sb, "  %4d  %s\n", lineNo, hl)
		if sites := annotations[lineNo]; len(sites) > 0 {
			fmt.Fprintf(sb, "%s%s\n", annotIndentStr, annotSeparator(sites))
			for _, s := range sites {
				fmt.Fprintf(sb, "%s%s\n", annotIndentStr, annotationStyle(s))
			}
		}
	}
}

// renderContext shows ±contextLines source lines around each site (for var blocks,
// init expressions, and other non-function escape sites).
func (m *escapeViewModel) renderContext(sb *strings.Builder, absFile string, sites []engine.EscapeSite) {
	const ctx = 4
	lines := m.fileLines(absFile)
	annotIndentStr := strings.Repeat(" ", annotIndent)

	// Build per-line annotation map.
	annotations := make(map[int][]engine.EscapeSite)
	for _, s := range sites {
		annotations[s.Line] = append(annotations[s.Line], s)
	}

	// Collect the union of line ranges to display, merging overlapping windows.
	type span struct{ start, end int }
	var spans []span
	for _, s := range sites {
		lo := s.Line - ctx
		if lo < 1 {
			lo = 1
		}
		hi := s.Line + ctx
		if len(lines) > 0 && hi > len(lines) {
			hi = len(lines)
		}
		if lo > hi {
			lo = hi // site line is past EOF
		}
		if len(spans) > 0 && lo <= spans[len(spans)-1].end+1 {
			if hi > spans[len(spans)-1].end {
				spans[len(spans)-1].end = hi
			}
		} else {
			spans = append(spans, span{lo, hi})
		}
	}

	if len(lines) == 0 {
		for _, s := range sites {
			fmt.Fprintf(sb, "  %4d  %s\n", s.Line, s.Message)
		}
		return
	}

	for i, sp := range spans {
		if i > 0 {
			fmt.Fprintf(sb, "  %4s  …\n", "")
		}
		raw := make([]string, sp.end-sp.start+1)
		for j, l := range lines[sp.start-1 : sp.end] {
			raw[j] = expandTabs(l, 4)
		}
		highlighted := highlightGoLines(raw)
		for j, hl := range highlighted {
			lineNo := sp.start + j
			fmt.Fprintf(sb, "  %4d  %s\n", lineNo, hl)
			if sitesOnLine := annotations[lineNo]; len(sitesOnLine) > 0 {
				fmt.Fprintf(sb, "%s%s\n", annotIndentStr, annotSeparator(sitesOnLine))
				for _, s := range sitesOnLine {
					fmt.Fprintf(sb, "%s%s\n", annotIndentStr, annotationStyle(s))
				}
			}
		}
	}
}

func (m *escapeViewModel) filteredSites() []engine.EscapeSite {
	var out []engine.EscapeSite
	for _, s := range m.sites {
		if m.projectOnly && !isProjectFile(m.sourcesRoot, s.File) {
			continue
		}
		if m.showAll || s.IsHeapEscape() || s.IsLeakingParam() {
			out = append(out, s)
		}
	}
	return out
}

func isProjectFile(sourcesRoot, file string) bool {
	if !filepath.IsAbs(file) {
		// relative paths are emitted relative to sourcesRoot; paths starting with ".." are outside
		return !strings.HasPrefix(filepath.Clean(file), "..")
	}
	rel, err := filepath.Rel(sourcesRoot, file)
	return err == nil && !strings.HasPrefix(rel, "..")
}

func (m *escapeViewModel) fileLines(absFile string) []string {
	if lines, ok := m.fileCache[absFile]; ok {
		return lines
	}
	rel, err := filepath.Rel(m.sourcesRoot, absFile)
	if err != nil || strings.HasPrefix(rel, "..") {
		rel = absFile
	}
	data, err := os.ReadFile(filepath.Join(m.sourcesRoot, rel))
	if err != nil {
		// try absolute path directly
		data, err = os.ReadFile(absFile)
	}
	var lines []string
	if err == nil {
		for _, l := range strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n") {
			lines = append(lines, l)
		}
	}
	m.fileCache[absFile] = lines
	return lines
}

func (m *escapeViewModel) funcBoundaries(absFile string) []engine.FuncBoundary {
	if bounds, ok := m.funcCache[absFile]; ok {
		return bounds
	}
	rel, err := filepath.Rel(m.sourcesRoot, absFile)
	if err != nil || strings.HasPrefix(rel, "..") {
		rel = absFile
	}
	fullPath := filepath.Join(m.sourcesRoot, rel)
	bounds, _ := engine.ParseFuncBoundaries(fullPath)
	if bounds == nil {
		// also try absolute path
		bounds, _ = engine.ParseFuncBoundaries(absFile)
	}
	m.funcCache[absFile] = bounds
	return bounds
}

func (m *escapeViewModel) relPath(file string) string {
	if !filepath.IsAbs(file) {
		return filepath.Clean(file)
	}
	rel, err := filepath.Rel(m.sourcesRoot, file)
	if err != nil || strings.HasPrefix(rel, "..") {
		return file
	}
	return rel
}

var (
	heapEscapeStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("214")) // orange
	leakingParamStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))  // cyan
	otherNoteStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("245")) // dim gray
)

const annotSepWidth = 36

// "  NNNN  " prefix = 8 chars; annotations indent 2 more.
const lineNumWidth = 8
const annotIndent = lineNumWidth + 2

func annotSeparator(sites []engine.EscapeSite) string {
	base := otherNoteStyle
	for _, s := range sites {
		if s.IsHeapEscape() {
			base = heapEscapeStyle
			break
		}
		if s.IsLeakingParam() {
			base = leakingParamStyle
		}
	}
	return base.Render(strings.Repeat("─", annotSepWidth))
}

// expandTabs replaces tab characters with spaces aligned to tabWidth stops.
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

var subjectStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255")) // bright white

func annotationStyle(s engine.EscapeSite) string {
	prefix, subject, suffix := splitSubject(s.Message)
	var base lipgloss.Style
	switch {
	case s.IsHeapEscape():
		base = heapEscapeStyle
	case s.IsLeakingParam():
		base = leakingParamStyle
	default:
		base = otherNoteStyle
	}
	return "↑ " + base.Render(prefix) + subjectStyle.Render(subject) + base.Render(suffix)
}

// splitSubject splits a compiler escape message into (prefix, subject, suffix)
// so the subject (variable/expression name) can be bolded independently.
//
//	"data escapes to heap"     → ("", "data", " escapes to heap")
//	"moved to heap: result"    → ("moved to heap: ", "result", "")
//	"leaking param: buf"       → ("leaking param: ", "buf", "")
func splitSubject(msg string) (prefix, subject, suffix string) {
	// "X escapes to heap"
	if i := strings.Index(msg, " escapes to heap"); i > 0 {
		return "", msg[:i], msg[i:]
	}
	// "moved to heap: X" or "leaking param[...]: X"
	if strings.HasPrefix(msg, "moved to heap: ") || strings.HasPrefix(msg, "leaking param") {
		if i := strings.LastIndex(msg, ": "); i >= 0 {
			return msg[:i+2], msg[i+2:], ""
		}
	}
	// fallback: bold first word
	if i := strings.Index(msg, " "); i > 0 {
		return "", msg[:i], msg[i:]
	}
	return "", msg, ""
}

// highlightGoLines syntax-highlights a slice of Go source lines using chroma
// and returns one highlighted string per input line. Falls back to the original
// lines if highlighting fails (e.g. not a terminal).
func highlightGoLines(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}
	source := strings.Join(lines, "\n")
	lexer := cmp.Or(lexers.Get("go"), lexers.Fallback)
	style := cmp.Or(styles.Get("monokai"), styles.Fallback)

	var buf bytes.Buffer
	it, err := lexer.Tokenise(nil, source)
	if err != nil {
		return lines
	}
	if err := formatters.TTY16m.Format(&buf, style, it); err != nil {
		return lines
	}

	// chroma emits a trailing newline; split and trim the extra empty entry.
	out := strings.Split(buf.String(), "\n")
	if len(out) > len(lines) {
		out = out[:len(lines)]
	}
	// Ensure each line ends with a reset so open color codes don't bleed into
	// whatever is rendered after (e.g. line numbers on the next row).
	const reset = "\x1b[0m"
	for i, l := range out {
		if !strings.HasSuffix(l, reset) {
			out[i] = l + reset
		}
	}
	return out
}
