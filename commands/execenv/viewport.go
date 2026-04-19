package execenv

import (
	"bytes"
	"io"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

var _ tea.Model = &viewportModel{}
var _ io.Writer = &viewportModel{}

type ViewportWriter interface {
	io.Writer
	SetContent(content string)
}

type viewportModel struct {
	content  bytes.Buffer
	ready    bool
	viewport viewport.Model
	program  *tea.Program
}

func (v *viewportModel) Init() tea.Cmd {
	return nil
}

type refreshMsg struct{}

func (v *viewportModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return v, tea.Quit
		}

	case tea.WindowSizeMsg:
		if !v.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			v.viewport = viewport.New(msg.Width, msg.Height)
			v.ready = true
			v.viewport.SetContent(v.content.String())
		} else {
			v.viewport.Width = msg.Width
			v.viewport.Height = msg.Height
		}

	case refreshMsg:
		if v.ready {
			v.viewport.SetContent(v.content.String())
		}
	}

	// Handle keyboard and mouse events in the viewport
	v.viewport, cmd = v.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return v, tea.Batch(cmds...)
}

func (v *viewportModel) View() string {
	if !v.ready {
		return "\n  Initializing..."
	}
	return v.viewport.View()
}

func (v *viewportModel) SetContent(content string) {
	v.content.Reset()
	v.content.WriteString(content)
	if !v.ready {
		v.viewport.SetContent(v.content.String())
	} else {
		v.program.Send(refreshMsg{})
	}
}

func (v *viewportModel) Write(p []byte) (int, error) {
	n, err := v.content.Write(p)
	if err != nil {
		return n, err
	}
	if !v.ready {
		v.viewport.SetContent(v.content.String())
	} else {
		v.program.Send(refreshMsg{})
	}
	return n, nil
}

// plainViewport is used when stdout is not a terminal: content is buffered and
// flushed directly to Out by the runFn returned from Viewport().
type plainViewport struct {
	bytes.Buffer
}

func (p *plainViewport) SetContent(content string) {
	p.Buffer.Reset()
	p.Buffer.WriteString(content)
}
