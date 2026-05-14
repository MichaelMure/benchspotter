package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
)

// ── Command & run ─────────────────────────────────────────────────────────────

func newCompareEscapeCommand(env *execenv.Env) *cobra.Command {
	options := compareAnalysisOptions{}
	cmd := &cobra.Command{
		Use:   "escape",
		Short: "Diff compiler escape analysis between two sessions",
		Long: `Diff compiler escape analysis between two sessions.

Shows variables where heap escape status changed between sessions. New heap
escapes are regressions (amber); fixed escapes are improvements (green).

In the viewer: [a] toggles showing unchanged sites, [p] toggles project-only
scope, [f] toggles the data-flow chain, [v] switches between stacked and
side-by-side layout, [⇥]/[⇤] jump between functions, [ctrl+f] search.

Any interactive prompt can be bypassed with the corresponding flags.`,
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCompareEscape(env, options)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&options.baseSession, "base", "", "Base session ID or name")
	flags.StringVar(&options.newSession, "new", "", "New session ID or name")
	flags.BoolVar(&options.all, "all", false, "Show unchanged sites in addition to diffs")
	flags.BoolVar(&options.includeDeps, "deps", false, "Include stdlib and dependencies")
	return cmd
}

func runCompareEscape(env *execenv.Env, options compareAnalysisOptions) error {
	baseInfo, newInfo, err := resolveAnalysisSessions(env, options, engine.ProfileEscape, "compare_escape")
	if err != nil {
		return err
	}
	baseSites, err := engine.ReadEscapeAnalysis(env.Repo.Storage(), baseInfo.Path)
	if err != nil {
		return fmt.Errorf("reading base escape analysis: %w", err)
	}
	newSites, err := engine.ReadEscapeAnalysis(env.Repo.Storage(), newInfo.Path)
	if err != nil {
		return fmt.Errorf("reading new escape analysis: %w", err)
	}
	diffs := engine.DiffEscapeAnalysis(baseSites, newSites)

	sourcesRoot := env.Repo.Sources().Root()
	switch env.Format {
	case execenv.FormatText:
		model := &compareEscapeViewModel{
			compareBase: newCompareBase(env, sourcesRoot, baseInfo, newInfo, options),
			diffs:       diffs,
		}
		return env.ViewportWithKeys(model)()
	case execenv.FormatJSON:
		return outputEscapeDiffJSON(env, diffs)
	default:
		return fmt.Errorf("unsupported format %v for compare escape (text, json)", env.Format)
	}
}

// ── escapeDiffAnnot ───────────────────────────────────────────────────────────

type escapeDiffAnnot struct {
	engine.EscapeSite
	status engine.DiffStatus
}

// ── compareEscapeViewModel ────────────────────────────────────────────────────

type compareEscapeViewModel struct {
	compareBase
	diffs    []engine.DiffEscapeSite
	showFlow bool
}

func (m *compareEscapeViewModel) HandleKey(key string) bool {
	if key == "f" {
		m.showFlow = !m.showFlow
		return true
	}
	return m.handleBaseKey(key)
}

func (m *compareEscapeViewModel) Status() string {
	flow := "off"
	if m.showFlow {
		flow = "on"
	}
	return m.baseStatusLine(fmt.Sprintf("[f] flow: %s    ", flow))
}

func (m *compareEscapeViewModel) Render() string {
	visible := m.filteredEscapeDiffs()
	if m.splitView && m.viewportWidth > 0 {
		return m.renderEscapeSplit(visible)
	}
	return m.renderEscapeStacked(visible)
}

func (m *compareEscapeViewModel) filteredEscapeDiffs() []engine.DiffEscapeSite {
	var out []engine.DiffEscapeSite
	for _, d := range m.diffs {
		if !m.showAll && d.Status() == engine.DiffSame {
			continue
		}
		if m.projectOnly && !isProjectFile(m.sourcesRoot, d.Site().File) {
			continue
		}
		out = append(out, d)
	}
	return out
}

func (m *compareEscapeViewModel) renderEscapeStacked(diffs []engine.DiffEscapeSite) string {
	var newAnnots, baseAnnots []escapeDiffAnnot
	addedN, removedN := 0, 0
	for _, d := range diffs {
		switch d.Status() {
		case engine.DiffAdded:
			newAnnots = append(newAnnots, escapeDiffAnnot{*d.New, d.Status()})
			addedN++
		case engine.DiffSame:
			newAnnots = append(newAnnots, escapeDiffAnnot{*d.New, d.Status()})
		case engine.DiffRemoved:
			baseAnnots = append(baseAnnots, escapeDiffAnnot{*d.Base, d.Status()})
			removedN++
		}
	}
	newGroups := buildGroups(&m.sourceViewBase, newAnnots,
		func(a escapeDiffAnnot) string { return a.File },
		func(a escapeDiffAnnot) int { return a.Line },
	)
	baseGroups := buildGroups(&m.baseSrc, baseAnnots,
		func(a escapeDiffAnnot) string { return a.File },
		func(a escapeDiffAnnot) int { return a.Line },
	)
	var sb strings.Builder
	renderStackedGroups(&m.compareBase, &sb, newGroups, baseGroups, addedN, removedN, len(diffs), "escape site",
		m.renderEscapeGroup,
	)
	content := sb.String()
	m.updateHeaderLines(content)
	return content
}

