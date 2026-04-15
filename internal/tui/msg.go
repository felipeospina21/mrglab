// Package tui provides shared types, messages, keys, and utilities for the TUI layer.
package tui

import "github.com/felipeospina21/mrglab/internal/gitlab"

// TaskStatus tracks the lifecycle of an async task.
type TaskStatus uint

// Task lifecycle constants.
const (
	TaskIdle TaskStatus = iota
	TaskStarted
	TaskFinished
)

// MRListFetchedMsg is sent when the merge request list has been fetched.
type MRListFetchedMsg struct {
	Mrs gitlab.MergeRequestConnection
	Err error
}

// MRDetailsFetchedMsg is sent when a single merge request's details have been fetched.
type MRDetailsFetchedMsg struct {
	MRId        string
	Discussions []gitlab.DiscussionNode
	Stages      []gitlab.CiStageNode
	Branches    [2]string
	Approvals   []gitlab.ApprovalRule
	Err         error
}

// MRMergedMsg is sent after a merge request accept mutation completes.
type MRMergedMsg struct {
	Errors []string
	Err    error
}

// NoteCreatedMsg is sent after a note (comment) has been created.
type NoteCreatedMsg struct {
	Errors []string
	Err    error
}

// MRCreatedMsg is sent after a merge request has been created.
type MRCreatedMsg struct {
	Errors []string
	Err    error
}

// MRTemplatesFetchedMsg is sent when MR description templates have been fetched.
type MRTemplatesFetchedMsg struct {
	Templates     []gitlab.MRDescriptionTemplate
	DefaultBranch string
	SourceBranch  string
	Err           error
}

// PipelineListFetchedMsg is sent when the pipeline list has been fetched.
type PipelineListFetchedMsg struct {
	Pipelines gitlab.PipelineConnection
	Err       error
}

// PipelineRetryMsg is sent after a pipeline retry mutation completes.
type PipelineRetryMsg struct {
	Errors []string
	Err    error
}

// JobPlayMsg is sent after a job play mutation completes.
type JobPlayMsg struct {
	Errors []string
	Err    error
}

// JobRetryMsg is sent after a job retry mutation completes.
type JobRetryMsg struct {
	Errors []string
	Err    error
}
