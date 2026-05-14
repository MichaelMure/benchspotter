package commands

import (
	"fmt"
	"io"
	"strings"

	"charm.land/lipgloss/v2"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

// ── Options & session resolution ──────────────────────────────────────────────

type compareAnalysisOptions struct {
	baseSession string
	newSession  string
	all         bool
	includeDeps bool
}

func resolveAnalysisSessions(
	env *execenv.Env,
	options compareAnalysisOptions,
	profileType engine.Profile,
	recallPrefix string,
) (*engine.SessionInfo, *engine.SessionInfo, error) {
	baseRecallKey := recallPrefix + "_base"
	newRecallKey := recallPrefix + "_new"
	profileFilter := func(info *engine.SessionInfo) bool { return info.HasProfile(profileType) }

	var baseInfo, newInfo *engine.SessionInfo
	var err error

	if options.baseSession == "" {
		baseInfo, err = inputs.SelectSession(env, "Select base session", env.Repo.GetRecall(baseRecallKey), profileFilter)
		if err != nil {
			return nil, nil, err
		}
		if err = env.Repo.SetRecall(baseRecallKey, baseInfo.Id); err != nil {
			return nil, nil, err
		}
	} else {
		baseInfo, err = engine.LocateSession(env.Repo.Storage(), options.baseSession)
		if err != nil {
			return nil, nil, err
		}
	}
	if !baseInfo.HasProfile(profileType) {
		return nil, nil, fmt.Errorf("session %q has no analysis of this type recorded", baseInfo.HumanName)
	}

	if options.newSession == "" {
		newInfo, err = inputs.SelectSession(env, "Select new session", env.Repo.GetRecall(newRecallKey), profileFilter)
		if err != nil {
			return nil, nil, err
		}
		if err = env.Repo.SetRecall(newRecallKey, newInfo.Id); err != nil {
			return nil, nil, err
		}
	} else {
		newInfo, err = engine.LocateSession(env.Repo.Storage(), options.newSession)
		if err != nil {
			return nil, nil, err
		}
	}
	if !newInfo.HasProfile(profileType) {
		return nil, nil, fmt.Errorf("session %q has no analysis of this type recorded", newInfo.HumanName)
	}

	return baseInfo, newInfo, nil
}

func loadSessionSrcView(env *execenv.Env, sourcesRoot string, session *engine.SessionInfo) sourceViewBase {
	var gitDiff []byte
	if session.HasGitDiff() {
		if f, err := session.OpenFile(engine.GitDiffFilename); err == nil {
			gitDiff, _ = io.ReadAll(io.LimitReader(f, 10*1024*1024))
			_ = f.Close()
		}
	}
	return sourceViewBase{
		env:         env,
		sourcesRoot: sourcesRoot,
		gitCommit:   session.GitCommit,
		gitDiff:     gitDiff,
		git:         env.Repo,
		rawCache:    make(map[string][]byte),
		fileCache:   make(map[string][]string),
		funcCache:   make(map[string][]engine.FuncBoundary),
	}
}

func newCompareBase(
	env *execenv.Env,
	sourcesRoot string,
	baseInfo, newInfo *engine.SessionInfo,
	options compareAnalysisOptions,
) compareBase {
	return compareBase{
		sourceViewBase: loadSessionSrcView(env, sourcesRoot, newInfo),
		baseSrc:        loadSessionSrcView(env, sourcesRoot, baseInfo),
		showAll:        options.all,
		projectOnly:    !options.includeDeps,
		splitView:      true,
		baseLabel:      baseInfo.HumanName,
		newLabel:       newInfo.HumanName,
	}
}

// ── compareBase ───────────────────────────────────────────────────────────────

// compareBase holds state and source-reading infrastructure shared by both
// compare inline and compare escape view models.
type compareBase struct {
	sourceViewBase                // new session's source — also owns navigation state
	baseSrc        sourceViewBase // base session's source

	showAll     bool
	projectOnly bool
	splitView   bool
	baseLabel   string
	newLabel    string

	viewportWidth int // set via SetViewportWidth; sizes split panels
}

// SetViewportWidth implements execenv.ViewportWidthSetter.
func (b *compareBase) SetViewportWidth(width int) {
	b.viewportWidth = width
}

// SidebarWidth delegates to the underlying source view in both modes.
func (b *compareBase) SidebarWidth(total int) int {
	return b.sourceViewBase.SidebarWidth(total)
}

// RenderSidebar delegates to the underlying source view in both modes.
func (b *compareBase) RenderSidebar(width, height int) string {
	return b.sourceViewBase.RenderSidebar(width, height)
}

func (b *compareBase) handleBaseKey(key string) bool {
	if b.handleSearchKey(key) {
		return false
	}
	switch key {
	case "a":
		b.showAll = !b.showAll
		return true
	case "p":
		b.projectOnly = !b.projectOnly
		return true
	case "v":
		b.splitView = !b.splitView
		return true
	case "right", "left":
		return true // suppress horizontal viewport scrolling
	case "tab":
		b.jumpNext()
	case "shift+tab":
		b.jumpPrev()
	}
	return false
}

// baseStatusLine builds the footer status line. extras are inserted before the
// layout toggle, e.g. "[f] flow: off    ".
func (b *compareBase) baseStatusLine(extras ...string) string {
	if s := b.searchStatusLine(); s != "" {
		return s
	}
	filter := "changed"
	if b.showAll {
		filter = "all"
	}
	scope := "project"
	if !b.projectOnly {
		scope = "all (incl. deps)"
	}
	layout := "[v] split"
	if b.splitView {
		layout = "[v] stacked"
	}
	extra := strings.Join(extras, "")
	return fmt.Sprintf("[a] filter: %s    [p] scope: %s    %s%s    [⇥] next  [⇤] prev    [ctrl+f] search    [q] quit",
		filter, scope, extra, layout)
}

func (b *compareBase) writeAnalysisSummary(sb *strings.Builder, addedN, removedN, total int, unit string) {
	switch {
	case addedN > 0 || removedN > 0:
		fmt.Fprintf(sb, "%d new, %d fixed %s(s)\n", addedN, removedN, unit)
	case b.showAll && total > 0:
		fmt.Fprintf(sb, "%d %s(s) (no changes)\n", total, unit)
	default:
		sb.WriteString("(no changes)\n")
	}
	fmt.Fprintf(sb, "base: %s\nnew:  %s\n", b.baseLabel, b.newLabel)
}

// writeSplitSummary writes the summary header used in split mode (column labels
// instead of stacked base/new lines).
func (b *compareBase) writeSplitSummary(sb *strings.Builder, addedN, removedN, total int, unit string, halfW int, style execenv.Style) {
	switch {
	case addedN > 0 || removedN > 0:
		fmt.Fprintf(sb, "%d new, %d fixed %s(s)\n", addedN, removedN, unit)
	case b.showAll && total > 0:
		fmt.Fprintf(sb, "%d %s(s) (no changes)\n", total, unit)
	default:
		sb.WriteString("(no changes)\n")
	}
	sb.WriteByte('\n')
	sb.WriteString(combinePanels(
		style.TonedDown("base: "+b.baseLabel),
		style.TonedDown("new:  "+b.newLabel),
		halfW, style,
	))
	sb.WriteByte('\n')
}

// ── Generic group rendering helpers ──────────────────────────────────────────

// renderGroupAnnotated renders a function body with per-line annotation callbacks.
func renderGroupAnnotated[A any](
	src *sourceViewBase,
	sb *strings.Builder,
	absFile string,
	bound *engine.FuncBoundary,
	sites []A,
	lineOf func(A) int,
	messageOf func(A) string,
	renderAnnots func(*strings.Builder, []A),
) {
	annotations := make(map[int][]A)
	for _, s := range sites {
		annotations[lineOf(s)] = append(annotations[lineOf(s)], s)
	}
	src.renderFunctionBody(sb, absFile, bound,
		func(sb *strings.Builder) {
			for _, s := range sites {
				fmt.Fprintf(sb, "  %4d  %s\n", lineOf(s), messageOf(s))
			}
		},
		func(sb *strings.Builder, lineNo int) { renderAnnots(sb, annotations[lineNo]) },
	)
}

// renderGroupContext renders context lines around sites with per-line annotation callbacks.
func renderGroupContext[A any](
	src *sourceViewBase,
	sb *strings.Builder,
	absFile string,
	sites []A,
	lineOf func(A) int,
	messageOf func(A) string,
	renderAnnots func(*strings.Builder, []A),
) {
	annotations := make(map[int][]A)
	lineNums := make([]int, len(sites))
	for i, s := range sites {
		annotations[lineOf(s)] = append(annotations[lineOf(s)], s)
		lineNums[i] = lineOf(s)
	}
	src.renderContextLines(sb, absFile, lineNums,
		func(sb *strings.Builder) {
			for _, s := range sites {
				fmt.Fprintf(sb, "  %4d  %s\n", lineOf(s), messageOf(s))
			}
		},
		func(sb *strings.Builder, lineNo int) { renderAnnots(sb, annotations[lineNo]) },
	)
}

// renderStackedGroups writes the full stacked layout: summary, new groups, optional
// "fixed" separator and base groups, and updates b.headers.
func renderStackedGroups[A any](
	b *compareBase,
	sb *strings.Builder,
	newGroups, baseGroups []siteGroup[A],
	addedN, removedN, total int,
	unit string,
	renderGroupFn func(*strings.Builder, *sourceViewBase, *siteGroup[A]),
) {
	b.writeAnalysisSummary(sb, addedN, removedN, total, unit)
	for i := range newGroups {
		g := &newGroups[i]
		sb.WriteByte('\n')
		b.writeGroupHeader(sb, b.relPath(g.key.file), g.key.fn, len(g.sites))
		sb.WriteByte('\n')
		renderGroupFn(sb, &b.sourceViewBase, g)
	}
	if len(baseGroups) > 0 {
		fmt.Fprintf(sb, "\n%s\n",
			b.env.Style.TonedDown(strings.Repeat("─", 20)+" fixed in new session "+strings.Repeat("─", 38)))
		for i := range baseGroups {
			g := &baseGroups[i]
			sb.WriteByte('\n')
			b.writeGroupHeader(sb, b.relPath(g.key.file), g.key.fn, len(g.sites))
			sb.WriteByte('\n')
			renderGroupFn(sb, &b.baseSrc, g)
		}
	}
	b.headers = make([]headerEntry, 0, len(newGroups)+len(baseGroups))
	for i := range newGroups {
		g := &newGroups[i]
		b.headers = append(b.headers, headerEntry{file: b.relPath(g.key.file), fn: g.key.fn, count: len(g.sites)})
	}
	for i := range baseGroups {
		g := &baseGroups[i]
		b.headers = append(b.headers, headerEntry{file: b.relPath(g.key.file), fn: g.key.fn, count: len(g.sites)})
	}
}

// renderSplitGroups writes the full split layout: summary column labels, then each
// merged group side-by-side, and updates b.headers.
func renderSplitGroups[A any](
	b *compareBase,
	sb *strings.Builder,
	newGroups, baseGroups []siteGroup[A],
	addedN, removedN, total int,
	unit string,
	halfW int,
	renderGroupFn func(*strings.Builder, *sourceViewBase, *siteGroup[A]),
) {
	merged := mergeSiteGroups(baseGroups, newGroups, b.relPath)
	b.writeSplitSummary(sb, addedN, removedN, total, unit, halfW, b.env.Style)

	b.headers = b.headers[:0]
	for _, mg := range merged {
		baseN, newN := 0, 0
		if mg.base != nil {
			baseN = len(mg.base.sites)
		}
		if mg.new_ != nil {
			newN = len(mg.new_.sites)
		}
		b.headers = append(b.headers, headerEntry{file: mg.relFile, fn: mg.fn, count: max(baseN, newN)})
		b.renderSplitGroup(sb, mg.relFile, mg.fn, baseN, newN, halfW,
			func(leftSb *strings.Builder) {
				if mg.base != nil {
					renderGroupFn(leftSb, &b.baseSrc, mg.base)
				} else {
					b.renderSourceOnly(leftSb, &b.baseSrc, mg.new_.key.file, mg.fn)
				}
			},
			func(rightSb *strings.Builder) {
				if mg.new_ != nil {
					renderGroupFn(rightSb, &b.sourceViewBase, mg.new_)
				} else {
					b.renderSourceOnly(rightSb, &b.sourceViewBase, mg.base.key.file, mg.fn)
				}
			},
		)
	}
}

// ── Split rendering helpers ───────────────────────────────────────────────────

// findFuncByName returns the FuncBoundary whose Name matches, or nil.
func findFuncByName(bounds []engine.FuncBoundary, name string) *engine.FuncBoundary {
	for i := range bounds {
		if bounds[i].Name == name {
			return &bounds[i]
		}
	}
	return nil
}

// mergedSiteGroup pairs matching base/new groups for aligned side-by-side rendering.
type mergedSiteGroup[S any] struct {
	relFile string
	fn      string
	base    *siteGroup[S] // nil when site only exists in new
	new_    *siteGroup[S] // nil when site only exists in base
}

// mergeSiteGroups aligns base and new groups by (relFile, fn) key.
// Groups that appear in new come first (in new order); unmatched base-only
// groups are appended at the end.
func mergeSiteGroups[S any](
	baseGroups, newGroups []siteGroup[S],
	relPathOf func(string) string,
) []mergedSiteGroup[S] {
	type key struct{ relFile, fn string }
	baseByKey := make(map[key]int, len(baseGroups))
	for i := range baseGroups {
		g := &baseGroups[i]
		baseByKey[key{relPathOf(g.key.file), g.key.fn}] = i
	}
	usedBase := make(map[key]bool)
	var result []mergedSiteGroup[S]
	for i := range newGroups {
		g := &newGroups[i]
		k := key{relPathOf(g.key.file), g.key.fn}
		mg := mergedSiteGroup[S]{relFile: k.relFile, fn: k.fn, new_: g}
		if bi, ok := baseByKey[k]; ok {
			mg.base = &baseGroups[bi]
			usedBase[k] = true
		}
		result = append(result, mg)
	}
	for i := range baseGroups {
		g := &baseGroups[i]
		k := key{relPathOf(g.key.file), g.key.fn}
		if !usedBase[k] {
			result = append(result, mergedSiteGroup[S]{relFile: k.relFile, fn: k.fn, base: g})
		}
	}
	return result
}

// renderSplitGroup writes a merged group's header and side-by-side body into sb.
// renderLeft and renderRight write the annotated (or plain) body for each side.
func (b *compareBase) renderSplitGroup(
	sb *strings.Builder,
	relFile, fn string,
	baseN, newN int,
	halfW int,
	renderLeft, renderRight func(*strings.Builder),
) {
	var leftSb, rightSb strings.Builder
	b.writeGroupHeader(&leftSb, relFile, fn, baseN)
	b.writeGroupHeader(&rightSb, relFile, fn, newN)
	leftSb.WriteByte('\n')
	rightSb.WriteByte('\n')
	renderLeft(&leftSb)
	renderRight(&rightSb)
	left := strings.TrimRight(leftSb.String(), "\n")
	right := strings.TrimRight(rightSb.String(), "\n")
	sb.WriteByte('\n')
	sb.WriteString(combinePanels(left, right, halfW, b.env.Style))
	sb.WriteByte('\n')
}

// renderSourceOnly renders the function body from src without any annotations,
// used to fill the empty side of a split panel when one session has no sites
// for a function.
func (b *compareBase) renderSourceOnly(sb *strings.Builder, src *sourceViewBase, absFile, fn string) {
	if fn == "" {
		return
	}
	if bound := findFuncByName(src.funcBoundaries(absFile), fn); bound != nil {
		src.renderFunctionBody(sb, absFile, bound,
			func(*strings.Builder) {},
			func(*strings.Builder, int) {},
		)
	}
}

// combinePanels joins left and right content side by side with a │ divider.
// left is padded/truncated to leftWidth visual columns per line.
func combinePanels(left, right string, leftWidth int, style execenv.Style) string {
	leftLines := strings.Split(left, "\n")
	rightLines := strings.Split(right, "\n")
	n := len(leftLines)
	if len(rightLines) > n {
		n = len(rightLines)
	}
	divider := style.SidebarDivider("│")
	var sb strings.Builder
	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteByte('\n')
		}
		l := ""
		if i < len(leftLines) {
			l = leftLines[i]
		}
		r := ""
		if i < len(rightLines) {
			r = rightLines[i]
		}
		padded := lipgloss.NewStyle().Width(leftWidth).MaxWidth(leftWidth).Render(l)
		sb.WriteString(padded + divider + r)
	}
	return sb.String()
}
