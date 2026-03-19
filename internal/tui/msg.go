package tui

import "github.com/felipeospina21/mrglab/internal/gitlab"

// Task state tracking
type TaskStatus uint

const (
	TaskIdle TaskStatus = iota
	TaskStarted
	TaskFinished
)

// Typed messages — replace the generic TaskMsg wrapper

type MRListFetchedMsg struct {
	Mrs gitlab.MergeRequestConnection
	Err error
}

type MRDetailsFetchedMsg struct {
	Discussions []gitlab.DiscussionNode
	Stages      []gitlab.CiStageNode
	Branches    [2]string
	Approvals   []gitlab.ApprovalRule
	Err         error
}

type MRMergedMsg struct {
	Errors []string
	Err    error
}
