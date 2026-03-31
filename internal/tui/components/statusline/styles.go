package statusline

import (
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	// TODO: update colors with tokens
	StatusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#C1C6B2")).
			Background(lipgloss.Color("#353533")).
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

	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(style.Violet[400]))
)
