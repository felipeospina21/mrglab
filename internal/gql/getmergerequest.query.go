package gql

import "time"

type GetMergeRequest struct {
	Project struct {
		MergeRequest MergeRequestResponse `graphql:"mergeRequest(iid: $mrIID)"`
	} `graphql:"project(fullPath: $fullPath)"`
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
	ApprovedBy        struct {
		Nodes []struct {
			Name string
		}
	}
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

// Pipelines
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
