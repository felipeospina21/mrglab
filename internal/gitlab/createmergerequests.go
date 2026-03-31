package gitlab

import (
	"context"
	"time"
)

// CreateMergeRequest creates a new merge request via the GitLab API.
func (c *Client) CreateMergeRequest(
	projectID string,
	input CreateMergeRequestInput,
) (CreateMergeRequestResponse, error) {
	if c.devMode {
		c.sleep(500 * time.Millisecond)
		return CreateMergeRequestResponse{}, nil
	}

	var mutation createMergeRequestMutation
	input.ProjectPath = c.projectFullPath(projectID)
	variables := createMergeRequestVariables(input)

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return CreateMergeRequestResponse{}, err
	}

	return mutation.MergeRequestCreate, nil
}
