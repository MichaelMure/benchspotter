package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/google/pprof/profile"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
	"golang.org/x/sys/execabs"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/commands/tabwriter"
	"benchspotter/engine"
)

// compareProfileViewModel implements execenv.InteractiveModel for profile diffs.
// [s] cycles through five sort orders. [m] cycles mem metrics (mem only).
// [o] opens the diff in pprof's web UI (go tool pprof -http -diff_base).
type compareProfileViewModel struct {
	baseProf    *profile.Profile
	newProf     *profile.Profile
	profileType engine.Profile
	metric      engine.MemMetric
	sort        engine.DiffSortOrder
	top         int
	baseLabel   string
	newLabel    string
	style       execenv.Style
	startPprof  func() (string, error)
	pprofURL    string
	pprofErr    string
}

func (m *compareProfileViewModel) Render() string {
	baseIdx, newIdx := m.valueIndices()
	funcs := engine.DiffProfileFuncs(m.baseProf, m.newProf, baseIdx, newIdx, m.sort)
	header := fmt.Sprintf("base: %s\nnew:  %s\n", m.baseLabel, m.newLabel)
	return header + "\n" + renderDiffTable(funcs, m.top, m.fmtValue(), m.sort, m.style)
}

func (m *compareProfileViewModel) Status() string {
	pprofHint := ""
	switch {
	case m.pprofURL != "":
		pprofHint = "    pprof: " + m.pprofURL
	case m.pprofErr != "":
		pprofHint = "    pprof error: " + m.pprofErr
	case m.startPprof != nil:
		pprofHint = "    [o] open in pprof"
	}
	if m.profileType == engine.ProfileMem {
		return fmt.Sprintf("[s] sort: %s    [m] metric: %s%s    [q] quit", m.sort, m.metric, pprofHint)
	}
	return fmt.Sprintf("[s] sort: %s%s    [q] quit", m.sort, pprofHint)
}

func (m *compareProfileViewModel) HandleKey(key string) bool {
	switch key {
	case "s":
		m.sort = m.sort.Next()
		return true
	case "m":
		if m.profileType == engine.ProfileMem {
			m.metric = m.metric.Next()
			return true
		}
	case "o":
		if m.startPprof != nil && m.pprofURL == "" {
			url, err := m.startPprof()
			if err != nil {
				m.pprofErr = err.Error()
			} else {
				m.pprofURL = url
			}
			return true
		}
	}
	return false
}

func (m *compareProfileViewModel) valueIndices() (baseIdx, newIdx int) {
	if m.profileType == engine.ProfileMem {
		return engine.MemMetricIndex(m.baseProf, m.metric), engine.MemMetricIndex(m.newProf, m.metric)
	}
	return engine.PickValueIndex(m.baseProf, m.profileType), engine.PickValueIndex(m.newProf, m.profileType)
}

func (m *compareProfileViewModel) fmtValue() func(time.Duration) string {
	if m.profileType == engine.ProfileMem && m.metric.IsCount() {
		return func(d time.Duration) string { return fmt.Sprintf("%d", int64(d)) }
	}
	p := m.profileType
	return func(d time.Duration) string { return formatDuration(d, p) }
}

var diffSortOrderIds = map[engine.DiffSortOrder][]string{
	engine.DiffSortAbsFlat:  {"|flat|", "absflat"},
	engine.DiffSortSignFlat: {"flat"},
	engine.DiffSortAbsCum:   {"|cum|", "abscum"},
	engine.DiffSortSignCum:  {"cum"},
	engine.DiffSortName:     {"name"},
}

type compareProfileOptions struct {
	baseSession string
	newSession  string
	bench       string
	metric      engine.MemMetric
	sort        engine.DiffSortOrder
	profileType engine.Profile
	top         int
	web         bool
}

func newCompareCPUCommand(env *execenv.Env) *cobra.Command {
	return newCompareProfileCommand(env, engine.ProfileCPU, "cpu",
		"Diff CPU profiles between two sessions",
		`Diff CPU profiles between two sessions.

Shows per-function flat and cumulative time deltas. Positive values (red) are
regressions; negative values (green) are improvements.

In the viewer: [s] cycles sort order (|flat|, flat, |cum|, cum, name — absolute
sorts surface the biggest movers first; signed sorts put regressions at the top),
[o] opens the diff in the pprof web UI (go tool pprof -http -diff_base) for a
full interactive call-graph view.

Any interactive prompt can be bypassed with the corresponding flags.`)
}

