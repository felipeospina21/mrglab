package gitlab

import "time"

var mergeRequestConnectionMock = MergeRequestConnection{
	Count: 1,
	Edges: []MergeRequestEdge{
		{
			Node: MergeRequestNode{
				IID:                 "1",
				DiffHeadSha:         "some-sha",
				Title:               "Mocked Title 1",
				CreatedAt:           time.Now(),
				UpdatedAt:           time.Now(),
				Draft:               false,
				Author:              Author{Name: "Mock User"},
				DetailedMergeStatus: "ci_still_running",
				Conflicts:           true,
				UserNotesCount:      0,
				WebURL:              "",
				Description:         "Dummy description",
				ApprovalsRequired:   3,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 10,
					Deletions: 50,
					Changes:   100,
					FileCount: 5,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                 "2",
				Title:               "Mocked Title 2",
				CreatedAt:           time.Now(),
				UpdatedAt:           time.Now(),
				Draft:               false,
				Author:              Author{Name: "Mock User"},
				DetailedMergeStatus: "blocked_status",
				Conflicts:           true,
				UserNotesCount:      10,
				WebURL:              "",
				Description:         "Dummy description",
				ApprovalsRequired:   3,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 10,
					Deletions: 50,
					Changes:   100,
					FileCount: 5,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                 "3",
				Title:               "Mocked Title 3",
				CreatedAt:           time.Now(),
				UpdatedAt:           time.Now(),
				Draft:               false,
				Author:              Author{Name: "Mock User"},
				DetailedMergeStatus: "conflict",
				Conflicts:           true,
				UserNotesCount:      0,
				WebURL:              "",
				Description:         "Dummy description",
				ApprovalsRequired:   3,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 10,
					Deletions: 50,
					Changes:   100,
					FileCount: 5,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                 "4",
				Title:               "Mocked Title 4",
				CreatedAt:           time.Now(),
				UpdatedAt:           time.Now(),
				Draft:               false,
				Author:              Author{Name: "Mock User"},
				DetailedMergeStatus: "mergeable",
				Conflicts:           false,
				UserNotesCount:      0,
				WebURL:              "",
				Description:         "Dummy description",
				ApprovalsRequired:   1,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 10,
					Deletions: 50,
					Changes:   100,
					FileCount: 5,
				},
				ApprovalState: MergeRequestApprovalState{
					Rules: []ApprovalRule{
						{
							Name:              "Rule 1",
							ApprovalsRequired: 1,
							Approved:          false,
							ApprovedBy: ApprovedBy{
								Nodes: []ApprovedByNode{
									{Name: "user"},
								},
							},
						},
					},
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                 "5",
				Title:               "Mocked Title 5",
				CreatedAt:           time.Now(),
				UpdatedAt:           time.Now(),
				Draft:               false,
				Author:              Author{Name: "Mock User"},
				DetailedMergeStatus: "external_status_checks",
				Conflicts:           false,
				UserNotesCount:      0,
				WebURL:              "",
				Description:         "Dummy description",
				ApprovalsRequired:   3,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 10,
					Deletions: 50,
					Changes:   100,
					FileCount: 5,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                 "6",
				Title:               "Mocked Title 6",
				CreatedAt:           time.Now(),
				UpdatedAt:           time.Now(),
				Draft:               false,
				Author:              Author{Name: "Mock User"},
				DetailedMergeStatus: "need_rebase",
				Conflicts:           false,
				UserNotesCount:      0,
				WebURL:              "",
				Description:         "Dummy description",
				ApprovalsRequired:   3,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 10,
					Deletions: 50,
					Changes:   100,
					FileCount: 5,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                 "7",
				Title:               "Mocked Title 7",
				CreatedAt:           time.Now(),
				UpdatedAt:           time.Now(),
				Draft:               false,
				Author:              Author{Name: "Mock User"},
				DetailedMergeStatus: "draft_status",
				Conflicts:           false,
				UserNotesCount:      0,
				WebURL:              "",
				Description:         "Dummy description",
				ApprovalsRequired:   3,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 10,
					Deletions: 50,
					Changes:   100,
					FileCount: 5,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                 "8",
				Title:               "Mocked Title 8",
				CreatedAt:           time.Now(),
				UpdatedAt:           time.Now(),
				Draft:               false,
				Author:              Author{Name: "Mock User"},
				DetailedMergeStatus: "checking",
				Conflicts:           false,
				UserNotesCount:      0,
				WebURL:              "",
				Description:         "Dummy description",
				ApprovalsRequired:   3,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 10,
					Deletions: 50,
					Changes:   100,
					FileCount: 5,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                 "9",
				Title:               "Mocked Title 9",
				CreatedAt:           time.Now(),
				UpdatedAt:           time.Now(),
				Draft:               false,
				Author:              Author{Name: "Mock User"},
				DetailedMergeStatus: "external_status_checks",
				Conflicts:           false,
				UserNotesCount:      0,
				WebURL:              "",
				Description:         "Dummy description",
				ApprovalsRequired:   3,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 10,
					Deletions: 50,
					Changes:   100,
					FileCount: 5,
				},
			},
		},
	},
}

