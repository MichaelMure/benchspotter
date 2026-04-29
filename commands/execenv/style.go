package execenv

import (
	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
)

type Style struct {
	*huh.Styles
	themeFunc huh.ThemeFunc
}

func (s Style) TonedDown(strs ...string) string {
	return s.Styles.Focused.Description.Render(strs...)
}

func (s Style) Bold(text string) string {
	return lipgloss.NewStyle().Bold(true).Render(text)
}

func (s Style) Accent(text string) string {
	return s.Styles.Focused.Title.Render(text)
}

func (s Style) Positive(text string) string {
	return lipgloss.NewStyle().Foreground(compat.AdaptiveColor{
		Light: lipgloss.Color("28"),
		Dark:  lipgloss.Color("10"),
	}).Render(text)
}

func (s Style) Negative(text string) string {
	return lipgloss.NewStyle().Foreground(compat.AdaptiveColor{
		Light: lipgloss.Color("160"),
		Dark:  lipgloss.Color("9"),
	}).Render(text)
}