func newCompareMemCommand(env *execenv.Env) *cobra.Command {
	return newCompareProfileCommand(env, engine.ProfileMem, "mem",
		"Diff memory profiles between two sessions",
		`Diff memory profiles between two sessions.

Shows per-function flat and cumulative allocation deltas. Four metrics are
available — alloc_space and alloc_objects count everything allocated over the
lifetime of the benchmark; inuse_space and inuse_objects show what was still
live at the time the profile was taken.

In the viewer: [m] cycles through the four metrics, [s] cycles sort order
(|flat|, flat, |cum|, cum, name), [o] opens the diff in the pprof web UI.

Any interactive prompt can be bypassed with the corresponding flags.`)
}

func newCompareMutexCommand(env *execenv.Env) *cobra.Command {
	return newCompareProfileCommand(env, engine.ProfileMutex, "mutex",
		"Diff mutex contention profiles between two sessions",
		`Diff mutex contention profiles between two sessions.

Shows per-function flat and cumulative mutex wait-time deltas. Positive values
(red) indicate increased contention; negative values (green) indicate less.

In the viewer: [s] cycles sort order (|flat|, flat, |cum|, cum, name), [o]
opens the diff in the pprof web UI.

Any interactive prompt can be bypassed with the corresponding flags.`)
}

func newCompareBlockCommand(env *execenv.Env) *cobra.Command {
	return newCompareProfileCommand(env, engine.ProfileBlock, "block",
		"Diff blocking contention profiles between two sessions",
		`Diff blocking contention profiles between two sessions.

Shows per-function flat and cumulative blocking wait-time deltas. Positive
values (red) indicate more blocking; negative values (green) indicate less.

In the viewer: [s] cycles sort order (|flat|, flat, |cum|, cum, name), [o]
opens the diff in the pprof web UI.

Any interactive prompt can be bypassed with the corresponding flags.`)
}

func newCompareProfileCommand(env *execenv.Env, profileType engine.Profile, use, short, long string) *cobra.Command {
	options := compareProfileOptions{
		profileType: profileType,
		top:         20,
	}

	cmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCompareProfile(env, options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.baseSession, "base", "", "Base session ID or name")
	flags.StringVar(&options.newSession, "new", "", "New session ID or name")
	flags.StringVar(&options.bench, "bench", "", "Benchmark name to compare (default: use common benchmark)")
	flags.Var(enumflag.New(&options.metric, "metric", memMetricIds, enumflag.EnumCaseInsensitive),
		"metric", "mem metric (alloc_space, alloc_objects, inuse_space, inuse_objects)")
	flags.Var(enumflag.New(&options.sort, "sort", diffSortOrderIds, enumflag.EnumCaseInsensitive),
		"sort", "sort order (|flat|, flat, |cum|, cum, name)")
	flags.IntVar(&options.top, "top", options.top, "Number of top functions to show")
	flags.BoolVar(&options.web, "web", false, "Open profile diff in pprof web UI instead of TUI")

	return cmd
}

