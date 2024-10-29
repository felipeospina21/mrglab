package gql

type GetMrPipeline struct {
	Project struct {
		MergeRequest struct {
			Id           string
			HeadPipeline MergeRequestHeadPipelineConnection
		} `graphql:"mergeRequest(iid: $mrIID)"`
	} `graphql:"project(fullPath: $fullPath)"`
}

type MergeRequestHeadPipelineConnection struct {
	Stages CiStageConnection
}

type CiStageConnection struct {
	Nodes []CiStageNode
}

type CiStageNode struct {
	Name   string
	Status string
	Jobs   JobsConnection
}

type JobsConnection struct {
	Nodes JobsNode
}

type JobsNode struct {
	Name     string
	Status   string
	Duration int
}
