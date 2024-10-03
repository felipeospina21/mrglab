package app

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	isLeftPanelFocused := m.ctx.FocusedPanel == context.LeftPanel
	isMainPanelFocused := m.ctx.FocusedPanel == context.MainPanel
	isRightPanelFocused := m.ctx.FocusedPanel == context.RightPanel

	switch msg := msg.(type) {
	case details.IsDetailsResponseReady:
		m.Details.SetResponseContent(content)

	case error:
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(msg.Error())

	case tea.KeyMsg:
		// TODO: Delete this cmd
		if msg.String() == "a" {
			m.Statusline.Status = "test"
			m.Statusline.Content = "test"
		}
		// TODO: Delete this cmd
		if msg.String() == "b" {
			m.Statusline.Status = statusline.ModesEnum.Loading
			cmds = append(cmds, cmd)
		}
		switch {
		case key.Matches(msg, tui.GlobalKeys.Quit):
			return m, tea.Quit

		case key.Matches(msg, tui.GlobalKeys.ThrowError):
			cmds = append(cmds, func() tea.Msg {
				return errors.New("mocked")
			})

		case key.Matches(msg, projects.Keybinds.ToggleSidePanel):
			m.toggleLeftPanel()
		}

		if isLeftPanelFocused {
			switch {
			case key.Matches(msg, projects.Keybinds.MRList):
				cb := func() tea.Cmd {
					m.toggleLeftPanel()
					m.Projects.SelectProject()
					return m.MergeRequests.GetListCmd()
				}
				cmds = append(cmds, m.startCommand(cb))
			}
		}

		if isMainPanelFocused {
			switch {
			// TODO: replace keybinds
			case key.Matches(msg, projects.Keybinds.MRList):
				// TODO: move to separate command function (duplicated in toggleLeftPanel)
				m.ctx.IsRightPanelOpen = !m.ctx.IsRightPanelOpen
				m.MergeRequests.Table.SetWidth(lipgloss.Width(m.MergeRequests.Table.View()))
				m.MergeRequests.Table.UpdateViewport()
				// m.Details.SetFocus()

				// TODO: move to separate command function
				viewportWidth := m.ctx.Window.Width - lipgloss.Width(m.MergeRequests.Table.View())

				c := m.Details.SetViewportViewSize(
					tea.WindowSizeMsg{Width: viewportWidth, Height: m.ctx.Window.Height},
				)

				cmd = func() tea.Msg {
					return details.IsDetailsResponseReady(true)
				}
				// m.Details.Viewport.Update(msg)
				cmds = append(cmds, c, cmd)
			}
		}

		if isRightPanelFocused {
			// TODO: right panel cmds
		}

	case spinner.TickMsg:
		cmd = m.updateSpinnerViewCommand(msg)
		cmds = append(cmds, cmd)

	case tea.WindowSizeMsg:
		// Sets window in context
		m.ctx.Window = msg

		m.setLeftPanelHeight()
		m.setStatuslineWidth()

	case task.TaskFinishedMsg:
		// TODO: Rethink this logic
		if msg.SectionType == "mrs" {
			t := endCommand[table.Model](
				&m,
				msg,
				m.MergeRequests.GetTableModel(msg),
			)

			m.MergeRequests.SetFocus()
			m.MergeRequests.Table = t
		}
	}

	m.Projects.List, cmd = m.Projects.List.Update(msg)
	m.MergeRequests.Table, cmd = m.MergeRequests.Table.Update(msg)
	m.Details.Viewport, cmd = m.Details.Viewport.Update(msg)
	return m, tea.Batch(cmds...)
}
