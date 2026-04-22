package execenv

import "github.com/charmbracelet/huh"

type Style struct {
	*huh.Theme
}

func (s Style) TonedDown(strs ...string) string {
	return s.Theme.Focused.Description.Render(strs...)
}
