package gql

type GetProjectMrs struct {
	Project Project `graphql:"project(fullPath: $fullPath)"`
}