var mergeRequestResponseMock = MergeRequestResponse{
	Id: "1",
	ApprovalState: MergeRequestApprovalState{
		Rules: []ApprovalRule{
			{
				Name:              "Rule 1",
				ApprovalsRequired: 1,
				Approved:          false,
				ApprovedBy: ApprovedBy{
					Nodes: []ApprovedByNode{
						{Name: "user"},
					},
				},
			},
		},
	},
	SourceBranch: "feature/mock-branch",
	TargetBranch: "develop",
	HeadPipeline: MergeRequestHeadPipelineConnection{
		Stages: CiStageConnection{
			Nodes: []CiStageNode{
				{Name: "stage1", Status: "success", Jobs: JobsConnection{
					Nodes: []JobsNode{
						{Name: "job1", Status: "success"},
					},
				}},
			},
		},
	},
	Discussions: MergeRequestDiscussionsConnection{
		Nodes: []DiscussionNode{
			{
				Resolvable: true,
				Resolved:   false,
				Notes: NoteConnection{
					Nodes: []Note{
						{
							Resolvable: true,
							Author:     Author{Name: "Mock User"},
							Body:       "Question about these functions that were a `useCallback`: \n\ni.e `handleToggleOpen` is a prop for an `Accordion` component, don't we need to keep the `useCallback` and do something like\n\n```\n  const handleToggleOpen = useCallback(\n    (panelId: string, nextValue: boolean): void => {\n      togglePanel(panelId, nextValue, openPanels, setOpenPanels);\n    },\n    [openPanels, setOpenPanels]\n  );\n```",
							CreatedAt:  time.Now(),
						},
						{
							Resolvable: true,
							Author:     Author{Name: "Mock User"},
							Body:       "not really. The useCallback only makes sense if the function is defined within a React component. Since we extracted the function it won't be re-declared when the component re-renders, since it will always be pointing to the function reference (in another file). So basically when you import a function and used it in the component it has the same behavior as if it was declared inside the function within a useCallback (regarding the preservation of the reference).",
							CreatedAt:  time.Now(),
						},
						{
							Resolvable: true,
							Author:     Author{Name: "Mock User"},
							Body:       "to be more specific, the getHandleToggleOpen function returns a function, so handleToggleOpen is just a reference/pointer.",
							CreatedAt:  time.Now(),
						},
					},
				},
			},
			{
				Resolvable: true,
				Resolved:   true,
				Notes: NoteConnection{
					Nodes: []Note{
						{
							Resolvable: true,
							Author:     Author{Name: "Mock User"},
							Body:       "question",
							CreatedAt:  time.Now(),
						},
						{
							Resolvable: true,
							Author:     Author{Name: "Mock User"},
							Body:       "response 1",
							CreatedAt:  time.Now(),
						},
						{
							Resolvable: false,
							Author:     Author{Name: "Mock User"},
							Body:       "add commit 1 [in versin 4] ()",
							CreatedAt:  time.Now(),
						},
						{
							Resolvable: true,
							Author:     Author{Name: "Mock User"},
							Body:       "response 2",
							CreatedAt:  time.Now(),
						},
					},
				},
			},
		},
	},
}

const MarkdownContentMock = `
# Today's Menu

## Appetizers

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you'd like |

## Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] René Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appétit!
`
