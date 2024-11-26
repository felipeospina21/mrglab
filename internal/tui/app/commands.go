package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) getMergeRequestModel(msg task.TaskMsg) func() table.Model {
	return func() table.Model {
		mrMsg := msg.Msg.(message.MergeRequestsListFetchedMsg)
		rows := mergerequests.GetTableRows(mrMsg)
		mainPanelHeaderHeight := 1
		return table.InitModel(table.InitModelParams{
			Rows:   rows,
			Colums: mergerequests.GetTableColums(m.ctx.Window.Width),
			StyleFunc: table.StyleIconsColumns(
				table.Styles(table.DefaultStyle()),
				mergerequests.IconCols(),
			),
			Height: m.ctx.PanelHeight - mainPanelHeaderHeight,
		})
	}
}

func (m Model) getMergeRequestDetails(msg task.TaskMsg) func() details.MergeRequestDetails {
	return func() details.MergeRequestDetails {
		mrMsg := msg.Msg.(message.MergeRequestFetchedMsg)

		return details.MergeRequestDetails{
			Pipelines:   mrMsg.Stages,
			Discussions: mrMsg.Discussions,
			Branches:    mrMsg.Branches,
			Approvals:   mrMsg.Approvals,
		}
	}
}

func (m Model) acceptMergeRequest() tea.Cmd {
	m.SelectMR()
	return m.MergeRequests.AcceptMergeRequest()
}

func (m Model) fetchMergeRequestsList() tea.Cmd {
	m.Projects.SelectProject()
	return m.Projects.GetListCmd()
}

func (m Model) fetchSingleMergeRequest() tea.Cmd {
	m.SelectMR()
	return m.MergeRequests.FetchMergeRequest()
}
