package table

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	DefaultStyle = func() table.Styles {
		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(style.Violet[400])).
			BorderBottom(true).
			Bold(false)

		s.Selected = s.Selected.
			Foreground(lipgloss.Color(style.Violet[50])).
			Background(lipgloss.Color(style.Violet[800])).
			Bold(false)

		s.Cell = s.Cell

		return s
	}
	TitleStyle = lipgloss.NewStyle().
			Margin(0, 0, 0, 1).
			Foreground(lipgloss.Color(style.Violet[300])).
			Bold(true)

	EmptyMsg = TitleStyle.Align(lipgloss.Center, lipgloss.Center)
	DocStyle = lipgloss.NewStyle().
		// Margin(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(style.DarkGray)) // TODO: update color with tokens

)
