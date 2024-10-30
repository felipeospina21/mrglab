package api

import (
	"context"
	"slices"

	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/data"
	"github.com/felipeospina21/mrglab/internal/gql"
	"github.com/hasura/go-graphql-client"
)

func GetProjectMergeRequestsGQL(
	projectID string,
	vars gql.MergeRequestsQueryVariables,
) (gql.MergeRequestConnection, error) {
	cfg := &config.GlobalConfig

	if cfg.DevMode {
		return data.GQLMergeRequestMock, nil
	}

	var query gql.GetProjectMrs
	configProjects := config.GlobalConfig.Filters.Projects
	projectIdx := slices.IndexFunc(configProjects, func(p config.Project) bool {
		return p.ID == projectID
	})

	vars.ProjectFullPath = graphql.ID(configProjects[projectIdx].FullPath)

	variables := gql.GetMergeRequestsVariables(vars)

	client := newClient()

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return gql.MergeRequestConnection{}, err
	}

	return query.Project.MergeRequests, nil
}

func GetMergeRequest(
	projectID string,
	vars gql.MergeRequestQueryVariables,
) (gql.MergeRequestResponse, error) {
	cfg := &config.GlobalConfig

	if cfg.DevMode {
		return data.GQLDiscussionsMock, nil
	}

	var query gql.GetMergeRequest
	configProjects := config.GlobalConfig.Filters.Projects
	projectIdx := slices.IndexFunc(configProjects, func(p config.Project) bool {
		return p.ID == projectID
	})

	vars.ProjectFullPath = graphql.ID(configProjects[projectIdx].FullPath)

	variables := gql.MergeRequestVariables(vars)

	client := newClient()

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return gql.MergeRequestResponse{}, err
	}

	return query.Project.MergeRequest, nil
}
