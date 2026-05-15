package execenv

import (
	"context"
	"fmt"
	"io"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"charm.land/huh/v2/spinner"
	"charm.land/lipgloss/v2"
	"github.com/muesli/termenv"

	"benchspotter/repository"
)

// Env is the environment of a command
type Env struct {
	Ctx      context.Context
	Repo     *repository.Repository
	In       In
	Out      Out
	Err      Out
	Style    Style
	Format   Format
	NoPrompt bool
	Dir      string
}

func NewEnv(ctx context.Context) *Env {
	tf := huh.ThemeFunc(huh.ThemeCharm)
	isDark := lipgloss.HasDarkBackground(os.Stdin, os.Stdout)
	return &Env{
		Ctx:    ctx,
		Repo:   nil,
		In:     in{Reader: os.Stdin},
		Out:    out{out: termenv.NewOutput(os.Stdout)},
		Err:    out{out: termenv.NewOutput(os.Stderr)},
		Style:  NewStyle(tf, isDark),
		Format: FormatText,
	}
}

func (e Env) Form(flag string, groups ...*huh.Group) Form {
	if e.NoPrompt {
		return &errForm{flag: flag}
	}
	return huh.NewForm(groups...).
		WithInput(e.In.Raw()).
		WithOutput(e.Out.Raw()).
		WithTheme(e.Style.themeFunc)
}

func (e Env) FormSingle(flag string, field huh.Field) Form {
	if e.NoPrompt {
		return &errForm{flag: flag}
	}
	return huh.NewForm(huh.NewGroup(field)).
		WithInput(e.In.Raw()).
		WithOutput(e.Out.Raw()).
		WithTheme(e.Style.themeFunc)
}

func (e Env) Spinner() *spinner.Spinner {
	s := spinner.New().WithOutput(e.Err.Raw())
	if !e.Out.IsTerminal() {
		s = s.WithAccessible(true)
	}
	return s
}

// RunModel runs m as a full-screen TUI. It is the caller's responsibility to
// handle the non-terminal case before calling this.
func (e Env) RunModel(m tea.Model) error {
	_, err := e.newProgram(e.Ctx, m).Run()
	return err
}

func (e Env) Viewport() (ViewportWriter, func() error) {
	if !e.Out.IsTerminal() {
		p := &plainViewport{}
		return p, func() error {
			_, err := io.Copy(e.Out, &p.Buffer)
			return err
		}
	}

	v := &viewportModel{}
	v.program = e.newProgram(e.Ctx, v)

	return v, func() error {
		_, err := v.program.Run()
		return err
	}
}

// ViewportWithKeys runs m as an interactive TUI viewport. In non-terminal mode
// it renders once and writes directly to Out. Flags should set the initial
// model state for headless use.
func (e Env) ViewportWithKeys(m InteractiveModel) func() error {
	if !e.Out.IsTerminal() {
		return func() error {
			_, err := fmt.Fprint(e.Out, m.Render())
			return err
		}
	}
	return func() error {
		return e.RunModel(&interactiveViewportModel{model: m})
	}
}

func (e Env) newProgram(ctx context.Context, m tea.Model) *tea.Program {
	return tea.NewProgram(m,
		tea.WithContext(ctx),
		tea.WithInput(e.In.Raw()),
		tea.WithOutput(e.Out.Raw()),
	)
}
