package gql

import "github.com/hasura/go-graphql-client"

type MergeRequestOptions struct {
	State           string
	ProjectFullPath graphql.ID
	MRIID           string
}

func GetMRVariables(opts MergeRequestOptions) map[string]any {
	return map[string]any{
		"fullPath": opts.ProjectFullPath,
		"mrIID":    opts.MRIID,
		// "state":     graphql.String(opts.State),
	}
}
