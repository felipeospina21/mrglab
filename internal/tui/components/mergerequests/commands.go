package mergerequests

import (
	"errors"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/api"
	"github.com/felipeospina21/mrglab/internal/gql"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m *Model) FetchMergeRequest() tea.Cmd {
	return func() tea.Msg {
		mr, err := api.GetMergeRequest(m.ctx.SelectedProject.ID, gql.MergeRequestQueryVariables{
			MRIID: m.ctx.SelectedMR.IID,
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
				Branches:    [2]string{mr.SourceBranch, mr.TargetBranch},
				Approvals:   mr.ApprovalState.Rules,
			},
		}
	}
}

func (m *Model) AcceptMergeRequest() tea.Cmd {
	if m.ctx.SelectedMR.Status != strings.ToLower("mergeable") {
		return func() tea.Msg {
			return task.TaskMsg{
				TaskID:      task.MergeMR,
				SectionType: task.TaskSectionMR,
				Msg:         gql.AcceptMergeRequestResponse{},
				Err: errors.New(fmt.Sprintf(
					"Mr can't be merged, its current status is: %s",
					m.ctx.SelectedMR.Status,
				)),
			}
		}
	}

	return func() tea.Msg {
		res, err := api.AcceptMergeRequest(m.ctx.SelectedProject.ID, gql.MergeRequestAcceptInput{
			Sha:                      m.ctx.SelectedMR.Sha,
			IID:                      m.ctx.SelectedMR.IID,
			Squash:                   true,
			ShouldRemoveSourceBranch: true,
		})

		return task.TaskMsg{
			TaskID:      task.MergeMR,
			SectionType: task.TaskSectionMR,
			Err:         err,
			Msg:         res,
		}
	}
}
