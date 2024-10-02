package app

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) View() string {
	if m.ctx.IsLeftPanelOpen && m.ctx.TaskStatus != task.TaskFinished {
		w, h := m.getEmptyTableSize()
		return m.renderLayout(LayoutComponents{
			body: table.EmptyMsg.Width(w).Height(h).Render("Select Project"),
		})
	}

	if m.ctx.IsRightPanelOpen {
		return m.renderLayout(LayoutComponents{
			header: table.TitleStyle.Render(
				fmt.Sprintf("%s - %s", m.ctx.SelectedProject.Name, "Merge Requests"),
			),
			body: table.DocStyle.Render(m.MergeRequests.Table.View()),
		})
	}

	return m.renderLayout(LayoutComponents{
		header: table.TitleStyle.Render(
			fmt.Sprintf("%s - %s", m.ctx.SelectedProject.Name, "Merge Requests"),
		),
		body: table.DocStyle.Render(m.MergeRequests.Table.View()),
	})
}

type LayoutComponents struct {
	header string
	body   string
}

func (m Model) renderLayout(c LayoutComponents) string {
	left := projects.DocStyle.Render(m.Projects.List.View())
	main := lipgloss.JoinVertical(0, c.header, c.body)
	body := lipgloss.JoinHorizontal(0, left, main)
	sl := m.Statusline.View()
	if !m.ctx.IsLeftPanelOpen {
		h := m.ctx.Window.Height - lipgloss.Height(c.header) - lipgloss.Height(c.body) - MainFrameStyle.GetVerticalFrameSize()
		sl = lipgloss.PlaceVertical(h, lipgloss.Bottom, m.Statusline.View())
		body = main
	}

	if m.ctx.IsRightPanelOpen {
		render := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Render
		right := render(
			fmt.Sprintf(
				"%s\n%s\n%s",
				m.Details.HeaderView("Header"),
				m.Details.Viewport.View(),
				m.Details.FooterView(),
			),
		)
		body = lipgloss.JoinHorizontal(0, main, right)
	}

	return MainFrameStyle.Render(lipgloss.JoinVertical(0, body, sl))
}
