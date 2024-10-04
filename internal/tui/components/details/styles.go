package details

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	MdTitle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	MdInfo = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return MdTitle.Copy().BorderStyle(b)
	}()

	PanelStyle = lipgloss.NewStyle().
			MarginTop(1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(style.DarkGray))
)
