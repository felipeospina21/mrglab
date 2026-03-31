package projects

import (
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	DocStyle = lipgloss.NewStyle().
			PaddingRight(4).
		// Margin(0, 0).
		MarginBottom(2).
		Foreground(lipgloss.Color(style.Violet[300])).
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color(style.DarkGray)).
		Width(30) // TODO: Set width in config file

	TitleStyle = lipgloss.NewStyle().
			MarginTop(0).
			Foreground(lipgloss.Color(style.Blue[400]))
	ItemStyle = lipgloss.NewStyle().
			MarginTop(0).
			Foreground(lipgloss.Color(style.Violet[300]))
	// SelectedItemStyle = lipgloss.NewStyle().
	// 			MarginLeft(2).
	// 			MarginTop(1).
	// 			PaddingLeft(2).
	// 			Foreground(lipgloss.Color(style.Violet[50])).
	// 			Background(lipgloss.Color(style.Violet[800]))
	// PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	// HelpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

var (
	hl = style.Violet[400]
	fg = style.Violet[50]
)

type DefaultItemStyles struct {
	// The Normal state.
	NormalTitle lipgloss.Style
	NormalDesc  lipgloss.Style

	// The selected item state.
	SelectedTitle lipgloss.Style
	SelectedDesc  lipgloss.Style

	// The dimmed state, for when the filter input is initially activated.
	DimmedTitle lipgloss.Style
	DimmedDesc  lipgloss.Style

	// Characters matching the current filter, if any.
	FilterMatch lipgloss.Style
}

func NewDefaultItemStyles() (s DefaultItemStyles) {
	s.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(fg)).
		Padding(0, 0, 0, 2)

	s.NormalDesc = s.NormalTitle.
		Foreground(lipgloss.Color("#777777"))

	s.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.Color("#AD58B4")).
		Foreground(lipgloss.Color(hl)).
		Padding(0, 0, 0, 1)

	s.SelectedDesc = s.SelectedTitle.
		Foreground(lipgloss.Color(hl))

	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#777777")).
		Padding(0, 0, 0, 2)

	s.DimmedDesc = s.DimmedTitle.
		Foreground(lipgloss.Color("#4D4D4D"))

	s.FilterMatch = lipgloss.NewStyle().Underline(true)

	return s
}
