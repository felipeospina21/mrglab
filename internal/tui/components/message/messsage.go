package message

import "github.com/felipeospina21/mrglab/internal/gql"

type MergeRequestsFetchedMsg struct {
	Mrs    gql.MergeRequestConnection
	TaskId string
}
