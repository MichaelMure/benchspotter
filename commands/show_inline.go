package commands

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/MichaelMure/benchspotter/commands/execenv"
	"github.com/MichaelMure/benchspotter/commands/inputs"
	"github.com/MichaelMure/benchspotter/engine"
)

type showInlineOptions struct {
	session     string
	all         bool
	includeDeps bool
	funcs       []string
}

func newShowInlineCommand(env *execenv.Env) *cobra.Command {
	options := showInlineOptions{}

	cmd := &cobra.Command{
		Use:   "inline",
		Short: "Show compiler inlining decisions recorded for a session",
		Long: `Show the compiler's inlining decisions, annotated against your source code.

When the compiler inlines a function call it substitutes the callee's body
directly at the call site. This eliminates function call overhead and often
enables further optimisations (such as better escape analysis). When a function
cannot be inlined, the compiler emits a "cannot inline" message with a reason.

By default only "cannot inline" sites are shown, since those are the ones worth
investigating. In the viewer: [a] toggles showing all decisions (including
successful inlines and "can inline" marks), [p] toggles including stdlib and
dependencies, [Tab]/[Shift+Tab] jump between sites.

Any interactive prompt can be bypassed with the corresponding flags.`,
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShowInline(env, options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.session, "session", "", "Session ID to inspect")
	flags.BoolVar(&options.all, "all", false, "Show all inlining decisions, not just 'cannot inline'")
	flags.BoolVar(&options.includeDeps, "deps", false, "Include stdlib and dependencies (default: project code only)")
	flags.StringArrayVar(&options.funcs, "func", nil, "Filter to entries whose function name contains this string (repeatable, case-insensitive)")

	return cmd
}

func runShowInline(env *execenv.Env, options showInlineOptions) error {
	var selection *engine.SessionInfo

	if options.session == "" {
		const recallKey = "show_inline_session"
		preSelected := env.Repo.GetRecall(recallKey)
		var err error
		selection, err = inputs.SelectSession(env, "--session", "Select session", preSelected,
			func(info *engine.SessionInfo) bool { return info.HasProfile(engine.ProfileInline) })
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

	if !selection.HasProfile(engine.ProfileInline) {
		return fmt.Errorf("session %q has no inline analysis recorded", selection.HumanName)
	}

	sites, err := engine.ReadInlineAnalysis(env.Repo.Storage(), selection.Path)
	if err != nil {
		return err
	}

	sourcesRoot := env.Repo.Sources().Root()

	switch env.Format {
	case execenv.FormatRaw:
		f, err := selection.OpenFile(engine.InlineFilename)
		if err != nil {
			return fmt.Errorf("opening inline data: %w", err)
		}
		defer f.Close()
		_, err = io.Copy(env.Out, f)
		return err

	case execenv.FormatJSON:
		type jsonSummary struct {
			CannotInline int `json:"cannot_inline"`
			InliningCall int `json:"inlining_call"`
			CanInline    int `json:"can_inline"`
		}
		type jsonSite struct {
			Line     int    `json:"line"`
			Kind     string `json:"kind"`
			Function string `json:"function"`
			Reason   string `json:"reason,omitempty"`
		}
		type jsonFileGroup struct {
			File  string     `json:"file"`
			Sites []jsonSite `json:"sites"`
		}
		type jsonOutput struct {
			Summary jsonSummary     `json:"summary"`
			Files   []jsonFileGroup `json:"files"`
		}
		kindStr := func(k engine.InlineKind) string {
			switch k {
			case engine.InlineCannotInline:
				return "cannot_inline"
			case engine.InlineInliningCall:
				return "inlining_call"
			default:
				return "can_inline"
			}
		}
		var fileOrder []string
		fileSites := make(map[string][]jsonSite)
		var summary jsonSummary
		for _, s := range sites {
			kind := s.Kind()
			if !options.all && kind != engine.InlineCannotInline {
				continue
			}
			// cannot_inline is always shown regardless of the deps flag: a dep
			// function that can't be inlined is real overhead for every project
			// call site. The scope filter only suppresses noisy can_inline /
			// inlining_call messages from dependencies.
			if kind != engine.InlineCannotInline && !options.includeDeps && !isProjectFile(sourcesRoot, s.File) {
				continue
			}
			fn, reason := inlineMessageParts(s.Message)
			if !matchesFuncFilter(fn, options.funcs) {
				continue
			}
			switch kind {
			case engine.InlineCannotInline:
				summary.CannotInline++
			case engine.InlineInliningCall:
				summary.InliningCall++
			default:
				summary.CanInline++
			}
			if _, seen := fileSites[s.File]; !seen {
				fileOrder = append(fileOrder, s.File)
			}
			fileSites[s.File] = append(fileSites[s.File], jsonSite{
				Line:     s.Line,
				Kind:     kindStr(kind),
				Function: fn,
				Reason:   reason,
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
		model := &inlineViewModel{
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
		return fmt.Errorf("unsupported format %v for show inline (raw, json, text)", env.Format)
	}
}

type inlineViewModel struct {
	sourceViewBase
	machineHeader string
	sites         []engine.InlineSite
	showAll       bool
	projectOnly   bool
	funcs         []string
}

func (m *inlineViewModel) HandleKey(key string) bool {
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
	case "tab":
		m.jumpNext()
	case "shift+tab":
		m.jumpPrev()
	}
	return false
}

func (m *inlineViewModel) Status() string {
	if s := m.searchStatusLine(); s != "" {
		return s
	}
	filter := "cannot inline"
	if m.showAll {
		filter = "all decisions"
	}
	// The scope toggle [p] only affects can_inline / inlining_call entries;
	// cannot_inline is always shown from all scopes, so don't surface [p] in
	// the default (cannot-inline-only) mode where it would have no effect.
	if !m.showAll {
		return fmt.Sprintf("[a] filter: %s    [⇥] next  [⇤] prev    [ctrl+f] search    [q] quit", filter)
	}
	scope := "project"
	if !m.projectOnly {
		scope = "all (incl. deps)"
	}
	return fmt.Sprintf("[a] filter: %s    [p] scope: %s    [⇥] next  [⇤] prev    [ctrl+f] search    [q] quit", filter, scope)
}

func (m *inlineViewModel) Render() string {
	visible := m.filteredSites()
	if len(visible) == 0 {
		m.headers = nil
		if m.showAll && !m.projectOnly {
			return "(no inline analysis output)\n"
		}
		var hints []string
		if !m.showAll {
			// cannot_inline is shown from all scopes, so no [p] hint here.
			hints = append(hints, "[a] to show all decisions")
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
		func(s engine.InlineSite) string { return s.File },
		func(s engine.InlineSite) int { return s.Line },
	)

	m.headers = make([]headerEntry, len(groups))
	for i, g := range groups {
		m.headers[i] = headerEntry{
			file:  m.relPath(g.key.file),
			fn:    g.key.fn,
			count: len(g.sites),
		}
	}

	cannotCount := 0
	for _, s := range visible {
		if s.Kind() == engine.InlineCannotInline {
			cannotCount++
		}
	}
	var sb strings.Builder
	if cannotCount > 0 {
		fmt.Fprintf(&sb, "%d cannot-inline site(s) across %d function(s)\n", cannotCount, len(groups))
	} else {
		fmt.Fprintf(&sb, "%d decision(s) across %d function(s)\n", len(visible), len(groups))
	}

	for _, g := range groups {
		sb.WriteByte('\n')
		m.writeGroupHeader(&sb, m.relPath(g.key.file), g.key.fn, len(g.sites))
		sb.WriteByte('\n')
		if g.bound != nil {
			m.renderInlineFunction(&sb, g.key.file, g.bound, g.sites)
		} else {
			m.renderInlineContext(&sb, g.key.file, g.sites)
		}
	}

	content := sb.String()
	if m.machineHeader != "" {
		content = m.machineHeader + "\n\n" + content
	}
	m.updateHeaderLines(content)
	return content
}

func (m *inlineViewModel) renderInlineFunction(sb *strings.Builder, absFile string, bound *engine.FuncBoundary, sites []engine.InlineSite) {
	annotations := make(map[int][]engine.InlineSite)
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
			m.renderInlineAnnotations(sb, annotations[lineNo])
		},
	)
}

func (m *inlineViewModel) renderInlineContext(sb *strings.Builder, absFile string, sites []engine.InlineSite) {
	annotations := make(map[int][]engine.InlineSite)
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
			m.renderInlineAnnotations(sb, annotations[lineNo])
		},
	)
}

func (m *inlineViewModel) filteredSites() []engine.InlineSite {
	var out []engine.InlineSite
	for _, s := range m.sites {
		kind := s.Kind()
		if !m.showAll && kind != engine.InlineCannotInline {
			continue
		}
		// cannot_inline is always included regardless of scope.
		if kind != engine.InlineCannotInline && m.projectOnly && !isProjectFile(m.sourcesRoot, s.File) {
			continue
		}
		fn, _ := inlineMessageParts(s.Message)
		if !matchesFuncFilter(fn, m.funcs) {
			continue
		}
		out = append(out, s)
	}
	return out
}

// ── Inline-specific annotation rendering ─────────────────────────────────────

func (m *inlineViewModel) renderInlineAnnotations(sb *strings.Builder, sites []engine.InlineSite) {
	if len(sites) == 0 {
		return
	}
	indent := strings.Repeat(" ", annotIndent)
	fmt.Fprintf(sb, "%s%s\n", indent, inlineAnnotSeparator(m.env.Style, sites))
	for _, s := range sites {
		fmt.Fprintf(sb, "%s%s\n", indent, inlineAnnotationStyle(m.env.Style, s))
	}
}

func inlineAnnotSeparator(style execenv.Style, sites []engine.InlineSite) string {
	render := func(s string) string { return style.TonedDown(s) }
	for _, s := range sites {
		switch s.Kind() {
		case engine.InlineCannotInline:
			render = style.Warning
		case engine.InlineInliningCall:
			render = style.Info
		}
	}
	return render(strings.Repeat("─", annotSepWidth))
}

func inlineAnnotationStyle(style execenv.Style, s engine.InlineSite) string {
	prefix, subject, suffix := splitInlineSubject(s.Message)
	render := func(t string) string { return style.TonedDown(t) }
	switch s.Kind() {
	case engine.InlineCannotInline:
		render = style.Warning
	case engine.InlineInliningCall:
		render = style.Info
	}
	return "↑ " + render(prefix) + style.Subject(subject) + render(suffix)
}

// inlineMessageParts extracts the function name and optional reason from a
// compiler inline message.
//
//	"cannot inline Foo: too complex"   → ("Foo", "too complex")
//	"inlining call to pkg.Foo"         → ("pkg.Foo", "")
//	"can inline Foo with cost 64 as:"  → ("Foo", "with cost 64 as:")
func inlineMessageParts(msg string) (fn, reason string) {
	_, fn, suffix := splitInlineSubject(msg)
	if r, ok := strings.CutPrefix(suffix, ": "); ok {
		reason = r
	} else {
		reason = strings.TrimPrefix(suffix, " ")
	}
	// For can_inline the compiler appends "as: <inlined body>" which is very
	// verbose. Strip everything from " as:" onward — the cost is sufficient.
	if before, _, found := strings.Cut(reason, " as:"); found {
		reason = before
	}
	return fn, reason
}

// splitInlineSubject splits an inline message so the function name is bolded.
//
//	"cannot inline Foo: too complex"  → ("cannot inline ", "Foo", ": too complex")
//	"inlining call to pkg.Foo"        → ("inlining call to ", "pkg.Foo", "")
//	"can inline Foo with cost 64 as:" → ("can inline ", "Foo", " with cost 64 as:")
func splitInlineSubject(msg string) (prefix, subject, suffix string) {
	for _, pfx := range []string{"cannot inline ", "inlining call to ", "can inline "} {
		after, ok := strings.CutPrefix(msg, pfx)
		if !ok {
			continue
		}
		// Subject ends at first space or colon.
		end := strings.IndexAny(after, " :")
		if end < 0 {
			return pfx, after, ""
		}
		return pfx, after[:end], after[end:]
	}
	if i := strings.Index(msg, " "); i > 0 {
		return "", msg[:i], msg[i:]
	}
	return "", msg, ""
}
