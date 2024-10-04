package style

import (
	"github.com/charmbracelet/lipgloss"
)

var MainFrameStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color(DarkGray))
