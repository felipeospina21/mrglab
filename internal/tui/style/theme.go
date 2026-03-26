package style

import (
	"github.com/charmbracelet/lipgloss"
)

// MainFrameStyle is the outer border style for the main application frame.
var MainFrameStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color(DarkGray))
