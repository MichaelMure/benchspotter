package commands

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/google/pprof/profile"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

var memMetricIds = map[engine.MemMetric][]string{
	engine.MetricAllocSpace:   {"alloc_space"},
	engine.MetricAllocObjects: {"alloc_objects"},
	engine.MetricInuseSpace:   {"inuse_space"},
	engine.MetricInuseObjects: {"inuse_objects"},
}

// profileViewModel implements execenv.InteractiveModel for all profile types.
// For mem profiles, [m] cycles through the four available metrics.
// [s] cycles through sort orders for all profile types.
type profileViewModel struct {
	prof        *profile.Profile
	profileType engine.Profile
	metric      engine.MemMetric // only meaningful for ProfileMem
	sort        engine.SortOrder
	top         int
}

func (m *profileViewModel) Render() string {
	var valueIdx int
	if m.profileType == engine.ProfileMem {
		valueIdx = engine.MemMetricIndex(m.prof, m.metric)
	} else {
		valueIdx = engine.PickValueIndex(m.prof, m.profileType)
	}
	funcs := engine.AggregateFuncs(m.prof, valueIdx, m.sort)
	return renderProfileTable(funcs, m.top, m.fmtValue(), m.sort)
}

func (m *profileViewModel) Status() string {
	if m.profileType == engine.ProfileMem {
		return fmt.Sprintf("[s] sort: %s    [m] metric: %s    [q] quit", m.sort, m.metric)
	}
	return fmt.Sprintf("[s] sort: %s    [q] quit", m.sort)
}

func (m *profileViewModel) HandleKey(key string) bool {
	switch key {
	case "s":
		m.sort = m.sort.Next()
		return true
	case "m":
		if m.profileType == engine.ProfileMem {
			m.metric = m.metric.Next()
			return true
		}
	}
	return false
}

// fmtValue returns the appropriate value formatter for the current metric.
func (m *profileViewModel) fmtValue() func(time.Duration) string {
	if m.profileType == engine.ProfileMem && m.metric.IsCount() {
		return func(d time.Duration) string { return fmt.Sprintf("%d", int64(d)) }
	}
	p := m.profileType
	return func(d time.Duration) string { return formatDuration(d, p) }
}

var sortOrderIds = map[engine.SortOrder][]string{
	engine.SortFlat:       {"flat"},
	engine.SortCumulative: {"cumulative"},
	engine.SortName:       {"name"},
}

type showProfileOptions struct {
	session     string
	bench       string
	metric      engine.MemMetric
	sort        engine.SortOrder
	profileType engine.Profile
	top         int
}

func newShowCPUCommand(env *execenv.Env) *cobra.Command {
	return newShowProfileCommand(env, engine.ProfileCPU, "cpu", "Show hot functions from a CPU profile")
}

func newShowMemCommand(env *execenv.Env) *cobra.Command {
	return newShowProfileCommand(env, engine.ProfileMem, "mem", "Show allocation sites from a memory profile")
}

func newShowBlockCommand(env *execenv.Env) *cobra.Command {
	return newShowProfileCommand(env, engine.ProfileBlock, "block", "Show blocking contention hotspots")
}

func newShowMutexCommand(env *execenv.Env) *cobra.Command {
	return newShowProfileCommand(env, engine.ProfileMutex, "mutex", "Show mutex contention hotspots")
}

