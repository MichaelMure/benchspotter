package execenv

import (
	"context"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"

	"benchspotter/repository"
)

// Env is the environment of a command
type Env struct {
	Repo  *repository.Repository
	In    In
	Out   Out
	Err   Out
	Style Style
}

func NewEnv() *Env {
	return &Env{
		Repo:  nil,
		In:    in{Reader: os.Stdin},
		Out:   out{Writer: os.Stdout},
		Err:   out{Writer: os.Stderr},
		Style: Style{huh.ThemeCharm()},
	}
}

func (e Env) Form(groups ...*huh.Group) *huh.Form {
	return huh.NewForm(groups...).
		WithInput(e.In.Raw()).
		WithOutput(e.Out.Raw()).
		WithTheme(e.Style.Theme)
}

func (e Env) FormSingle(field huh.Field) *huh.Form {
	group := huh.NewGroup(field)
	return e.Form(group).WithShowHelp(false)
}

func (e Env) Spinner() *spinner.Spinner {
	return spinner.New().Output(e.Out.Raw())
}

func (e Env) Viewport(ctx context.Context) (ViewportWriter, func() error) {
	v := &viewportModel{}

	v.program = tea.NewProgram(v,
		tea.WithContext(ctx),
		tea.WithInput(e.In.Raw()),
		tea.WithOutput(e.Out.Raw()),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	return v, func() error {
		_, err := v.program.Run()
		return err
	}
}

type Style struct {
	*huh.Theme
}

func (s Style) TonedDown(strs ...string) string {
	return s.Theme.Focused.Description.Render(strs...)
}
