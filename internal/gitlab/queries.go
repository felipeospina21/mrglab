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

// MergeRequestsQueryVariables holds the variables for the merge requests list query.
type MergeRequestsQueryVariables struct {
	State           string
	ProjectFullPath graphql.ID
}

// MergeRequestQueryVariables holds the variables for a single merge request query.
type MergeRequestQueryVariables struct {
	MRIID string
	MergeRequestsQueryVariables
}

// MergeRequestAcceptInput holds the input for the accept merge request mutation.
type MergeRequestAcceptInput struct {
	ProjectPath              graphql.ID
	IID                      string
	Sha                      string
	ShouldRemoveSourceBranch bool
	Squash                   bool
}

// NoteableID is a typed GraphQL ID for noteable resources.
type NoteableID string

// GetGraphQLType returns the GraphQL type name for NoteableID.
func (NoteableID) GetGraphQLType() string { return "NoteableID" }

// DiscussionID is a typed GraphQL ID for discussions.
type DiscussionID string

// GetGraphQLType returns the GraphQL type name for DiscussionID.
func (DiscussionID) GetGraphQLType() string { return "DiscussionID" }

// CreateNoteInput holds the input for creating a note on a discussion.
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
