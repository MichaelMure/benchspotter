package commands

import (
	"context"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
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
		newSessionUntagCommand(env),
		newSessionRenameCommand(env),
		newSessionRmCommand(env),
	)

	return cmd
}

func newSessionLsCommand(env *execenv.Env) *cobra.Command {
	var tag string
	cmd := &cobra.Command{
		Use:     "ls",
		Short:   "List benchmark sessions",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSessionLs(cmd.Context(), env, tag)
		},
	}
	cmd.Flags().StringVar(&tag, "tag", "", "Filter sessions by tag")
	return cmd
}

func runSessionLs(ctx context.Context, env *execenv.Env, tag string) error {
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

	if tag != "" {
		filtered := sessions[:0]
		for _, s := range sessions {
			for _, t := range s.Tags {
				if t == tag {
					filtered = append(filtered, s)
					break
				}
			}
		}
		sessions = filtered
	}

	switch env.Format {
	case execenv.FormatText:
		w := tabwriter.NewWriter(env.Out, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tTIME\tCOMMIT\tBENCH\tCPU\tMEM\tBLOCK\tMUTEX\tESCAPE\tBENCHMARKS\tTAGS")
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
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\n",
				s.HumanName,
				s.Time.Format("2006-01-02 15:04"),
				commit,
				check(s.HasBench()),
				check(s.HasProfile(engine.ProfileCPU)),
				check(s.HasProfile(engine.ProfileMem)),
				check(s.HasProfile(engine.ProfileBlock)),
				check(s.HasProfile(engine.ProfileMutex)),
				check(s.HasProfile(engine.ProfileEscape)),
				len(s.Benches),
				strings.Join(s.Tags, ", "),
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
			Tags       []string  `json:"tags"`
			Benchmarks []string  `json:"benchmarks"`
		}
		out := make([]sessionJSON, len(sessions))
		for i, s := range sessions {
			tags := s.Tags
			if tags == nil {
				tags = []string{}
			}
			out[i] = sessionJSON{
				ID:         s.Id,
				Name:       s.Name,
				HumanName:  s.HumanName,
				Time:       s.Time,
				Commit:     s.GitCommit,
				HasDiff:    s.HasGitDiff(),
				Profiles:   sessionProfiles(s),
				Tags:       tags,
				Benchmarks: s.Benches,
			}
		}
		return env.Out.PrintJSON(out)
	default:
		return fmt.Errorf("unsupported format %v for session ls (text, json)", env.Format)
	}
}

func sessionProfiles(s *engine.SessionInfo) []string {
	var out []string
	for _, p := range []engine.Profile{engine.ProfileCPU, engine.ProfileMem, engine.ProfileBlock, engine.ProfileMutex} {
		if s.HasProfile(p) {
			out = append(out, engine.ProfileDir(p))
		}
	}
	if s.HasProfile(engine.ProfileEscape) {
		out = append(out, "escape")
	}
	return out
}

func newSessionTagCommand(env *execenv.Env) *cobra.Command {
	return &cobra.Command{
		Use:     "tag [id] [tag]",
		Short:   "Tag a session",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSessionTag(cmd.Context(), env, args)
		},
	}
}

func runSessionTag(ctx context.Context, env *execenv.Env, args []string) error {
	var sessionID, tag string

	if len(args) >= 2 {
		sessionID, tag = args[0], args[1]
	} else {
		selection, err := inputs.SelectSession(ctx, env, env.Repo.GetRecall("session_tag_session"), nil)
		if err != nil {
			return err
		}
		if err = env.Repo.SetRecall("session_tag_session", selection.Id); err != nil {
			return err
		}
		sessionID = selection.Id

		if len(args) == 1 {
			tag = args[0]
		} else {
			err = env.FormSingle(huh.NewInput().
				Title("Tag name").
				Value(&tag)).
				RunWithContext(ctx)
			if err != nil {
				return err
			}
		}
	}

	if tag == "" {
		return fmt.Errorf("tag name cannot be empty")
	}

	sessions, err := engine.LocateSessions(env.Repo.Storage())
	if err != nil {
		return err
	}
	var target *engine.SessionInfo
	for _, s := range sessions {
		if s.Id == sessionID {
			target = s
			break
		}
	}
	if target == nil {
		return fmt.Errorf("session %q not found", sessionID)
	}

	return engine.TagSession(env.Repo.Storage(), target.Path, tag)
}

