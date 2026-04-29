package commands

import (
	"context"
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/NimbleMarkets/ntcharts/v2/linechart"
	"github.com/charmbracelet/x/ansi"
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/commands/tabwriter"
	"benchspotter/engine"
)

type trendOptions struct {
	bench      string
	tags       []string
	last       int
	confidence float64
}

func newTrendCommand(env *execenv.Env) *cobra.Command {
	opts := trendOptions{confidence: 0.95}

	cmd := &cobra.Command{
		Use:     "trend",
		Short:   "Show benchmark trends over time",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTrend(cmd.Context(), env, opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.bench, "bench", "", "open detail view for a specific benchmark directly")
	flags.StringArrayVar(&opts.tags, "tag", nil, "filter sessions to those with this tag (repeatable)")
	flags.IntVar(&opts.last, "last", 0, "limit to last N sessions (0 = all)")
	flags.Float64Var(&opts.confidence, "confidence", 0.95, "confidence level for CI range")

	return cmd
}

func runTrend(ctx context.Context, env *execenv.Env, opts trendOptions) error {
	var sessions []*engine.SessionInfo
	var data *engine.TrendData

	err := env.Spinner().Title("Loading trend data").ActionWithErr(func(ctx context.Context) error {
		var err error
		sessions, err = engine.LocateSessions(env.Repo.Storage(), opts.tags...)
		if err != nil {
			return err
		}
		if opts.last > 0 && len(sessions) > opts.last {
			sessions = sessions[len(sessions)-opts.last:]
		}
		data, err = engine.LoadTrendData(sessions, opts.confidence)
		return err
	}).Context(ctx).Run()
	if err != nil {
		return err
	}

	if len(data.BenchNames) == 0 {
		return fmt.Errorf("no benchmark results found in selected sessions")
	}

	switch env.Format {
	case execenv.FormatJSON:
		return renderTrendJSON(env, data, opts.bench)
	case execenv.FormatText:
		if !env.Out.IsTerminal() {
			return renderTrendTextPlain(env, data, opts)
		}
		return env.RunModel(ctx, newTrendModel(data, opts, env.Style))
	default:
		return fmt.Errorf("unsupported format %v for trend (text, json)", env.Format)
	}
}

// ── JSON output ───────────────────────────────────────────────────────────────

type trendJSONOutput struct {
	Benchmarks []trendJSONBench `json:"benchmarks"`
}

type trendJSONBench struct {
	Name  string          `json:"name"`
	Units []trendJSONUnit `json:"units"`
}

type trendJSONUnit struct {
	Unit   string           `json:"unit"`
	Points []trendJSONPoint `json:"points"`
}

type trendJSONPoint struct {
	SessionID string    `json:"session_id"`
	HumanName string    `json:"human_name"`
	Time      time.Time `json:"time"`
	Center    float64   `json:"center"`
	Lo        float64   `json:"lo"`
	Hi        float64   `json:"hi"`
	N         int       `json:"n"`
}

func renderTrendJSON(env *execenv.Env, data *engine.TrendData, bench string) error {
	var benches []string
	if bench != "" {
		benches = []string{bench}
	} else {
		benches = data.BenchNames
	}

	out := trendJSONOutput{}
	for _, b := range benches {
		jb := trendJSONBench{Name: b}
		for _, unit := range data.Units {
			pts := data.Points[b][unit]
			if len(pts) == 0 {
				continue
			}
			ju := trendJSONUnit{Unit: unit}
			for _, p := range pts {
				ju.Points = append(ju.Points, trendJSONPoint{
					SessionID: p.Session.Id,
					HumanName: p.Session.HumanName,
					Time:      p.Session.Time,
					Center:    p.Center,
					Lo:        p.Lo,
					Hi:        p.Hi,
					N:         p.N,
				})
			}
			jb.Units = append(jb.Units, ju)
		}
		out.Benchmarks = append(out.Benchmarks, jb)
	}
	return env.Out.PrintJSON(out)
}

// ── Plain-text output (non-terminal) ─────────────────────────────────────────

func renderTrendTextPlain(env *execenv.Env, data *engine.TrendData, opts trendOptions) error {
	if opts.bench != "" {
		return renderTrendDetailText(env, data, opts.bench)
	}
	return renderTrendOverviewText(env, data)
}

func renderTrendOverviewText(env *execenv.Env, data *engine.TrendData) error {
	unit := data.Units[0]
	mc := engine.NewMachineContext(data.Sessions)
	tw := tabwriter.NewWriter(env.Out, 0, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "%-30s", "benchmark")
	for _, s := range data.Sessions {
		col := s.HumanName
		if mc != nil {
			if n := mc.Labels[s.Id]; n != 0 {
				col = fmt.Sprintf("⚙%d ", n) + col
			}
		}
		fmt.Fprintf(tw, "\t%s", col)
	}
	fmt.Fprintln(tw)

	for _, bench := range data.BenchNames {
		pts := data.Points[bench][unit]
		ptMap := make(map[string]engine.TrendPoint)
		for _, p := range pts {
			ptMap[p.Session.Id] = p
		}
		display := strings.TrimPrefix(bench, "Benchmark")
		fmt.Fprintf(tw, "%-30s", display)
		for i, s := range data.Sessions {
			p, ok := ptMap[s.Id]
			if !ok {
				fmt.Fprintf(tw, "\t%s", "-")
				continue
			}
			val := formatMetricValue(p.Center, unit)
			if i > 0 {
				if prev, hasPrev := ptMap[data.Sessions[i-1].Id]; hasPrev {
					val += trendArrow(p.Center, prev.Center)
				}
			}
			fmt.Fprintf(tw, "\t%s", val)
		}
		fmt.Fprintln(tw)
	}
	if err := tw.Flush(); err != nil {
		return err
	}
	if mc != nil {
		fmt.Fprintln(env.Out)
		for i, e := range mc.Entries {
			fmt.Fprintf(env.Out, "  ⚙%d  %s\n", i+1, engine.FormatMachineLine(&e.Machine, e.GoVersion))
		}
	}
	return nil
}

func renderTrendDetailText(env *execenv.Env, data *engine.TrendData, bench string) error {
	fmt.Fprintf(env.Out, "Benchmark: %s\n\n", bench)
	mc := engine.NewMachineContext(data.Sessions)
	tw := tabwriter.NewWriter(env.Out, 0, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "%-20s\t%-12s", "session", "date")
	for _, unit := range data.Units {
		if len(data.Points[bench][unit]) > 0 {
			fmt.Fprintf(tw, "\t%s", unit)
		}
	}
	fmt.Fprintln(tw)

	type ptKey struct{ unit, sessionID string }
	ptMap := make(map[ptKey]engine.TrendPoint)
	for _, unit := range data.Units {
		for _, p := range data.Points[bench][unit] {
			ptMap[ptKey{unit, p.Session.Id}] = p
		}
	}

	var prevSID string
	for _, s := range data.Sessions {
		hasAny := false
		for _, unit := range data.Units {
			if _, ok := ptMap[ptKey{unit, s.Id}]; ok {
				hasAny = true
				break
			}
		}
		if !hasAny {
			continue
		}

		sessionName := s.HumanName
		if mc != nil {
			if n := mc.Labels[s.Id]; n != 0 {
				sessionName = fmt.Sprintf("⚙%d ", n) + sessionName
			}
		}
		fmt.Fprintf(tw, "%-20s\t%-12s", ansi.Truncate(sessionName, 19, "…"), s.Time.Format("06-Jan-02"))
		for _, unit := range data.Units {
			pts := data.Points[bench][unit]
			if len(pts) == 0 {
				continue
			}
			p, ok := ptMap[ptKey{unit, s.Id}]
			if !ok {
				fmt.Fprintf(tw, "\t%s", "-")
				continue
			}
			val := formatMetricValue(p.Center, unit)
			if prevSID != "" {
				if prev, hasPrev := ptMap[ptKey{unit, prevSID}]; hasPrev {
					val += trendArrow(p.Center, prev.Center)
				}
			}
			fmt.Fprintf(tw, "\t%s", val)
		}
		fmt.Fprintln(tw)
		prevSID = s.Id
	}
	if err := tw.Flush(); err != nil {
		return err
	}
	if mc != nil {
		fmt.Fprintln(env.Out)
		for i, e := range mc.Entries {
			fmt.Fprintf(env.Out, "  ⚙%d  %s\n", i+1, engine.FormatMachineLine(&e.Machine, e.GoVersion))
		}
	}
	return nil
}

// ── TUI model ─────────────────────────────────────────────────────────────────

type trendMode int

const (
	trendModeOverview trendMode = iota
	trendModeDetail
)

// Overview layout constants: name column width and per-session column width.
const (
	overviewNameW = 24
	overviewColW  = 13
)

type trendViewModel struct {
	data  *engine.TrendData
	opts  trendOptions
	style execenv.Style

	width, height int
	ready         bool

	mode    trendMode
	unitIdx int

	// overview state
	cursor    int // selected benchmark index
	scrollOff int // first visible row index
	colOff    int // first visible session column index

	// detail state
	bench string
	chart linechart.Model

	// machine filter: nil when single machine; machineFilter=-1 means all
	machineCtx    *engine.MachineContext
	machineFilter int
}

func newTrendModel(data *engine.TrendData, opts trendOptions, style execenv.Style) trendViewModel {
	m := trendViewModel{
		data:          data,
		opts:          opts,
		style:         style,
		machineCtx:    engine.NewMachineContext(data.Sessions),
		machineFilter: -1,
	}
	if opts.bench != "" {
		m.bench = opts.bench
		m.mode = trendModeDetail
		m.cursor = m.benchIdx()
	}
	return m
}

func (m trendViewModel) filteredSessions() []*engine.SessionInfo {
	if m.machineCtx == nil || m.machineFilter < 0 {
		return m.data.Sessions
	}
	n := m.machineFilter + 1
	out := make([]*engine.SessionInfo, 0, len(m.data.Sessions))
	for _, s := range m.data.Sessions {
		if m.machineCtx.Labels[s.Id] == n {
			out = append(out, s)
		}
	}
	return out
}

func (m trendViewModel) filteredPoints(bench, unit string) []engine.TrendPoint {
	pts := m.data.Points[bench][unit]
	if m.machineCtx == nil || m.machineFilter < 0 {
		return pts
	}
	n := m.machineFilter + 1
	out := make([]engine.TrendPoint, 0, len(pts))
	for _, p := range pts {
		if m.machineCtx.Labels[p.Session.Id] == n {
			out = append(out, p)
		}
	}
	return out
}

func (m trendViewModel) Init() tea.Cmd { return nil }

func (m trendViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		if m.mode == trendModeDetail {
			m = m.rebuildChart()
		}
		m = m.clampScroll() // also clamps colOff

	case tea.KeyPressMsg:
		k := msg.String()
		if k == "ctrl+c" || k == "q" {
			return m, tea.Quit
		}

		switch m.mode {
		case trendModeOverview:
			switch k {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
					m = m.clampScroll()
				}
			case "down", "j":
				if m.cursor < len(m.data.BenchNames)-1 {
					m.cursor++
					m = m.clampScroll()
				}
			case "enter":
				if len(m.data.BenchNames) > 0 {
					m.bench = m.data.BenchNames[m.cursor]
					m.mode = trendModeDetail
					if m.ready {
						m = m.rebuildChart()
					}
				}
			case "left", "h":
				if m.colOff > 0 {
					m.colOff--
				}
			case "right", "l":
				sessions := m.filteredSessions()
				maxOff := len(sessions) - m.maxCols()
				if maxOff < 0 {
					maxOff = 0
				}
				if m.colOff < maxOff {
					m.colOff++
				}
			case "m":
				if len(m.data.Units) > 0 {
					m.unitIdx = (m.unitIdx + 1) % len(m.data.Units)
				}
			case "f":
				if m.machineCtx != nil {
					m.machineFilter++
					if m.machineFilter >= len(m.machineCtx.Entries) {
						m.machineFilter = -1
					}
					m.colOff = 0
					m = m.clampScroll()
				}
			case "esc":
				return m, tea.Quit
			}

		case trendModeDetail:
			switch k {
			case "esc", "b", "backspace":
				m.mode = trendModeOverview
			case "left", "h":
				if m.cursor > 0 {
					m.cursor--
					m.bench = m.data.BenchNames[m.cursor]
					if m.ready {
						m = m.rebuildChart()
					}
				}
			case "right", "l":
				if m.cursor < len(m.data.BenchNames)-1 {
					m.cursor++
					m.bench = m.data.BenchNames[m.cursor]
					if m.ready {
						m = m.rebuildChart()
					}
				}
			case "m":
				if len(m.data.Units) > 0 {
					m.unitIdx = (m.unitIdx + 1) % len(m.data.Units)
					if m.ready {
						m = m.rebuildChart()
					}
				}
			case "f":
				if m.machineCtx != nil {
					m.machineFilter++
					if m.machineFilter >= len(m.machineCtx.Entries) {
						m.machineFilter = -1
					}
					if m.ready {
						m = m.rebuildChart()
					}
				}
			}
		}
	}
	return m, nil
}

