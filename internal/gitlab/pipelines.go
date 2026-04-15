package gitlab

import (
	"context"
	"time"
)

// GetProjectPipelines fetches the pipelines for a project.
func (c *Client) GetProjectPipelines(
	projectID string,
	vars PipelinesQueryVariables,
) (PipelineConnection, error) {
	if c.devMode {
		c.sleep(1 * time.Second)
		return pipelineConnectionMock, nil
	}

	var query getProjectPipelines
	vars.ProjectFullPath = c.projectFullPath(projectID)
	variables := pipelinesVariables(vars)

	err := c.gql.Query(context.Background(), &query, variables)
	if err != nil {
		return PipelineConnection{}, err
	}

	return query.Project.Pipelines, nil
}
