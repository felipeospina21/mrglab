package projects

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/api"
	"github.com/felipeospina21/mrglab/internal/gql"
	"github.com/felipeospina21/mrglab/internal/tui"
)

func (m *Model) GetListCmd() tea.Cmd {
	return func() tea.Msg {
		mrs, err := api.GetProjectMergeRequestsGQL(m.ctx.SelectedProject.ID, gql.MergeRequestsQueryVariables{
			State: "opened",
		})

		return tui.MRListFetchedMsg{
			Mrs: mrs,
			Err: err,
		}
	}
}