func (m trendViewModel) benchIdx() int {
	for i, b := range m.data.BenchNames {
		if b == m.bench {
			return i
		}
	}
	return 0
}

func (m trendViewModel) overviewBodyRows() int {
	// rows available for benchmark list: total height minus title, header, separator, status
	n := m.height - 4
	if n < 1 {
		n = 1
	}
	return n
}

func (m trendViewModel) maxCols() int {
	n := (m.width - overviewNameW) / overviewColW
	if n < 1 {
		n = 1
	}
	return n
}

func (m trendViewModel) clampScroll() trendViewModel {
	visible := m.overviewBodyRows()
	if m.cursor < m.scrollOff {
		m.scrollOff = m.cursor
	}
	if m.cursor >= m.scrollOff+visible {
		m.scrollOff = m.cursor - visible + 1
	}
	if m.scrollOff < 0 {
		m.scrollOff = 0
	}
	maxOff := len(m.filteredSessions()) - m.maxCols()
	if maxOff < 0 {
		maxOff = 0
	}
	if m.colOff > maxOff {
		m.colOff = maxOff
	}
	if m.colOff < 0 {
		m.colOff = 0
	}
	return m
}

func (m trendViewModel) chartHeight() int {
	h := m.height - m.detailTableRows() - 4
	if h < 5 {
		h = 5
	}
	if h > 24 {
		h = 24
	}
	return h
}

