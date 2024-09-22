package mergerequests

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/api"
	"github.com/felipeospina21/mrglab/internal/tui/task"
	"github.com/xanzy/go-gitlab"
)

func (m *Model) GetMRListCmd() tea.Cmd {
	return func() tea.Msg {
		p := &gitlab.ListProjectMergeRequestsOptions{
			State: gitlab.Ptr("opened"),
		}

		mrs, err := api.GetProjectMergeRequests(m.ctx.SelectedProject.ID, p)
		if err != nil {
			return err
		}
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
