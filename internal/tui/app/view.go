package app

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
)

func (m Model) View() string {
	if m.Projects.IsOpen {
		t := table.TitleStyle.Render("Select Project")
		p := projects.DocStyle.Render(m.Projects.List.View())

		return lipgloss.JoinHorizontal(0, p, t)
	}
	table := table.DocStyle.Render(m.MergeRequests.Table.View())
	h := m.ctx.Window.Height - lipgloss.Height(table)
	sl := lipgloss.PlaceVertical(h, lipgloss.Bottom, m.Statusline.View())
	return lipgloss.JoinVertical(0, table, sl)
}