func (m trendViewModel) detailTableRows() int {
	pts := m.filteredPoints(m.bench, m.data.Units[m.unitIdx])
	n := len(pts) + 2
	if n > 12 {
		n = 12
	}
	return n
}

func (m trendViewModel) rebuildChart() trendViewModel {
	if len(m.data.Units) == 0 || m.bench == "" {
		return m
	}
	unit := m.data.Units[m.unitIdx]
	pts := m.filteredPoints(m.bench, unit)

	chartW := m.width - 2
	if chartW < 10 {
		chartW = 10
	}

	n := len(pts)
	if n == 0 {
		return m
	}

	minV, maxV := math.MaxFloat64, -math.MaxFloat64
	for _, p := range pts {
		if p.Center < minV {
			minV = p.Center
		}
		if p.Center > maxV {
			maxV = p.Center
		}
	}
	pad := (maxV - minV) * 0.15
	if pad == 0 {
		pad = math.Abs(minV) * 0.1
	}
	if pad == 0 {
		pad = 1
	}

	chart := linechart.New(chartW, m.chartHeight(), 1, float64(n), minV-pad, maxV+pad,
		linechart.WithXYSteps(1, 2),
		linechart.WithXLabelFormatter(func(_ int, _ float64) string { return "" }),
	)

	chart.DrawXYAxisAndLabel()
	for i := 1; i < n; i++ {
		chart.DrawBrailleLine(
			canvas.Float64Point{X: float64(i), Y: pts[i-1].Center},
			canvas.Float64Point{X: float64(i + 1), Y: pts[i].Center},
		)
	}

	// drawXLabel positions labels using increment=rangeSz/graphWidth, but ScaleFloat64Point
	// (which places the braille dots) uses graphWidth-1 as the divisor. The two formulas
	// diverge: the last data point lands at column graphWidth-1, but the built-in label for
	// that value would need i=graphWidth which the loop never reaches. We work around this by
	// disabling the built-in formatter and writing labels ourselves at the columns
	// ScaleFloat64Point actually returns for each session's x value.
	origin := chart.Origin()
	for i := 0; i < n; i++ {
		scaled := chart.ScaleFloat64Point(canvas.Float64Point{X: float64(i + 1), Y: 0})
		col := int(math.Round(scaled.X))
		chart.Canvas.SetStringWithStyle(
			canvas.Point{X: origin.X + 1 + col, Y: origin.Y + 1},
			strconv.Itoa(i+1), chart.LabelStyle,
		)
	}

	m.chart = chart
	return m
}

