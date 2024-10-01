package mergerequests

import (
	"fmt"
	"strconv"

	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

type Model struct {
	Table table.Model
	ctx   *context.AppContext
}

var Cols = []table.TableCol{
	{
		Name:  "created_at",
		Title: icon.Clock,
		Width: 4,
	},
	{
		Name:  "is_draft",
		Title: icon.Edit,
		Width: 4,
	},
	{
		Name:  "title",
		Title: "Title",
		Width: 25,
	},
	{
		Name:  "author",
		Title: "Author",
		Width: 8,
	},
	{
		Name:  "status",
		Title: "Status",
		Width: 4,
	},
	{
		Name:  "is_mergeable",
		Title: icon.Merge,
		Width: 4,
	},
	{
		Name:  "approvals",
		Title: icon.Approval,
		Width: 4,
	},
	{
		Name:  "discussions",
		Title: icon.Discussion,
		Width: 4,
	},
	{
		Name:  "diffs",
		Title: icon.Diff,
		Width: 8,
	},
	{
		Name:  "updated_at",
		Title: icon.Update,
		Width: 4,
	},
	{
		Name:  "url",
		Title: "Url",
		Width: 0,
	},
	{
		Name:  "description",
		Title: "Description",
		Width: 0,
	},
	{
		Name:  "id",
		Title: "Id",
		Width: 0,
	},
}

var IconCols = func() []int {
	return []int{
		table.GetColIndex(Cols, "is_draft"),
		table.GetColIndex(Cols, "status"),
		table.GetColIndex(Cols, "is_mergeable"),
		// table.GetColIndex(Cols, "approvals"),
		// table.GetColIndex(Cols, "diffs"),
	}
}

func New(ctx *context.AppContext) Model {
	return Model{
		Table: table.Model{},
		ctx:   ctx,
	}
}

func GetTableColums(width int) []table.Column {
	w := table.ColWidth
	columns := []table.Column{}
	for _, col := range Cols {
		columns = append(columns, table.Column{Title: col.Title, Width: w(width, col.Width)})
	}
	return columns
}

func getTableRows(msg task.TaskFinishedMsg) []table.Row {
	var rows []table.Row
	mrs := msg.Msg.(MergeRequestsFetchedMsg)

	for _, edge := range mrs.Mrs.Edges {
		node := edge.Node
		r := table.Row{
			table.FormatTime(node.CreatedAt),
			table.RenderIcon(node.Draft, icon.Check),
			node.Title,
			node.Author.Name,
			// node.DetailedMergeStatus,
			detailedStatus(node.DetailedMergeStatus),
			isMergeable(node.DetailedMergeStatus, node.Conflicts),
			approvals(node.ApprovalsLeft, node.ApprovalsRequired),
			strconv.Itoa(node.UserNotesCount),
			fmt.Sprintf("+%v / -%v", node.DiffStatsSummary.Additions, node.DiffStatsSummary.Deletions),
			table.FormatTime(node.UpdatedAt),
			node.WebURL,
			node.Description,
			node.IID,
		}

		rows = append(rows, r)
	}
	return rows
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
