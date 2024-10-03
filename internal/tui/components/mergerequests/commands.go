package mergerequests

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/api"
	"github.com/felipeospina21/mrglab/internal/gql"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m *Model) GetListCmd() tea.Cmd {
	return func() tea.Msg {
		mrs, err := api.GetProjectMergeRequestsGQL(m.ctx.SelectedProject.ID, gql.MergeRequestOptions{
			State: "opened",
		})

		return task.TaskFinishedMsg{
			TaskID:      "",
			SectionID:   0,
			SectionType: "mrs",
			Err:         err,
			Msg: MergeRequestsFetchedMsg{
				Mrs:    mrs,
				TaskId: "fetch_Mrs",
			},
		}
	}
}

func (m Model) GetTableModel(msg task.TaskFinishedMsg) func() table.Model {
	return func() table.Model {
		rows := getTableRows(msg)
		mainPanelHeaderHeight := 1
		return table.InitModel(table.InitModelParams{
			Rows:   rows,
			Colums: GetTableColums(m.ctx.Window.Width),
			StyleFunc: table.StyleIconsColumns(
				table.Styles(table.DefaultStyle()),
				IconCols(),
			),
			Height: m.ctx.PanelHeight - mainPanelHeaderHeight,
		})
	}
}