func (m trendViewModel) View() tea.View {
	var content string
	if !m.ready {
		content = "\n  Initializing..."
	} else if len(m.data.Units) == 0 {
		content = "  No benchmark data available."
	} else {
		switch m.mode {
		case trendModeOverview:
			content = m.viewOverview()
		case trendModeDetail:
			content = m.viewDetail()
		}
	}
	view := tea.NewView(content)
	view.AltScreen = true
	return view
}

func (m trendViewModel) viewOverview() string {
	unit := m.data.Units[m.unitIdx]

	// Determine which sessions fit as columns given terminal width.
	allSessions := m.filteredSessions()
	maxCols := m.maxCols()
	colEnd := m.colOff + maxCols
	if colEnd > len(allSessions) {
		colEnd = len(allSessions)
	}
	sessions := allSessions[m.colOff:colEnd]

	// Build per-bench point maps once.
	type benchRow struct {
		display string
		ptMap   map[string]engine.TrendPoint
	}
	rows := make([]benchRow, len(m.data.BenchNames))
	for i, bench := range m.data.BenchNames {
		ptMap := make(map[string]engine.TrendPoint)
		for _, p := range m.data.Points[bench][unit] {
			ptMap[p.Session.Id] = p
		}
		rows[i] = benchRow{
			display: ansi.Truncate(strings.TrimPrefix(bench, "Benchmark"), overviewNameW-1, "…"),
			ptMap:   ptMap,
		}
	}

	var sb strings.Builder

	// Title
	sb.WriteString(m.style.Bold("Trend") + m.style.TonedDown("  unit: ") + m.style.Accent(unit) + "\n")

	// Header row — machine tag styled separately so it isn't swallowed by TonedDown.
	header := m.style.TonedDown(fmt.Sprintf("%-*s", overviewNameW, "benchmark"))
	for _, s := range sessions {
		var machineTag string
		if m.machineCtx != nil {
			if n := m.machineCtx.Labels[s.Id]; n != 0 {
				machineTag = m.style.Info(fmt.Sprintf("⚙%d", n))
			}
		}
		tagW := lipgloss.Width(machineTag)
		nameW := overviewColW - 2 - tagW
		if tagW > 0 {
			nameW-- // space between tag and name
		}
		if nameW < 1 {
			nameW = 1
		}
		col := ansi.Truncate(s.HumanName, nameW, "…")
		if machineTag != "" {
			header += "  " + machineTag + " " + m.style.TonedDown(fmt.Sprintf("%-*s", nameW, col))
		} else {
			header += "  " + m.style.TonedDown(fmt.Sprintf("%-*s", overviewColW-2, col))
		}
	}
	sb.WriteString(header + "\n")

	sb.WriteString(strings.Repeat("─", m.width) + "\n")

	// Benchmark rows (windowed by scroll offset)
	visible := m.overviewBodyRows()
	end := m.scrollOff + visible
	if end > len(rows) {
		end = len(rows)
	}

	selBg := m.style.SelectionBg()

	withBg := func(text string, bg color.Color) string { return m.style.WithBg(bg, text) }

	for i := m.scrollOff; i < end; i++ {
		row := rows[i]
		selected := i == m.cursor

		defaultBg := func() color.Color {
			if selected {
				return selBg
			}
			return lipgloss.NoColor{}
		}

		line := withBg(fmt.Sprintf("%-*s", overviewNameW, row.display), defaultBg())
		for j, s := range sessions {
			p, ok := row.ptMap[s.Id]
			if !ok {
				line += withBg("  "+padRight("-", overviewColW-2), defaultBg())
				continue
			}
			val := formatMetricValue(p.Center, unit)
			cellBg := defaultBg()
			var arrow string
			if j > 0 {
				if prev, hasPrev := row.ptMap[sessions[j-1].Id]; hasPrev {
					if bg, ok := m.style.TrendCellBg(p.Center / prev.Center); ok {
						cellBg = bg
					}
					arrow = m.style.TrendArrow(p.Center, prev.Center, cellBg)
				}
			}
			arrowW := lipgloss.Width(arrow)
			line += withBg("  "+fmt.Sprintf("%-*s", overviewColW-2-arrowW, val), cellBg) + arrow
		}
		if vis := lipgloss.Width(line); vis < m.width {
			line += withBg(strings.Repeat(" ", m.width-vis), defaultBg())
		}
		sb.WriteString(line + "\n")
	}

	// Status bar (pinned to bottom via remaining blank lines)
	linesUsed := 3 + (end - m.scrollOff) // title + header + separator + rows
	for i := linesUsed; i < m.height-1; i++ {
		sb.WriteString("\n")
	}

	statusParts := []string{
		"[↑↓jk] select",
		"[↵] detail",
		"[←→hl] scroll cols",
		"[m] cycle unit",
	}
	if m.machineCtx != nil {
		machineLabel := "all"
		if m.machineFilter >= 0 {
			e := m.machineCtx.Entries[m.machineFilter]
			machineLabel = fmt.Sprintf("⚙%d", m.machineFilter+1) + " " + e.Machine.Summary()
		}
		statusParts = append(statusParts, "[f] machine: "+machineLabel)
	}
	statusParts = append(statusParts, "[q] quit")
	status := strings.Join(statusParts, "  ")
	if len(allSessions) > maxCols {
		status += fmt.Sprintf("  cols %d-%d/%d", m.colOff+1, colEnd, len(allSessions))
	}
	sb.WriteString(m.style.TonedDown(status))

	return sb.String()
}

