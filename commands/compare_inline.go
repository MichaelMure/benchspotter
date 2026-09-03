package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/MichaelMure/benchspotter/commands/execenv"
	"github.com/MichaelMure/benchspotter/engine"
)

// ── Command & run ─────────────────────────────────────────────────────────────

func newCompareInlineCommand(env *execenv.Env) *cobra.Command {
	options := compareAnalysisOptions{}
	cmd := &cobra.Command{
		Use:   "inline",
		Short: "Diff compiler inlining decisions between two sessions",
		Long: `Diff compiler inlining decisions between two sessions.

Shows functions where inlining changed between sessions. New "cannot inline"
sites are regressions (amber); fixed sites are improvements (green). New
"inlining call" sites are improvements (cyan); lost ones are regressions.

In the viewer: [a] toggles showing unchanged sites, [p] toggles project-only
scope, [v] switches between stacked and side-by-side layout, [⇥]/[⇤] jump
between functions, [ctrl+f] fuzzy-searches file/function names.

Any interactive prompt can be bypassed with the corresponding flags.`,
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCompareInline(env, options)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&options.baseSession, "base", "", "Base session ID or name")
	flags.StringVar(&options.newSession, "new", "", "New session ID or name")
	flags.BoolVar(&options.all, "all", false, "Show unchanged sites in addition to diffs")
	flags.BoolVar(&options.includeDeps, "deps", false, "Include stdlib and dependencies")
	flags.StringArrayVar(&options.funcs, "func", nil, "Filter to entries whose function name contains this string (repeatable, case-insensitive)")
	return cmd
}

func runCompareInline(env *execenv.Env, options compareAnalysisOptions) error {
	baseInfo, newInfo, err := resolveAnalysisSessions(env, options, engine.ProfileInline, "compare_inline")
	if err != nil {
		return err
	}
	baseSites, err := engine.ReadInlineAnalysis(env.Repo.Storage(), baseInfo.Path)
	if err != nil {
		return fmt.Errorf("reading base inline analysis: %w", err)
	}
	newSites, err := engine.ReadInlineAnalysis(env.Repo.Storage(), newInfo.Path)
	if err != nil {
		return fmt.Errorf("reading new inline analysis: %w", err)
	}
	diffs := engine.DiffInlineAnalysis(baseSites, newSites)

	sourcesRoot := env.Repo.Sources().Root()
	switch env.Format {
	case execenv.FormatText:
		model := &compareInlineViewModel{
			compareBase: newCompareBase(env, sourcesRoot, baseInfo, newInfo, options),
			diffs:       diffs,
			funcs:       options.funcs,
		}
		return env.ViewportWithKeys(model)()
	case execenv.FormatJSON:
		return outputInlineDiffJSON(env, diffs, options)
	case execenv.FormatRaw:
		return outputInlineDiffRaw(env, diffs)
	default:
		return fmt.Errorf("unsupported format %v for compare inline (raw, json, text)", env.Format)
	}
}

// ── inlineDiffAnnot ───────────────────────────────────────────────────────────

type inlineDiffAnnot struct {
	engine.InlineSite
	status engine.DiffStatus
}

// ── compareInlineViewModel ────────────────────────────────────────────────────

type compareInlineViewModel struct {
	compareBase
	diffs []engine.DiffInlineSite
	funcs []string
}

func (m *compareInlineViewModel) HandleKey(key string) bool {
	return m.handleBaseKey(key)
}

func (m *compareInlineViewModel) Status() string {
	return m.baseStatusLine()
}

func (m *compareInlineViewModel) Render() string {
	visible := m.filteredInlineDiffs()
	if m.splitView && m.viewportWidth > 0 {
		return m.renderInlineSplit(visible)
	}
	return m.renderInlineStacked(visible)
}

func (m *compareInlineViewModel) filteredInlineDiffs() []engine.DiffInlineSite {
	var out []engine.DiffInlineSite
	for _, d := range m.diffs {
		if !m.showAll && d.Status() == engine.DiffSame {
			continue
		}
		// cannot_inline diffs are always included regardless of scope.
		if d.Site().Kind() != engine.InlineCannotInline && m.projectOnly && !isProjectFile(m.sourcesRoot, d.Site().File) {
			continue
		}
		fn, _ := inlineMessageParts(d.Site().Message)
		if !matchesFuncFilter(fn, m.funcs) {
			continue
		}
		out = append(out, d)
	}
	return out
}

