package app

import (
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) GetMergeRequestModel(msg task.TaskMsg) func() table.Model {
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

func (m Model) GetMergeRequestDetails(msg task.TaskMsg) func() details.MergeRequestDetails {
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
