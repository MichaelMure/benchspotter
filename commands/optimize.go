package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/NimbleMarkets/ntcharts/v2/linechart"
	"github.com/charmbracelet/x/ansi"
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

type optimizeOptions struct {
	bench     string
	params    []string
	count     int
	maxTrials int
	metric    string
	strategy  string
	maximize  bool
}

func newOptimizeCommand(env *execenv.Env) *cobra.Command {
	opts := optimizeOptions{
		count: 1,
	}

	cmd := &cobra.Command{
		Use:     "optimize",
		Short:   "Search for optimal benchinput parameter values",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runOptimize(cmd.Context(), env, opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.bench, "bench", "b", "", "benchmark to optimize (prompt if not set)")
	flags.StringArrayVarP(&opts.params, "param", "p", nil, "input parameter name(s) to sweep (repeatable)")
	flags.IntVarP(&opts.count, "count", "c", 1, "benchmark iterations per trial")
	flags.IntVar(&opts.maxTrials, "max-trials", 0, "stop after this many trials (0 = unlimited)")
	flags.StringVarP(&opts.metric, "metric", "m", "", "metric to optimize (prompt if not set)")
	strategyKeys := make([]string, len(engine.StrategyDefs))
	for i, d := range engine.StrategyDefs {
		strategyKeys[i] = d.Key
	}
	flags.StringVarP(&opts.strategy, "strategy", "s", "", "strategy: "+strings.Join(strategyKeys, ", ")+" (prompt if not set)")
	flags.BoolVar(&opts.maximize, "maximize", false, "maximize the metric instead of minimizing")

	return cmd
}

func runOptimize(ctx context.Context, env *execenv.Env, opts optimizeOptions) error {
	// Discover benchinput parameters.
	var allInputs []engine.InputInfo
	err := env.Spinner().Title("Discovering inputs").ActionWithErr(func(ctx context.Context) error {
		var err error
		allInputs, err = engine.Inputs(env.Repo.Sources().Root())
		return err
	}).Context(ctx).Run()
	if err != nil {
		return err
	}
	if len(allInputs) == 0 {
		return fmt.Errorf("no benchinput parameters found – add benchinput.Bool/Int/Float calls to test files")
	}

	// Select parameters.
	var selectedInputs []engine.InputInfo
	if len(opts.params) > 0 {
		for _, input := range allInputs {
			for _, p := range opts.params {
				if input.Name() == p {
					selectedInputs = append(selectedInputs, input)
					break
				}
			}
		}
		if len(selectedInputs) == 0 {
			return fmt.Errorf("no inputs found matching: %v", opts.params)
		}
	} else {
		const recallKey = "optimize_params"
		preSelected := env.Repo.GetRecalls(recallKey)

		err = env.FormSingle(huh.NewMultiSelect[engine.InputInfo]().
			Title("Select parameters to optimize").
			OptionsFunc(func() []huh.Option[engine.InputInfo] {
				options := make([]huh.Option[engine.InputInfo], len(allInputs))
				for i, input := range allInputs {
					label := input.Name() + " " + input.Label() +
						env.Style.TonedDown(" — "+input.Package())
					options[i] = huh.NewOption(label, input).
						Selected(slices.Contains(preSelected, input.Name()))
				}
				return options
			}, nil).
			Value(&selectedInputs)).
			RunWithContext(ctx)
		if err != nil {
			return err
		}

		_ = env.Repo.SetRecalls(recallKey, func(yield func(string) bool) {
			for _, inp := range selectedInputs {
				if !yield(inp.Name()) {
					return
				}
			}
		})
	}

	// Select benchmark.
	var bench engine.BenchInfo
	if opts.bench != "" {
		var allBenches []engine.BenchInfo
		err = env.Spinner().Title("Finding benchmarks").ActionWithErr(func(ctx context.Context) error {
			var err error
			allBenches, err = engine.LocateBenchmarks(ctx, env.Repo.Sources())
			return err
		}).Context(ctx).Run()
		if err != nil {
			return err
		}
		for _, b := range allBenches {
			if b.Name == opts.bench {
				bench = b
				break
			}
		}
		if bench.Name == "" {
			return fmt.Errorf("benchmark %q not found", opts.bench)
		}
	} else {
		const recallKey = "optimize_bench"
		bench, err = inputs.SelectBenchmark(ctx, env, env.Repo.GetRecall(recallKey))
		if err != nil {
			return err
		}
		_ = env.Repo.SetRecall(recallKey, bench.Name)
	}

	// Select strategy.
	if opts.strategy == "" {
		const recallKey = "optimize_strategy"
		strat := env.Repo.GetRecall(recallKey)
		err = env.FormSingle(huh.NewSelect[string]().
			Title("Optimization strategy").
			OptionsFunc(func() []huh.Option[string] {
				options := make([]huh.Option[string], len(engine.StrategyDefs))
				for i, def := range engine.StrategyDefs {
					label := strings.ToUpper(def.Key[:1]) + def.Key[1:] + " – " + def.Description
					options[i] = huh.NewOption(label, def.Key)
				}
				return options
			}, nil).
			Value(&strat)).
			RunWithContext(ctx)
		if err != nil {
			return err
		}
		opts.strategy = strat
		_ = env.Repo.SetRecall(recallKey, strat)
	}

	// Select metric.
	if opts.metric == "" {
		const recallKey = "optimize_metric"
		selected := engine.MetricNsPerOp
		if recalled := env.Repo.GetRecall(recallKey); recalled != "" {
			for _, m := range engine.KnownBenchMetrics {
				if m.String() == recalled {
					selected = m
					break
				}
			}
		}
		err = env.FormSingle(huh.NewSelect[engine.BenchMetric]().
			Title("Metric to optimize").
			OptionsFunc(func() []huh.Option[engine.BenchMetric] {
				opts := make([]huh.Option[engine.BenchMetric], len(engine.KnownBenchMetrics))
				for i, m := range engine.KnownBenchMetrics {
					opts[i] = huh.NewOption(m.String()+" – "+m.Label(), m)
				}
				return opts
			}, nil).
			Value(&selected)).
			RunWithContext(ctx)
		if err != nil {
			return err
		}
		opts.metric = selected.String()
		_ = env.Repo.SetRecall(recallKey, opts.metric)
	}

	strategy := buildOptimizeStrategy(opts, selectedInputs)
	minimize := !opts.maximize

	optCtx, cancelOpt := context.WithCancel(ctx)
	defer cancelOpt()

	resultCh := engine.RunOptimize(optCtx, bench, strategy, opts.count, opts.maxTrials)

	switch env.Format {
	case execenv.FormatText:
		if !env.Out.IsTerminal() {
			return runOptimizePlain(env, resultCh, opts.metric, minimize, opts.maxTrials)
		}
		var modelResult optimizeModelResult
		err = env.RunModel(ctx, newOptimizeModel(bench, selectedInputs, opts.metric, minimize, strategy, opts.maxTrials, resultCh, cancelOpt, env.Style, &modelResult))
		if err != nil {
			return err
		}
		if modelResult.print {
			printOptimizeReport(env, modelResult.results, opts.metric, minimize)
		}
		return nil
	case execenv.FormatJSON:
		return runOptimizeJSON(env, bench, strategy, resultCh, opts.metric, minimize)
	default:
		return fmt.Errorf("unsupported format for optimize (text only)")
	}
}

func buildOptimizeStrategy(opts optimizeOptions, inputs []engine.InputInfo) engine.Strategy {
	for _, def := range engine.StrategyDefs {
		if def.Key == opts.strategy {
			return def.Build(inputs, opts.metric, !opts.maximize)
		}
	}
	return engine.StrategyDefs[0].Build(inputs, opts.metric, !opts.maximize)
}

// ── Plain-text (non-terminal) output ─────────────────────────────────────────

func runOptimizePlain(env *execenv.Env, resultCh <-chan engine.EvaluatedPoint, metric string, minimize bool, maxTrials int) error {
	var results []engine.EvaluatedPoint
	for p := range resultCh {
		results = append(results, p)
		n := len(results)

		prefix := fmt.Sprintf("Trial %d", n)
		if maxTrials > 0 {
			prefix = fmt.Sprintf("Trial %d/%d", n, maxTrials)
		}

		params := candidateStr(p.Candidate)
		if p.Err != nil {
			env.Out.Printf("%s: %s → error: %v\n", prefix, params, p.Err)
			continue
		}

		v, ok := p.Metrics[metric]
		metricStr := "-"
		if ok {
			metricStr = formatMetricValue(v, metric)
		}

		bestIdx, _ := engine.BestPoint(results, metric, minimize)
		isBest := bestIdx == len(results)-1

		line := fmt.Sprintf("%s: %s → %s", prefix, params, metricStr)
		if isBest {
			line += " ← best"
		}
		env.Out.Printf("%s\n", line)
	}

	if bestIdx, ok := engine.BestPoint(results, metric, minimize); ok {
		b := results[bestIdx]
		env.Out.Printf("\nBest: %s (%s)\n", candidateStr(b.Candidate), formatMetricValue(b.Metrics[metric], metric))
	}
	return nil
}

// ── JSON output ───────────────────────────────────────────────────────────────

type optimizeJSONOutput struct {
	Benchmark   string                        `json:"benchmark"`
	Strategy    string                        `json:"strategy"`
	Metric      string                        `json:"metric"`
	Minimize    bool                          `json:"minimize"`
	Trials      []optimizeJSONTrial           `json:"trials"`
	Best        *optimizeJSONTrial            `json:"best,omitempty"`
	Correlation map[string]map[string]float64 `json:"correlation,omitempty"`
}

type optimizeJSONTrial struct {
	N       int                `json:"n"`
	Params  map[string]string  `json:"params"`
	Metrics map[string]float64 `json:"metrics,omitempty"`
	Error   string             `json:"error,omitempty"`
}

func runOptimizeJSON(env *execenv.Env, bench engine.BenchInfo, strategy engine.Strategy, resultCh <-chan engine.EvaluatedPoint, metric string, minimize bool) error {
	var results []engine.EvaluatedPoint
	for p := range resultCh {
		results = append(results, p)
	}

	out := optimizeJSONOutput{
		Benchmark: bench.Name,
		Strategy:  strategy.Name(),
		Metric:    metric,
		Minimize:  minimize,
		Trials:    make([]optimizeJSONTrial, len(results)),
	}

	for i, r := range results {
		trial := optimizeJSONTrial{
			N:      i + 1,
			Params: make(map[string]string, len(r.Candidate)),
		}
		for _, a := range r.Candidate {
			trial.Params[a.Input.Name()] = a.Value
		}
		if r.Err != nil {
			trial.Error = r.Err.Error()
		} else {
			trial.Metrics = r.Metrics
		}
		out.Trials[i] = trial
	}

	if bestIdx, ok := engine.BestPoint(results, metric, minimize); ok {
		r := results[bestIdx]
		best := optimizeJSONTrial{
			N:       bestIdx + 1,
			Params:  make(map[string]string, len(r.Candidate)),
			Metrics: r.Metrics,
		}
		for _, a := range r.Candidate {
			best.Params[a.Input.Name()] = a.Value
		}
		out.Best = &best
	}

	unitSet := make(map[string]struct{})
	for _, r := range results {
		for u := range r.Metrics {
			unitSet[u] = struct{}{}
		}
	}
	units := preferredOptimizeUnits(unitSet, metric)
	if len(results) >= 3 {
		out.Correlation = engine.InputCorrelations(results, units)
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	env.Out.Printf("%s\n", data)
	return nil
}

// ── TUI model ─────────────────────────────────────────────────────────────────

// optimizeModelResult is written by the model before it quits so the caller
// can act on it (e.g. print a full report when the user presses [p]).
type optimizeModelResult struct {
	print   bool
	results []engine.EvaluatedPoint
}

type optimizeModel struct {
	bench     engine.BenchInfo
	inputs    []engine.InputInfo
	metric    string
	minimize  bool
	strategy  engine.Strategy
	maxTrials int

	results     []engine.EvaluatedPoint
	done        bool
	stopped     bool // true when the user pressed [s]
	tableScroll int

	width, height int
	ready         bool
	resultCh      <-chan engine.EvaluatedPoint
	stopFn        func() // cancels the optimization goroutine

	chart    linechart.Model
	hasChart bool

	style  execenv.Style
	result *optimizeModelResult // written before tea.Quit
}

func newOptimizeModel(
	bench engine.BenchInfo,
	inputs []engine.InputInfo,
	metric string,
	minimize bool,
	strategy engine.Strategy,
	maxTrials int,
	resultCh <-chan engine.EvaluatedPoint,
	stopFn func(),
	style execenv.Style,
	result *optimizeModelResult,
) optimizeModel {
	return optimizeModel{
		bench:     bench,
		inputs:    inputs,
		metric:    metric,
		minimize:  minimize,
		strategy:  strategy,
		maxTrials: maxTrials,
		resultCh:  resultCh,
		stopFn:    stopFn,
		style:     style,
		result:    result,
	}
}

type optimizePointMsg engine.EvaluatedPoint
type optimizeDoneMsg struct{}

func listenForOptimizeResult(ch <-chan engine.EvaluatedPoint) tea.Cmd {
	return func() tea.Msg {
		p, ok := <-ch
		if !ok {
			return optimizeDoneMsg{}
		}
		return optimizePointMsg(p)
	}
}

func (m optimizeModel) Init() tea.Cmd {
	return listenForOptimizeResult(m.resultCh)
}

func (m optimizeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		m = m.rebuildChart()

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "p":
			m.result.print = true
			m.result.results = m.results
			return m, tea.Quit
		case "s":
			if !m.done {
				m.stopped = true
				m.stopFn()
			}
		case "up", "k":
			if m.tableScroll > 0 {
				m.tableScroll--
			}
		case "down", "j":
			maxScroll := len(m.results) - m.tableBodyRows()
			if m.tableScroll < maxScroll {
				m.tableScroll++
			}
		case "pgup":
			m.tableScroll -= m.tableBodyRows()
			m = m.clampScroll()
		case "pgdown":
			m.tableScroll += m.tableBodyRows()
			m = m.clampScroll()
		}

	case optimizePointMsg:
		m.results = append(m.results, engine.EvaluatedPoint(msg))
		if m.ready && m.is1D() {
			m = m.rebuildChart()
		}
		return m, listenForOptimizeResult(m.resultCh)

	case optimizeDoneMsg:
		m.done = true
		// Scroll so the best result is centred in the table.
		if bestIdx, ok := engine.BestPoint(m.results, m.metric, m.minimize); ok {
			sorted := m.sortedResultIndices()
			for dispIdx, resIdx := range sorted {
				if resIdx == bestIdx {
					m.tableScroll = dispIdx - m.tableBodyRows()/2
					break
				}
			}
		}
		m = m.clampScroll()
	}
	return m, nil
}

