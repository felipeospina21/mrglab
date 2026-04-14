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
