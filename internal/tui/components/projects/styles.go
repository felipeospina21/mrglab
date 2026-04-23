package projects

import (
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/tuishell/style"
)

var pkgTheme style.Theme

// SetTheme sets the theme used by the projects package and refreshes derived styles.
func SetTheme(t style.Theme) {
	pkgTheme = t
	DocStyle = lipgloss.NewStyle().
		PaddingRight(4).
		MarginBottom(2).
		Foreground(t.Primary).
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(t.Border).
		Width(30)
	TitleStyle = lipgloss.NewStyle().
		MarginTop(0).
		Foreground(t.Info)
	ItemStyle = lipgloss.NewStyle().
		MarginTop(0).
		Foreground(t.Primary)
}

var (
	DocStyle = lipgloss.NewStyle().
			PaddingRight(4).
			MarginBottom(2).
			Foreground(lipgloss.Color("#b8a6ff")).
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(lipgloss.Color("#3f4145")).
			Width(30)

	TitleStyle = lipgloss.NewStyle().
			MarginTop(0).
			Foreground(lipgloss.Color("#3ac4d9"))
	ItemStyle = lipgloss.NewStyle().
			MarginTop(0).
			Foreground(lipgloss.Color("#b8a6ff"))
)

type DefaultItemStyles struct {
	NormalTitle    lipgloss.Style
	NormalDesc     lipgloss.Style
	SelectedTitle  lipgloss.Style
	SelectedDesc   lipgloss.Style
	DimmedTitle    lipgloss.Style
	DimmedDesc     lipgloss.Style
	FilterMatch    lipgloss.Style
}

func NewDefaultItemStyles() (s DefaultItemStyles) {
	t := pkgTheme
	hl := t.PrimaryBright
	fg := t.PrimaryFg
	if hl == nil {
		hl = lipgloss.Color("#9673ff")
	}
	if fg == nil {
		fg = lipgloss.Color("#f2f0ff")
	}

	s.NormalTitle = lipgloss.NewStyle().
		Foreground(fg).
		Padding(0, 0, 0, 2)

	s.NormalDesc = s.NormalTitle.
		Foreground(lipgloss.Color("#777777"))

	selBorder := t.SelectionBorder
	if selBorder == nil {
		selBorder = lipgloss.Color("#AD58B4")
	}

	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(selBorder).
		Foreground(hl).
		Padding(0, 0, 0, 1)

	s.SelectedDesc = s.SelectedTitle.
		Foreground(hl)

	dimmed := t.TextDimmed
	if dimmed == nil {
		dimmed = lipgloss.Color("#777777")
	}

	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(dimmed).
		Padding(0, 0, 0, 2)

	s.DimmedDesc = s.DimmedTitle.
		Foreground(lipgloss.Color("#4D4D4D"))

	s.FilterMatch = lipgloss.NewStyle().Underline(true)

	return s
}
