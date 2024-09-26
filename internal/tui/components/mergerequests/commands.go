package mergerequests

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/api"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
	"github.com/xanzy/go-gitlab"
)

func (m *Model) GetListCmd() tea.Cmd {
	return func() tea.Msg {
		p := &gitlab.ListProjectMergeRequestsOptions{
			State: gitlab.Ptr("opened"),
		}

		mrs, err := api.GetProjectMergeRequests(m.ctx.SelectedProject.ID, p)

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
		m.ctx.IsLeftPanelOpen = false
		rows := getTableRows(msg)
		return table.InitModel(table.InitModelParams{
			Rows:   rows,
			Colums: GetTableColums(m.ctx.Window.Width - 10),
			StyleFunc: table.StyleIconsColumns(
				table.Styles(table.DefaultStyle()),
				IconCols(),
			),
		})
	}
}
