package mergerequests

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/api"
	"github.com/felipeospina21/mrglab/internal/gql"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m *Model) GetMRNotesCmd() tea.Cmd {
	return func() tea.Msg {
		d, err := api.GetMergeRequestDiscussions(m.ctx.SelectedProject.ID, gql.NotesQueryVariables{
			MRIID: m.ctx.SelectedMRID,
		})

		var discussions []gql.DiscussionNode
		for _, item := range d.Nodes {
			discussions = append(discussions, item)
		}

		return task.TaskMsg{
			TaskID:      task.FetchDiscussions,
			SectionType: task.TaskSectionMR,
			Err:         err,
			Msg: message.MergeRequestNotesFetchedMsg{
				Discussions: discussions,
			},
		}
	}
}
