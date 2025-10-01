package inputs

import (
	"context"
	"fmt"
	"slices"

	"github.com/charmbracelet/huh"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
)

func SelectSessions(ctx context.Context, env *execenv.Env, preSelect []string) ([]engine.SessionInfo, error) {
	var sessions []engine.SessionInfo

	err := env.Spinner().Title("Finding sessions").
		ActionWithErr(func(ctx context.Context) error {
			var err error
			sessions, err = engine.LocateSessions(env.Repo.Storage())
			return err
		}).Context(ctx).Run()
	if err != nil {
		return nil, fmt.Errorf("failed to discover sessions: %w", err)
	}

	if len(sessions) == 0 {
		return nil, fmt.Errorf("no sessions to select, use `benchspotter bench` to run benchmarks")
	}

	var selection []engine.SessionInfo

	err = env.FormSingle(huh.NewMultiSelect[engine.SessionInfo]().
		Title("Select session").
		OptionsFunc(func() []huh.Option[engine.SessionInfo] {
			opts := make([]huh.Option[engine.SessionInfo], len(sessions))
			for i, session := range sessions {
				opts[i] = huh.NewOption(session.Id, session).
					Selected(slices.Contains(preSelect, session.Id))
			}
			return opts
		}, nil).
		Value(&selection)).
		RunWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return selection, nil
}
