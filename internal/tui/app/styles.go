package app

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var MainFrameStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color(style.DarkGray))
