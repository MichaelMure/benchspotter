package inputs

import (
	"context"

	"github.com/charmbracelet/huh"

	"benchspotter/commands/execenv"
)

// SelectProfileBench prompts the user to choose a benchmark from the available profiles.
func SelectProfileBench(ctx context.Context, env *execenv.Env, benches []string, preSelect string) (string, error) {
	opts := make([]huh.Option[string], len(benches))
	for i, b := range benches {
		opts[i] = huh.NewOption(b, b).Selected(b == preSelect)
	}

	var selection string
	err := env.FormSingle(huh.NewSelect[string]().
		Title("Select benchmark").
		Options(opts...).
		Value(&selection)).
		RunWithContext(ctx)
	if err != nil {
		return "", err
	}
	return selection, nil
}
