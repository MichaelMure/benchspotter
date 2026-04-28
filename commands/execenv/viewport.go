package execenv

import (
	"bytes"
	"io"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
	case tea.KeyPressMsg:
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
			v.viewport = viewport.New(viewport.WithWidth(msg.Width), viewport.WithHeight(msg.Height))
			v.ready = true
			v.viewport.SetContent(v.content.String())
		} else {
			v.viewport.SetWidth(msg.Width)
			v.viewport.SetHeight(msg.Height)
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

func (v *viewportModel) View() tea.View {
	var content string
	if !v.ready {
		content = "\n  Initializing..."
	} else {
		content = v.viewport.View()
	}
	view := tea.NewView(content)
	view.AltScreen = true
	view.MouseMode = tea.MouseModeCellMotion
	return view
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

// InteractiveModel is implemented by commands that want a TUI viewport with
// dynamic content. The viewport calls Render on init and after any key that
// HandleKey consumes; Status is rendered as a fixed footer line.
type InteractiveModel interface {
	Render() string
	Status() string
	HandleKey(key string) bool
}

// YOffsetAdjuster may optionally be implemented by an InteractiveModel to
// preserve a meaningful scroll position when content changes after a key press.
// oldContent and newContent are the full rendered strings before and after;
// oldOffset is the viewport's YOffset and viewportHeight is its visible line
// count before the change. The returned value becomes the new YOffset.
type YOffsetAdjuster interface {
	AdjustYOffset(oldContent, newContent string, oldOffset, viewportHeight int) int
}

// YOffsetSetter may optionally be implemented by an InteractiveModel to
// receive the current viewport position before Status() is called each frame,
// allowing the status bar to reflect the current scroll position (e.g. breadcrumb).
type YOffsetSetter interface {
	SetCurrentYOffset(yOffset, viewportHeight int)
}

// JumpRequester may optionally be implemented by an InteractiveModel to
// request a specific scroll position after a key press that does not change
// content (e.g. jumping between groups).
type JumpRequester interface {
	ConsumeJumpOffset() (int, bool)
}

// ModalModel may optionally be implemented by an InteractiveModel to signal
// that it is in a transient modal state (e.g. a search prompt) where Esc
// should be forwarded to the model rather than quitting the application.
type ModalModel interface {
	IsModal() bool
}

// SidebarProvider may optionally be implemented by an InteractiveModel to
// supply a fixed-width left panel. SidebarWidth returns the desired total
// sidebar width (including its divider column) for the given terminal width,
// or 0 to hide it. RenderSidebar must produce exactly height lines each
// visually width characters wide.
type SidebarProvider interface {
	SidebarWidth(totalWidth int) int
	RenderSidebar(width, height int) string
}

var _ tea.Model = &interactiveViewportModel{}

type interactiveViewportModel struct {
	model        InteractiveModel
	viewport     viewport.Model
	ready        bool
	lastContent  string
	sidebarWidth int
}

func (v *interactiveViewportModel) Init() tea.Cmd { return nil }

func (v *interactiveViewportModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		k := msg.String()
		if k == "ctrl+c" || k == "q" {
			return v, tea.Quit
		}
		if k == "esc" {
			if mm, ok := v.model.(ModalModel); !ok || !mm.IsModal() {
				return v, tea.Quit
			}
		}
		if v.model.HandleKey(k) {
			oldContent := v.lastContent
			oldOffset := v.viewport.YOffset()
			newContent := v.model.Render()
			v.lastContent = newContent
			v.viewport.SetContent(newContent)
			if adj, ok := v.model.(YOffsetAdjuster); ok {
				v.viewport.SetYOffset(adj.AdjustYOffset(oldContent, newContent, oldOffset, v.viewport.Height()))
			}
			return v, nil
		}
		if jr, ok := v.model.(JumpRequester); ok {
			if offset, ok := jr.ConsumeJumpOffset(); ok {
				v.viewport.SetYOffset(offset)
				return v, nil
			}
		}
	case tea.WindowSizeMsg:
		if sp, ok := v.model.(SidebarProvider); ok {
			v.sidebarWidth = sp.SidebarWidth(msg.Width)
		}
		vpWidth := msg.Width - v.sidebarWidth
		// Reserve one line for the status footer.
		if !v.ready {
			v.viewport = viewport.New(viewport.WithWidth(vpWidth), viewport.WithHeight(msg.Height-1))
			v.lastContent = v.model.Render()
			v.viewport.SetContent(v.lastContent)
			v.ready = true
		} else {
			v.viewport.SetWidth(vpWidth)
			v.viewport.SetHeight(msg.Height - 1)
		}
	}
	var cmd tea.Cmd
	v.viewport, cmd = v.viewport.Update(msg)
	return v, cmd
}

func (v *interactiveViewportModel) View() tea.View {
	var content string
	if !v.ready {
		content = "\n  Initializing..."
	} else {
		if setter, ok := v.model.(YOffsetSetter); ok {
			setter.SetCurrentYOffset(v.viewport.YOffset(), v.viewport.Height())
		}
		main := v.viewport.View()
		if sp, ok := v.model.(SidebarProvider); ok && v.sidebarWidth > 0 {
			sidebar := sp.RenderSidebar(v.sidebarWidth, v.viewport.Height())
			main = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, main)
		}
		content = main + "\n" + v.model.Status()
	}
	view := tea.NewView(content)
	view.AltScreen = true
	view.MouseMode = tea.MouseModeCellMotion
	return view
}
