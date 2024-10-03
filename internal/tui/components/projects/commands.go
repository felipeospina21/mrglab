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
		mrs, err := api.GetProjectMergeRequestsGQL(m.ctx.SelectedProject.ID, gql.MergeRequestOptions{
			State: "opened",
		})

		return task.TaskFinishedMsg{
			TaskID:      "",
			SectionID:   0,
			SectionType: "mrs",
			Err:         err,
			Msg: message.MergeRequestsFetchedMsg{
				Mrs:    mrs,
				TaskId: "fetch_Mrs",
			},
		}
	}
}
