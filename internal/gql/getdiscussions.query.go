package gql

import "time"

type GetMrDiscussions struct {
	Project struct {
		MergeRequest struct {
			Id    string
			Notes MergeRequestNotesConnection `graphql:"notes(filter: ONLY_COMMENTS)"`
		} `graphql:"mergeRequest(iid: $mrIID)"`
	} `graphql:"project(fullPath: $fullPath)"`
}

type MergeRequestNotesConnection struct {
	Count int
	Nodes []DiscussionNode
}

type DiscussionNode struct {
	Discussion Discussion
}

type Discussion struct {
	Notes NoteConnection
}

type NoteConnection struct {
	Count int
	Nodes []Note
}

type Note struct {
	Author    Author
	Body      string
	CreatedAt time.Time
	Resolved  bool
}
