package pipelines

import (
	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui"
)

// FetchPipelines returns a command that fetches pipelines for the selected project.
func (m *Model) FetchPipelines() tea.Cmd {
	return func() tea.Msg {
		pipelines, err := m.client.GetProjectPipelines(m.ctx.SelectedProject.ID, gitlab.PipelinesQueryVariables{})
		return tui.PipelineListFetchedMsg{
			Pipelines: pipelines,
			Err:       err,
		}
	}
}
