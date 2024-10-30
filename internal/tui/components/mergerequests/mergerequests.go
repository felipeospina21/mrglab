package mergerequests

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
)

type Model struct {
	Table table.Model
	ctx   *context.AppContext
}

type ColName struct {
	CreatedAt   string
	IsDraft     string
	Title       string
	Author      string
	Status      string
	IsMergeable string
	Approvals   string
	Discussions string
	Diffs       string
	UpdatedAt   string
	URL         string
	Description string
	ID          string
}

var ColNames = ColName{
	CreatedAt:   "created_at",
	IsDraft:     "is_draft",
	Title:       "title",
	Author:      "author",
	Status:      "status",
	IsMergeable: "is_mergeable",
	Approvals:   "approvals",
	Discussions: "discussions",
	Diffs:       "diffs",
	UpdatedAt:   "updated_at",
	URL:         "url",
	Description: "description",
	ID:          "id",
}

var Cols = []table.Column{
	{
		Name:  ColNames.CreatedAt,
		Title: icon.Clock,
		Width: 2,
	},
	{
		Name:     ColNames.IsDraft,
		Title:    "",
		Width:    2,
		Centered: true,
	},
	{
		Name:  ColNames.Title,
		Title: "Title",
		Width: 25,
	},
	{
		Name:  ColNames.Author,
		Title: "Author",
		Width: 8,
	},
	{
		Name:     ColNames.Status,
		Title:    "Status",
		Width:    4,
		Centered: true,
	},
	{
		Name:     ColNames.IsMergeable,
		Title:    icon.Merge,
		Width:    4,
		Centered: true,
	},
	{
		Name:     ColNames.Approvals,
		Title:    icon.Approval,
		Width:    4,
		Centered: true,
	},
	{
		Name:     ColNames.Discussions,
		Title:    icon.Discussion,
		Width:    4,
		Centered: true,
	},
	{
		Name:     ColNames.Diffs,
		Title:    icon.Diff,
		Width:    8,
		Centered: true,
	},
	{
		Name:  ColNames.UpdatedAt,
		Title: icon.UserUpdate,
		Width: 4,
	},
	{
		Name:  ColNames.URL,
		Title: "Url",
		Width: 0,
	},
	{
		Name:  ColNames.Description,
		Title: "Description",
		Width: 0,
	},
	{
		Name:  ColNames.ID,
		Title: "Id",
		Width: 0,
	},
}

var IconCols = func() []int {
	return []int{
		table.GetColIndex(Cols, "is_draft"),
		table.GetColIndex(Cols, "status"),
		table.GetColIndex(Cols, "is_mergeable"),
		table.GetColIndex(Cols, "approvals"),
		table.GetColIndex(Cols, "diffs"),
	}
}

func New(ctx *context.AppContext) Model {
	return Model{
		Table: table.Model{
			EmptyMessage: "Select A Project",
		},
		ctx: ctx,
	}
}

func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.MainPanel
}

func GetTableColums(width int) []table.Column {
	w := table.ColWidth
	columns := []table.Column{}
	for _, col := range Cols {
		col.Width = w(width, col.Width)
		columns = append(columns, table.Column(col))
	}
	return columns
}

func GetTableRows(msg message.MergeRequestsListFetchedMsg) []table.Row {
	var rows []table.Row

	for _, edge := range msg.Mrs.Edges {
		node := edge.Node
		r := table.Row{
			table.FormatTime(node.CreatedAt),
			table.RenderIcon(node.Draft, icon.Edit),
			node.Title,
			node.Author.Name,
			// node.DetailedMergeStatus,
			detailedStatus(node.DetailedMergeStatus),
			isMergeable(node.DetailedMergeStatus, node.Conflicts),
			approvals(node.ApprovalsLeft, node.ApprovalsRequired),
			strconv.Itoa(node.UserNotesCount),
			diff(node.DiffStatsSummary.Additions, node.DiffStatsSummary.Deletions),
			table.FormatTime(node.UpdatedAt),
			node.WebURL,
			node.Description,
			node.IID,
		}

		rows = append(rows, r)
	}
	return rows
}

func GetColIndex(colName string) int {
	return slices.IndexFunc(Cols, func(c table.Column) bool {
		return c.Name == colName
	})
}

// approvals_syncing: The merge request’s approvals are syncing.
// blocked_status: Blocked by another merge request.
// checking: Git is testing if a valid merge is possible.
// ci_must_pass: A CI/CD pipeline must succeed before merge.
// ci_still_running: A CI/CD pipeline is still running.
// conflict: Conflicts exist between the source and target branches.
// discussions_not_resolved: All discussions must be resolved before merge.
// draft_status: Can’t merge because the merge request is a draft.
// external_status_checks: All status checks must pass before merge.
// jira_association_missing: The title or description must reference a Jira issue. To configure, see Require associated Jira issue for merge requests to be merged.
// mergeable: The branch can merge cleanly into the target branch.
// need_rebase: The merge request must be rebased.
// not_approved: Approval is required before merge.
// not_open: The merge request must be open before merge.
// requested_changes: The merge request has reviewers who have requested changes.
// unchecked: Git has not yet tested if a valid merge is possible.
func detailedStatus(status string) string {
	s := map[string]string{
		"not_approved":             icon.Empty,
		"unchecked":                icon.Dash,
		"mergeable":                icon.Check,
		"checking":                 icon.Clock,
		"need_rebase":              icon.Rebase,
		"conflict":                 icon.Cross,
		"blocked_status":           icon.Alert,
		"discussions_not_resolved": icon.Discussion,
		"ci_still_running":         icon.Time,
		"draft_status":             icon.Edit,
	}
	return s[status]
}

func isMergeable(status string, hasConflicts bool) string {
	if hasConflicts {
		return icon.CircleCross
	}

	if status == "mergeable" {
		return icon.CircleCheck
	}

	return icon.Dash
}

func approvals(left int, total int) string {
	if left == total {
		return icon.Check
	}

	return fmt.Sprintf("%v/%v", left, total)
}

func diff(additions int, deletions int) string {
	return fmt.Sprintf("+%v / -%v", additions, deletions)
}