func (m *compareEscapeViewModel) renderEscapeSplit(diffs []engine.DiffEscapeSite) string {
	halfW := (m.viewportWidth - 1) / 2
	var baseAnnots, newAnnots []escapeDiffAnnot
	addedN, removedN := 0, 0
	for _, d := range diffs {
		switch d.Status() {
		case engine.DiffAdded:
			newAnnots = append(newAnnots, escapeDiffAnnot{*d.New, d.Status()})
			addedN++
		case engine.DiffRemoved:
			baseAnnots = append(baseAnnots, escapeDiffAnnot{*d.Base, d.Status()})
			removedN++
		case engine.DiffSame:
			baseAnnots = append(baseAnnots, escapeDiffAnnot{*d.Base, d.Status()})
			newAnnots = append(newAnnots, escapeDiffAnnot{*d.New, d.Status()})
		}
	}
	baseGroups := buildGroups(&m.baseSrc, baseAnnots,
		func(a escapeDiffAnnot) string { return a.File },
		func(a escapeDiffAnnot) int { return a.Line },
	)
	newGroups := buildGroups(&m.sourceViewBase, newAnnots,
		func(a escapeDiffAnnot) string { return a.File },
		func(a escapeDiffAnnot) int { return a.Line },
	)
	var sb strings.Builder
	renderSplitGroups(&m.compareBase, &sb, newGroups, baseGroups, addedN, removedN, len(diffs), "escape site", halfW,
		m.renderEscapeGroup,
	)
	content := sb.String()
	m.updateHeaderLines(content)
	return content
}

func (m *compareEscapeViewModel) renderEscapeGroup(sb *strings.Builder, src *sourceViewBase, g *siteGroup[escapeDiffAnnot]) {
	if g.bound != nil {
		renderGroupAnnotated(src, sb, g.key.file, g.bound, g.sites,
			func(a escapeDiffAnnot) int { return a.Line },
			func(a escapeDiffAnnot) string { return a.Message },
			m.renderEscapeDiffAnnots,
		)
	} else {
		renderGroupContext(src, sb, g.key.file, g.sites,
			func(a escapeDiffAnnot) int { return a.Line },
			func(a escapeDiffAnnot) string { return a.Message },
			m.renderEscapeDiffAnnots,
		)
	}
}

func (m *compareEscapeViewModel) renderEscapeDiffAnnots(sb *strings.Builder, sites []escapeDiffAnnot) {
	if len(sites) == 0 {
		return
	}
	style := m.env.Style
	indent := strings.Repeat(" ", annotIndent)

	sepRender := func(s string) string { return style.TonedDown(s) }
	for _, s := range sites {
		switch s.status {
		case engine.DiffAdded:
			if s.IsHeapEscape() {
				sepRender = style.Warning
			}
		case engine.DiffRemoved:
			if s.IsHeapEscape() {
				sepRender = style.Positive
			}
		}
	}
	fmt.Fprintf(sb, "%s%s\n", indent, sepRender(strings.Repeat("─", annotSepWidth)))

	splitMode := m.splitView
	for _, s := range sites {
		prefix, subject, suffix := splitEscapeSubject(s.Message)
		var render func(string) string
		var marker string
		switch s.status {
		case engine.DiffAdded:
			if !splitMode {
				marker = "▲ "
			}
			if s.IsHeapEscape() {
				render = style.Warning
			} else if s.IsLeakingParam() {
				render = style.Info
			} else {
				render = func(s string) string { return style.TonedDown(s) }
			}
		case engine.DiffRemoved:
			if !splitMode {
				marker = "▼ "
			}
			if s.IsHeapEscape() || s.IsLeakingParam() {
				render = style.Positive
			} else {
				render = func(s string) string { return style.TonedDown(s) }
			}
		default: // Same
			if splitMode {
				render = func(s string) string { return style.TonedDown(s) }
			} else if s.IsHeapEscape() {
				render = style.Warning
			} else if s.IsLeakingParam() {
				render = style.Info
			} else {
				render = func(s string) string { return style.TonedDown(s) }
			}
		}
		fmt.Fprintf(sb, "%s%s↑ %s%s%s\n", indent, marker, render(prefix), style.Subject(subject), render(suffix))
		if m.showFlow {
			for _, f := range s.FlowChain {
				fmt.Fprintf(sb, "%s  %s\n", indent, style.FlowLine(f))
			}
		}
	}
}

// ── JSON output ───────────────────────────────────────────────────────────────

func outputEscapeDiffJSON(env *execenv.Env, diffs []engine.DiffEscapeSite) error {
	type jsonDiff struct {
		File         string   `json:"file"`
		Line         int      `json:"line"`
		Col          int      `json:"col"`
		Message      string   `json:"message"`
		HeapEscape   bool     `json:"heap_escape"`
		LeakingParam bool     `json:"leaking_param"`
		Status       string   `json:"status"`
		FlowChain    []string `json:"flow_chain,omitempty"`
		BaseLine     int      `json:"base_line,omitempty"`
		NewLine      int      `json:"new_line,omitempty"`
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
	out := make([]jsonDiff, len(diffs))
	for i, d := range diffs {
		s := d.Site()
		j := jsonDiff{
			File:         s.File,
			Line:         s.Line,
			Col:          s.Col,
			Message:      s.Message,
			HeapEscape:   s.IsHeapEscape(),
			LeakingParam: s.IsLeakingParam(),
			Status:       statusStr(d.Status()),
			FlowChain:    s.FlowChain,
		}
		if d.Base != nil {
			j.BaseLine = d.Base.Line
		}
		if d.New != nil {
			j.NewLine = d.New.Line
		}
		out[i] = j
	}
	return env.Out.PrintJSON(out)
}
