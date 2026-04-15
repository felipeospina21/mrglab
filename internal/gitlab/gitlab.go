package gitlab

// GitLabAPI defines the interface for GitLab API operations.
// Consumers should depend on this interface to enable testing with mocks.
type GitLabAPI interface {
	GetProjectMergeRequests(projectID string, vars MergeRequestsQueryVariables) (MergeRequestConnection, error)
	GetMergeRequest(projectID string, vars MergeRequestQueryVariables) (MergeRequestResponse, error)
	GetProjectInfo(projectID string) (ProjectInfo, error)
	GetMRDescriptionTemplates(projectID string) ([]MRDescriptionTemplate, error)
	AcceptMergeRequest(projectID string, input MergeRequestAcceptInput) (AcceptMergeRequestResponse, error)
	CreateNote(input CreateNoteInput) (CreateNoteResponse, error)
	CreateMergeRequest(projectID string, input CreateMergeRequestInput) (CreateMergeRequestResponse, error)
	GetProjectPipelines(projectID string, vars PipelinesQueryVariables) (PipelineConnection, error)
	RetryPipeline(id string) (PipelineRetryResponse, error)
	PlayJob(id string) (*JobPlayResponse, error)
}

// Compile-time check that *Client satisfies GitLabAPI.
var _ GitLabAPI = (*Client)(nil)
