package gql

import "time"

type GetProjectMrs struct {
	Project Project `graphql:"project(fullPath: $fullPath)"`
}

type Project struct {
	Name string
	ID   string
	// TODO: check how to have multiple MergeRequests values each with different states
	// ex: MergeRequests MergeRequestConnection `graphql:"mergeRequests(state: closed)"`
	MergeRequests MergeRequestConnection `graphql:"mergeRequests(state: opened)"`
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
