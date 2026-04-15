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

// RetryPipeline returns a command that retries failed jobs in a pipeline.
func (m *Model) RetryPipeline(pipelineID string) tea.Cmd {
	return func() tea.Msg {
		res, err := m.client.RetryPipeline(pipelineID)
		return tui.PipelineRetryMsg{
			Errors: res.Errors,
			Err:    err,
		}
	}
}

// PlayJob returns a command that triggers a manual CI job.
func (m *Model) PlayJob(jobID string) tea.Cmd {
	return func() tea.Msg {
		res, err := m.client.PlayJob(jobID)
		return tui.JobPlayMsg{
			Errors: res.Errors,
			Err:    err,
		}
	}
}

// RetryJob returns a command that retries a CI job.
func (m *Model) RetryJob(jobID string) tea.Cmd {
	return func() tea.Msg {
		res, err := m.client.RetryJob(jobID)
		return tui.JobRetryMsg{
			Errors: res.Errors,
			Err:    err,
		}
	}
}
