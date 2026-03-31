package projects

import (
	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui"
)

func (m *Model) GetListCmd() tea.Cmd {
	return func() tea.Msg {
		mrs, err := m.client.GetProjectMergeRequests(m.ctx.SelectedProject.ID, gitlab.MergeRequestsQueryVariables{
			State: "opened",
		})

		return tui.MRListFetchedMsg{
			Mrs: mrs,
			Err: err,
		}
	}
}
