package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/google/pprof/profile"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
	"golang.org/x/sys/execabs"

	"github.com/MichaelMure/benchspotter/commands/execenv"
	"github.com/MichaelMure/benchspotter/commands/inputs"
	"github.com/MichaelMure/benchspotter/commands/tabwriter"
	"github.com/MichaelMure/benchspotter/engine"
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
// [a] toggles source line annotations (when sourcesRoot is set).
// [o] opens the profile in pprof's web UI (go tool pprof -http).
type profileViewModel struct {
	prof          *profile.Profile
	profileType   engine.Profile
	metric        engine.MemMetric // only meaningful for ProfileMem
	sort          engine.SortOrder
	top           int
	machineHeader string
	style         execenv.Style
	// source annotation
	sourcesRoot string
	sourceCache map[string][]byte // relPath → file content (nil = not found)
	annotate    bool
	// pprof web UI; lifecycle is managed by the context passed to startPprof.
	startPprof func() (url string, err error) // nil if unavailable
	pprofURL   string                         // set once pprof is running
	pprofErr   string                         // last error launching pprof
}

func (m *profileViewModel) Render() string {
	var valueIdx int
	if m.profileType == engine.ProfileMem {
		valueIdx = engine.MemMetricIndex(m.prof, m.metric)
	} else {
		valueIdx = engine.PickValueIndex(m.prof, m.profileType)
	}
	funcs := engine.AggregateFuncs(m.prof, valueIdx, m.sort)
	var getSource func(string, int64) (string, []string)
	if m.annotate && m.sourcesRoot != "" {
		getSource = m.getSourceLines
	}
	content := renderProfileTable(funcs, m.top, m.fmtValue(), m.sort, m.style, getSource)
	if m.machineHeader != "" {
		return m.machineHeader + "\n\n" + content
	}
	return content
}

// getSourceLines returns a short path to display and up to 3 source lines
// starting at startLine for the given file path, reading from disk. It tries
// sourcesRoot-relative first, then falls back to absFile directly (for stdlib
// and dependencies).
func (m *profileViewModel) getSourceLines(absFile string, startLine int64) (string, []string) {
	content, seen := m.sourceCache[absFile]
	if !seen {
		rel, err := filepath.Rel(m.sourcesRoot, absFile)
		if err == nil && !strings.HasPrefix(rel, "..") {
			content, _ = os.ReadFile(filepath.Join(m.sourcesRoot, rel))
		}
		if content == nil {
			content, _ = os.ReadFile(absFile)
		}
		m.sourceCache[absFile] = content // nil if not found — skip next time
	}
	if content == nil {
		return "", nil
	}
	lines := bytes.Split(content, []byte("\n"))
	start := int(startLine) - 1
	if start < 0 || start >= len(lines) {
		return "", nil
	}
	end := start + 3
	if end > len(lines) {
		end = len(lines)
	}
	result := make([]string, end-start)
	for i, l := range lines[start:end] {
		result[i] = strings.TrimRight(string(l), "\r")
	}
	return m.displayPath(absFile), result
}

// displayPath shortens a source path for display. Files inside the project are
// shown relative to its root, so the path can be opened as-is. Everything else
// — the standard library, module dependencies — is reduced to its file name:
// the FUNCTION column above already spells out the package, and the absolute
// path is long enough to push the annotation off the side of the terminal.
func (m *profileViewModel) displayPath(absFile string) string {
	if m.sourcesRoot != "" {
		rel, err := filepath.Rel(m.sourcesRoot, absFile)
		if err == nil && !strings.HasPrefix(rel, "..") {
			return filepath.ToSlash(rel)
		}
	}
	return filepath.Base(absFile)
}

func (m *profileViewModel) Status() string {
	annotateHint := ""
	if m.sourcesRoot != "" {
		if m.annotate {
			annotateHint = "    [a] source: on"
		} else {
			annotateHint = "    [a] source: off"
		}
	}
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
		return fmt.Sprintf("[s] sort: %s    [m] metric: %s%s%s    [q] quit", m.sort, m.metric, annotateHint, pprofHint)
	}
	return fmt.Sprintf("[s] sort: %s%s%s    [q] quit", m.sort, annotateHint, pprofHint)
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
	case "a":
		if m.sourcesRoot != "" {
			m.annotate = !m.annotate
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
	web         bool
}

