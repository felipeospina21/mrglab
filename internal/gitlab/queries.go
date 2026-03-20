package gitlab

import "github.com/hasura/go-graphql-client"

// Query structs

type getProjectMrs struct {
	Project project `graphql:"project(fullPath: $fullPath)"`
}

type project struct {
	Name          string
	ID            string
	MergeRequests MergeRequestConnection `graphql:"mergeRequests(state: opened)"`
}

type getMergeRequest struct {
	Project struct {
		MergeRequest MergeRequestResponse `graphql:"mergeRequest(iid: $mrIID)"`
	} `graphql:"project(fullPath: $fullPath)"`
}

// Mutation structs

type acceptMergeRequestMutation struct {
	MergeRequestAccept AcceptMergeRequestResponse `graphql:"mergeRequestAccept(input:{shouldRemoveSourceBranch:$shouldRemoveSourceBranch,squash:$squash,sha:$sha,projectPath:$projectPath,iid:$iid})"`
}

type createNoteMutation struct {
	CreateNote CreateNoteResponse `graphql:"createNote(input:{noteableId:$noteableId,discussionId:$discussionId,body:$body})"`
}

// Input types

type MergeRequestsQueryVariables struct {
	State           string
	ProjectFullPath graphql.ID
}

type MergeRequestQueryVariables struct {
	MRIID string
	MergeRequestsQueryVariables
}

type MergeRequestAcceptInput struct {
	ProjectPath              graphql.ID
	IID                      string
	Sha                      string
	ShouldRemoveSourceBranch bool
	Squash                   bool
}

type NoteableID string

func (NoteableID) GetGraphQLType() string { return "NoteableID" }

type DiscussionID string

func (DiscussionID) GetGraphQLType() string { return "DiscussionID" }

type CreateNoteInput struct {
	NoteableId   NoteableID
	DiscussionId DiscussionID
	Body         string
}

// Variable builders

func mergeRequestsVariables(vars MergeRequestsQueryVariables) map[string]any {
	return map[string]any{
		"fullPath": vars.ProjectFullPath,
	}
}

func mergeRequestVariables(vars MergeRequestQueryVariables) map[string]any {
	return map[string]any{
		"fullPath": vars.ProjectFullPath,
		"mrIID":    vars.MRIID,
	}
}

func acceptMergeRequestVariables(input MergeRequestAcceptInput) map[string]any {
	return map[string]any{
		"sha":                      input.Sha,
		"iid":                      input.IID,
		"projectPath":              input.ProjectPath,
		"shouldRemoveSourceBranch": input.ShouldRemoveSourceBranch,
		"squash":                   input.Squash,
	}
}

func createNoteVariables(input CreateNoteInput) map[string]any {
	return map[string]any{
		"noteableId":   input.NoteableId,
		"discussionId": input.DiscussionId,
		"body":         input.Body,
	}
}
