package mergerequests

import (
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
		Width: 50,
	},
	{
		Name:  "author",
		Title: "Author",
		Width: 20,
	},
	{
		Name:  "status",
		Title: "Status",
		Width: 4,
	},
	{
		Name:  "conflicts",
		Title: "Conflicts",
		Width: 4,
	},
	{
		Name:  "discussions",
		Title: "Discussions",
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
		table.GetColIndex(Cols, "conflicts"),
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
	ml := msg.Msg.(MergeRequestsFetchedMsg)
	for _, mr := range ml.Mrs {
		r := table.Row{
			table.FormatTime(*mr.CreatedAt),
			table.RenderIcon(mr.Draft, icon.Check),
			mr.Title,
			mr.Author.Name,
			// mr.DetailedMergeStatus,
			mapMergeStatus(mr.DetailedMergeStatus),
			table.RenderIcon(mr.HasConflicts, icon.Cross),
			strconv.Itoa(mr.UserNotesCount),
			mr.WebURL,
			mr.Description,
			strconv.Itoa(mr.IID),
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
func mapMergeStatus(status string) string {
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