func newShowCPUCommand(env *execenv.Env) *cobra.Command {
	return newShowProfileCommand(env, engine.ProfileCPU, "cpu", "Show hot functions from a CPU profile", `Show hot functions from a CPU profile.

A CPU profile tells you where the program is spending wall-clock time. The
table lists each function with its flat time (time spent inside the function
itself) and cumulative time (including all callees), both as absolute values
and as a percentage of total.

In the viewer: [s] cycles the sort order (flat / cumulative / name), [a]
toggles inline source-line annotations, [o] opens the profile in the pprof
web UI (go tool pprof -http) for a full call-graph view.

Any interactive prompt can be bypassed with the corresponding flags.`)
}

func newShowMemCommand(env *execenv.Env) *cobra.Command {
	return newShowProfileCommand(env, engine.ProfileMem, "mem", "Show allocation sites from a memory profile", `Show allocation sites from a memory profile.

A memory profile tells you what is allocating heap memory. Four metrics are
available — alloc_space and alloc_objects count everything allocated over the
lifetime of the benchmark; inuse_space and inuse_objects show what was still
live at the time the profile was taken.

In the viewer: [m] cycles through the four metrics, [s] cycles the sort order
(flat / cumulative / name), [a] toggles inline source-line annotations, [o]
opens the profile in the pprof web UI.

Any interactive prompt can be bypassed with the corresponding flags.`)
}

func newShowBlockCommand(env *execenv.Env) *cobra.Command {
	return newShowProfileCommand(env, engine.ProfileBlock, "block", "Show blocking contention hotspots", `Show blocking contention hotspots.

A blocking profile records goroutines waiting on synchronisation primitives —
channel sends/receives, select statements, and sync package calls. The values
are cumulative wait durations, so the top entries are where your goroutines
spend the most time blocked.

In the viewer: [s] cycles the sort order, [a] toggles inline source-line
annotations, [o] opens the profile in the pprof web UI.

Any interactive prompt can be bypassed with the corresponding flags.`)
}

func newShowMutexCommand(env *execenv.Env) *cobra.Command {
	return newShowProfileCommand(env, engine.ProfileMutex, "mutex", "Show mutex contention hotspots", `Show mutex contention hotspots.

A mutex profile records time spent waiting to acquire sync.Mutex and sync.RWMutex
locks. Unlike the blocking profile, it only covers mutex contention, not channels
or other waiting, which makes it easier to pinpoint lock-related bottlenecks.

In the viewer: [s] cycles the sort order, [a] toggles inline source-line
annotations, [o] opens the profile in the pprof web UI.

Any interactive prompt can be bypassed with the corresponding flags.`)
}

func newShowProfileCommand(env *execenv.Env, profileType engine.Profile, use, short, long string) *cobra.Command {
	options := showProfileOptions{
		profileType: profileType,
		top:         20,
	}

	cmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShowProfile(env, options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.session, "session", "", "Session ID to inspect")
	flags.StringVar(&options.bench, "bench", "", "Benchmark name to show (default: merge all)")
	// Only a memory profile carries several sample types to choose between;
	// for the others the value index is derived from the profile type.
	if profileType == engine.ProfileMem {
		flags.Var(enumflag.New(&options.metric, "metric", memMetricIds, enumflag.EnumCaseInsensitive),
			"metric", "which allocation metric to report (alloc_space, alloc_objects, inuse_space, inuse_objects)")
	}
	flags.Var(enumflag.New(&options.sort, "sort", sortOrderIds, enumflag.EnumCaseInsensitive), "sort", "sort order (flat, cumulative, name)")
	flags.IntVar(&options.top, "top", options.top, "Number of top functions to show")
	flags.BoolVar(&options.web, "web", false, "Open profile in pprof web UI instead of TUI")

	return cmd
}

