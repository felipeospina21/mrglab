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
	if c.demoMode {
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
	if c.demoMode {
		c.sleep(500 * time.Millisecond)
		return PipelineRetryResponse{}, nil
	}

	var mutation pipelineRetryMutation
	variables := map[string]any{
		"id": CiPipelineID(id),
	}

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return PipelineRetryResponse{}, err
	}

	return mutation.PipelineRetry, nil
}

// PlayJob triggers a manual CI job.
func (c *Client) PlayJob(id string) (*JobPlayResponse, error) {
	if c.demoMode {
		c.sleep(500 * time.Millisecond)
		return &JobPlayResponse{}, nil
	}

	var mutation jobPlayMutation
	variables := map[string]any{
		"id": CiProcessableID(id),
	}

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return &JobPlayResponse{}, err
	}

	return &mutation.JobPlay, nil
}

// RetryJob retries a CI job.
func (c *Client) RetryJob(id string) (*JobRetryResponse, error) {
	if c.demoMode {
		c.sleep(500 * time.Millisecond)
		return &JobRetryResponse{}, nil
	}

	var mutation jobRetryMutation
	variables := map[string]any{
		"id": CiBuildID(id),
	}

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return &JobRetryResponse{}, err
	}

	return &mutation.JobRetry, nil
}

// CancelPipeline cancels a running pipeline.
func (c *Client) CancelPipeline(id string) (PipelineCancelResponse, error) {
	if c.demoMode {
		c.sleep(500 * time.Millisecond)
		return PipelineCancelResponse{}, nil
	}

	var mutation pipelineCancelMutation
	variables := map[string]any{
		"id": CiPipelineID(id),
	}

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return PipelineCancelResponse{}, err
	}

	return mutation.PipelineCancel, nil
}

// CancelJob cancels a running CI job.
func (c *Client) CancelJob(id string) (*JobCancelResponse, error) {
	if c.demoMode {
		c.sleep(500 * time.Millisecond)
		return &JobCancelResponse{}, nil
	}

	var mutation jobCancelMutation
	variables := map[string]any{
		"id": CiBuildID(id),
	}

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return &JobCancelResponse{}, err
	}

	return &mutation.JobCancel, nil
}
