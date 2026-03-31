package statusline

import (
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	// TODO: update colors with tokens
	StatusBarStyle = lipgloss.NewStyle().
			Margin(0, 0)

	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color(style.StatuslineText)).
			Padding(0, 1)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(style.StatuslineText)).
			Background(lipgloss.Color(style.StatuslineMode)).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.
			Background(lipgloss.Color(style.StatuslineEncoding)).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle()

	helpText = lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center)

	projectStyle = statusNugget.Background(lipgloss.Color(style.StatuslineProject))

	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(style.Violet[400]))
)
