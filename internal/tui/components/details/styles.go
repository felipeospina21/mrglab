package details

import "github.com/charmbracelet/lipgloss"

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
)
