package mergerequests

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/api"
	"github.com/felipeospina21/mrglab/internal/gql"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m *Model) FetchMergeRequest() tea.Cmd {
	return func() tea.Msg {
		mr, err := api.GetMergeRequest(m.ctx.SelectedProject.ID, gql.MergeRequestQueryVariables{
			MRIID: m.ctx.SelectedMRID,
		})

		var discussions []gql.DiscussionNode
		for _, item := range mr.Discussions.Nodes {
			discussions = append(discussions, item)
		}

		return task.TaskMsg{
			TaskID:      task.FetchDiscussions,
			SectionType: task.TaskSectionMR,
			Err:         err,
			Msg: message.MergeRequestFetchedMsg{
				Discussions: discussions,
				Stages:      mr.HeadPipeline.Stages.Nodes,
			},
		}
	}
}
