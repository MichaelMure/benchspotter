package commands

import (
	"context"
	"fmt"
	"strings"
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
		fmt.Fprintln(w, "NAME\tTIME\tCOMMIT\tBENCHMARKS")
		for _, s := range sessions {
			commit := ""
			if s.GitCommit != "" {
				commit = s.GitCommit[:7]
				if s.HasGitDiff() {
					commit += "±"
				}
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
				s.HumanName,
				s.Time.Format("2006-01-02 15:04"),
				commit,
				strings.Join(s.Benches, ", "),
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
				Benchmarks: s.Benches,
			}
		}
		return env.Out.PrintJSON(out)
	default:
		return fmt.Errorf("unsupported format %v for session ls (text, json)", env.Format)
	}
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
