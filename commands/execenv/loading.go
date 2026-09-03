package execenv

import (
	"github.com/spf13/cobra"

	"github.com/MichaelMure/benchspotter/repository"
)

// LoadRepo is a pre-run function that load the repository for use in a command
func LoadRepo(env *Env) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var repo *repository.Repository
		var err error

		if env.Dir != "" {
			repo, err = repository.Open(env.Dir)
		} else {
			repo, err = repository.AutoDetect()
		}
		if err != nil {
			return err
		}
		env.Repo = repo
		return nil
	}
}
