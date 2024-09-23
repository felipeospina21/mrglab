package api

import (
	"log"

	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/xanzy/go-gitlab"
)

func GetProjectMergeRequests(projectID string, opts *gitlab.ListProjectMergeRequestsOptions) ([]*gitlab.MergeRequest, error) {
	cfg := &config.GlobalConfig
	url := buildURL(cfg)

	client, err := gitlab.NewClient(
		cfg.APIToken,
		gitlab.WithBaseURL(url),
	)
	if err != nil {
		// TODO: handle error
		log.Fatal(err)
	}

	l, _, err := client.MergeRequests.ListProjectMergeRequests(projectID, opts)

	return l, err
}
