package statusline

import (
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

// These are kept for any mrglab code that references the local style vars directly.
// The tuishell statusline uses its own theme-aware styles internally.

var (
	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color(style.StatuslineText)).
			Padding(0, 1)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(style.StatuslineText)).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.
			Background(lipgloss.Color(style.StatuslineEncoding)).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Foreground(lipgloss.Color(style.StatuslineText))

	helpText = lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center)

	projectStyle = statusNugget.Background(lipgloss.Color(style.StatuslineProject))
)
