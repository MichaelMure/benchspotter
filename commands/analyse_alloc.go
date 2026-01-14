package commands

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

type analyseAllocOptions struct {
	session string
	sample  engine.Sample
}

var allocSampleIds = map[engine.Sample][]string{
	engine.SampleInuseSpace:   {"inuse_space"},
	engine.SampleInuseObjects: {"inuse_objects"},
	engine.SampleAllocSpace:   {"alloc_space"},
	engine.SampleAllocObjects: {"alloc_objects"},
}

var allocSampleHelp = map[engine.Sample]string{
	engine.SampleInuseSpace:   "Memory in use (space)",
	engine.SampleInuseObjects: "Memory in use (objects)",
	engine.SampleAllocSpace:   "Memory allocated (space)",
	engine.SampleAllocObjects: "Memory allocated (objects)",
}

func newAnalyseAllocCommand(env *execenv.Env) *cobra.Command {
	options := analyseAllocOptions{}

	cmd := &cobra.Command{
		Use:     "alloc",
		Short:   "Analyse memory allocations",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAnalyseAlloc(cmd.Context(), env, options)
		},
	}

	flags := cmd.Flags()

	sampleEnum := enumflag.New(&options.sample, "sample", allocSampleIds, enumflag.EnumCaseInsensitive)
	flags.VarP(sampleEnum, "sample", "s", "Memory sample to analyse")
	err := sampleEnum.RegisterCompletion(cmd, "sample", allocSampleHelp)
	if err != nil {
		panic(err)
	}

	return cmd
}

func runAnalyseAlloc(ctx context.Context, env *execenv.Env, options analyseAllocOptions) error {
	if len(options.session) == 0 {
		const recallKey = "analyse_alloc_session"
		preSelected := env.Repo.GetRecall(recallKey)

		selection, err := inputs.SelectSession(ctx, env, preSelected, nil)
		if err != nil {
			return err
		}
		if selection == nil {
			return fmt.Errorf("no session selected")
		}

		err = env.Repo.SetRecall(recallKey, selection.Id)
		if err != nil {
			return err
		}
	}

	viewport, runFn := env.Viewport(ctx)

	viewport.SetContent("dfshkjl")

	return runFn()
}
