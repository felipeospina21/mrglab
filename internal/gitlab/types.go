package gitlab

import "time"

type Author struct {
	Name string
}

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

type MergeRequestConnection struct {
	Count int
	Edges []MergeRequestEdge
}

type MergeRequestEdge struct {
	Cursor string
	Node   MergeRequestNode
}

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
	IID                 string
	Labels              Labels
	Title               string
	UserNotesCount      int
	WebURL              string
}

type DiffStatsSummary struct {
	Additions int
	Changes   int
	Deletions int
	FileCount int
}

type PageInfo struct {
	StartCursor string
	EndCursor   string
	HasNextPage bool
}

type MergeRequestResponse struct {
	Id            string
	SourceBranch  string
	TargetBranch  string
	ApprovalState MergeRequestApprovalState
	Discussions   MergeRequestDiscussionsConnection
	HeadPipeline  MergeRequestHeadPipelineConnection
}

type MergeRequestApprovalState struct {
	Rules []ApprovalRule
}

type ApprovalRule struct {
	Name              string
	ApprovalsRequired int
	Approved          bool
	ApprovedBy        ApprovedBy
}

type ApprovedBy struct {
	Nodes []ApprovedByNode
}

type ApprovedByNode struct {
	Name string
}

type MergeRequestDiscussionsConnection struct {
	Nodes []DiscussionNode
}

type DiscussionNode struct {
	Resolvable bool
	Resolved   bool
	ResolvedAt time.Time
	Notes      NoteConnection
}

type NoteConnection struct {
	Nodes []Note
}

type Note struct {
	Author     Author
	Body       string
	CreatedAt  time.Time
	Resolvable bool
}

type MergeRequestHeadPipelineConnection struct {
	Stages CiStageConnection
}

type CiStageConnection struct {
	Nodes []CiStageNode
}

type CiStageNode struct {
	Name   string
	Status string
	Jobs   JobsConnection
}

type JobsConnection struct {
	Nodes []JobsNode
}

type JobsNode struct {
	Name     string
	Status   string
	Duration int
}

type AcceptMergeRequestResponse struct {
	ClientMutationId string
	Errors           []string
}

type CreateNoteResponse struct {
	Errors []string
}