func (m optimizeModel) is1D() bool {
	if len(m.inputs) != 1 {
		return false
	}
	_, isBool := m.inputs[0].(engine.BoolType)
	return !isBool
}

func (m optimizeModel) chartHeight() int {
	h := m.height / 3
	if h < 8 {
		h = 8
	}
	if h > 18 {
		h = 18
	}
	return h
}

func (m optimizeModel) corrHeight() int {
	if len(m.results) < 3 {
		return 0
	}
	return 1 + len(m.inputs) // metric header row + one row per input
}

func (m optimizeModel) headerRows() int {
	return 3 + m.corrHeight() // title + status + corr matrix + separator
}

func (m optimizeModel) visibleUnits() []string {
	unitSet := make(map[string]struct{})
	for _, r := range m.results {
		for u := range r.Metrics {
			unitSet[u] = struct{}{}
		}
	}
	return preferredOptimizeUnits(unitSet, m.metric)
}

func (m optimizeModel) tableBodyRows() int {
	used := m.headerRows() + 1 + 1 // column header + status bar
	if m.is1D() && m.hasChart {
		used += m.chartHeight() + 1 // chart + newline
	}
	h := m.height - used
	if h < 2 {
		h = 2
	}
	return h
}

func (m optimizeModel) clampScroll() optimizeModel {
	max := len(m.results) - m.tableBodyRows()
	if max < 0 {
		max = 0
	}
	if m.tableScroll > max {
		m.tableScroll = max
	}
	if m.tableScroll < 0 {
		m.tableScroll = 0
	}
	return m
}

