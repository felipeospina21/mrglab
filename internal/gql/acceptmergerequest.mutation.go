package gql

type AcceptMergeRequest struct {
	MergeRequestAccept AcceptMergeRequestResponse `graphql:"mergeRequestAccept(input:{shouldRemoveSourceBranch:$shouldRemoveSourceBranch,squash:$squash,sha:$sha,projectPath:$projectPath,iid:$iid})"`
}
type AcceptMergeRequestResponse struct {
	ClientMutationId string
	Errors           []string
}