func newShowProfileCommand(env *execenv.Env, profileType engine.Profile, use, short string) *cobra.Command {
	options := showProfileOptions{
		profileType: profileType,
		top:         20,
	}

	cmd := &cobra.Command{
		Use:     use,
		Short:   short,
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShowProfile(cmd.Context(), env, options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.session, "session", "", "Session ID to inspect")
	flags.StringVar(&options.bench, "bench", "", "Benchmark name to show (default: merge all)")
	flags.Var(enumflag.New(&options.metric, "metric", memMetricIds, enumflag.EnumCaseInsensitive), "metric", "mem metric (alloc_space, alloc_objects, inuse_space, inuse_objects)")
	flags.Var(enumflag.New(&options.sort, "sort", sortOrderIds, enumflag.EnumCaseInsensitive), "sort", "sort order (flat, cumulative, name)")
	flags.IntVar(&options.top, "top", options.top, "Number of top functions to show")

	return cmd
}

func runShowProfile(ctx context.Context, env *execenv.Env, options showProfileOptions) error {
	var selection *engine.SessionInfo
	var err error

	sessionRecallKey := "show_profile_" + engine.ProfileDir(options.profileType) + "_session"
	benchRecallKey := "show_profile_" + engine.ProfileDir(options.profileType) + "_bench"
	interactive := options.session == ""

	if interactive {
		preSelected := env.Repo.GetRecall(sessionRecallKey)
		selection, err = inputs.SelectSession(ctx, env, preSelected,
			func(info *engine.SessionInfo) bool {
				return info.HasProfile(options.profileType)
			})
		if err != nil {
			return err
		}
		if err = env.Repo.SetRecall(sessionRecallKey, selection.Id); err != nil {
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

	if !selection.HasProfile(options.profileType) {
		return fmt.Errorf("session %q has no %s profile", selection.HumanName, engine.ProfileDir(options.profileType))
	}

	if options.bench == "" {
		benches, err := engine.ListProfileBenchmarks(env.Repo.Storage(), selection.Path, options.profileType)
		if err != nil {
			return err
		}
		switch {
		case len(benches) == 1:
			options.bench = benches[0]
		case interactive:
			summaries, err := engine.ListProfileBenchmarkSummaries(env.Repo.Storage(), selection.Path, options.profileType)
			if err != nil {
				return err
			}
			benchOpts := make([]inputs.BenchOption, len(summaries))
			for i, s := range summaries {
				benchOpts[i] = inputs.BenchOption{Name: s.Name, Label: formatBenchSummary(s, options.profileType)}
			}
			preSelected := env.Repo.GetRecall(benchRecallKey)
			options.bench, err = inputs.SelectProfileBench(ctx, env, benchOpts, preSelected)
			if err != nil {
				return err
			}
			if err = env.Repo.SetRecall(benchRecallKey, options.bench); err != nil {
				return err
			}
		default:
			return fmt.Errorf("session has multiple %s profiles, specify --bench: %s",
				engine.ProfileDir(options.profileType), strings.Join(benches, ", "))
		}
	}

	switch env.Format {
	case execenv.FormatRaw:
		data, err := engine.ReadProfileRaw(env.Repo.Storage(), selection.Path, options.profileType, options.bench)
		if err != nil {
			return err
		}
		_, err = env.Out.Write(data)
		return err

	case execenv.FormatJSON:
		funcs, err := engine.ReadProfileFunctions(env.Repo.Storage(), selection.Path, options.profileType, options.bench)
		if err != nil {
			return err
		}
		type jsonFunc struct {
			Name    string  `json:"name"`
			Flat    int64   `json:"flat_ns"`
			Cum     int64   `json:"cum_ns"`
			FlatPct float64 `json:"flat_pct"`
			CumPct  float64 `json:"cum_pct"`
		}
		top := min(options.top, len(funcs))
		out := make([]jsonFunc, top)
		for i, f := range funcs[:top] {
			out[i] = jsonFunc{
				Name:    f.Name,
				Flat:    int64(f.Flat),
				Cum:     int64(f.Cumulative),
				FlatPct: f.FlatPct,
				CumPct:  f.CumPct,
			}
		}
		return env.Out.PrintJSON(out)

	case execenv.FormatText:
		prof, err := engine.ReadParsedProfile(env.Repo.Storage(), selection.Path, options.profileType, options.bench)
		if err != nil {
			return err
		}
		model := &profileViewModel{
			prof:        prof,
			profileType: options.profileType,
			metric:      options.metric,
			sort:        options.sort,
			top:         options.top,
		}
		return env.ViewportWithKeys(ctx, model)()

	default:
		return fmt.Errorf("unsupported format %v for show %s (text, json, raw)", env.Format, engine.ProfileDir(options.profileType))
	}
}

func renderProfileTable(funcs []engine.ProfileFunc, top int, fmtValue func(time.Duration) string, order engine.SortOrder) string {
	top = min(top, len(funcs))
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	flat, flatPct, cum, cumPct, name := "FLAT", "FLAT%", "CUM", "CUM%", "FUNCTION"
	switch order {
	case engine.SortFlat:
		flat += " ▼"
		flatPct += " ▼"
	case engine.SortCumulative:
		cum += " ▼"
		cumPct += " ▼"
	case engine.SortName:
		name += " ▼"
	}
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", flat, flatPct, cum, cumPct, name)

	for _, f := range funcs[:top] {
		fmt.Fprintf(w, "%s\t%.2f%%\t%s\t%.2f%%\t%s\n",
			fmtValue(f.Flat),
			f.FlatPct,
			fmtValue(f.Cumulative),
			f.CumPct,
			f.Name,
		)
	}
	_ = w.Flush()
	return buf.String()
}

func formatBenchSummary(s engine.ProfileBenchSummary, p engine.Profile) string {
	if p == engine.ProfileMem {
		return fmt.Sprintf("%s  %s inuse / %s alloc",
			s.Name,
			formatDuration(time.Duration(s.Primary), p),
			formatDuration(time.Duration(s.AllocBytes), p))
	}
	return fmt.Sprintf("%s  %s", s.Name, formatDuration(time.Duration(s.Primary), p))
}

func formatDuration(d time.Duration, p engine.Profile) string {
	switch p {
	case engine.ProfileCPU, engine.ProfileBlock, engine.ProfileMutex:
		if d >= time.Second {
			return fmt.Sprintf("%.2fs", d.Seconds())
		}
		if d >= time.Millisecond {
			return fmt.Sprintf("%.2fms", float64(d)/float64(time.Millisecond))
		}
		return fmt.Sprintf("%.2fµs", float64(d)/float64(time.Microsecond))
	default:
		// mem: values are bytes, stored as nanoseconds numerically
		b := int64(d)
		switch {
		case b >= 1024*1024*1024:
			return fmt.Sprintf("%.2fGB", float64(b)/float64(1024*1024*1024))
		case b >= 1024*1024:
			return fmt.Sprintf("%.2fMB", float64(b)/float64(1024*1024))
		case b >= 1024:
			return fmt.Sprintf("%.2fKB", float64(b)/float64(1024))
		default:
			return fmt.Sprintf("%dB", b)
		}
	}
}