func (m optimizeModel) rebuildChart() optimizeModel {
	if !m.is1D() || len(m.results) == 0 {
		return m
	}

	type pt struct{ x, y float64 }
	var pts []pt
	for _, r := range m.results {
		if r.Err != nil {
			continue
		}
		y, ok := r.Metrics[m.metric]
		if !ok {
			continue
		}
		pts = append(pts, pt{r.Candidate[0].FloatVal, y})
	}
	if len(pts) == 0 {
		return m
	}

	sort.Slice(pts, func(i, j int) bool { return pts[i].x < pts[j].x })

	minX, maxX := pts[0].x, pts[len(pts)-1].x
	minY, maxY := pts[0].y, pts[0].y
	for _, p := range pts[1:] {
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// Ensure non-zero x range for single-point case.
	if minX == maxX {
		minX -= 1
		maxX += 1
	}

	pad := (maxY - minY) * 0.15
	if pad == 0 {
		pad = math.Abs(minY) * 0.1
	}
	if pad == 0 {
		pad = 1
	}

	chartW := m.width - 2
	if chartW < 10 {
		chartW = 10
	}

	ch := linechart.New(chartW, m.chartHeight(), minX, maxX, minY-pad, maxY+pad,
		linechart.WithXYSteps(1, 2),
	)
	ch.DrawXYAxisAndLabel()

	for i := 1; i < len(pts); i++ {
		ch.DrawBrailleLine(
			canvas.Float64Point{X: pts[i-1].x, Y: pts[i-1].y},
			canvas.Float64Point{X: pts[i].x, Y: pts[i].y},
		)
	}
	// Always mark each point individually (covers the single-point case too).
	for _, p := range pts {
		ch.DrawBrailleLine(
			canvas.Float64Point{X: p.x, Y: p.y},
			canvas.Float64Point{X: p.x, Y: p.y},
		)
	}

	m.chart = ch
	m.hasChart = true
	return m
}

func (m optimizeModel) View() tea.View {
	if !m.ready {
		return tea.NewView("\n  Initializing...")
	}

	var sb strings.Builder

	// Title
	paramNames := make([]string, len(m.inputs))
	for i, inp := range m.inputs {
		paramNames[i] = inp.Name()
	}
	sb.WriteString(
		m.style.Bold("Optimize: "+m.bench.Name) +
			m.style.TonedDown("  params: "+strings.Join(paramNames, ", ")) +
			m.style.TonedDown("  strategy: "+m.strategy.Name()) + "\n",
	)

	// Status / progress line
	progress := fmt.Sprintf("%d trials", len(m.results))
	if m.maxTrials > 0 {
		progress = fmt.Sprintf("%d / %d", len(m.results), m.maxTrials)
	}
	runStatus := m.style.TonedDown("running… " + progress)
	if m.done {
		if m.stopped {
			runStatus = m.style.Bold("stopped") + m.style.TonedDown("  "+progress)
		} else {
			runStatus = m.style.Bold("done") + m.style.TonedDown("  "+progress)
		}
	}

	bestStr := "—"
	if bestIdx, ok := engine.BestPoint(m.results, m.metric, m.minimize); ok {
		bv := m.results[bestIdx].Metrics[m.metric]
		bestStr = formatMetricValue(bv, m.metric) + "  @ " + candidateStr(m.results[bestIdx].Candidate)
	}
	sb.WriteString(runStatus + "  " + m.style.TonedDown("best: ") + m.style.Accent(bestStr) + "\n")

	// Correlation matrix (shown in header once enough data is available).
	units := m.visibleUnits()
	sb.WriteString(m.renderCorr(units))

	sb.WriteString(strings.Repeat("─", m.width) + "\n")

	// Chart (1D only)
	if m.is1D() && m.hasChart {
		sb.WriteString(m.chart.View() + "\n")
	}

	// Table
	sb.WriteString(m.renderTable(units))

	// Footer
	footerParts := []string{"[↑↓jk⇞⇟] scroll"}
	if !m.done {
		footerParts = append(footerParts, "[s] stop")
	}
	footerParts = append(footerParts, "[p] print+quit", "[q] quit")
	sb.WriteString(m.style.TonedDown(strings.Join(footerParts, "  ")))

	view := tea.NewView(sb.String())
	view.AltScreen = true
	return view
}

func (m optimizeModel) sortedResultIndices() []int { return sortResultIndices(m.results) }

func (m optimizeModel) renderTable(units []string) string {
	bestIdx, _ := engine.BestPoint(m.results, m.metric, m.minimize)
	lastResIdx := len(m.results) - 1
	sorted := m.sortedResultIndices()
	lastBg := m.style.SelectionBg()

	var sb strings.Builder
	sb.WriteString(m.style.TonedDown(resultTableHeader(m.inputs, units)) + "\n")

	start := m.tableScroll
	if start < 0 {
		start = 0
	}
	end := start + m.tableBodyRows()
	if end > len(sorted) {
		end = len(sorted)
	}

	for dispIdx := start; dispIdx < end; dispIdx++ {
		resIdx := sorted[dispIdx]
		line := resultTableRow(m.results[resIdx], units, resIdx)
		if resIdx == bestIdx {
			line += "  " + m.style.Accent("← best")
		}
		if resIdx == lastResIdx {
			line = m.style.WithBg(lastBg, line)
		}
		sb.WriteString(line + "\n")
	}

	return sb.String()
}

func (m optimizeModel) renderCorr(units []string) string {
	if m.corrHeight() == 0 || len(units) == 0 {
		return ""
	}

	corr := engine.InputCorrelations(m.results, units)

	nameW := len("correlation")
	for _, inp := range m.inputs {
		if l := len(inp.Name()); l > nameW {
			nameW = l
		}
	}
	const colW = 6 // "+0.82" = 5 chars, padded to 6

	var sb strings.Builder

	// Metric header row.
	row := padRight(ansi.Truncate("correlation", nameW, "…"), nameW)
	for _, u := range units {
		row += "  " + padRight(ansi.Truncate(u, colW, "…"), colW)
	}
	sb.WriteString(m.style.TonedDown(row) + "\n")

	// One row per input.
	for _, inp := range m.inputs {
		mc := corr[inp.Name()]
		row = padRight(inp.Name(), nameW)
		for _, u := range units {
			r, ok := mc[u]
			if !ok {
				row += "  " + padRight("—", colW)
			} else {
				row += "  " + m.style.Corr(r, padRight(fmt.Sprintf("%+.2f", r), colW))
			}
		}
		sb.WriteString(row + "\n")
	}

	if m.strategy.Name() != "random" {
		sb.WriteString(m.style.TonedDown("  ↳ correlation is biased for " + m.strategy.Name() + " (not uniform sampling)\n"))
	}

	return sb.String()
}

func printOptimizeReport(env *execenv.Env, results []engine.EvaluatedPoint, metric string, minimize bool) {
	if len(results) == 0 {
		return
	}

	unitSet := make(map[string]struct{})
	for _, r := range results {
		for u := range r.Metrics {
			unitSet[u] = struct{}{}
		}
	}
	units := preferredOptimizeUnits(unitSet, metric)
	bestIdx, _ := engine.BestPoint(results, metric, minimize)

	inputs := make([]engine.InputInfo, len(results[0].Candidate))
	for i, a := range results[0].Candidate {
		inputs[i] = a.Input
	}

	env.Out.Printf("%s\n", resultTableHeader(inputs, units))
	for _, resIdx := range sortResultIndices(results) {
		line := resultTableRow(results[resIdx], units, resIdx)
		if resIdx == bestIdx {
			line += "  ← best"
		}
		env.Out.Printf("%s\n", line)
	}

	if bestIdx >= 0 {
		b := results[bestIdx]
		env.Out.Printf("\nBest: %s (%s)\n", candidateStr(b.Candidate), formatMetricValue(b.Metrics[metric], metric))
	}

	// Correlation matrix.
	if len(results) >= 3 {
		corr := engine.InputCorrelations(results, units)
		nameW := len("correlation")
		for _, inp := range inputs {
			if l := len(inp.Name()); l > nameW {
				nameW = l
			}
		}
		const colW = 6
		env.Out.Printf("\n")
		header := padRight("correlation", nameW)
		for _, u := range units {
			header += "  " + padRight(ansi.Truncate(u, colW, "…"), colW)
		}
		env.Out.Printf("%s\n", header)
		for _, inp := range inputs {
			mc := corr[inp.Name()]
			line := padRight(inp.Name(), nameW)
			for _, u := range units {
				r, ok := mc[u]
				if !ok {
					line += "  " + padRight("—", colW)
				} else {
					line += "  " + padRight(fmt.Sprintf("%+.2f", r), colW)
				}
			}
			env.Out.Printf("%s\n", line)
		}
		env.Out.Printf("  note: correlation is reliable only with the random strategy\n")
	}
}

// ── Helpers ───────────────────────────────────────────────────────────────────

// sortResultIndices returns result indices sorted by parameter values
// (numerically, left-to-right across params). Evaluation order is preserved
// as a tiebreak so the sort is stable.
func sortResultIndices(results []engine.EvaluatedPoint) []int {
	indices := make([]int, len(results))
	for i := range indices {
		indices[i] = i
	}
	sort.SliceStable(indices, func(a, b int) bool {
		ca, cb := results[indices[a]].Candidate, results[indices[b]].Candidate
		for j := 0; j < len(ca) && j < len(cb); j++ {
			va, vb := ca[j].FloatVal, cb[j].FloatVal
			if va != vb {
				return va < vb
			}
		}
		return false
	})
	return indices
}

func resultTableHeader(inputs []engine.InputInfo, units []string) string {
	h := fmt.Sprintf("%-5s", "#")
	for _, inp := range inputs {
		h += fmt.Sprintf("  %-14s", ansi.Truncate(inp.Name(), 14, "…"))
	}
	for _, u := range units {
		h += fmt.Sprintf("  %-12s", u)
	}
	return h
}

// resultTableRow builds a plain (unstyled) result row. Callers append their own
// best-marker and any terminal styling on top.
func resultTableRow(r engine.EvaluatedPoint, units []string, resIdx int) string {
	line := fmt.Sprintf("%-5d", resIdx+1)
	for _, a := range r.Candidate {
		line += "  " + padRight(displayParamValue(a), 14)
	}
	if r.Err != nil {
		line += fmt.Sprintf("  error: %v", r.Err)
	} else {
		for _, u := range units {
			v, ok := r.Metrics[u]
			cell := "-"
			if ok {
				cell = formatMetricValue(v, u)
			}
			line += "  " + padRight(cell, 12)
		}
	}
	return line
}

// displayParamValue formats a float parameter value for fixed-width table columns.
// Uses %e notation with 3 decimal places so the string is always exactly 10 chars
// (e.g. "1.234e+02"), giving stable column alignment regardless of magnitude.
// Non-float types are returned as-is (their natural representation fits easily).
func displayParamValue(a engine.ParamAssignment) string {
	switch a.Input.(type) {
	case engine.FloatType, engine.FloatLogType:
		return fmt.Sprintf("%.4e", a.FloatVal)
	}
	return a.Value
}

func candidateStr(c engine.Candidate) string {
	parts := make([]string, len(c))
	for i, a := range c {
		parts[i] = a.Input.Name() + "=" + a.Value
	}
	return strings.Join(parts, ", ")
}

func preferredOptimizeUnits(unitSet map[string]struct{}, primary string) []string {
	var units []string
	add := func(u string) {
		if _, ok := unitSet[u]; !ok {
			return
		}
		for _, existing := range units {
			if existing == u {
				return
			}
		}
		units = append(units, u)
	}
	add(primary)
	for _, u := range []string{"ns/op", "B/op", "allocs/op"} {
		add(u)
	}
	for u := range unitSet {
		add(u)
	}
	return units
}
