package commands

import (
	"context"
	"fmt"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
)

func newSessionCommand(env *execenv.Env) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Manage benchmark sessions",
	}

	cmd.AddCommand(
		newSessionLsCommand(env),
		newSessionTagCommand(env),
		newSessionRmCommand(env),
	)

	return cmd
}

func newSessionLsCommand(env *execenv.Env) *cobra.Command {
	return &cobra.Command{
		Use:     "ls",
		Short:   "List benchmark sessions",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSessionLs(cmd.Context(), env)
		},
	}
}

func runSessionLs(ctx context.Context, env *execenv.Env) error {
	var sessions []*engine.SessionInfo

	err := env.Spinner().Title("Finding sessions").
		ActionWithErr(func(ctx context.Context) error {
			var err error
			sessions, err = engine.LocateSessions(env.Repo.Storage())
			return err
		}).Context(ctx).Run()
	if err != nil {
		return err
	}

	switch env.Format {
	case execenv.FormatText:
		w := tabwriter.NewWriter(env.Out, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tTIME\tCOMMIT\tBENCH\tCPU\tMEM\tBLOCK\tMUTEX\tBENCHMARKS")
		for _, s := range sessions {
			commit := ""
			if s.GitCommit != "" {
				commit = s.GitCommit[:7]
				if s.HasGitDiff() {
					commit += "±"
				}
			}
			check := func(ok bool) string {
				if ok {
					return "✓"
				}
				return "-"
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\n",
				s.HumanName,
				s.Time.Format("2006-01-02 15:04"),
				commit,
				check(s.HasBench()),
				check(s.HasProfile(engine.ProfileCPU)),
				check(s.HasProfile(engine.ProfileMem)),
				check(s.HasProfile(engine.ProfileBlock)),
				check(s.HasProfile(engine.ProfileMutex)),
				len(s.Benches),
			)
		}
		return w.Flush()
	case execenv.FormatJSON:
		type sessionJSON struct {
			ID         string    `json:"id"`
			Name       string    `json:"name,omitempty"`
			HumanName  string    `json:"human_name"`
			Time       time.Time `json:"time"`
			Commit     string    `json:"commit,omitempty"`
			HasDiff    bool      `json:"has_diff"`
			Profiles   []string  `json:"profiles"`
			Benchmarks []string  `json:"benchmarks"`
		}
		out := make([]sessionJSON, len(sessions))
		for i, s := range sessions {
			out[i] = sessionJSON{
				ID:         s.Id,
				Name:       s.Name,
				HumanName:  s.HumanName,
				Time:       s.Time,
				Commit:     s.GitCommit,
				HasDiff:    s.HasGitDiff(),
				Profiles:   sessionProfiles(s),
				Benchmarks: s.Benches,
			}
		}
		return env.Out.PrintJSON(out)
	default:
		return fmt.Errorf("unsupported format %v for session ls (text, json)", env.Format)
	}
}

func sessionProfiles(s *engine.SessionInfo) []string {
	all := []engine.Profile{engine.ProfileCPU, engine.ProfileMem, engine.ProfileBlock, engine.ProfileMutex}
	var out []string
	for _, p := range all {
		if s.HasProfile(p) {
			out = append(out, engine.ProfileDir(p))
		}
	}
	return out
}

func newSessionTagCommand(env *execenv.Env) *cobra.Command {
	return &cobra.Command{
		Use:   "tag <id> <tag>",
		Short: "Tag a session",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("not yet implemented")
		},
	}
}

func newSessionRmCommand(env *execenv.Env) *cobra.Command {
	return &cobra.Command{
		Use:   "rm <id>",
		Short: "Delete a session",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("not yet implemented")
		},
	}
}