func (m trendViewModel) viewDetail() string {
	if m.bench == "" {
		return "  No benchmark selected."
	}

	unit := m.data.Units[m.unitIdx]
	pts := m.filteredPoints(m.bench, unit)

	var sb strings.Builder

	// Title line
	benchDisplay := strings.TrimPrefix(m.bench, "Benchmark")
	nav := fmt.Sprintf("(%d/%d)", m.cursor+1, len(m.data.BenchNames))
	title := m.style.Bold(benchDisplay) + "  " + m.style.TonedDown(nav) +
		m.style.TonedDown("  unit: ") + m.style.Accent(unit)
	sb.WriteString(title + "\n")

	// Chart
	sb.WriteString(m.chart.View() + "\n")

	// Metrics table
	sb.WriteString(strings.Repeat("─", m.width) + "\n")

	colFmt := "%-3s  %-20s  %-12s  %-14s  %-8s  %-8s  %-4s\n"
	sb.WriteString(fmt.Sprintf(colFmt, "#", "session", "date", unit, "B/op", "allocs/op", "n"))

	tableMax := m.detailTableRows() - 2
	start := 0
	if len(pts) > tableMax {
		start = len(pts) - tableMax
	}

	bMap := pointMap(m.filteredPoints(m.bench, "B/op"))
	aMap := pointMap(m.filteredPoints(m.bench, "allocs/op"))

	var prevID string
	for i := start; i < len(pts); i++ {
		p := pts[i]
		sid := p.Session.Id

		nsVal := formatMetricValue(p.Center, unit)
		if prevID != "" {
			for _, prev := range pts {
				if prev.Session.Id == prevID {
					nsVal += " " + m.style.TrendArrow(p.Center, prev.Center)
					break
				}
			}
		}

		bVal := "-"
		if bp, ok := bMap[sid]; ok {
			bVal = formatMetricValue(bp.Center, "B/op")
		}
		aVal := "-"
		if ap, ok := aMap[sid]; ok {
			aVal = formatMetricValue(ap.Center, "allocs/op")
		}

		var machineTag string
		if m.machineCtx != nil {
			if n := m.machineCtx.Labels[p.Session.Id]; n != 0 {
				machineTag = m.style.Info(fmt.Sprintf("⚙%d ", n))
			}
		}
		tagW := lipgloss.Width(machineTag)
		nameCell := machineTag + ansi.Truncate(p.Session.HumanName, 19-tagW, "…")

		sb.WriteString(
			padRight(fmt.Sprintf("%d", i+1), 3) +
				"  " + padRight(nameCell, 20) +
				"  " + padRight(p.Session.Time.Format("06-Jan-02"), 12) +
				"  " + padRight(nsVal, 14) +
				"  " + padRight(bVal, 8) +
				"  " + padRight(aVal, 8) +
				"  " + fmt.Sprintf("%d", p.N) + "\n",
		)
		prevID = sid
	}

	// Status bar
	statusParts := []string{
		"[←→hl] cycle benchmark",
		"[m] cycle unit",
	}
	if m.machineCtx != nil {
		machineLabel := "all"
		if m.machineFilter >= 0 {
			e := m.machineCtx.Entries[m.machineFilter]
			machineLabel = fmt.Sprintf("⚙%d", m.machineFilter+1) + " " + e.Machine.Summary()
		}
		statusParts = append(statusParts, "[f] machine: "+machineLabel)
	}
	statusParts = append(statusParts, "[b/esc] back", "[q] quit")
	sb.WriteString(m.style.TonedDown(strings.Join(statusParts, "  ")))

	return sb.String()
}

