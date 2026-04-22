package execenv

import (
	"context"
	"fmt"
	"io"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/muesli/termenv"

	"benchspotter/repository"
)

// Env is the environment of a command
type Env struct {
	Repo   *repository.Repository
	In     In
	Out    Out
	Err    Out
	Style  Style
	Format Format
}

func NewEnv() *Env {
	return &Env{
		Repo:   nil,
		In:     in{Reader: os.Stdin},
		Out:    out{out: termenv.NewOutput(os.Stdout)},
		Err:    out{out: termenv.NewOutput(os.Stderr)},
		Style:  Style{huh.ThemeCharm()},
		Format: FormatText,
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

// RunModel runs m as a full-screen TUI. It is the caller's responsibility to
// handle the non-terminal case before calling this.
func (e Env) RunModel(ctx context.Context, m tea.Model) error {
	_, err := e.newProgram(ctx, m).Run()
	return err
}

func (e Env) Viewport(ctx context.Context) (ViewportWriter, func() error) {
	if !e.Out.IsTerminal() {
		p := &plainViewport{}
		return p, func() error {
			_, err := io.Copy(e.Out, &p.Buffer)
			return err
		}
	}

	v := &viewportModel{}
	v.program = e.newProgram(ctx, v)

	return v, func() error {
		_, err := v.program.Run()
		return err
	}
}

// ViewportWithKeys runs m as an interactive TUI viewport. In non-terminal mode
// it renders once and writes directly to Out. Flags should set the initial
// model state for headless use.
func (e Env) ViewportWithKeys(ctx context.Context, m InteractiveModel) func() error {
	if !e.Out.IsTerminal() {
		return func() error {
			_, err := fmt.Fprint(e.Out, m.Render())
			return err
		}
	}
	return func() error {
		return e.RunModel(ctx, &interactiveViewportModel{model: m})
	}
}

func (e Env) newProgram(ctx context.Context, m tea.Model) *tea.Program {
	return tea.NewProgram(m,
		tea.WithContext(ctx),
		tea.WithInput(e.In.Raw()),
		tea.WithOutput(e.Out.Raw()),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
}
