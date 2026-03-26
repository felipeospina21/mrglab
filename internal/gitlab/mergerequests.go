package gitlab

import (
	"context"
	"time"
)

// GetProjectMergeRequests fetches the open merge requests for a project.
func (c *Client) GetProjectMergeRequests(
	projectID string,
	vars MergeRequestsQueryVariables,
) (MergeRequestConnection, error) {
	if c.devMode {
		time.Sleep(1 * time.Second)
		return mergeRequestConnectionMock, nil
	}

	var query getProjectMrs
	vars.ProjectFullPath = c.projectFullPath(projectID)
	variables := mergeRequestsVariables(vars)

	err := c.gql.Query(context.Background(), &query, variables)
	if err != nil {
		return MergeRequestConnection{}, err
	}

	return query.Project.MergeRequests, nil
}

// GetMergeRequest fetches a single merge request by IID.
func (c *Client) GetMergeRequest(
	projectID string,
	vars MergeRequestQueryVariables,
) (MergeRequestResponse, error) {
	if c.devMode {
		time.Sleep(800 * time.Millisecond)
		return mergeRequestResponseMock, nil
	}

	var query getMergeRequest
	vars.ProjectFullPath = c.projectFullPath(projectID)
	variables := mergeRequestVariables(vars)

	err := c.gql.Query(context.Background(), &query, variables)
	if err != nil {
		return MergeRequestResponse{}, err
	}

	return query.Project.MergeRequest, nil
}

// AcceptMergeRequest merges a merge request via the GitLab API.
func (c *Client) AcceptMergeRequest(
	projectID string,
	input MergeRequestAcceptInput,
) (AcceptMergeRequestResponse, error) {
	if c.devMode {
		time.Sleep(500 * time.Millisecond)
		return AcceptMergeRequestResponse{}, nil
	}

	var mutation acceptMergeRequestMutation
	input.ProjectPath = c.projectFullPath(projectID)
	variables := acceptMergeRequestVariables(input)

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return AcceptMergeRequestResponse{}, err
	}

	return mutation.MergeRequestAccept, nil
}
