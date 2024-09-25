package app

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) View() string {
	if m.Projects.IsOpen && m.ctx.TaskStatus != task.TaskFinished {
		w, h := getWindowFrameSize(m.ctx.Window.Width, m.ctx.Window.Height)
		body := lipgloss.JoinHorizontal(
			0,
			projects.DocStyle.Render(m.Projects.List.View()),
			table.EmptyMsg.Width(w).Height(h).Render("Select Project"),
		)

		return m.renderLayout(LayoutComponents{
			header: "",
			body:   body,
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
	h := m.ctx.Window.Height - lipgloss.Height(c.header) - lipgloss.Height(c.body)
	sl := lipgloss.PlaceVertical(h, lipgloss.Bottom, m.Statusline.View())
	return lipgloss.JoinVertical(0, c.header, c.body, sl)
}

func getWindowFrameSize(width, height int) (int, int) {
	h, v := projects.GetFrameSize()

	return width - h, height - v
}
