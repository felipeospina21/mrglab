package modal

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var headerStyle = table.TitleStyle.
	Foreground(lipgloss.Color(style.White)).
	Background(lipgloss.Color(style.Violet[700])).
	Padding(0, 1).
	MarginBottom(1)

var helpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(style.DarkGray)).
	MarginTop(1)

var boxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#666666")).
	Padding(1)

var dimStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#444444"))
