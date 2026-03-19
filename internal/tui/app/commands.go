package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/exec"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
)

func (m Model) getMergeRequestModel(msg tui.MRListFetchedMsg) func() table.Model {
	return func() table.Model {
		rows := mergerequests.GetTableRows(msg.Mrs)
		tableW := m.layout.MainPanel.Width - table.DocStyle.GetHorizontalFrameSize() - tableBorderX
		return table.InitModel(table.InitModelParams{
			Rows:   rows,
			Colums: mergerequests.GetTableColums(tableW),
			StyleFunc: table.StyleIconsColumns(
				table.Styles(table.DefaultStyle()),
				mergerequests.IconCols(),
			),
			Height: m.layout.ContentH - mainPanelHeaderLines - tableBorderY - tableViewOverhead,
		})
	}
}

func (m Model) getMergeRequestDetails(msg tui.MRDetailsFetchedMsg) func() details.MergeRequestDetails {
	return func() details.MergeRequestDetails {
		return details.MergeRequestDetails{
			Pipelines:   msg.Stages,
			Discussions: msg.Discussions,
			Branches:    msg.Branches,
			Approvals:   msg.Approvals,
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

func (m *Model) openInBrowser() {
	colIdx := mergerequests.GetColIndex(mergerequests.ColNames.URL)
	url := m.MergeRequests.Table.SelectedRow()[colIdx]
	exec.Openbrowser(url)
}
