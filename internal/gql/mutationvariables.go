package gql

import "github.com/hasura/go-graphql-client"

type MergeRequestAcceptInput struct {
	ProjectPath              graphql.ID
	IID                      string
	Sha                      string
	ShouldRemoveSourceBranch bool
	Squash                   bool
}

func AcceptMergeRequestVariables(input MergeRequestAcceptInput) map[string]any {
	return map[string]any{
		"sha":                      input.Sha,
		"iid":                      input.IID,
		"projectPath":              input.ProjectPath,
		"shouldRemoveSourceBranch": input.ShouldRemoveSourceBranch,
		"squash":                   input.Squash,
	}
}
