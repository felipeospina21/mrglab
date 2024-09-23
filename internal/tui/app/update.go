package app

import (
	"errors"
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
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
		logger.Debug("err", func() {
			log.Println(msg)
		})

	case tea.KeyMsg:
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
				return m.MergeRequests.GetMRListCmd()
			}
			cmds = append(cmds, m.startCommand(cb))
		}

	case spinner.TickMsg:
		m.Statusline.Spinner, cmd = m.Statusline.Spinner.Update(msg)
		cmds = append(cmds, cmd)

	case tea.WindowSizeMsg:
		m.ctx.Window = msg
		h, v := projects.ItemStyle.GetFrameSize()
		th, tv := projects.TitleStyle.GetFrameSize()
		m.Projects.List.SetSize(msg.Width-h-th, msg.Height-v-tv)
		m.Statusline.Width = msg.Width - statusline.StatusBarStyle.GetHorizontalFrameSize() - table.DocStyle.GetHorizontalFrameSize()

	case task.TaskFinishedMsg:
		// TODO: Rethink this logic
		if msg.SectionType == "mrs" {
			var rows []table.Row
			if msg.Err != nil {
				endCommand[error](
					&m,
					endCommandStatus{isError: true, error: msg.Err},
					func() error {
						return msg.Err
					},
				)
			} else {
				rows = endCommand[[]table.Row](
					&m,
					endCommandStatus{isSuccess: true},
					func() []table.Row {
						return mergerequests.GetMRTableRows(msg)
					},
				)
			}

			m.MergeRequests.Table = table.InitModel(table.InitModelParams{
				Rows:   rows,
				Colums: mergerequests.GetMergeReqsColums(m.ctx.Window.Width - 10),
				// StyleFunc: mergerequests.StyleIconsColumns(table.Styles(table.DefaultStyle()), table.MergeReqsIconCols),
			})
		}
	}

	m.Projects.List, cmd = m.Projects.List.Update(msg)
	m.MergeRequests.Table, cmd = m.MergeRequests.Table.Update(msg)
	return m, tea.Batch(cmds...)
}
