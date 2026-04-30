package commands

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"charm.land/huh/v2"
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/commands/tabwriter"
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

type sessionLsOptions struct {
	tag string
}

func newSessionLsCommand(env *execenv.Env) *cobra.Command {
	var opts sessionLsOptions
	cmd := &cobra.Command{
		Use:     "ls",
		Short:   "List benchmark sessions",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSessionLs(env, opts)
		},
	}
	cmd.Flags().StringVar(&opts.tag, "tag", "", "Filter sessions by tag")
	return cmd
}

func runSessionLs(env *execenv.Env, opts sessionLsOptions) error {
	var sessions []*engine.SessionInfo

	err := env.Spinner().Title("Finding sessions").
		ActionWithErr(func(ctx context.Context) error {
			var err error
			// HACK: bubbletea or huh suffers at the moment from an issue where, if the app runs
			// and stops too fast, a race can happen where bubbletea query the terminal for info,
			// the app stops, then the terminal print the response visibly to stdin as garbage.
			// A small delay fix that.
			time.Sleep(100 * time.Millisecond)
			sessions, err = engine.LocateSessions(env.Repo.Storage())
			return err
		}).Context(env.Ctx).Run()
	if err != nil {
		return err
	}

	if opts.tag != "" {
		filtered := sessions[:0]
		for _, s := range sessions {
			if slices.Contains(s.Tags, opts.tag) {
				filtered = append(filtered, s)
			}
		}
		sessions = filtered
	}

	mc := engine.NewMachineContext(sessions)

	switch env.Format {
	case execenv.FormatText:
		w := tabwriter.NewWriter(env.Out, 0, 0, 2, ' ', 0)
		header := "NAME\tTIME\tCOMMIT\tBENCH\tCPU\tMEM\tBLOCK\tMUTEX\tESCAPE\tINLINE\tBENCHMARKS\tTAGS"
		if mc != nil {
			header += "\tMACHINE"
		}
		_, _ = fmt.Fprintln(w, header)
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
					return env.Style.Positive("✓")
				}
				return env.Style.Negative("-")
			}
			row := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s",
				s.HumanName,
				s.Time.Format("2006-01-02 15:04"),
				commit,
				check(s.HasBench()),
				check(s.HasProfile(engine.ProfileCPU)),
				check(s.HasProfile(engine.ProfileMem)),
				check(s.HasProfile(engine.ProfileBlock)),
				check(s.HasProfile(engine.ProfileMutex)),
				check(s.HasProfile(engine.ProfileEscape)),
				check(s.HasProfile(engine.ProfileInline)),
				len(s.Benches),
				strings.Join(s.Tags, ", "),
			)
			if mc != nil {
				if n := mc.Labels[s.Id]; n != 0 {
					row += fmt.Sprintf("\t⚙%d", n)
				} else {
					row += "\t"
				}
			}
			_, _ = fmt.Fprintln(w, row)
		}
		if err := w.Flush(); err != nil {
			return err
		}
		if mc != nil {
			_, _ = fmt.Fprintln(env.Out)
			for i, e := range mc.Entries {
				_, _ = fmt.Fprintf(env.Out, "  ⚙%d  %s\n", i+1, engine.FormatMachineLine(&e.Machine, e.GoVersion))
			}
		}
		return nil
	case execenv.FormatJSON:
		type sessionJSON struct {
			ID         string              `json:"id"`
			Name       string              `json:"name,omitempty"`
			HumanName  string              `json:"human_name"`
			Time       time.Time           `json:"time"`
			Commit     string              `json:"commit,omitempty"`
			HasDiff    bool                `json:"has_diff"`
			Profiles   []string            `json:"profiles"`
			Tags       []string            `json:"tags"`
			Benchmarks []string            `json:"benchmarks"`
			Machine    *engine.MachineInfo `json:"machine,omitempty"`
			GoVersion  string              `json:"go_version,omitempty"`
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
				Machine:    s.Machine,
				GoVersion:  s.GoVersion,
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
	if s.HasProfile(engine.ProfileInline) {
		out = append(out, "inline")
	}
	return out
}

func newSessionTagCommand(env *execenv.Env) *cobra.Command {
	return &cobra.Command{
		Use:     "tag [id] [tag]",
		Short:   "Tag a session",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSessionTag(env, args)
		},
	}
}

func runSessionTag(env *execenv.Env, args []string) error {
	var sessionID, tag string

	if len(args) >= 2 {
		sessionID, tag = args[0], args[1]
	} else {
		selection, err := inputs.SelectSession(env, env.Repo.GetRecall("session_tag_session"), nil)
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
				RunWithContext(env.Ctx)
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
			return runSessionUntag(env, args)
		},
	}
}

func runSessionUntag(env *execenv.Env, args []string) error {
	var sessionID, tag string

	if len(args) >= 2 {
		sessionID, tag = args[0], args[1]
	} else {
		selection, err := inputs.SelectSession(env, env.Repo.GetRecall("session_untag_session"),
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
				RunWithContext(env.Ctx)
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
			return runSessionRename(env, args)
		},
	}
}

func runSessionRename(env *execenv.Env, args []string) error {
	var sessionID, name string

	if len(args) >= 2 {
		sessionID, name = args[0], args[1]
	} else {
		selection, err := inputs.SelectSession(env, env.Repo.GetRecall("session_rename_session"), nil)
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
				RunWithContext(env.Ctx)
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

type sessionRmOptions struct {
	skipConfirmation bool
}

func newSessionRmCommand(env *execenv.Env) *cobra.Command {
	var opts sessionRmOptions
	cmd := &cobra.Command{
		Use:     "rm [id]",
		Short:   "Delete a session",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSessionRm(env, args, opts)
		},
	}
	cmd.Flags().BoolVarP(&opts.skipConfirmation, "skip-confirmation", "y", false, "Skip confirmation")
	return cmd
}

func runSessionRm(env *execenv.Env, args []string, opts sessionRmOptions) error {
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
		target, err = inputs.SelectSession(env, env.Repo.GetRecall("session_rm_session"), nil)
		if err != nil {
			return err
		}
	}

	if !opts.skipConfirmation {
		var confirmed bool
		err := env.FormSingle(huh.NewConfirm().
			Title(fmt.Sprintf("Delete session %q?", target.HumanName)).
			Value(&confirmed)).
			RunWithContext(env.Ctx)
		if err != nil {
			return err
		}
		if !confirmed {
			return nil
		}
	}

	return engine.RemoveSession(env.Repo.Storage(), target.Path)
}
