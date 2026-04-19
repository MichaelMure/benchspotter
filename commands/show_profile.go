package commands

import (
	"context"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

type showProfileOptions struct {
	session     string
	bench       string
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
			preSelected := env.Repo.GetRecall(benchRecallKey)
			options.bench, err = inputs.SelectProfileBench(ctx, env, benches, preSelected)
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
		funcs, err := engine.ReadProfileFunctions(env.Repo.Storage(), selection.Path, options.profileType, options.bench)
		if err != nil {
			return err
		}
		top := min(options.top, len(funcs))
		viewport, runFn := env.Viewport(ctx)
		w := tabwriter.NewWriter(viewport, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "FLAT\tFLAT%\tCUM\tCUM%\tFUNCTION")
		for _, f := range funcs[:top] {
			fmt.Fprintf(w, "%s\t%.2f%%\t%s\t%.2f%%\t%s\n",
				formatDuration(f.Flat, options.profileType),
				f.FlatPct,
				formatDuration(f.Cumulative, options.profileType),
				f.CumPct,
				f.Name,
			)
		}
		if err := w.Flush(); err != nil {
			return err
		}
		return runFn()

	default:
		return fmt.Errorf("unsupported format %v for show %s (text, json, raw)", env.Format, engine.ProfileDir(options.profileType))
	}
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
