package gql

type CreateNote struct {
	CreateNote CreateNoteResponse `graphql:"createNote(input:{noteableId:$noteableId,discussionId:$discussionId,body:$body})"`
}

type CreateNoteResponse struct {
	Errors []string
}
