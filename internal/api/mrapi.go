package api

import (
	"context"
	"slices"

	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/data"
	"github.com/felipeospina21/mrglab/internal/gql"
)

func GetProjectMergeRequestsGQL(projectID string, opts gql.MergeRequestOptions) (gql.MergeRequestConnection, error) {
	cfg := &config.GlobalConfig

	if cfg.DevMode {
		return data.GQLMergeRequestMock, nil
	}

	var query gql.GetProjectMrs
	configProjects := config.GlobalConfig.Filters.Projects
	projectIdx := slices.IndexFunc(configProjects, func(p config.Project) bool {
		return p.ID == projectID
	})

	opts.FullPaths = []string{configProjects[projectIdx].FullPath}

	variables := gql.GetMRVariables(opts)

	client := newClient()

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return gql.MergeRequestConnection{}, err
	}

	// Since it is filtering by one project always return the first result
	return query.Projects.Edges[0].Node.MergeRequests, nil
}