func (m *compareInlineViewModel) renderInlineStacked(diffs []engine.DiffInlineSite) string {
	var newAnnots, baseAnnots []inlineDiffAnnot
	addedN, removedN := 0, 0
	for _, d := range diffs {
		switch d.Status() {
		case engine.DiffAdded:
			newAnnots = append(newAnnots, inlineDiffAnnot{*d.New, d.Status()})
			addedN++
		case engine.DiffSame:
			newAnnots = append(newAnnots, inlineDiffAnnot{*d.New, d.Status()})
		case engine.DiffRemoved:
			baseAnnots = append(baseAnnots, inlineDiffAnnot{*d.Base, d.Status()})
			removedN++
		}
	}
	newGroups := buildGroups(&m.sourceViewBase, newAnnots,
		func(a inlineDiffAnnot) string { return a.File },
		func(a inlineDiffAnnot) int { return a.Line },
	)
	baseGroups := buildGroups(&m.baseSrc, baseAnnots,
		func(a inlineDiffAnnot) string { return a.File },
		func(a inlineDiffAnnot) int { return a.Line },
	)
	var sb strings.Builder
	renderStackedGroups(&m.compareBase, &sb, newGroups, baseGroups, addedN, removedN, len(diffs), "inline site",
		m.renderInlineGroup,
	)
	content := sb.String()
	m.updateHeaderLines(content)
	return content
}

func (m *compareInlineViewModel) renderInlineSplit(diffs []engine.DiffInlineSite) string {
	halfW := (m.viewportWidth - 1) / 2
	var baseAnnots, newAnnots []inlineDiffAnnot
	addedN, removedN := 0, 0
	for _, d := range diffs {
		switch d.Status() {
		case engine.DiffAdded:
			newAnnots = append(newAnnots, inlineDiffAnnot{*d.New, d.Status()})
			addedN++
		case engine.DiffRemoved:
			baseAnnots = append(baseAnnots, inlineDiffAnnot{*d.Base, d.Status()})
			removedN++
		case engine.DiffSame:
			baseAnnots = append(baseAnnots, inlineDiffAnnot{*d.Base, d.Status()})
			newAnnots = append(newAnnots, inlineDiffAnnot{*d.New, d.Status()})
		}
	}
	baseGroups := buildGroups(&m.baseSrc, baseAnnots,
		func(a inlineDiffAnnot) string { return a.File },
		func(a inlineDiffAnnot) int { return a.Line },
	)
	newGroups := buildGroups(&m.sourceViewBase, newAnnots,
		func(a inlineDiffAnnot) string { return a.File },
		func(a inlineDiffAnnot) int { return a.Line },
	)
	var sb strings.Builder
	renderSplitGroups(&m.compareBase, &sb, newGroups, baseGroups, addedN, removedN, len(diffs), "inline site", halfW,
		m.renderInlineGroup,
	)
	content := sb.String()
	m.updateHeaderLines(content)
	return content
}

func (m *compareInlineViewModel) renderInlineGroup(sb *strings.Builder, src *sourceViewBase, g *siteGroup[inlineDiffAnnot]) {
	if g.bound != nil {
		renderGroupAnnotated(src, sb, g.key.file, g.bound, g.sites,
			func(a inlineDiffAnnot) int { return a.Line },
			func(a inlineDiffAnnot) string { return a.Message },
			m.renderInlineDiffAnnots,
		)
	} else {
		renderGroupContext(src, sb, g.key.file, g.sites,
			func(a inlineDiffAnnot) int { return a.Line },
			func(a inlineDiffAnnot) string { return a.Message },
			m.renderInlineDiffAnnots,
		)
	}
}

