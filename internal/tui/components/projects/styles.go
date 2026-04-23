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
	DocStyle   = lipgloss.NewStyle()
	TitleStyle = lipgloss.NewStyle()
	ItemStyle  = lipgloss.NewStyle()
)

type DefaultItemStyles struct {
	NormalTitle   lipgloss.Style
	NormalDesc    lipgloss.Style
	SelectedTitle lipgloss.Style
	SelectedDesc  lipgloss.Style
	DimmedTitle   lipgloss.Style
	DimmedDesc    lipgloss.Style
	FilterMatch   lipgloss.Style
}

func NewDefaultItemStyles() (s DefaultItemStyles) {
	t := pkgTheme
	s.NormalTitle = lipgloss.NewStyle().
		Foreground(t.PrimaryFg).
		Padding(0, 0, 0, 2)
	s.NormalDesc = s.NormalTitle.
		Foreground(t.TextDimmed)
	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(t.SelectionBorder).
		Foreground(t.PrimaryBright).
		Padding(0, 0, 0, 1)
	s.SelectedDesc = s.SelectedTitle.
		Foreground(t.PrimaryBright)
	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(t.TextDimmed).
		Padding(0, 0, 0, 2)
	s.DimmedDesc = s.DimmedTitle.
		Foreground(t.Dim)
	s.FilterMatch = lipgloss.NewStyle().Underline(true)
	return s
}
