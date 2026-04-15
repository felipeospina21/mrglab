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

type getRepositoryTree struct {
	Project struct {
		Repository struct {
			Tree struct {
				Blobs struct {
					Nodes []struct {
						Name string
						Path string
					}
				}
			} `graphql:"tree(path: $path, ref: $ref)"`
		}
	} `graphql:"project(fullPath: $fullPath)"`
}

type getRepositoryBlobs struct {
	Project struct {
		Repository struct {
			Blobs struct {
				Nodes []struct {
					Name        string
					RawTextBlob string
				}
			} `graphql:"blobs(paths: $paths, ref: $ref)"`
		}
	} `graphql:"project(fullPath: $fullPath)"`
}

type getProjectPipelines struct {
	Project struct {
		Pipelines PipelineConnection `graphql:"pipelines(first: 30)"`
	} `graphql:"project(fullPath: $fullPath)"`
}

// Mutation structs

type acceptMergeRequestMutation struct {
	MergeRequestAccept AcceptMergeRequestResponse `graphql:"mergeRequestAccept(input:{shouldRemoveSourceBranch:$shouldRemoveSourceBranch,squash:$squash,sha:$sha,projectPath:$projectPath,iid:$iid})"`
}

type createNoteMutation struct {
	CreateNote CreateNoteResponse `graphql:"createNote(input:{noteableId:$noteableId,discussionId:$discussionId,body:$body})"`
}

type createMergeRequestMutation struct {
	MergeRequestCreate CreateMergeRequestResponse `graphql:"mergeRequestCreate(input:{projectPath:$projectPath,sourceBranch:$sourceBranch,targetBranch:$targetBranch,title:$title,description:$description})"`
}

type pipelineRetryMutation struct {
	PipelineRetry PipelineRetryResponse `graphql:"pipelineRetry(input:{id:$id})"`
}

type jobPlayMutation struct {
	JobPlay JobPlayResponse `graphql:"jobPlay(input:{id:$id})"`
}

type jobRetryMutation struct {
	JobRetry JobRetryResponse `graphql:"jobRetry(input:{id:$id})"`
}

type pipelineCancelMutation struct {
	PipelineCancel PipelineCancelResponse `graphql:"pipelineCancel(input:{id:$id})"`
}

type jobCancelMutation struct {
	JobCancel JobCancelResponse `graphql:"jobCancel(input:{id:$id})"`
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

// CreateMergeRequestInput holds the input for the create merge request mutation.
type CreateMergeRequestInput struct {
	ProjectPath  graphql.ID
	SourceBranch string
	TargetBranch string
	Title        string
	Description  string
}

// PipelinesQueryVariables holds the variables for the pipelines list query.
type PipelinesQueryVariables struct {
	ProjectFullPath graphql.ID
}

// Variable builders

func pipelinesVariables(vars PipelinesQueryVariables) map[string]any {
	return map[string]any{
		"fullPath": vars.ProjectFullPath,
	}
}

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

func createMergeRequestVariables(input CreateMergeRequestInput) map[string]any {
	return map[string]any{
		"projectPath":  input.ProjectPath,
		"sourceBranch": input.SourceBranch,
		"targetBranch": input.TargetBranch,
		"title":        input.Title,
		"description":  input.Description,
	}
}
