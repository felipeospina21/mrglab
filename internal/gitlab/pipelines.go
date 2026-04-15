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

// RetryPipeline retries all failed jobs in a pipeline.
func (c *Client) RetryPipeline(id string) (PipelineRetryResponse, error) {
	if c.devMode {
		c.sleep(500 * time.Millisecond)
		return PipelineRetryResponse{}, nil
	}

	var mutation pipelineRetryMutation
	variables := pipelineRetryVariables(CiPipelineID(id))

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return PipelineRetryResponse{}, err
	}

	return mutation.PipelineRetry, nil
}
