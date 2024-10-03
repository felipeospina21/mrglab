package app

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) View() string {
	left := projects.DocStyle.Render(m.Projects.List.View())
	if m.ctx.TaskStatus == task.TaskIdle {
		m.MergeRequests.Table.W, m.MergeRequests.Table.H = m.getEmptyTableSize()
		body := lipgloss.JoinHorizontal(0, left, table.DocStyle.Render(m.MergeRequests.Table.View()))
		sl := m.Statusline.View()
		return MainFrameStyle.Render(lipgloss.JoinVertical(0, body, sl))

	}

	header := table.TitleStyle.Render(
		fmt.Sprintf("%s - %s", m.ctx.SelectedProject.Name, "Merge Requests"),
	)
	body := lipgloss.JoinVertical(0, header, table.DocStyle.Render(m.MergeRequests.Table.View()))
	if m.ctx.IsLeftPanelOpen {
		main := lipgloss.JoinHorizontal(0, left, body)
		sl := m.Statusline.View()
		return MainFrameStyle.Render(lipgloss.JoinVertical(0, main, sl))

	}

	h := m.ctx.Window.Height - lipgloss.Height(body) - MainFrameStyle.GetVerticalFrameSize()
	sl := lipgloss.PlaceVertical(h, lipgloss.Bottom, m.Statusline.View())
	if m.ctx.IsRightPanelOpen {
		render := lipgloss.NewStyle().MarginTop(1).Border(lipgloss.NormalBorder()).Render
		right := render(
			fmt.Sprintf(
				"%s\n%s\n%s",
				m.Details.HeaderView("Header"),
				m.Details.Viewport.View(),
				m.Details.FooterView(),
			),
		)
		main := lipgloss.JoinHorizontal(0, body, right)
		return MainFrameStyle.Render(lipgloss.JoinVertical(0, main, sl))
	}

	return MainFrameStyle.Render(lipgloss.JoinVertical(0, body, sl))
}
