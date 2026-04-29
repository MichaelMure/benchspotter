package commands

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

type showInlineOptions struct {
	session     string
	all         bool
	includeDeps bool
}

func newShowInlineCommand(env *execenv.Env) *cobra.Command {
	options := showInlineOptions{}

	cmd := &cobra.Command{
		Use:     "inline",
		Short:   "Show compiler inlining decisions recorded for a session",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShowInline(cmd.Context(), env, options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.session, "session", "", "Session ID to inspect")
	flags.BoolVar(&options.all, "all", false, "Show all inlining decisions, not just 'cannot inline'")
	flags.BoolVar(&options.includeDeps, "deps", false, "Include stdlib and dependencies (default: project code only)")

	return cmd
}

func runShowInline(ctx context.Context, env *execenv.Env, options showInlineOptions) error {
	var selection *engine.SessionInfo

	if options.session == "" {
		const recallKey = "show_inline_session"
		preSelected := env.Repo.GetRecall(recallKey)
		var err error
		selection, err = inputs.SelectSession(ctx, env, preSelected,
			func(info *engine.SessionInfo) bool { return info.HasProfile(engine.ProfileInline) })
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

	if !selection.HasProfile(engine.ProfileInline) {
		return fmt.Errorf("session %q has no inline analysis recorded", selection.HumanName)
	}

	sites, err := engine.ReadInlineAnalysis(env.Repo.Storage(), selection.Path)
	if err != nil {
		return err
	}

	sourcesRoot := env.Repo.Sources().Root()

	switch env.Format {
	case execenv.FormatJSON:
		type jsonSite struct {
			File    string `json:"file"`
			Line    int    `json:"line"`
			Col     int    `json:"col"`
			Message string `json:"message"`
			Kind    string `json:"kind"`
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
		out := make([]jsonSite, 0, len(sites))
		for _, s := range sites {
			if !options.all && s.Kind() != engine.InlineCannotInline {
				continue
			}
			if !options.includeDeps && !isProjectFile(sourcesRoot, s.File) {
				continue
			}
			out = append(out, jsonSite{
				File:    s.File,
				Line:    s.Line,
				Col:     s.Col,
				Message: s.Message,
				Kind:    kindStr(s.Kind()),
			})
		}
		return env.Out.PrintJSON(out)

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
				style:       env.Style,
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
		}
		return env.ViewportWithKeys(ctx, model)()

	default:
		return fmt.Errorf("unsupported format %v for show inline (text, json)", env.Format)
	}
}

type inlineViewModel struct {
	sourceViewBase
	machineHeader string
	sites         []engine.InlineSite
	showAll       bool
	projectOnly   bool
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
	scope := "project"
	if !m.projectOnly {
		scope = "all (incl. deps)"
	}
	return fmt.Sprintf("[a] filter: %s    [p] scope: %s    [⇥] next  [⇤] prev    [ctrl+s] search    [q] quit", filter, scope)
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
			hints = append(hints, "[a] to show all decisions")
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
		if m.projectOnly && !isProjectFile(m.sourcesRoot, s.File) {
			continue
		}
		if m.showAll || s.Kind() == engine.InlineCannotInline {
			out = append(out, s)
		}
	}
	return out
}

// ── Inline-specific annotation rendering ─────────────────────────────────────

func (m *inlineViewModel) renderInlineAnnotations(sb *strings.Builder, sites []engine.InlineSite) {
	if len(sites) == 0 {
		return
	}
	indent := strings.Repeat(" ", annotIndent)
	fmt.Fprintf(sb, "%s%s\n", indent, inlineAnnotSeparator(m.style, sites))
	for _, s := range sites {
		fmt.Fprintf(sb, "%s%s\n", indent, inlineAnnotationStyle(m.style, s))
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
