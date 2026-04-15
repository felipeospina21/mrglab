package gitlab

import "time"

var now = time.Now()

var mergeRequestConnectionMock = MergeRequestConnection{
	Count: 8,
	Edges: []MergeRequestEdge{
		{
			Node: MergeRequestNode{
				IID:                        "482",
				DiffHeadSha:                "a1b2c3d4e5f6",
				Title:                      "feat: add keyboard shortcuts for discussion navigation",
				CreatedAt:                  now.Add(-2 * 24 * time.Hour),
				UpdatedAt:                  now.Add(-30 * time.Minute),
				Draft:                      false,
				Author:                     Author{Name: "Sarah Chen"},
				DetailedMergeStatus:        "mergeable",
				Conflicts:                  false,
				HeadPipeline:               &HeadPipelineStatus{Status: "SUCCESS"},
				UserNotesCount:      3,
				WebURL:              "https://gitlab.com/my-group/my-project/-/merge_requests/482",
				Description:         "Adds `n`/`N` keybindings to navigate between resolvable discussions in the details panel.\n\n## Changes\n- Navigate forward/backward through unresolved threads\n- Viewport auto-scrolls to the selected discussion\n- Status bar shows current discussion index",
				ApprovalsRequired:   2,
				ApprovalState: MergeRequestApprovalState{
					Rules: []ApprovalRule{
						{
							Name:              "Code Review",
							ApprovalsRequired: 2,
							Approved:          true,
							ApprovedBy: ApprovedBy{
								Nodes: []ApprovedByNode{
									{Name: "James Park"},
									{Name: "Maria Lopez"},
								},
							},
						},
					},
				},
				DiffStatsSummary: DiffStatsSummary{
					Additions: 147,
					Deletions: 23,
					Changes:   170,
					FileCount: 6,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                        "479",
				DiffHeadSha:                "f7e8d9c0b1a2",
				Title:                      "fix: resolve pipeline status icon not updating after retry",
				CreatedAt:                  now.Add(-5 * 24 * time.Hour),
				UpdatedAt:                  now.Add(-3 * time.Hour),
				Draft:                      false,
				Author:                     Author{Name: "James Park"},
				DetailedMergeStatus:        "ci_still_running",
				Conflicts:                  false,
				HeadPipeline:               &HeadPipelineStatus{Status: "RUNNING"},
				UserNotesCount:      7,
				WebURL:              "https://gitlab.com/my-group/my-project/-/merge_requests/479",
				Description:         "Pipeline stage icons were stuck on the previous status after a manual retry. This patch re-fetches the head pipeline when the details panel is opened.\n\nCloses #312",
				ApprovalsRequired:   2,
				ApprovalState: MergeRequestApprovalState{
					Rules: []ApprovalRule{
						{
							Name:              "Code Review",
							ApprovalsRequired: 2,
							Approved:          false,
							ApprovedBy: ApprovedBy{
								Nodes: []ApprovedByNode{
									{Name: "Sarah Chen"},
								},
							},
						},
					},
				},
				DiffStatsSummary: DiffStatsSummary{
					Additions: 34,
					Deletions: 12,
					Changes:   46,
					FileCount: 3,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                        "476",
				DiffHeadSha:                "1a2b3c4d5e6f",
				Title:                      "refactor: extract layout computation into dedicated module",
				CreatedAt:                  now.Add(-7 * 24 * time.Hour),
				UpdatedAt:                  now.Add(-1 * 24 * time.Hour),
				Draft:                      false,
				Author:                     Author{Name: "Maria Lopez"},
				DetailedMergeStatus:        "discussions_not_resolved",
				Conflicts:                  false,
				HeadPipeline:               &HeadPipelineStatus{Status: "SUCCESS"},
				UserNotesCount:      12,
				WebURL:              "https://gitlab.com/my-group/my-project/-/merge_requests/476",
				Description:         "Moves panel sizing logic from `app/model.go` into `app/layout.go`.\n\nThis separates layout concerns from state management and makes the resize behavior easier to test.\n\n**No functional changes** — all layout calculations produce identical results.",
				ApprovalsRequired:   2,
				ApprovalState: MergeRequestApprovalState{
					Rules: []ApprovalRule{
						{
							Name:              "Code Review",
							ApprovalsRequired: 2,
							Approved:          false,
							ApprovedBy: ApprovedBy{
								Nodes: []ApprovedByNode{},
							},
						},
					},
				},
				DiffStatsSummary: DiffStatsSummary{
					Additions: 210,
					Deletions: 185,
					Changes:   395,
					FileCount: 8,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                        "474",
				DiffHeadSha:                "c0d1e2f3a4b5",
				Title:                      "WIP: support custom color themes via config",
				CreatedAt:                  now.Add(-10 * 24 * time.Hour),
				UpdatedAt:                  now.Add(-2 * 24 * time.Hour),
				Draft:                      true,
				Author:                     Author{Name: "Alex Rivera"},
				DetailedMergeStatus:        "draft_status",
				Conflicts:                  false,
				HeadPipeline:               &HeadPipelineStatus{Status: "FAILED"},
				UserNotesCount:      2,
				WebURL:              "https://gitlab.com/my-group/my-project/-/merge_requests/474",
				Description:         "Early exploration of user-defined color themes.\n\nReads an optional `[theme]` section from the config TOML and overrides the default palette. Still needs work on validation and fallback behavior.",
				ApprovalsRequired:   2,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 89,
					Deletions: 4,
					Changes:   93,
					FileCount: 5,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                        "471",
				DiffHeadSha:                "b5a4f3e2d1c0",
				Title:                      "fix: table column widths overflow on narrow terminals",
				CreatedAt:                  now.Add(-12 * 24 * time.Hour),
				UpdatedAt:                  now.Add(-6 * 24 * time.Hour),
				Draft:                      false,
				Author:                     Author{Name: "Sarah Chen"},
				DetailedMergeStatus:        "conflict",
				Conflicts:                  true,
				HeadPipeline:               &HeadPipelineStatus{Status: "SUCCESS"},
				UserNotesCount:      4,
				WebURL:              "https://gitlab.com/my-group/my-project/-/merge_requests/471",
				Description:         "On terminals narrower than ~100 columns, the title column could push other columns off-screen. This adds a minimum width clamp and distributes remaining space proportionally.\n\nFixes #298",
				ApprovalsRequired:   1,
				ApprovalState: MergeRequestApprovalState{
					Rules: []ApprovalRule{
						{
							Name:              "Code Review",
							ApprovalsRequired: 1,
							Approved:          true,
							ApprovedBy: ApprovedBy{
								Nodes: []ApprovedByNode{
									{Name: "James Park"},
								},
							},
						},
					},
				},
				DiffStatsSummary: DiffStatsSummary{
					Additions: 52,
					Deletions: 18,
					Changes:   70,
					FileCount: 2,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                        "468",
				DiffHeadSha:                "e6f7a8b9c0d1",
				Title:                      "chore: bump bubbletea to v1.3.0 and update viewport API",
				CreatedAt:                  now.Add(-14 * 24 * time.Hour),
				UpdatedAt:                  now.Add(-10 * 24 * time.Hour),
				Draft:                      false,
				Author:                     Author{Name: "James Park"},
				DetailedMergeStatus:        "not_approved",
				Conflicts:                  false,
				HeadPipeline:               &HeadPipelineStatus{Status: "CANCELED"},
				UserNotesCount:      0,
				WebURL:              "https://gitlab.com/my-group/my-project/-/merge_requests/468",
				Description:         "Updates bubbletea from v1.2.4 to v1.3.0.\n\n### Breaking changes handled\n- `viewport.New()` signature changed — updated in `details/details.go`\n- Deprecated `tea.WindowSizeMsg` fields removed\n\n```\ngo get github.com/charmbracelet/bubbletea@v1.3.0\n```",
				ApprovalsRequired:   2,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 28,
					Deletions: 31,
					Changes:   59,
					FileCount: 4,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                        "465",
				DiffHeadSha:                "d2c3b4a5f6e7",
				Title:                      "feat: copy MR URL to clipboard with keybinding",
				CreatedAt:                  now.Add(-18 * 24 * time.Hour),
				UpdatedAt:                  now.Add(-15 * 24 * time.Hour),
				Draft:                      false,
				Author:                     Author{Name: "Alex Rivera"},
				DetailedMergeStatus:        "need_rebase",
				Conflicts:                  false,
				HeadPipeline:               &HeadPipelineStatus{Status: "SUCCESS"},
				UserNotesCount:      1,
				WebURL:              "https://gitlab.com/my-group/my-project/-/merge_requests/465",
				Description:         "Adds a `y` keybinding in the merge requests panel to copy the selected MR's web URL to the system clipboard.\n\nUses the existing `exec.CopyToClipboard` helper.",
				ApprovalsRequired:   1,
				ApprovalState: MergeRequestApprovalState{
					Rules: []ApprovalRule{
						{
							Name:              "Code Review",
							ApprovalsRequired: 1,
							Approved:          true,
							ApprovedBy: ApprovedBy{
								Nodes: []ApprovedByNode{
									{Name: "Maria Lopez"},
								},
							},
						},
					},
				},
				DiffStatsSummary: DiffStatsSummary{
					Additions: 15,
					Deletions: 2,
					Changes:   17,
					FileCount: 2,
				},
			},
		},
		{
			Node: MergeRequestNode{
				IID:                        "460",
				DiffHeadSha:                "a8b9c0d1e2f3",
				Title:                      "docs: update README with keybindings and config reference",
				CreatedAt:                  now.Add(-21 * 24 * time.Hour),
				UpdatedAt:                  now.Add(-20 * 24 * time.Hour),
				Draft:                      false,
				Author:                     Author{Name: "Maria Lopez"},
				DetailedMergeStatus:        "checking",
				Conflicts:                  false,
				HeadPipeline:               nil,
				UserNotesCount:      0,
				WebURL:              "https://gitlab.com/my-group/my-project/-/merge_requests/460",
				Description:         "Rewrites the README to include:\n- Full keybinding tables for each panel\n- Configuration reference with all TOML options\n- Usage examples for dev mode",
				ApprovalsRequired:   1,
				DiffStatsSummary: DiffStatsSummary{
					Additions: 95,
					Deletions: 42,
					Changes:   137,
					FileCount: 1,
				},
			},
		},
	},
}

var dur120 = 120
var dur345 = 345
var dur42 = 42
var dur780 = 780
var finishedAt1 = now.Add(-1 * time.Hour)
var finishedAt2 = now.Add(-4 * time.Hour)
var finishedAt3 = now.Add(-2 * 24 * time.Hour)
var finishedAt4 = now.Add(-6 * time.Hour)

var pipelineConnectionMock = PipelineConnection{
	Count: 6,
	Nodes: []PipelineNode{
		{
			ID:           "gid://gitlab/Ci::Pipeline/1001",
			IID:          "301",
			Path:         "/acme-corp/payments-api/-/pipelines/301",
			Commit:       PipelineCommit{ShortId: "a1b2c3d4", Title: "feat: add keyboard shortcuts for discussion navigation"},
			MergeRequest: &PipelineMergeRequest{IID: "482", SourceBranch: "feat/discussion-navigation"},
			Jobs:         PipelineJobsConnection{Count: 3, Nodes: []PipelineJobNode{{ID: "gid://gitlab/Ci::Build/2001", Name: "compile", Status: "SUCCESS", Stage: PipelineJobStage{Name: "build"}}, {ID: "gid://gitlab/Ci::Build/2002", Name: "lint", Status: "SUCCESS", Stage: PipelineJobStage{Name: "build"}}, {ID: "gid://gitlab/Ci::Build/2003", Name: "unit-tests", Status: "SUCCESS", Stage: PipelineJobStage{Name: "test"}}}},
			Status:       "SUCCESS",
			Source:       "merge_request_event",
			CreatedAt:    now.Add(-2 * time.Hour),
			UpdatedAt:    now.Add(-1 * time.Hour),
			FinishedAt:   &finishedAt1,
			Duration:     &dur120,
			User:         Author{Name: "Sarah Chen"},
		},
		{
			ID:           "gid://gitlab/Ci::Pipeline/1002",
			IID:          "300",
			Path:         "/acme-corp/payments-api/-/pipelines/300",
			Commit:       PipelineCommit{ShortId: "f7e8d9c0", Title: "fix: resolve pipeline status icon not updating after retry"},
			MergeRequest: &PipelineMergeRequest{IID: "479", SourceBranch: "fix/pipeline-status-icon"},
			Jobs:         PipelineJobsConnection{Count: 4, Nodes: []PipelineJobNode{{ID: "gid://gitlab/Ci::Build/2004", Name: "compile", Status: "SUCCESS", Stage: PipelineJobStage{Name: "build"}}, {ID: "gid://gitlab/Ci::Build/2005", Name: "lint", Status: "SUCCESS", Stage: PipelineJobStage{Name: "build"}}, {ID: "gid://gitlab/Ci::Build/2006", Name: "unit-tests", Status: "RUNNING", Stage: PipelineJobStage{Name: "test"}}, {ID: "gid://gitlab/Ci::Build/2007", Name: "integration-tests", Status: "PENDING", Stage: PipelineJobStage{Name: "test"}}}},
			Status:       "RUNNING",
			Source:       "merge_request_event",
			CreatedAt:    now.Add(-30 * time.Minute),
			UpdatedAt:    now.Add(-5 * time.Minute),
			Duration:     &dur345,
			User:         Author{Name: "James Park"},
		},
		{
			ID:         "gid://gitlab/Ci::Pipeline/1003",
			IID:        "299",
			Path:       "/acme-corp/payments-api/-/pipelines/299",
			Commit:     PipelineCommit{ShortId: "1a2b3c4d", Title: "refactor: extract layout computation into dedicated module"},
			Jobs:       PipelineJobsConnection{Count: 3, Nodes: []PipelineJobNode{{ID: "gid://gitlab/Ci::Build/2008", Name: "compile", Status: "SUCCESS", Stage: PipelineJobStage{Name: "build"}}, {ID: "gid://gitlab/Ci::Build/2009", Name: "lint", Status: "FAILED", Stage: PipelineJobStage{Name: "build"}}, {ID: "gid://gitlab/Ci::Build/2010", Name: "unit-tests", Status: "SKIPPED", Stage: PipelineJobStage{Name: "test"}}}},
			Status:     "FAILED",
			Source:     "push",
			CreatedAt:  now.Add(-5 * time.Hour),
			UpdatedAt:  now.Add(-4 * time.Hour),
			FinishedAt: &finishedAt2,
			Duration:   &dur42,
			User:       Author{Name: "Maria Lopez"},
		},
		{
			ID:           "gid://gitlab/Ci::Pipeline/1004",
			IID:          "298",
			Path:         "/acme-corp/payments-api/-/pipelines/298",
			Commit:       PipelineCommit{ShortId: "c0d1e2f3", Title: "WIP: support custom color themes via config"},
			MergeRequest: &PipelineMergeRequest{IID: "474", SourceBranch: "feat/custom-themes"},
			Jobs:         PipelineJobsConnection{Count: 2, Nodes: []PipelineJobNode{{ID: "gid://gitlab/Ci::Build/2011", Name: "compile", Status: "CANCELED", Stage: PipelineJobStage{Name: "build"}}, {ID: "gid://gitlab/Ci::Build/2012", Name: "lint", Status: "CANCELED", Stage: PipelineJobStage{Name: "build"}}}},
			Status:       "CANCELED",
			Source:       "web",
			CreatedAt:    now.Add(-3 * 24 * time.Hour),
			UpdatedAt:    now.Add(-2 * 24 * time.Hour),
			FinishedAt:   &finishedAt3,
			Duration:     &dur780,
			User:         Author{Name: "Alex Rivera"},
		},
		{
			ID:        "gid://gitlab/Ci::Pipeline/1005",
			IID:       "297",
			Path:      "/acme-corp/payments-api/-/pipelines/297",
			Commit:    PipelineCommit{ShortId: "b5a4f3e2", Title: "chore: scheduled nightly pipeline"},
			Jobs:      PipelineJobsConnection{Count: 0},
			Status:    "PENDING",
			Source:    "schedule",
			CreatedAt: now.Add(-10 * time.Minute),
			UpdatedAt: now.Add(-10 * time.Minute),
			User:      Author{Name: "Sarah Chen"},
		},
		{
			ID:         "gid://gitlab/Ci::Pipeline/1006",
			IID:        "296",
			Path:       "/acme-corp/payments-api/-/pipelines/296",
			Commit:     PipelineCommit{ShortId: "e6f7a8b9", Title: "feat: copy MR URL to clipboard with keybinding"},
			Jobs:       PipelineJobsConnection{Count: 1, Nodes: []PipelineJobNode{{ID: "gid://gitlab/Ci::Build/2013", Name: "deploy-staging", Status: "MANUAL", Stage: PipelineJobStage{Name: "deploy"}}}},
			Status:     "MANUAL",
			Source:     "push",
			CreatedAt:  now.Add(-7 * time.Hour),
			UpdatedAt:  now.Add(-6 * time.Hour),
			FinishedAt: &finishedAt4,
			User:       Author{Name: "James Park"},
		},
	},
}

var mergeRequestResponseMock = MergeRequestResponse{
	Id:           "gid://gitlab/MergeRequest/482",
	SourceBranch: "feat/discussion-navigation",
	TargetBranch: "main",
	ApprovalState: MergeRequestApprovalState{
		Rules: []ApprovalRule{
			{
				Name:              "Code Review",
				ApprovalsRequired: 2,
				Approved:          true,
				ApprovedBy: ApprovedBy{
					Nodes: []ApprovedByNode{
						{Name: "James Park"},
						{Name: "Maria Lopez"},
					},
				},
			},
		},
	},
	HeadPipeline: MergeRequestHeadPipelineConnection{
		Stages: CiStageConnection{
			Nodes: []CiStageNode{
				{
					Name:   "build",
					Status: "SUCCESS",
					Jobs: JobsConnection{
						Nodes: []JobsNode{
							{Name: "compile", Status: "SUCCESS", Duration: 45},
							{Name: "lint", Status: "SUCCESS", Duration: 22},
						},
					},
				},
				{
					Name:   "test",
					Status: "SUCCESS",
					Jobs: JobsConnection{
						Nodes: []JobsNode{
							{Name: "unit-tests", Status: "SUCCESS", Duration: 120},
							{Name: "integration-tests", Status: "SUCCESS", Duration: 340},
						},
					},
				},
				{
					Name:   "deploy",
					Status: "MANUAL",
					Jobs: JobsConnection{
						Nodes: []JobsNode{
							{Name: "deploy-staging", Status: "MANUAL", Duration: 0},
						},
					},
				},
			},
		},
	},
	Discussions: MergeRequestDiscussionsConnection{
		Nodes: []DiscussionNode{
			{
				Id:         "gid://gitlab/Discussion/a1b2c3d4",
				Resolvable: true,
				Resolved:   false,
				Notes: NoteConnection{
					Nodes: []Note{
						{
							Resolvable: true,
							Author:     Author{Name: "James Park"},
							Body:       "Should we debounce the `n`/`N` key presses? If someone holds the key down, it could rapidly cycle through all discussions and cause a lot of re-renders.",
							CreatedAt:  now.Add(-26 * time.Hour),
						},
						{
							Resolvable: true,
							Author:     Author{Name: "Sarah Chen"},
							Body:       "Good point. I tested it and bubbletea's key repeat rate is already throttled by the terminal, so it feels fine in practice. But I can add a small cooldown if you think it's worth it.",
							CreatedAt:  now.Add(-24 * time.Hour),
						},
						{
							Resolvable: true,
							Author:     Author{Name: "James Park"},
							Body:       "Fair enough — let's leave it as-is for now and revisit if anyone reports issues.",
							CreatedAt:  now.Add(-22 * time.Hour),
						},
					},
				},
			},
			{
				Id:         "gid://gitlab/Discussion/e5f6a7b8",
				Resolvable: true,
				Resolved:   true,
				ResolvedAt: now.Add(-20 * time.Hour),
				Notes: NoteConnection{
					Nodes: []Note{
						{
							Resolvable: true,
							Author:     Author{Name: "Maria Lopez"},
							Body:       "Nit: `resolvableDiscussions()` allocates a new slice on every call. Since we call it from both `nextDiscussion` and `prevDiscussion`, could we cache the indices?",
							CreatedAt:  now.Add(-28 * time.Hour),
						},
						{
							Resolvable: true,
							Author:     Author{Name: "Sarah Chen"},
							Body:       "The list is typically small (< 20 discussions), so the allocation is negligible. I'd rather keep it simple than introduce cache invalidation. Resolved in `a1b2c3d`.",
							CreatedAt:  now.Add(-25 * time.Hour),
						},
					},
				},
			},
			{
				Id:         "gid://gitlab/Discussion/c9d0e1f2",
				Resolvable: true,
				Resolved:   false,
				Notes: NoteConnection{
					Nodes: []Note{
						{
							Resolvable: true,
							Author:     Author{Name: "James Park"},
							Body:       "The viewport auto-scroll offset seems slightly off — the selected discussion header ends up right at the top edge. Could we add a small margin (2-3 lines) above it so there's some context visible?",
							CreatedAt:  now.Add(-18 * time.Hour),
						},
					},
				},
			},
			{
				Id:         "gid://gitlab/Discussion/g3h4i5j6",
				Resolvable: false,
				Resolved:   false,
				Notes: NoteConnection{
					Nodes: []Note{
						{
							Resolvable: false,
							Author:     Author{Name: "Maria Lopez"},
							Body:       "Nice work! The navigation feels really smooth. 🎉",
							CreatedAt:  now.Add(-16 * time.Hour),
						},
					},
				},
			},
		},
	},
}

var mrDescriptionTemplatesMock = []MRDescriptionTemplate{
	{
		Name: "Default",
		Content: `## Summary

Add a quick summary of what this MR is addressing.

### Related Issue

Add a link to an existing issue or bug here.

### Checklist

- [ ] Code follows the project style guidelines
- [ ] Tests have been added or updated
- [ ] Documentation has been updated
- [ ] Changes have been tested locally
`,
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