func newSessionUntagCommand(env *execenv.Env) *cobra.Command {
	return &cobra.Command{
		Use:     "untag [id] [tag]",
		Short:   "Remove a tag from a session",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSessionUntag(cmd.Context(), env, args)
		},
	}
}

func runSessionUntag(ctx context.Context, env *execenv.Env, args []string) error {
	var sessionID, tag string

	if len(args) >= 2 {
		sessionID, tag = args[0], args[1]
	} else {
		selection, err := inputs.SelectSession(ctx, env, env.Repo.GetRecall("session_untag_session"),
			func(s *engine.SessionInfo) bool { return len(s.Tags) > 0 })
		if err != nil {
			return err
		}
		if err = env.Repo.SetRecall("session_untag_session", selection.Id); err != nil {
			return err
		}
		sessionID = selection.Id

		if len(args) == 1 {
			tag = args[0]
		} else {
			opts := make([]huh.Option[string], len(selection.Tags))
			for i, t := range selection.Tags {
				opts[i] = huh.NewOption(t, t)
			}
			err = env.FormSingle(huh.NewSelect[string]().
				Title("Tag to remove").
				Options(opts...).
				Value(&tag)).
				RunWithContext(ctx)
			if err != nil {
				return err
			}
		}
	}

	sessions, err := engine.LocateSessions(env.Repo.Storage())
	if err != nil {
		return err
	}
	for _, s := range sessions {
		if s.Id == sessionID {
			return engine.UntagSession(env.Repo.Storage(), s.Path, tag)
		}
	}
	return fmt.Errorf("session %q not found", sessionID)
}

func newSessionRenameCommand(env *execenv.Env) *cobra.Command {
	return &cobra.Command{
		Use:     "rename [id] [name]",
		Short:   "Rename a session",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSessionRename(cmd.Context(), env, args)
		},
	}
}

func runSessionRename(ctx context.Context, env *execenv.Env, args []string) error {
	var sessionID, name string

	if len(args) >= 2 {
		sessionID, name = args[0], args[1]
	} else {
		selection, err := inputs.SelectSession(ctx, env, env.Repo.GetRecall("session_rename_session"), nil)
		if err != nil {
			return err
		}
		if err = env.Repo.SetRecall("session_rename_session", selection.Id); err != nil {
			return err
		}
		sessionID = selection.Id

		if len(args) == 1 {
			name = args[0]
		} else {
			name = selection.Name
			err = env.FormSingle(huh.NewInput().
				Title("New name").
				Value(&name)).
				RunWithContext(ctx)
			if err != nil {
				return err
			}
		}
	}

	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	sessions, err := engine.LocateSessions(env.Repo.Storage())
	if err != nil {
		return err
	}
	for _, s := range sessions {
		if s.Id == sessionID {
			return engine.RenameSession(env.Repo.Storage(), s.Path, name)
		}
	}
	return fmt.Errorf("session %q not found", sessionID)
}

func newSessionRmCommand(env *execenv.Env) *cobra.Command {
	var yes bool
	cmd := &cobra.Command{
		Use:     "rm [id]",
		Short:   "Delete a session",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSessionRm(cmd.Context(), env, args, yes)
		},
	}
	cmd.Flags().BoolVarP(&yes, "yes", "y", false, "Skip confirmation")
	return cmd
}

func runSessionRm(ctx context.Context, env *execenv.Env, args []string, yes bool) error {
	var target *engine.SessionInfo

	if len(args) >= 1 {
		sessions, err := engine.LocateSessions(env.Repo.Storage())
		if err != nil {
			return err
		}
		for _, s := range sessions {
			if s.Id == args[0] {
				target = s
				break
			}
		}
		if target == nil {
			return fmt.Errorf("session %q not found", args[0])
		}
	} else {
		var err error
		target, err = inputs.SelectSession(ctx, env, env.Repo.GetRecall("session_rm_session"), nil)
		if err != nil {
			return err
		}
	}

	if !yes {
		var confirmed bool
		err := env.FormSingle(huh.NewConfirm().
			Title(fmt.Sprintf("Delete session %q?", target.HumanName)).
			Value(&confirmed)).
			RunWithContext(ctx)
		if err != nil {
			return err
		}
		if !confirmed {
			return nil
		}
	}

	return engine.RemoveSession(env.Repo.Storage(), target.Path)
}
