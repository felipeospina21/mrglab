package message

import "github.com/felipeospina21/mrglab/internal/gql"

type MergeRequestsFetchedMsg struct {
	Mrs gql.MergeRequestConnection
}

type MergeRequestNotesFetchedMsg struct {
	Discussions []gql.DiscussionNode
}

type MergeRequestPipelineFetchedMsg struct {
	Stages []gql.CiStageNode
}
