package execenv

import (
	"github.com/spf13/cobra"

	"benchspotter/repository"
)

// LoadRepo is a pre-run function that load the repository for use in a command
func LoadRepo(env *Env) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		repo, err := repository.AutoDetect()
		if err != nil {
			return err
		}
		env.Repo = repo
		return nil
	}
}
