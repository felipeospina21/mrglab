package app

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/components/loader"
	"github.com/felipeospina21/mrglab/internal/tui/components/modal"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

func (m Model) View() tea.View {
	left := projects.DocStyle.Render(m.Projects.List.View())
	render := style.MainFrameStyle.Render
	screen := m.renderNormalScreen(left, render)

	if m.isModalOpen {
		m.setHelpKeys(modal.Keybinds)
		screen = m.Modal.View(screen)
	}

	v := tea.NewView(screen)
	v.AltScreen = true
	return v
}

func (m Model) renderNormalScreen(left string, render func(...string) string) string {
	isInitialScreen := m.taskStatus == taskIdle
	isFetching := m.taskStatus == taskStarted

	if isInitialScreen {
		if m.taskErr != nil {
			body := lipgloss.JoinHorizontal(0, left, m.taskErr.Error())
			sl := m.Statusline.View()
			return render(lipgloss.JoinVertical(0, body, sl))
		}
		m.MergeRequests.Table.W = m.layout.MainPanel.Width - table.DocStyle.GetHorizontalFrameSize() - tableBorderX
		m.MergeRequests.Table.H = m.layout.MainPanel.Height - tableBorderY
		tbl := table.DocStyle.Render(m.MergeRequests.Table.View())
		body := lipgloss.JoinHorizontal(0, left, tbl)
		sl := m.Statusline.View()
		return render(lipgloss.JoinVertical(0, body, sl))
	}

	body, sl := m.getMainPanelComponents()

	if m.isLeftOpen {
		if isFetching {
			body = loader.View(m.Spinner.View())
		}
		main := lipgloss.JoinHorizontal(0, left, body)
		sl := m.Statusline.View()
		return render(lipgloss.JoinVertical(0, main, sl))
	}

	if m.isRightOpen {
		right := m.Details.View()
		main := lipgloss.JoinHorizontal(0, body, right)
		return render(lipgloss.JoinVertical(0, main, sl))
	}

	return render(lipgloss.JoinVertical(0, body, sl))
}

func (m Model) getMainPanelComponents() (string, string) {
	header := table.TitleStyle.Render(
		fmt.Sprintf("%s - %s", m.ctx.SelectedProject.Name, "Merge Requests"),
	)
	body := lipgloss.JoinVertical(0, header, table.DocStyle.Render(m.MergeRequests.Table.View()))

	bodyHeight := lipgloss.Height(body)
	innerH := m.ctx.Window.Height - style.MainFrameStyle.GetVerticalFrameSize()
	remainingH := innerH - bodyHeight
	if remainingH < m.layout.Statusline.Height {
		remainingH = m.layout.Statusline.Height
	}
	statusline := lipgloss.PlaceVertical(remainingH, lipgloss.Bottom, m.Statusline.View())

	return body, statusline
}
