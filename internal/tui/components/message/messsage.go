package message

import "github.com/felipeospina21/mrglab/internal/gql"

type MergeRequestsListFetchedMsg struct {
	Mrs gql.MergeRequestConnection
}

type MergeRequestFetchedMsg struct {
	Discussions []gql.DiscussionNode
	Stages      []gql.CiStageNode
	Branches    [2]string
	Approvals   []gql.ApprovalRule
}

type MergeRequestMergedMsg = gql.AcceptMergeRequestResponse
