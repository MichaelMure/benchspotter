package execenv

import (
	"context"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"golang.org/x/sys/execabs"

	"benchspotter/repository"
)

// Env is the environment of a command
type Env struct {
	Repo *repository.Repository
	In   In
	Out  Out
	Err  Out
}

func NewEnv() *Env {
	return &Env{
		Repo: nil,
		In:   in{Reader: os.Stdin},
		Out:  out{Writer: os.Stdout},
		Err:  out{Writer: os.Stderr},
	}
}

func (e Env) Form(groups ...*huh.Group) *huh.Form {
	return huh.NewForm(groups...).
		WithInput(e.In.Raw()).
		WithOutput(e.Out.Raw())
}

func (e Env) FormSingle(field huh.Field) *huh.Form {
	group := huh.NewGroup(field)
	return e.Form(group).WithShowHelp(false)
}

func (e Env) Spinner() *spinner.Spinner {
	return spinner.New().Output(e.Out.Raw())
}

func (e Env) ExecGo(ctx context.Context, args ...string) *exec.Cmd {
	cmd := execabs.CommandContext(ctx, "go", args...)
	// cmd.Stdout = e.Out.Raw()
	// cmd.Stderr = e.Err.Raw()
	// cmd.Stdin = e.In.Raw()
	return cmd
}
