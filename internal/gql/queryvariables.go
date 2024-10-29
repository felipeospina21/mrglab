package gql

import "github.com/hasura/go-graphql-client"

type MergeRequestsQueryVariables struct {
	State           string
	ProjectFullPath graphql.ID
}

type MergeRequestQueryVariables struct {
	MRIID string
	MergeRequestsQueryVariables
}

func GetMergeRequestsVariables(vars MergeRequestsQueryVariables) map[string]any {
	return map[string]any{
		"fullPath": vars.ProjectFullPath,
		// "state":     graphql.String(opts.State),
	}
}

func MergeRequestVariables(vars MergeRequestQueryVariables) map[string]any {
	return map[string]any{
		"fullPath": vars.ProjectFullPath,
		"mrIID":    vars.MRIID,
		// "state":     graphql.String(opts.State),
	}
}
