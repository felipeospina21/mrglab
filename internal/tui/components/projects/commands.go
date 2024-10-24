package projects

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/api"
	"github.com/felipeospina21/mrglab/internal/gql"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m *Model) GetListCmd() tea.Cmd {
	return func() tea.Msg {
		mrs, err := api.GetProjectMergeRequestsGQL(m.ctx.SelectedProject.ID, gql.MergeRequestsQueryVariables{
			State: "opened",
		})

		return task.TaskFinishedMsg{
			TaskID:      task.FetchMRs,
			SectionType: task.TaskSectionMR,
			Err:         err,
			Msg: message.MergeRequestsFetchedMsg{
				Mrs: mrs,
			},
		}
	}
}