func runShowProfile(env *execenv.Env, options showProfileOptions) error {
	var selection *engine.SessionInfo
	var err error

	sessionRecallKey := "show_profile_" + engine.ProfileDir(options.profileType) + "_session"
	benchRecallKey := "show_profile_" + engine.ProfileDir(options.profileType) + "_bench"
	interactive := options.session == ""

	if interactive {
		preSelected := env.Repo.GetRecall(sessionRecallKey)
		selection, err = inputs.SelectSession(env, "--session", "Select session", preSelected,
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
		selection, err = engine.LocateSession(env.Repo.Storage(), options.session)
		if err != nil {
			return err
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
			options.bench, err = inputs.SelectProfileBench(env, benchOpts, preSelected)
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

	if options.web {
		profilePath, err := engine.ProfileFilePath(env.Repo.Storage(), selection.Path, options.profileType, options.bench)
		if err != nil {
			return err
		}
		url, err := launchPprof(env.Ctx, profilePath)
		if err != nil {
			return fmt.Errorf("launching pprof: %w", err)
		}
		fmt.Fprintf(env.Out, "session: %s\npprof:   %s\n\nPress Ctrl+C to stop.\n", selection.HumanName, url)
		<-env.Ctx.Done()
		return nil
	}

	switch env.Format {
	case execenv.FormatRaw:
		// Raw() bypasses the color-profile writer: a pprof profile is binary
		// protobuf, and ANSI stripping on a non-terminal would corrupt it.
		return engine.ReadProfileRaw(env.Repo.Storage(), selection.Path, options.profileType, options.bench, env.Out.Raw())

	case execenv.FormatJSON:
		funcs, err := engine.ReadProfileFunctions(env.Repo.Storage(), selection.Path, options.profileType, options.bench)
		if err != nil {
			return err
		}
		type jsonFunc struct {
			Name    string  `json:"name"`
			File    string  `json:"file,omitempty"`
			Line    int64   `json:"line,omitempty"`
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
				File:    f.File,
				Line:    f.StartLine,
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
			prof:          prof,
			profileType:   options.profileType,
			metric:        options.metric,
			sort:          options.sort,
			top:           options.top,
			machineHeader: engine.FormatMachineLine(selection.Machine, selection.GoVersion),
			style:         env.Style,
			sourcesRoot:   env.Repo.Sources().Root(),
			sourceCache:   make(map[string][]byte),
		}

		profilePath, _ := engine.ProfileFilePath(env.Repo.Storage(), selection.Path, options.profileType, options.bench)
		if profilePath != "" {
			model.startPprof = func() (url string, err error) {
				return launchPprof(env.Ctx, profilePath)
			}
		}

		return env.ViewportWithKeys(model)()

	default:
		return fmt.Errorf("unsupported format %v for show %s (text, json, raw)", env.Format, engine.ProfileDir(options.profileType))
	}
}

func renderProfileTable(funcs []engine.ProfileFunc, top int, fmtValue func(time.Duration) string, order engine.SortOrder, style execenv.Style, getSource func(file string, startLine int64) (string, []string)) string {
	top = min(top, len(funcs))
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	flat, flatPct, cum, cumPct, name := "FLAT", "FLAT%", "CUM", "CUM%", "FUNCTION"
	switch order {
	case engine.SortFlat:
		flat = style.Accent(flat + " ▼")
		flatPct = style.Accent(flatPct + " ▼")
	case engine.SortCumulative:
		cum = style.Accent(cum + " ▼")
		cumPct = style.Accent(cumPct + " ▼")
	case engine.SortName:
		name = style.Accent(name + " ▼")
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

	if getSource == nil {
		return buf.String()
	}

	// Post-process: insert source annotation lines after each function row.
	tableLines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	// Determine indentation for the FUNCTION column from the header. The header
	// is styled, so the byte offset of "FUNCTION" counts the ANSI escapes too —
	// measure the display width of what precedes it instead.
	funcColStart := 0
	if i := strings.Index(tableLines[0], "FUNCTION"); i > 0 {
		funcColStart = lipgloss.Width(tableLines[0][:i])
	}
	indent := strings.Repeat(" ", funcColStart)

	var out strings.Builder
	out.WriteString(tableLines[0])
	out.WriteByte('\n')
	for i, f := range funcs[:top] {
		if i+1 < len(tableLines) {
			out.WriteString(tableLines[i+1])
			out.WriteByte('\n')
		}
		if f.File == "" || f.StartLine <= 0 {
			continue
		}
		path, srcLines := getSource(f.File, f.StartLine)
		if len(srcLines) == 0 {
			continue
		}
		fmt.Fprintf(&out, "%s%s:%s\n", indent, path, strconv.FormatInt(f.StartLine, 10))
		for _, l := range srcLines {
			fmt.Fprintf(&out, "%s│ %s\n", indent, l)
		}
	}
	return out.String()
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

// launchPprof starts `go tool pprof -http` for profilePath on a random free
// port and returns the URL once the process is running.
func launchPprof(ctx context.Context, profilePath string) (string, error) {
	port, err := findFreePort()
	if err != nil {
		return "", fmt.Errorf("no free port: %w", err)
	}
	addr := fmt.Sprintf("localhost:%d", port)
	pprofBin, err := execabs.Command("go", "tool", "-n", "pprof").Output()
	if err != nil {
		return "", fmt.Errorf("locate pprof: %w", err)
	}

	cmd := execabs.CommandContext(ctx, strings.TrimSpace(string(pprofBin)), "-http="+addr, profilePath)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	return "http://" + addr, cmd.Start()
}

// findFreePort returns an available TCP port on localhost.
func findFreePort() (int, error) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	l.Close()
	return port, nil
}
