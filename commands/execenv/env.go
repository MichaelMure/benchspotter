package execenv

import (
	"context"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/muesli/termenv"

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
		Out:   out{out: termenv.NewOutput(os.Stdout)},
		Err:   out{out: termenv.NewOutput(os.Stderr)},
		Style: Style{huh.ThemeCharm()},
	}
}

// Form creates a huh form with multiple widgets
func (e Env) Form(groups ...*huh.Group) *huh.Form {
	return huh.NewForm(groups...).
		WithInput(e.In.Raw()).
		WithOutput(e.Out.Raw()).
		WithTheme(e.Style.Theme)
}

// FormSingle creates a huh form with a single widget
func (e Env) FormSingle(field huh.Field) *huh.Form {
	group := huh.NewGroup(field)
	return e.Form(group).WithShowHelp(false)
}

// Spinner creates a new spinner for the Env configuration
func (e Env) Spinner() *spinner.Spinner {
	return spinner.New().Output(e.Out.Raw())
}

// Viewport creates a new viewport for the Env configuration
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
