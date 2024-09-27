package app

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case error:
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(msg.Error())

	case tea.KeyMsg:
		if msg.String() == "a" {
			m.Statusline.Status = "test"
			m.Statusline.Content = "test"
		}
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

		// TODO: separate in views
		case key.Matches(msg, projects.Keybinds.MRList):
			cb := func() tea.Cmd {
				m.Projects.SelectProject()
				return m.MergeRequests.GetListCmd()
			}
			cmds = append(cmds, m.startCommand(cb))

		case key.Matches(msg, projects.Keybinds.ToggleSidePanel):
			m.ctx.IsLeftPanelOpen = !m.ctx.IsLeftPanelOpen
			logger.Debug("table", m.ctx.Window.Width, m.Projects.List.Width())
			// m.MergeRequests.Table.SetWidth(150)
		}

	case spinner.TickMsg:
		cmd = m.updateSpinnerViewCommand(msg)
		cmds = append(cmds, cmd)

	case tea.WindowSizeMsg:
		// Sets window in context
		m.ctx.Window = msg

		m.setLeftPanelHeight()

		m.setStatuslineWidth()

		// table
		// m.MergeRequests.Table.SetWidth(msg.Width - w)

	case task.TaskFinishedMsg:
		// TODO: Rethink this logic
		if msg.SectionType == "mrs" {
			t := endCommand[table.Model](
				&m,
				msg,
				m.MergeRequests.GetTableModel(msg),
			)

			m.MergeRequests.Table = t
		}
	}

	m.Projects.List, cmd = m.Projects.List.Update(msg)
	m.MergeRequests.Table, cmd = m.MergeRequests.Table.Update(msg)
	return m, tea.Batch(cmds...)
}
