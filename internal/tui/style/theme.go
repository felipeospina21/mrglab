package style

import (
	"charm.land/lipgloss/v2"
)

// MainFrameStyle is the outer border style for the main application frame.
var MainFrameStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color(DarkGray))
