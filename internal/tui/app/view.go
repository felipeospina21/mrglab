package app

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/components/modal"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/style"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) View() string {
	left := projects.DocStyle.Render(m.Projects.List.View())
	render := style.MainFrameStyle.Render
	isInitialScreen := m.ctx.Task.Status == task.TaskIdle
	isFetching := m.ctx.Task.Status == task.TaskStarted

	switch {
	case isInitialScreen:
		if m.ctx.Task.Err != nil {
			body := lipgloss.JoinHorizontal(0, left, m.ctx.Task.Err.Error())
			sl := m.Statusline.View()
			return render(lipgloss.JoinVertical(0, body, sl))

		}

		m.MergeRequests.Table.W, m.MergeRequests.Table.H = m.getEmptyTableSize()
		body := lipgloss.JoinHorizontal(0, left, table.DocStyle.Render(m.MergeRequests.Table.View()))
		sl := m.Statusline.View()
		return render(lipgloss.JoinVertical(0, body, sl))

	case m.ctx.IsModalOpen:
		m.setHelpKeys(modal.Keybinds)
		body := m.Modal.View()
		sl := m.Statusline.View()
		return render(lipgloss.JoinVertical(0, body, sl))

	default:
		body, sl := m.getMainPanelComponents()

		if m.ctx.IsLeftPanelOpen {
			if isFetching {
				textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(style.Violet[300])).Render
				body = fmt.Sprintf("\n %s%s%s\n\n",
					m.Spinner.View(),
					" ",
					textStyle("Fetching..."),
				)
			}
			main := lipgloss.JoinHorizontal(0, left, body)
			sl := m.Statusline.View()
			return render(lipgloss.JoinVertical(0, main, sl))

		}

		if m.ctx.IsRightPanelOpen {
			right := m.Details.View()
			main := lipgloss.JoinHorizontal(0, body, right)
			return render(lipgloss.JoinVertical(0, main, sl))
		}

		return render(lipgloss.JoinVertical(0, body, sl))
	}
}

func (m Model) getMainPanelComponents() (string, string) {
	header := table.TitleStyle.Render(
		fmt.Sprintf("%s - %s", m.ctx.SelectedProject.Name, "Merge Requests"),
	)
	body := lipgloss.JoinVertical(0, header, table.DocStyle.Render(m.MergeRequests.Table.View()))
	h := m.ctx.Window.Height - lipgloss.Height(body) - style.MainFrameStyle.GetVerticalFrameSize()
	statusline := lipgloss.PlaceVertical(h, lipgloss.Bottom, m.Statusline.View())

	return body, statusline
}