func (m *compareInlineViewModel) renderInlineDiffAnnots(sb *strings.Builder, sites []inlineDiffAnnot) {
	if len(sites) == 0 {
		return
	}
	style := m.env.Style
	indent := strings.Repeat(" ", annotIndent)

	sepRender := func(s string) string { return style.TonedDown(s) }
	for _, s := range sites {
		switch s.status {
		case engine.DiffAdded:
			if s.Kind() == engine.InlineCannotInline {
				sepRender = style.Warning
			}
		case engine.DiffRemoved:
			sepRender = style.Positive
		}
	}
	fmt.Fprintf(sb, "%s%s\n", indent, sepRender(strings.Repeat("─", annotSepWidth)))

	splitMode := m.splitView
	for _, s := range sites {
		prefix, subject, suffix := splitInlineSubject(s.Message)
		var render func(string) string
		var marker string
		switch s.status {
		case engine.DiffAdded:
			if !splitMode {
				marker = "▲ "
			}
			switch s.Kind() {
			case engine.InlineCannotInline:
				render = style.Warning
			case engine.InlineInliningCall:
				render = style.Info
			default:
				render = func(s string) string { return style.TonedDown(s) }
			}
		case engine.DiffRemoved:
			if !splitMode {
				marker = "▼ "
			}
			switch s.Kind() {
			case engine.InlineCannotInline:
				render = style.Positive
			case engine.InlineInliningCall:
				render = style.Negative
			default:
				render = func(s string) string { return style.TonedDown(s) }
			}
		default: // Same
			if splitMode {
				render = func(s string) string { return style.TonedDown(s) }
			} else {
				switch s.Kind() {
				case engine.InlineCannotInline:
					render = style.Warning
				case engine.InlineInliningCall:
					render = style.Info
				default:
					render = func(s string) string { return style.TonedDown(s) }
				}
			}
		}
		fmt.Fprintf(sb, "%s%s↑ %s%s%s\n", indent, marker, render(prefix), style.Subject(subject), render(suffix))
	}
}

// ── JSON output ───────────────────────────────────────────────────────────────

func outputInlineDiffRaw(env *execenv.Env, diffs []engine.DiffInlineSite) error {
	for _, d := range diffs {
		if d.Status() == engine.DiffSame {
			continue
		}
		s := d.Site()
		prefix := "+"
		if d.Status() == engine.DiffRemoved {
			prefix = "-"
		}
		fmt.Fprintf(env.Out, "%s %s:%d:%d: %s\n", prefix, s.File, s.Line, s.Col, s.Message)
	}
	return nil
}

func outputInlineDiffJSON(env *execenv.Env, diffs []engine.DiffInlineSite, options compareAnalysisOptions) error {
	showSame := options.all
	type jsonSummary struct {
		Added   int `json:"added"`
		Removed int `json:"removed"`
		Same    int `json:"same"`
	}
	type jsonSite struct {
		Line     int    `json:"line"`
		Kind     string `json:"kind"`
		Function string `json:"function"`
		Reason   string `json:"reason,omitempty"`
		Status   string `json:"status"`
		BaseLine int    `json:"base_line,omitempty"`
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
	statusStr := func(s engine.DiffStatus) string {
		switch s {
		case engine.DiffAdded:
			return "added"
		case engine.DiffRemoved:
			return "removed"
		default:
			return "same"
		}
	}
	var fileOrder []string
	fileSites := make(map[string][]jsonSite)
	var summary jsonSummary
	for _, d := range diffs {
		if !showSame && d.Status() == engine.DiffSame {
			continue
		}
		s := d.Site()
		fn, reason := inlineMessageParts(s.Message)
		if !matchesFuncFilter(fn, options.funcs) {
			continue
		}
		switch d.Status() {
		case engine.DiffAdded:
			summary.Added++
		case engine.DiffRemoved:
			summary.Removed++
		default:
			summary.Same++
		}
		// base_line only when a same-status site has shifted lines.
		var baseLine int
		if d.Status() == engine.DiffSame && d.Base != nil && d.Base.Line != s.Line {
			baseLine = d.Base.Line
		}
		entry := jsonSite{
			Line:     s.Line,
			Kind:     kindStr(s.Kind()),
			Function: fn,
			Reason:   reason,
			Status:   statusStr(d.Status()),
			BaseLine: baseLine,
		}
		if _, seen := fileSites[s.File]; !seen {
			fileOrder = append(fileOrder, s.File)
		}
		fileSites[s.File] = append(fileSites[s.File], entry)
	}
	files := make([]jsonFileGroup, 0, len(fileOrder))
	for _, f := range fileOrder {
		files = append(files, jsonFileGroup{File: f, Sites: fileSites[f]})
	}
	return env.Out.PrintJSON(jsonOutput{Summary: summary, Files: files})
}
