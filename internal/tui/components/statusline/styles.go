package statusline

import (
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	StatusBarStyle = lipgloss.NewStyle().
			Margin(0, 0)

	statusNugget = lipgloss.NewStyle().
			Inherit(statusText).
			Padding(0, 1)

	statusStyle = lipgloss.NewStyle().
			Inherit(statusText).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.
			Inherit(statusText).
			Background(lipgloss.Color(style.StatuslineEncoding)).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Foreground(lipgloss.Color(style.StatuslineText))

	helpText = lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center)

	projectStyle = statusNugget.Background(lipgloss.Color(style.StatuslineProject))

	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(style.Violet[400]))
)
