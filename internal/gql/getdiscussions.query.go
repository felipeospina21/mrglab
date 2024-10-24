package gql

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
	Edges []MergeRequestNotesEdges
}

type MergeRequestNotesEdges struct {
	Node MergeRequestNotesNode
}

type MergeRequestNotesNode struct {
	Id         string
	Discussion MergeRequestDiscussion
}

type MergeRequestDiscussion struct {
	Notes NoteConnection
}

type NoteConnection struct {
	Count int
	Edges []NoteEdges
}

type NoteEdges struct {
	Node Note
}

type Note struct {
	Author Author
	Body   string
}