func runCompareProfile(env *execenv.Env, options compareProfileOptions) error {
	profileDir := engine.ProfileDir(options.profileType)
	baseRecallKey := "compare_profile_" + profileDir + "_base"
	newRecallKey := "compare_profile_" + profileDir + "_new"
	benchRecallKey := "compare_profile_" + profileDir + "_bench"
	interactive := options.baseSession == "" || options.newSession == ""

	profileFilter := func(info *engine.SessionInfo) bool {
		return info.HasProfile(options.profileType)
	}

	var baseInfo, newInfo *engine.SessionInfo
	var err error

	if options.baseSession == "" {
		baseInfo, err = inputs.SelectSession(env, "Select base session", env.Repo.GetRecall(baseRecallKey), profileFilter)
		if err != nil {
			return err
		}
		if err = env.Repo.SetRecall(baseRecallKey, baseInfo.Id); err != nil {
			return err
		}
	} else {
		baseInfo, err = engine.LocateSession(env.Repo.Storage(), options.baseSession)
		if err != nil {
			return err
		}
	}
	if !baseInfo.HasProfile(options.profileType) {
		return fmt.Errorf("session %q has no %s profile", baseInfo.HumanName, profileDir)
	}

	if options.newSession == "" {
		newInfo, err = inputs.SelectSession(env, "Select new session", env.Repo.GetRecall(newRecallKey), profileFilter)
		if err != nil {
			return err
		}
		if err = env.Repo.SetRecall(newRecallKey, newInfo.Id); err != nil {
			return err
		}
	} else {
		newInfo, err = engine.LocateSession(env.Repo.Storage(), options.newSession)
		if err != nil {
			return err
		}
	}
	if !newInfo.HasProfile(options.profileType) {
		return fmt.Errorf("session %q has no %s profile", newInfo.HumanName, profileDir)
	}

	// Resolve bench: find common benchmarks across both sessions.
	if options.bench == "" {
		baseBenches, err := engine.ListProfileBenchmarks(env.Repo.Storage(), baseInfo.Path, options.profileType)
		if err != nil {
			return err
		}
		newBenches, err := engine.ListProfileBenchmarks(env.Repo.Storage(), newInfo.Path, options.profileType)
		if err != nil {
			return err
		}
		newSet := make(map[string]bool, len(newBenches))
		for _, b := range newBenches {
			newSet[b] = true
		}
		var common []string
		for _, b := range baseBenches {
			if newSet[b] {
				common = append(common, b)
			}
		}
		switch {
		case len(common) == 0:
			return fmt.Errorf("no common %s benchmark profiles between %q and %q",
				profileDir, baseInfo.HumanName, newInfo.HumanName)
		case len(common) == 1:
			options.bench = common[0]
		case interactive:
			benchOpts := make([]inputs.BenchOption, len(common))
			for i, name := range common {
				benchOpts[i] = inputs.BenchOption{Name: name, Label: name}
			}
			options.bench, err = inputs.SelectProfileBench(env, benchOpts, env.Repo.GetRecall(benchRecallKey))
			if err != nil {
				return err
			}
			if err = env.Repo.SetRecall(benchRecallKey, options.bench); err != nil {
				return err
			}
		default:
			return fmt.Errorf("sessions have multiple common %s profiles, specify --bench: %s",
				profileDir, strings.Join(common, ", "))
		}
	}

	// Web mode: launch pprof diff UI and block until Ctrl+C.
	if options.web {
		basePath, err := engine.ProfileFilePath(env.Repo.Storage(), baseInfo.Path, options.profileType, options.bench)
		if err != nil {
			return err
		}
		newPath, err := engine.ProfileFilePath(env.Repo.Storage(), newInfo.Path, options.profileType, options.bench)
		if err != nil {
			return err
		}
		url, err := launchPprofDiff(env.Ctx, basePath, newPath)
		if err != nil {
			return fmt.Errorf("launching pprof: %w", err)
		}
		fmt.Fprintf(env.Out, "base:  %s\nnew:   %s\npprof: %s\n\nPress Ctrl+C to stop.\n",
			baseInfo.HumanName, newInfo.HumanName, url)
		<-env.Ctx.Done()
		return nil
	}

	baseProf, err := engine.ReadParsedProfile(env.Repo.Storage(), baseInfo.Path, options.profileType, options.bench)
	if err != nil {
		return fmt.Errorf("reading base profile: %w", err)
	}
	newProf, err := engine.ReadParsedProfile(env.Repo.Storage(), newInfo.Path, options.profileType, options.bench)
	if err != nil {
		return fmt.Errorf("reading new profile: %w", err)
	}

	switch env.Format {
	case execenv.FormatRaw:
		// Raw() bypasses the color-profile writer: a pprof profile is binary
		// protobuf, and ANSI stripping on a non-terminal would corrupt it.
		return engine.DiffProfileRaw(env.Repo.Storage(), baseInfo.Path, newInfo.Path, options.profileType, options.bench, env.Out.Raw())

	case execenv.FormatText:
		model := &compareProfileViewModel{
			baseProf:    baseProf,
			newProf:     newProf,
			profileType: options.profileType,
			metric:      options.metric,
			sort:        options.sort,
			top:         options.top,
			baseLabel:   baseInfo.HumanName,
			newLabel:    newInfo.HumanName,
			style:       env.Style,
		}
		basePath, baseErr := engine.ProfileFilePath(env.Repo.Storage(), baseInfo.Path, options.profileType, options.bench)
		newPath, newErr := engine.ProfileFilePath(env.Repo.Storage(), newInfo.Path, options.profileType, options.bench)
		if baseErr == nil && newErr == nil {
			model.startPprof = func() (string, error) {
				return launchPprofDiff(env.Ctx, basePath, newPath)
			}
		}
		return env.ViewportWithKeys(model)()

	case execenv.FormatJSON:
		baseIdx, newIdx := diffValueIndices(baseProf, newProf, options.profileType, options.metric)
		funcs := engine.DiffProfileFuncs(baseProf, newProf, baseIdx, newIdx, engine.DiffSortAbsFlat)
		top := min(options.top, len(funcs))
		type jsonDiffFunc struct {
			Name         string  `json:"name"`
			File         string  `json:"file,omitempty"`
			Line         int64   `json:"line,omitempty"`
			BaseFlat     int64   `json:"base_flat_ns"`
			NewFlat      int64   `json:"new_flat_ns"`
			DeltaFlat    int64   `json:"delta_flat_ns"`
			DeltaFlatPct float64 `json:"delta_flat_pct"`
			BaseCum      int64   `json:"base_cum_ns"`
			NewCum       int64   `json:"new_cum_ns"`
			DeltaCum     int64   `json:"delta_cum_ns"`
			DeltaCumPct  float64 `json:"delta_cum_pct"`
		}
		out := make([]jsonDiffFunc, top)
		for i, f := range funcs[:top] {
			out[i] = jsonDiffFunc{
				Name:         f.Name,
				File:         f.File,
				Line:         f.StartLine,
				BaseFlat:     int64(f.BaseFlat),
				NewFlat:      int64(f.NewFlat),
				DeltaFlat:    int64(f.DeltaFlat),
				DeltaFlatPct: f.DeltaFlatPct,
				BaseCum:      int64(f.BaseCum),
				NewCum:       int64(f.NewCum),
				DeltaCum:     int64(f.DeltaCum),
				DeltaCumPct:  f.DeltaCumPct,
			}
		}
		return env.Out.PrintJSON(out)

	default:
		return fmt.Errorf("unsupported format %v for compare %s (raw, json, text)", env.Format, profileDir)
	}
}

