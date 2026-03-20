package mergerequests

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui"
)

func (m *Model) FetchMergeRequest() tea.Cmd {
	return func() tea.Msg {
		mr, err := m.client.GetMergeRequest(m.ctx.SelectedProject.ID, gitlab.MergeRequestQueryVariables{
			MRIID: m.ctx.SelectedMR.IID,
		})

		var discussions []gitlab.DiscussionNode
		for _, item := range mr.Discussions.Nodes {
			discussions = append(discussions, item)
		}

		return tui.MRDetailsFetchedMsg{
			MRId:        mr.Id,
			Discussions: discussions,
			Stages:      mr.HeadPipeline.Stages.Nodes,
			Branches:    [2]string{mr.SourceBranch, mr.TargetBranch},
			Approvals:   mr.ApprovalState.Rules,
			Err:         err,
		}
	}
}

func (m *Model) AcceptMergeRequest() tea.Cmd {
	if strings.ToLower(m.ctx.SelectedMR.Status) != "mergeable" {
		return func() tea.Msg {
			return tui.MRMergedMsg{
				Err: fmt.Errorf(
					"Mr can't be merged, its current status is: %s",
					m.ctx.SelectedMR.Status,
				),
			}
		}
	}

	return func() tea.Msg {
		res, err := m.client.AcceptMergeRequest(m.ctx.SelectedProject.ID, gitlab.MergeRequestAcceptInput{
			Sha:                      m.ctx.SelectedMR.Sha,
			IID:                      m.ctx.SelectedMR.IID,
			Squash:                   true,
			ShouldRemoveSourceBranch: true,
		})

		return tui.MRMergedMsg{
			Errors: res.Errors,
			Err:    err,
		}
	}
}

func (m *Model) CreateNote(input gitlab.CreateNoteInput) tea.Cmd {
	return func() tea.Msg {
		res, err := m.client.CreateNote(input)
		return tui.NoteCreatedMsg{
			Errors: res.Errors,
			Err:    err,
		}
	}
}
