package statusline

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	StatusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"}).
			Margin(0, 0)

	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusStyle = lipgloss.NewStyle().
			Inherit(StatusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.
			Background(lipgloss.Color("#A550DF")).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Inherit(StatusBarStyle)

	helpText = lipgloss.NewStyle().Inherit(StatusBarStyle).AlignHorizontal(lipgloss.Center)

	projectStyle = statusNugget.Background(lipgloss.Color("#6124DF"))

	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(style.Violet[400]))
)
