package inputs

import (
	"charm.land/huh/v2"

	"benchspotter/commands/execenv"
)

// BenchOption is a benchmark entry for display in the profile bench selector.
type BenchOption struct {
	Name  string // returned as the selection
	Label string // displayed in the prompt
}

// SelectProfileBench prompts the user to choose a benchmark from the available profiles.
func SelectProfileBench(env *execenv.Env, opts []BenchOption, preSelect string) (string, error) {
	hopts := make([]huh.Option[string], len(opts))
	for i, o := range opts {
		hopts[i] = huh.NewOption(o.Label, o.Name).Selected(o.Name == preSelect)
	}

	var selection string
	err := env.FormSingle(huh.NewSelect[string]().
		Title("Select a benchmark").
		Options(hopts...).
		Value(&selection)).
		RunWithContext(env.Ctx)
	if err != nil {
		return "", err
	}
	return selection, nil
}
