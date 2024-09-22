package mergerequests

import "github.com/xanzy/go-gitlab"

type MergeRequestsFetchedMsg struct {
	Mrs    []*gitlab.MergeRequest
	TaskId string
}
