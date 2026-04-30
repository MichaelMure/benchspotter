package inputs

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"charm.land/huh/v2"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
)

func SelectSession(env *execenv.Env, preSelect string, filter func(info *engine.SessionInfo) bool) (*engine.SessionInfo, error) {
	var sessions []*engine.SessionInfo

	err := env.Spinner().Title("Finding sessions").
		ActionWithErr(func(ctx context.Context) error {
			var err error
			sessions, err = engine.LocateSessions(env.Repo.Storage())
			return err
		}).Context(env.Ctx).Run()
	if err != nil {
		return nil, fmt.Errorf("failed to discover sessions: %w", err)
	}

	if len(sessions) == 0 {
		return nil, fmt.Errorf("no sessions to select, use `benchspotter bench` to run benchmarks")
	}

	var visible []*engine.SessionInfo
	for _, s := range sessions {
		if filter == nil || filter(s) {
			visible = append(visible, s)
		}
	}

	mc := engine.NewMachineContext(visible)

	var selection *engine.SessionInfo
	fields := []huh.Field{
		huh.NewSelect[*engine.SessionInfo]().
			Title("Select a session").
			OptionsFunc(func() []huh.Option[*engine.SessionInfo] {
				opts := make([]huh.Option[*engine.SessionInfo], 0, len(sessions))
				for _, session := range sessions {
					if filter != nil && !filter(session) {
						continue
					}
					line := formatSession(env, session, mc)
					opts = append(opts, huh.NewOption(line, session).
						Selected(preSelect == session.Id))
				}
				return opts
			}, nil).
			Value(&selection),
	}
	if mc != nil {
		fields = append([]huh.Field{machineLegendNote(env, mc)}, fields...)
	}
	err = env.Form(huh.NewGroup(fields...)).
		WithShowHelp(false).
		RunWithContext(env.Ctx)
	if err != nil {
		return nil, err
	}

	return selection, nil
}

func SelectSessions(env *execenv.Env, preSelect []string) ([]*engine.SessionInfo, error) {
	var sessions []*engine.SessionInfo

	err := env.Spinner().Title("Finding sessions").
		ActionWithErr(func(ctx context.Context) error {
			var err error
			sessions, err = engine.LocateSessions(env.Repo.Storage())
			return err
		}).Context(env.Ctx).Run()
	if err != nil {
		return nil, fmt.Errorf("failed to discover sessions: %w", err)
	}

	if len(sessions) == 0 {
		return nil, fmt.Errorf("no sessions to select, use `benchspotter bench` to run benchmarks")
	}

	mc := engine.NewMachineContext(sessions)

	var selection []*engine.SessionInfo
	fields := []huh.Field{
		huh.NewMultiSelect[*engine.SessionInfo]().
			Title("Select sessions (/ to filter)").
			OptionsFunc(func() []huh.Option[*engine.SessionInfo] {
				opts := make([]huh.Option[*engine.SessionInfo], len(sessions))
				for i, session := range sessions {
					line := formatSession(env, session, mc)
					opts[i] = huh.NewOption(line, session).
						Selected(slices.Contains(preSelect, session.Id))
				}
				return opts
			}, nil).
			Filterable(true).
			Value(&selection),
	}
	if mc != nil {
		fields = append([]huh.Field{machineLegendNote(env, mc)}, fields...)
	}
	err = env.Form(huh.NewGroup(fields...)).
		WithShowHelp(false).
		RunWithContext(env.Ctx)
	if err != nil {
		return nil, err
	}

	return selection, nil
}

func machineLabel(style execenv.Style, n int) string {
	return style.Info(fmt.Sprintf("⚙%d", n))
}

func formatSession(env *execenv.Env, session *engine.SessionInfo, mc *engine.MachineContext) string {
	var label string
	if mc != nil {
		if n := mc.Labels[session.Id]; n != 0 {
			label = fmt.Sprintf(" ⚙%d", n)
		}
	}
	var commit string
	if len(session.GitCommit) > 0 {
		if session.HasGitDiff() {
			commit = fmt.Sprintf(" - (⎇  %s ±)", session.GitCommit[:7])
		} else {
			commit = fmt.Sprintf(" - (⎇  %s)", session.GitCommit[:7])
		}
	}
	return fmt.Sprintf("%s%s%s%s",
		session.HumanName,
		label,
		commit,
		env.Style.TonedDown(" - "+session.Id),
	)
}

func machineLegendNote(env *execenv.Env, mc *engine.MachineContext) *huh.Note {
	var body strings.Builder
	for i, e := range mc.Entries {
		if i > 0 {
			body.WriteByte('\n')
		}
		_, _ = fmt.Fprintf(&body, "  %s  %s", machineLabel(env.Style, i+1), engine.FormatMachineLine(&e.Machine, e.GoVersion))
	}
	return huh.NewNote().Title("Machines:").Description(body.String())
}
