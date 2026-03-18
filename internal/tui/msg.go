package tui

import "github.com/felipeospina21/mrglab/internal/gql"

// Task state tracking
type TaskStatus uint

const (
	TaskIdle TaskStatus = iota
	TaskStarted
	TaskFinished
)

// Typed messages — replace the generic TaskMsg wrapper

type MRListFetchedMsg struct {
	Mrs gql.MergeRequestConnection
	Err error
}

type MRDetailsFetchedMsg struct {
	Discussions []gql.DiscussionNode
	Stages      []gql.CiStageNode
	Branches    [2]string
	Approvals   []gql.ApprovalRule
	Err         error
}

type MRMergedMsg struct {
	Errors []string
	Err    error
}
