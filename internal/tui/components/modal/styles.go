package modal

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	headerStyle = func(isError bool) lipgloss.Style {
		s := table.TitleStyle.
			Foreground(lipgloss.Color(style.White)).
			Padding(0, 1).
			Background(lipgloss.Color(style.Violet[600]))

		if isError {
			s = s.Background(lipgloss.Color(style.Red[600]))
		}

		return s
	}

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(style.DarkGray)).
			MarginTop(1)

	bodyStyle = func(h int) lipgloss.Style {
		return lipgloss.NewStyle().Padding(1).Height(h)
	}

	formStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderRight(true)

	inputLabel = lipgloss.NewStyle().
			Foreground(lipgloss.Color(style.Violet[400])).
			Bold(true)
)
