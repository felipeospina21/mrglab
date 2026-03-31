package mergerequests

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	execPkg "github.com/felipeospina21/mrglab/internal/exec"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui"
)

// FetchMergeRequest returns a command that fetches a single merge request's details.
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

// AcceptMergeRequest returns a command that merges the selected merge request.
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

// CreateNote returns a command that posts a comment on a discussion.
func (m *Model) CreateNote(input gitlab.CreateNoteInput) tea.Cmd {
	return func() tea.Msg {
		res, err := m.client.CreateNote(input)
		return tui.NoteCreatedMsg{
			Errors: res.Errors,
			Err:    err,
		}
	}
}

// CreateMergeRequest returns a command that creates a new merge request.
func (m *Model) CreateMergeRequest(input gitlab.CreateMergeRequestInput) tea.Cmd {
	return func() tea.Msg {
		res, err := m.client.CreateMergeRequest(m.ctx.SelectedProject.ID, input)
		return tui.MRCreatedMsg{
			Errors: res.Errors,
			Err:    err,
		}
	}
}

// FetchMRTemplates returns a command that fetches MR description templates and branch info.
func (m *Model) FetchMRTemplates() tea.Cmd {
	return func() tea.Msg {
		type tmplResult struct {
			templates []gitlab.MRDescriptionTemplate
			err       error
		}
		type infoResult struct {
			info gitlab.ProjectInfo
			err  error
		}

		tmplCh := make(chan tmplResult, 1)
		infoCh := make(chan infoResult, 1)

		go func() {
			t, err := m.client.GetMRDescriptionTemplates(m.ctx.SelectedProject.ID)
			tmplCh <- tmplResult{t, err}
		}()

		go func() {
			i, err := m.client.GetProjectInfo(m.ctx.SelectedProject.ID)
			infoCh <- infoResult{i, err}
		}()

		tr := <-tmplCh
		ir := <-infoCh

		sourceBranch := execPkg.CurrentGitBranch()

		defaultBranch := "main"
		if ir.err == nil && ir.info.DefaultBranch != "" {
			defaultBranch = ir.info.DefaultBranch
		}

		return tui.MRTemplatesFetchedMsg{
			Templates:     tr.templates,
			DefaultBranch: defaultBranch,
			SourceBranch:  sourceBranch,
			Err:           tr.err,
		}
	}
}
