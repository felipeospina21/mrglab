package gql

import "time"

type GetMrDiscussions struct {
	Project struct {
		MergeRequest struct {
			Id          string
			Discussions MergeRequestDiscussionsConnection
		} `graphql:"mergeRequest(iid: $mrIID)"`
	} `graphql:"project(fullPath: $fullPath)"`
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
