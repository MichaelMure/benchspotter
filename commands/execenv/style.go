package execenv

import (
	"image/color"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
)

type Style struct {
	*huh.Styles
	themeFunc huh.ThemeFunc
	isDark    bool
}

func NewStyle(tf huh.ThemeFunc, isDark bool) Style {
	return Style{Styles: tf.Theme(isDark), themeFunc: tf, isDark: isDark}
}

func (s Style) ld(light, dark string) color.Color {
	return lipgloss.LightDark(s.isDark)(lipgloss.Color(light), lipgloss.Color(dark))
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

// Positive renders text in green — improvement (smaller is better).
func (s Style) Positive(text string) string {
	return lipgloss.NewStyle().Foreground(s.ld("28", "10")).Render(text)
}

// Negative renders text in red — regression (larger is worse).
func (s Style) Negative(text string) string {
	return lipgloss.NewStyle().Foreground(s.ld("160", "9")).Render(text)
}

// Warning renders text in orange — heap escapes, cannot-inline, etc.
func (s Style) Warning(text string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Render(text)
}

// Info renders text in cyan/teal — leaking params, inlining calls, etc.
func (s Style) Info(text string) string {
	return lipgloss.NewStyle().Foreground(s.ld("33", "81")).Render(text)
}

// Subject renders a symbol or function name with bold emphasis.
func (s Style) Subject(text string) string {
	return lipgloss.NewStyle().Bold(true).Foreground(s.ld("16", "255")).Render(text)
}

// FlowLine renders a secondary flow-chain annotation line.
func (s Style) FlowLine(text string) string {
	return lipgloss.NewStyle().Foreground(s.ld("244", "240")).Render(text)
}

// Corr renders a correlation value with color reflecting strength and direction.
func (s Style) Corr(r float64, text string) string {
	abs := r
	if abs < 0 {
		abs = -abs
	}
	var st lipgloss.Style
	switch {
	case abs < 0.3:
		st = lipgloss.NewStyle().Foreground(s.ld("243", "245"))
	case r > 0.7:
		st = lipgloss.NewStyle().Foreground(s.ld("28", "82")).Bold(true)
	case r > 0:
		st = lipgloss.NewStyle().Foreground(s.ld("34", "76"))
	case r < -0.7:
		st = lipgloss.NewStyle().Foreground(s.ld("160", "196")).Bold(true)
	default:
		st = lipgloss.NewStyle().Foreground(s.ld("166", "214"))
	}
	return st.Render(text)
}

// WithBg renders text with the given background color.
func (s Style) WithBg(bg color.Color, text string) string {
	return lipgloss.NewStyle().Background(bg).Render(text)
}

// SelectionBg returns the background color for a selected/highlighted row.
func (s Style) SelectionBg() color.Color {
	return s.ld("254", "238")
}

// TrendCellBg returns a background color for a trend cell based on the curr/prev ratio.
func (s Style) TrendCellBg(ratio float64) (color.Color, bool) {
	switch {
	case ratio > 1.15:
		return s.ld("224", "88"), true
	case ratio > 1.05:
		return s.ld("217", "52"), true
	case ratio < 0.85:
		return s.ld("120", "28"), true
	case ratio < 0.95:
		return s.ld("157", "22"), true
	}
	return lipgloss.NoColor{}, false
}

// TrendArrow returns a colored directional arrow for the curr/prev ratio.
func (s Style) TrendArrow(curr, prev float64, bg ...color.Color) string {
	if prev == 0 {
		return ""
	}
	ratio := curr / prev
	var bgColor color.Color = lipgloss.NoColor{}
	if len(bg) > 0 {
		bgColor = bg[0]
	}
	red := lipgloss.NewStyle().Foreground(s.ld("160", "9")).Background(bgColor)
	green := lipgloss.NewStyle().Foreground(s.ld("28", "10")).Background(bgColor)
	switch {
	case ratio > 1.15:
		return red.Render("↑↑")
	case ratio > 1.05:
		return red.Render("↑")
	case ratio < 0.85:
		return green.Render("↓↓")
	case ratio < 0.95:
		return green.Render("↓")
	}
	return ""
}

// Sidebar methods — used by the source view sidebar in show_escape / show_inline.

func (s Style) SidebarFile(text string) string {
	return lipgloss.NewStyle().Foreground(s.ld("238", "250")).Render(text)
}

func (s Style) SidebarFn(text string) string {
	return lipgloss.NewStyle().Foreground(s.ld("242", "244")).Render(text)
}

func (s Style) SidebarActive(text string) string {
	return lipgloss.NewStyle().Bold(true).Foreground(s.ld("16", "255")).Render(text)
}

func (s Style) SidebarSelected(text string) string {
	return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214")).Render(text)
}

func (s Style) SidebarDim(text string) string {
	return lipgloss.NewStyle().Foreground(s.ld("250", "238")).Render(text)
}

func (s Style) SidebarDivider(text string) string {
	return lipgloss.NewStyle().Foreground(s.ld("244", "240")).Render(text)
}