func pointMap(pts []engine.TrendPoint) map[string]engine.TrendPoint {
	m := make(map[string]engine.TrendPoint, len(pts))
	for _, p := range pts {
		m[p.Session.Id] = p
	}
	return m
}

// ── Formatting helpers ────────────────────────────────────────────────────────

func formatMetricValue(v float64, unit string) string {
	switch unit {
	case "ns/op":
		switch {
		case v >= 1e9:
			return fmt.Sprintf("%.2fs", v/1e9)
		case v >= 1e6:
			return fmt.Sprintf("%.2fms", v/1e6)
		case v >= 1e3:
			return fmt.Sprintf("%.2fµs", v/1e3)
		default:
			return fmt.Sprintf("%.1fns", v)
		}
	case "B/op":
		switch {
		case v >= 1<<20:
			return fmt.Sprintf("%.1fMB", v/float64(1<<20))
		case v >= 1<<10:
			return fmt.Sprintf("%.1fKB", v/float64(1<<10))
		default:
			return fmt.Sprintf("%.0fB", v)
		}
	case "allocs/op":
		return fmt.Sprintf("%.0f", v)
	default:
		return fmt.Sprintf("%.4g", v)
	}
}

func trendArrow(curr, prev float64) string {
	if prev == 0 {
		return ""
	}
	ratio := curr / prev
	switch {
	case ratio > 1.10:
		return " ↑↑"
	case ratio > 1.02:
		return " ↑"
	case ratio < 0.90:
		return " ↓↓"
	case ratio < 0.98:
		return " ↓"
	}
	return ""
}

// padRight pads s to at least width visible characters, accounting for ANSI codes.
func padRight(s string, width int) string {
	vis := lipgloss.Width(s)
	if vis >= width {
		return s
	}
	return s + strings.Repeat(" ", width-vis)
}
