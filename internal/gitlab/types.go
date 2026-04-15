// Package gitlab provides a GraphQL client for the GitLab API,
// including queries and mutations for merge requests, notes, and pipelines.
package gitlab

import "time"

// Author represents a GitLab user who authored a resource.
type Author struct {
	Name string
}

// Labels holds the label metadata attached to a merge request.
type Labels struct {
	Count int
	Edges []struct {
		Node struct {
			Color     string
			Title     string
			TextColor string
			ID        string
		}
	}
}

// MergeRequestConnection is the paginated list of merge requests returned by the API.
type MergeRequestConnection struct {
	Count int
	Edges []MergeRequestEdge
}

// MergeRequestEdge wraps a single merge request node with its cursor for pagination.
type MergeRequestEdge struct {
	Cursor string
	Node   MergeRequestNode
}

// MergeRequestNode contains the core fields of a merge request.
type MergeRequestNode struct {
	ApprovalsRequired   int
	ApprovalState       MergeRequestApprovalState
	Author              Author
	Conflicts           bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DiffHeadSha         string
	Description         string
	DetailedMergeStatus string
	DiffStatsSummary    DiffStatsSummary
	Draft               bool
	HeadPipeline        *HeadPipelineStatus
	IID                 string
	Labels              Labels
	Title               string
	UserNotesCount      int
	WebURL              string
}

// HeadPipelineStatus holds the overall pipeline status for a merge request.
type HeadPipelineStatus struct {
	Status string
}

// DiffStatsSummary holds the diff statistics for a merge request.
type DiffStatsSummary struct {
	Additions int
	Changes   int
	Deletions int
	FileCount int
}

// PageInfo contains pagination cursors for GraphQL connections.
type PageInfo struct {
	StartCursor string
	EndCursor   string
	HasNextPage bool
}

// MergeRequestResponse is the detailed response for a single merge request query.
type MergeRequestResponse struct {
	Id            string
	SourceBranch  string
	TargetBranch  string
	ApprovalState MergeRequestApprovalState
	Discussions   MergeRequestDiscussionsConnection
	HeadPipeline  MergeRequestHeadPipelineConnection
}

// MergeRequestApprovalState holds the approval rules for a merge request.
type MergeRequestApprovalState struct {
	Rules []ApprovalRule
}

// ApprovalRule represents a single approval rule and its current state.
type ApprovalRule struct {
	Name              string
	ApprovalsRequired int
	Approved          bool
	ApprovedBy        ApprovedBy
}

// ApprovedBy contains the list of users who approved a rule.
type ApprovedBy struct {
	Nodes []ApprovedByNode
}

// ApprovedByNode represents a user who approved a merge request.
type ApprovedByNode struct {
	Name string
}

// MergeRequestDiscussionsConnection holds the discussions on a merge request.
type MergeRequestDiscussionsConnection struct {
	Nodes []DiscussionNode
}

// DiscussionNode represents a single discussion thread.
type DiscussionNode struct {
	Id         string
	Resolvable bool
	Resolved   bool
	ResolvedAt time.Time
	Notes      NoteConnection
}

// NoteConnection holds the notes (comments) within a discussion.
type NoteConnection struct {
	Nodes []Note
}

// Note represents a single comment in a discussion.
type Note struct {
	Author     Author
	Body       string
	CreatedAt  time.Time
	Resolvable bool
}

// MergeRequestHeadPipelineConnection holds the head pipeline for a merge request.
type MergeRequestHeadPipelineConnection struct {
	Stages CiStageConnection
}

// CiStageConnection holds the CI pipeline stages.
type CiStageConnection struct {
	Nodes []CiStageNode
}

// CiStageNode represents a single CI pipeline stage.
type CiStageNode struct {
	Name   string
	Status string
	Jobs   JobsConnection
}

// JobsConnection holds the jobs within a CI stage.
type JobsConnection struct {
	Nodes []JobsNode
}

// JobsNode represents a single CI job.
type JobsNode struct {
	Name     string
	Status   string
	Duration int
}

// PipelineConnection is the paginated list of pipelines returned by the API.
type PipelineConnection struct {
	Count int
	Nodes []PipelineNode
}

// PipelineNode contains the core fields of a pipeline.
type PipelineNode struct {
	ID           string
	IID          string
	Path         string
	Commit       PipelineCommit
	MergeRequest *PipelineMergeRequest
	Jobs         PipelineJobsConnection
	Status       string
	Source       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	FinishedAt   *time.Time
	Duration     *int
	User         Author
	Latest       bool
}

// PipelineCommit holds the commit info for a pipeline.
type PipelineCommit struct {
	ShortId string
	Title   string
}

// PipelineMergeRequest holds the merge request reference for a pipeline.
type PipelineMergeRequest struct {
	IID          string
	SourceBranch string
}

// PipelineJobsConnection holds the jobs within a pipeline.
type PipelineJobsConnection struct {
	Count int
	Nodes []PipelineJobNode
}

// PipelineJobNode represents a single CI job in a pipeline.
type PipelineJobNode struct {
	ID      string
	Name    string
	Status  string
	Retried bool
	Stage   PipelineJobStage
}

// PipelineJobStage holds the stage info for a pipeline job.
type PipelineJobStage struct {
	Name string
}

// AcceptMergeRequestResponse is the result of a merge request accept mutation.
type AcceptMergeRequestResponse struct {
	ClientMutationId string
	Errors           []string
}

// CreateNoteResponse is the result of a create note mutation.
type CreateNoteResponse struct {
	Errors []string
}

// CreateMergeRequestResponse is the result of a create merge request mutation.
type CreateMergeRequestResponse struct {
	Errors []string
}

// CiPipelineID is a typed GraphQL ID for CI pipelines.
type CiPipelineID string

// GetGraphQLType returns the GraphQL type name for CiPipelineID.
func (CiPipelineID) GetGraphQLType() string { return "CiPipelineID" }

// PipelineRetryResponse is the result of a pipeline retry mutation.
type PipelineRetryResponse struct {
	Errors []string
}

// CiBuildID is a typed GraphQL ID for CI jobs.
type CiBuildID string

// GetGraphQLType returns the GraphQL type name for CiBuildID.
func (CiBuildID) GetGraphQLType() string { return "CiBuildID" }

// JobPlayResponse is the result of a job play mutation.
type JobPlayResponse struct {
	Errors []string
}

// JobRetryResponse is the result of a job retry mutation.
type JobRetryResponse struct {
	Errors []string
}

// PipelineCancelResponse is the result of a pipeline cancel mutation.
type PipelineCancelResponse struct {
	Errors []string
}

// JobCancelResponse is the result of a job cancel mutation.
type JobCancelResponse struct {
	Errors []string
}
