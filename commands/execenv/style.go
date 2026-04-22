package execenv

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Style struct {
	*huh.Theme
}

func (s Style) TonedDown(strs ...string) string {
	return s.Theme.Focused.Description.Render(strs...)
}

func (s Style) Bold(text string) string {
	return lipgloss.NewStyle().Bold(true).Render(text)
}

func (s Style) Accent(text string) string {
	return s.Theme.Focused.Title.Render(text)
}
