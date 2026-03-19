package gitlab

import "context"

func (c *Client) GetProjectMergeRequests(
	projectID string,
	vars MergeRequestsQueryVariables,
) (MergeRequestConnection, error) {
	if c.devMode {
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

func (c *Client) GetMergeRequest(
	projectID string,
	vars MergeRequestQueryVariables,
) (MergeRequestResponse, error) {
	if c.devMode {
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

func (c *Client) AcceptMergeRequest(
	projectID string,
	input MergeRequestAcceptInput,
) (AcceptMergeRequestResponse, error) {
	if c.devMode {
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
