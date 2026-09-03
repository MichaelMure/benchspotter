package commands

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/MichaelMure/benchspotter/commands/execenv"
	"github.com/MichaelMure/benchspotter/commands/inputs"
	"github.com/MichaelMure/benchspotter/engine"
)

type showEscapeOptions struct {
	session     string
	all         bool
	includeDeps bool
	funcs       []string
}

func newShowEscapeCommand(env *execenv.Env) *cobra.Command {
	options := showEscapeOptions{}

	cmd := &cobra.Command{
		Use:   "escape",
		Short: "Show compiler escape analysis recorded for a session",
		Long: `Show the compiler's escape analysis output, annotated against your source code.

The Go compiler decides whether a variable lives on the stack or must be moved
to the heap. Variables that escape to the heap cause garbage collector pressure
and extra allocations. Escape analysis output (from -gcflags='-m') tells you
which variables escaped and why — the "why" is often a function call, interface
conversion, or closure capture that prevented stack allocation.

By default only heap escapes and leaking parameters are shown. In the viewer:
[a] toggles showing all compiler notes, [p] toggles including stdlib and
dependencies (default: project code only), [f] toggles the dataflow chain that
explains how each value escaped, [Tab]/[Shift+Tab] jump between sites.

Any interactive prompt can be bypassed with the corresponding flags.`,
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShowEscape(env, options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.session, "session", "", "Session ID to inspect")
	flags.BoolVar(&options.all, "all", false, "Show all compiler notes, not just heap escapes and leaking params")
	flags.BoolVar(&options.includeDeps, "deps", false, "Include stdlib and dependencies (default: project code only)")
	flags.StringArrayVar(&options.funcs, "func", nil, "Filter to entries whose message contains this string (repeatable, case-insensitive)")

	return cmd
}

func runShowEscape(env *execenv.Env, options showEscapeOptions) error {
	var selection *engine.SessionInfo

	if options.session == "" {
		const recallKey = "show_escape_session"
		preSelected := env.Repo.GetRecall(recallKey)
		var err error
		selection, err = inputs.SelectSession(env, "--session", "Select session", preSelected,
			func(info *engine.SessionInfo) bool { return info.HasProfile(engine.ProfileEscape) })
		if err != nil {
			return err
		}
		if err = env.Repo.SetRecall(recallKey, selection.Id); err != nil {
			return err
		}
	} else {
		var err error
		selection, err = engine.LocateSession(env.Repo.Storage(), options.session)
		if err != nil {
			return err
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
	case execenv.FormatRaw:
		f, err := selection.OpenFile(engine.EscapeFilename)
		if err != nil {
			return fmt.Errorf("opening escape data: %w", err)
		}
		defer f.Close()
		_, err = io.Copy(env.Out, f)
		return err

	case execenv.FormatJSON:
		type jsonSummary struct {
			HeapEscape   int `json:"heap_escape"`
			LeakingParam int `json:"leaking_param"`
			Other        int `json:"other"`
		}
		type jsonSite struct {
			Line    int    `json:"line"`
			Kind    string `json:"kind"`
			Message string `json:"message"`
		}
		type jsonFileGroup struct {
			File  string     `json:"file"`
			Sites []jsonSite `json:"sites"`
		}
		type jsonOutput struct {
			Summary jsonSummary     `json:"summary"`
			Files   []jsonFileGroup `json:"files"`
		}
		var fileOrder []string
		fileSites := make(map[string][]jsonSite)
		var summary jsonSummary
		for _, s := range sites {
			isKey := s.IsHeapEscape() || s.IsLeakingParam()
			if !options.all && !isKey {
				continue
			}
			// heap_escape and leaking_param are always shown regardless of the
			// deps flag — they represent real allocation pressure on any caller.
			// The scope filter only suppresses other (informational) notes from deps.
			if !isKey && !options.includeDeps && !isProjectFile(sourcesRoot, s.File) {
				continue
			}
			if !matchesFuncFilter(s.Message, options.funcs) {
				continue
			}
			kind := escapeKind(s)
			switch kind {
			case "heap_escape":
				summary.HeapEscape++
			case "leaking_param":
				summary.LeakingParam++
			default:
				summary.Other++
			}
			if _, seen := fileSites[s.File]; !seen {
				fileOrder = append(fileOrder, s.File)
			}
			fileSites[s.File] = append(fileSites[s.File], jsonSite{
				Line:    s.Line,
				Kind:    kind,
				Message: s.Message,
			})
		}
		files := make([]jsonFileGroup, 0, len(fileOrder))
		for _, f := range fileOrder {
			files = append(files, jsonFileGroup{File: f, Sites: fileSites[f]})
		}
		return env.Out.PrintJSON(jsonOutput{Summary: summary, Files: files})

	case execenv.FormatText:
		var gitDiff []byte
		if selection.HasGitDiff() {
			if f, err := selection.OpenFile(engine.GitDiffFilename); err == nil {
				gitDiff, _ = io.ReadAll(io.LimitReader(f, 10*1024*1024))
				_ = f.Close()
			}
		}
		model := &escapeViewModel{
			sourceViewBase: sourceViewBase{
				env:         env,
				sourcesRoot: sourcesRoot,
				gitCommit:   selection.GitCommit,
				gitDiff:     gitDiff,
				git:         env.Repo,
				rawCache:    make(map[string][]byte),
				fileCache:   make(map[string][]string),
				funcCache:   make(map[string][]engine.FuncBoundary),
			},
			machineHeader: engine.FormatMachineLine(selection.Machine, selection.GoVersion),
			sites:         sites,
			showAll:       options.all,
			projectOnly:   !options.includeDeps,
			funcs:         options.funcs,
		}
		return env.ViewportWithKeys(model)()

	default:
		return fmt.Errorf("unsupported format %v for show escape (raw, json, text)", env.Format)
	}
}

type escapeViewModel struct {
	sourceViewBase
	machineHeader string
	sites         []engine.EscapeSite
	showAll       bool
	projectOnly   bool
	showFlow      bool
	funcs         []string
}

func (m *escapeViewModel) HandleKey(key string) bool {
	if m.handleSearchKey(key) {
		return false
	}
	switch key {
	case "a":
		m.showAll = !m.showAll
		return true
	case "p":
		m.projectOnly = !m.projectOnly
		return true
	case "f":
		m.showFlow = !m.showFlow
		return true
	case "tab":
		m.jumpNext()
	case "shift+tab":
		m.jumpPrev()
	}
	return false
}

func (m *escapeViewModel) Status() string {
	if s := m.searchStatusLine(); s != "" {
		return s
	}
	filter := "heap + leaking params"
	if m.showAll {
		filter = "all notes"
	}
	flow := "off"
	if m.showFlow {
		flow = "on"
	}
	// The scope toggle [p] only affects "other" compiler notes; heap_escape and
	// leaking_param are always shown from all scopes, so don't surface [p] in
	// the default mode where it would have no effect.
	if !m.showAll {
		return fmt.Sprintf("[a] filter: %s    [f] flow: %s    [⇥] next  [⇤] prev    [ctrl+f] search    [q] quit", filter, flow)
	}
	scope := "project"
	if !m.projectOnly {
		scope = "all (incl. deps)"
	}
	return fmt.Sprintf("[a] filter: %s    [p] scope: %s    [f] flow: %s    [⇥] next  [⇤] prev    [ctrl+f] search    [q] quit", filter, scope, flow)
}

func (m *escapeViewModel) Render() string {
	visible := m.filteredSites()
	if len(visible) == 0 {
		m.headers = nil
		if m.showAll && !m.projectOnly {
			return "(no escape analysis output)\n"
		}
		var hints []string
		if !m.showAll {
			// heap/leaking shown from all scopes, so no [p] hint here.
			hints = append(hints, "[a] to show all compiler notes")
		} else if m.projectOnly {
			hints = append(hints, "[p] to include deps and stdlib")
		}
		msg := "(no results"
		if len(hints) > 0 {
			msg += " — press " + strings.Join(hints, " or ")
		}
		return msg + ")\n"
	}

	groups := buildGroups(&m.sourceViewBase, visible,
		func(s engine.EscapeSite) string { return s.File },
		func(s engine.EscapeSite) int { return s.Line },
	)

	m.headers = make([]headerEntry, len(groups))
	for i, g := range groups {
		m.headers[i] = headerEntry{
			file:  m.relPath(g.key.file),
			fn:    g.key.fn,
			count: len(g.sites),
		}
	}

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
		sb.WriteByte('\n')
		m.writeGroupHeader(&sb, m.relPath(g.key.file), g.key.fn, len(g.sites))
		sb.WriteByte('\n')
		if g.bound != nil {
			m.renderEscapeFunction(&sb, g.key.file, g.bound, g.sites)
		} else {
			m.renderEscapeContext(&sb, g.key.file, g.sites)
		}
	}

	content := sb.String()
	if m.machineHeader != "" {
		content = m.machineHeader + "\n\n" + content
	}
	m.updateHeaderLines(content)
	return content
}

func (m *escapeViewModel) renderEscapeFunction(sb *strings.Builder, absFile string, bound *engine.FuncBoundary, sites []engine.EscapeSite) {
	annotations := make(map[int][]engine.EscapeSite)
	for _, s := range sites {
		annotations[s.Line] = append(annotations[s.Line], s)
	}
	m.renderFunctionBody(sb, absFile, bound,
		func(sb *strings.Builder) {
			for _, s := range sites {
				fmt.Fprintf(sb, "  %4d  %s\n", s.Line, s.Message)
			}
		},
		func(sb *strings.Builder, lineNo int) {
			m.renderEscapeAnnotations(sb, annotations[lineNo])
		},
	)
}

func (m *escapeViewModel) renderEscapeContext(sb *strings.Builder, absFile string, sites []engine.EscapeSite) {
	annotations := make(map[int][]engine.EscapeSite)
	lineNums := make([]int, len(sites))
	for i, s := range sites {
		annotations[s.Line] = append(annotations[s.Line], s)
		lineNums[i] = s.Line
	}
	m.renderContextLines(sb, absFile, lineNums,
		func(sb *strings.Builder) {
			for _, s := range sites {
				fmt.Fprintf(sb, "  %4d  %s\n", s.Line, s.Message)
			}
		},
		func(sb *strings.Builder, lineNo int) {
			m.renderEscapeAnnotations(sb, annotations[lineNo])
		},
	)
}

func (m *escapeViewModel) renderEscapeAnnotations(sb *strings.Builder, sites []engine.EscapeSite) {
	if len(sites) == 0 {
		return
	}
	indent := strings.Repeat(" ", annotIndent)
	fmt.Fprintf(sb, "%s%s\n", indent, escapeAnnotSeparator(m.env.Style, sites))
	for _, s := range sites {
		fmt.Fprintf(sb, "%s%s\n", indent, escapeAnnotationStyle(m.env.Style, s))
		if m.showFlow {
			for _, f := range s.FlowChain {
				fmt.Fprintf(sb, "%s  %s\n", indent, m.env.Style.FlowLine(f))
			}
		}
	}
}

func (m *escapeViewModel) filteredSites() []engine.EscapeSite {
	var out []engine.EscapeSite
	for _, s := range m.sites {
		isKey := s.IsHeapEscape() || s.IsLeakingParam()
		if !m.showAll && !isKey {
			continue
		}
		// heap_escape and leaking_param are always included regardless of scope.
		if !isKey && m.projectOnly && !isProjectFile(m.sourcesRoot, s.File) {
			continue
		}
		if !matchesFuncFilter(s.Message, m.funcs) {
			continue
		}
		out = append(out, s)
	}
	return out
}

// ── Escape-specific annotation rendering ─────────────────────────────────────

// "  NNNN  " prefix = 8 chars; annotations indent 2 more.
const lineNumWidth = 8
const annotIndent = lineNumWidth + 2
const annotSepWidth = 36

func escapeAnnotSeparator(style execenv.Style, sites []engine.EscapeSite) string {
	render := func(s string) string { return style.TonedDown(s) }
	for _, s := range sites {
		if s.IsHeapEscape() {
			render = style.Warning
			break
		}
		if s.IsLeakingParam() {
			render = style.Info
		}
	}
	return render(strings.Repeat("─", annotSepWidth))
}

func escapeAnnotationStyle(style execenv.Style, s engine.EscapeSite) string {
	prefix, subject, suffix := splitEscapeSubject(s.Message)
	render := func(t string) string { return style.TonedDown(t) }
	switch {
	case s.IsHeapEscape():
		render = style.Warning
	case s.IsLeakingParam():
		render = style.Info
	}
	return "↑ " + render(prefix) + style.Subject(subject) + render(suffix)
}

func splitEscapeSubject(msg string) (prefix, subject, suffix string) {
	if i := strings.Index(msg, " escapes to heap"); i > 0 {
		return "", msg[:i], msg[i:]
	}
	if strings.HasPrefix(msg, "moved to heap: ") || strings.HasPrefix(msg, "leaking param") {
		if i := strings.LastIndex(msg, ": "); i >= 0 {
			return msg[:i+2], msg[i+2:], ""
		}
	}
	if i := strings.Index(msg, " "); i > 0 {
		return "", msg[:i], msg[i:]
	}
	return "", msg, ""
}

// escapeKind classifies an EscapeSite into one of three JSON kind strings.
func escapeKind(s engine.EscapeSite) string {
	if s.IsHeapEscape() {
		return "heap_escape"
	}
	if s.IsLeakingParam() {
		return "leaking_param"
	}
	return "other"
}

// matchesFuncFilter reports whether name matches any pattern (case-insensitive
// substring). An empty list matches everything.
func matchesFuncFilter(name string, funcs []string) bool {
	if len(funcs) == 0 {
		return true
	}
	lower := strings.ToLower(name)
	for _, f := range funcs {
		if strings.Contains(lower, strings.ToLower(f)) {
			return true
		}
	}
	return false
}

// isProjectFile reports whether file is within the project (not a dep or stdlib).
func isProjectFile(sourcesRoot, file string) bool {
	if !filepath.IsAbs(file) {
		return !strings.HasPrefix(filepath.Clean(file), "..")
	}
	rel, err := filepath.Rel(sourcesRoot, file)
	return err == nil && !strings.HasPrefix(rel, "..")
}
