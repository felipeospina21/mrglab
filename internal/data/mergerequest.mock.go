package data

import (
	"time"

	"github.com/xanzy/go-gitlab"
)

var (
	MockUser          = &gitlab.BasicUser{Name: "Dummy"}
	MergerequestsMock = []*gitlab.MergeRequest{
		{
			IID:                 1,
			Title:               "Mocked Title",
			CreatedAt:           &time.Time{},
			Draft:               true,
			Author:              MockUser,
			DetailedMergeStatus: "draft_status",
			HasConflicts:        false,
			UserNotesCount:      4,
			WebURL:              "",
			Description:         "Dummy description",
		},
		{
			IID:                 2,
			Title:               "Mocked Title 2",
			CreatedAt:           &time.Time{},
			Draft:               false,
			Author:              MockUser,
			DetailedMergeStatus: "conflict",
			HasConflicts:        true,
			UserNotesCount:      0,
			WebURL:              "",
			Description:         "Dummy description",
		},
		{
			IID:                 3,
			Title:               "Mocked Title 3",
			CreatedAt:           &time.Time{},
			Draft:               false,
			Author:              MockUser,
			DetailedMergeStatus: "need_rebase",
			HasConflicts:        true,
			UserNotesCount:      0,
			WebURL:              "",
			Description:         "Dummy description",
		},
		{
			IID:                 4,
			Title:               "Mocked Title 4",
			CreatedAt:           &time.Time{},
			Draft:               true,
			Author:              MockUser,
			DetailedMergeStatus: "mergeable",
			HasConflicts:        false,
			UserNotesCount:      4,
			WebURL:              "",
			Description:         "Dummy description",
		},
		{
			IID:                 5,
			Title:               "Mocked Title 5",
			CreatedAt:           &time.Time{},
			Draft:               false,
			Author:              MockUser,
			DetailedMergeStatus: "blocked_status",
			HasConflicts:        true,
			UserNotesCount:      0,
			WebURL:              "",
			Description:         "Dummy description",
		},
		{
			IID:                 6,
			Title:               "Mocked Title 6",
			CreatedAt:           &time.Time{},
			Draft:               false,
			Author:              MockUser,
			DetailedMergeStatus: "discussions_not_resolved",
			HasConflicts:        true,
			UserNotesCount:      0,
			WebURL:              "",
			Description:         "Dummy description",
		},
		{
			IID:                 7,
			Title:               "Mocked Title 7",
			CreatedAt:           &time.Time{},
			Draft:               false,
			Author:              MockUser,
			DetailedMergeStatus: "ci_still_running",
			HasConflicts:        true,
			UserNotesCount:      0,
			WebURL:              "",
			Description:         "Dummy description",
		},
	}
)
