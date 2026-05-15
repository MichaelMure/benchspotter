package commands

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"

	"benchspotter/commands/execenv"
)

const RootCmdName = "benchspotter"

func NewRootCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   RootCmdName,
		Short: "Your companion for go benchmarking",
		Long: RootCmdName + ` is a companion tool to handle the logistics of your go benchmarking, and guide you through powerful analysis.

A typical session is as follow:
 1. Select and run benchmarks. They get recorded on disk for later analysis.
 2. Analyse, compare, drill down with a collection of tools.
 3. Optionally, let ` + RootCmdName + ` optimize parameters.`,

		SilenceUsage:      true,
		DisableAutoGenTag: true,
	}

	const mainGroup = "main"
	const utilitiesGroup = "utilities"

	cmd.AddGroup(&cobra.Group{ID: mainGroup, Title: "Main commands"})
	cmd.AddGroup(&cobra.Group{ID: utilitiesGroup, Title: "Utilities"})

	addCmdWithGroup := func(child *cobra.Command, groupID string) {
		cmd.AddCommand(child)
		child.GroupID = groupID
	}

	if v, err := getVersion(); err == nil {
		cmd.Version = v
	}

	env := execenv.NewEnv(ctx)

	formatFlag := enumflag.New(&env.Format, "format", execenv.FormatIds, enumflag.EnumCaseInsensitive)
	cmd.PersistentFlags().VarP(formatFlag, "format", "f", "output format (not all formats supported by every command)")
	_ = formatFlag.RegisterCompletion(cmd, "format", execenv.FormatHelp)

	cmd.PersistentFlags().BoolVar(&env.NoPrompt, "no-prompt", false, "disable interactive prompts; all required values must be supplied via flags")

	addCmdWithGroup(newBenchCommand(env), mainGroup)
	addCmdWithGroup(newSessionCommand(env), mainGroup)
	addCmdWithGroup(newShowCommand(env), mainGroup)
	addCmdWithGroup(newCompareCommand(env), mainGroup)
	addCmdWithGroup(newTrendCommand(env), mainGroup)
	addCmdWithGroup(newOptimizeCommand(env), mainGroup)

	addCmdWithGroup(newVersionCommand(env), utilitiesGroup)
	cmd.SetHelpCommandGroupID(utilitiesGroup)
	cmd.SetCompletionCommandGroupID(utilitiesGroup)

	return cmd
}