func diffValueIndices(base, new *profile.Profile, p engine.Profile, metric engine.MemMetric) (baseIdx, newIdx int) {
	if p == engine.ProfileMem {
		return engine.MemMetricIndex(base, metric), engine.MemMetricIndex(new, metric)
	}
	return engine.PickValueIndex(base, p), engine.PickValueIndex(new, p)
}

func renderDiffTable(funcs []engine.DiffProfileFunc, top int, fmtValue func(time.Duration) string, order engine.DiffSortOrder, style execenv.Style) string {
	top = min(top, len(funcs))
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	dflat, dflatPct, dcum, dcumPct, name := "ΔFLAT", "ΔFLAT%", "ΔCUM", "ΔCUM%", "FUNCTION"
	switch order {
	case engine.DiffSortAbsFlat, engine.DiffSortSignFlat:
		dflat = style.Accent(dflat + " ▼")
		dflatPct = style.Accent(dflatPct + " ▼")
	case engine.DiffSortAbsCum, engine.DiffSortSignCum:
		dcum = style.Accent(dcum + " ▼")
		dcumPct = style.Accent(dcumPct + " ▼")
	case engine.DiffSortName:
		name = style.Accent(name + " ▼")
	}
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", dflat, dflatPct, dcum, dcumPct, name)

	for _, f := range funcs[:top] {
		flatStr, flatPctStr := fmtDeltaCell(f.DeltaFlat, f.DeltaFlatPct, fmtValue, style)
		cumStr, cumPctStr := fmtDeltaCell(f.DeltaCum, f.DeltaCumPct, fmtValue, style)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", flatStr, flatPctStr, cumStr, cumPctStr, f.Name)
	}
	_ = w.Flush()
	return buf.String()
}

// fmtDeltaCell formats a delta value and its percentage for a table cell, applying
// color: red for regressions (positive delta), green for improvements (negative).
func fmtDeltaCell(d time.Duration, pct float64, fmtValue func(time.Duration) string, style execenv.Style) (val, pctStr string) {
	v := fmtDeltaWithSign(d, fmtValue)
	if d == 0 {
		pctStr = "-"
	} else {
		pctStr = fmt.Sprintf("%+.1f%%", pct)
	}
	switch {
	case d > 0:
		return style.Negative(v), style.Negative(pctStr)
	case d < 0:
		return style.Positive(v), style.Positive(pctStr)
	default:
		return v, pctStr
	}
}

// fmtDeltaWithSign formats the absolute value of d via fmtValue, prepending "+"
// or "-" explicitly so that fmtValue always receives a non-negative duration
// (required for correct unit selection in formatDuration).
func fmtDeltaWithSign(d time.Duration, fmtValue func(time.Duration) string) string {
	switch {
	case d > 0:
		return "+" + fmtValue(d)
	case d < 0:
		return "-" + fmtValue(-d)
	default:
		return fmtValue(0)
	}
}

// launchPprofDiff starts `go tool pprof -http -diff_base=basePath newPath` on a
// random free port and returns the URL once the process is running.
func launchPprofDiff(ctx context.Context, basePath, newPath string) (string, error) {
	port, err := findFreePort()
	if err != nil {
		return "", fmt.Errorf("no free port: %w", err)
	}
	addr := fmt.Sprintf("localhost:%d", port)
	pprofBin, err := execabs.Command("go", "tool", "-n", "pprof").Output()
	if err != nil {
		return "", fmt.Errorf("locate pprof: %w", err)
	}
	cmd := execabs.CommandContext(ctx, strings.TrimSpace(string(pprofBin)),
		"-http="+addr, "-diff_base="+basePath, newPath)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return "http://" + addr, cmd.Start()
}
