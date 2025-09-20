package commands

import (
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
)

const RootCmdName = "benchspotter"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   RootCmdName,
		Short: "", // TODO
		Long:  ``, // TODO

		SilenceUsage:      true,
		DisableAutoGenTag: true,
	}

	env := execenv.NewEnv()

	cmd.AddCommand(newBenchCommand(env))
	cmd.AddCommand(newVersionCommand(env))

	return cmd
}
