package style

import (
	"charm.land/lipgloss/v2"
	tsstyle "github.com/felipeospina21/tuishell/style"
)

// DefaultTheme returns the default mrglab theme.
var DefaultTheme = tsstyle.DefaultTheme

// MainFrameStyle is the outer border style for the main application frame.
var MainFrameStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color(DarkGray))
