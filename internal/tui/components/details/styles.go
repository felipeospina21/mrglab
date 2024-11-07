package details

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	MdTitle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0).Margin(0)
	}()

	MdInfo = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return MdTitle.BorderStyle(b)
	}()

	PanelStyle = lipgloss.NewStyle().
			MarginTop(1).
			Border(lipgloss.NormalBorder(), true, false, true, true).
			BorderForeground(lipgloss.Color(style.DarkGray))

	iconStyle = func(color string) lipgloss.Style {
		return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).MarginLeft(2)
	}

	contentStyle = lipgloss.NewStyle().MarginLeft(LeftMargin)

	sectionTextStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color(style.White)).MarginLeft(1)
	sectionTitleStyle        = sectionTextStyle.Bold(true).MarginLeft(0)
	sectionIndentedTextStyle = sectionTextStyle.MarginLeft(LeftMargin).Foreground(lipgloss.Color(style.DarkGray))
)
