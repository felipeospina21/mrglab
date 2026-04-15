// Package pipelines implements the pipelines panel component.
package pipelines

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
)

// Model holds the state for the pipelines panel.
type Model struct {
	Table       table.Model
	Loading     bool
	SpinnerView string
	ctx         *context.AppContext
	client      *gitlab.Client
	width       int
	height      int
}

// ColName maps column identifiers to their string names.
type ColName struct {
	Status    string
	IID       string
	Commit    string
	MrIID     string
	Source    string
	Jobs      string
	Author    string
	Duration  string
	CreatedAt string
	Sha       string
	Branch    string
	Path      string
}

// ColNames contains the canonical column name strings.
var ColNames = ColName{
	Status:    "status",
	IID:       "iid",
	Commit:    "commit",
	MrIID:     "mr_iid",
	Source:    "source",
	Jobs:      "jobs",
	Author:    "author",
	Duration:  "duration",
	CreatedAt: "created_at",
	Sha:       "sha",
	Branch:    "branch",
	Path:      "path",
}

// Cols defines the default column layout for the pipelines table.
var Cols = []table.Column{
	{Name: ColNames.CreatedAt, Title: icon.Clock, Width: 4},
	{Name: ColNames.Status, Title: icon.Pipeline, Width: 2, Centered: true},
	{Name: ColNames.IID, Title: "#", Width: 3, Centered: true},
	{Name: ColNames.Commit, Title: "Commit", Width: 50},
	{Name: ColNames.Jobs, Title: "Jobs", Width: 3, Centered: true},
	{Name: ColNames.Author, Title: "Author", Width: 8},
	{Name: ColNames.MrIID, Title: icon.PR, Width: 3, Centered: true},
	{Name: ColNames.Branch, Title: icon.SourceBranch, Width: 8},
	{Name: ColNames.Source, Title: icon.Start, Width: 8},
	{Name: ColNames.Duration, Title: icon.Time, Width: 5},
	{Name: ColNames.Sha, Title: "Sha", Width: 0},
	{Name: ColNames.Path, Title: "Path", Width: 0},
}

// IconCols returns the indices of columns that contain icons.
var IconCols = func() []int {
	return []int{table.GetColIndex(Cols, ColNames.Status)}
}

// New creates a new pipelines model.
func New(ctx *context.AppContext, client *gitlab.Client) Model {
	return Model{
		Table: table.Model{
			EmptyMessage: "Select A Project",
		},
		ctx:    ctx,
		client: client,
	}
}

// SetFocus sets the focused panel to the main panel.
func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.MainPanel
}

// GetTableColums computes column widths proportionally for the given terminal width.
func GetTableColums(width int) []table.Column {
	w := table.ColWidth
	columns := []table.Column{}
	visibleCount := 0
	for _, col := range Cols {
		if col.Width > 0 {
			visibleCount++
		}
	}
	contentWidth := width - visibleCount*2
	used := 0
	commitIdx := -1
	for i, col := range Cols {
		col.Width = w(contentWidth, col.Width)
		if col.Name == ColNames.Commit {
			commitIdx = i
		}
		used += col.Width
		columns = append(columns, table.Column(col))
	}
	if commitIdx >= 0 {
		columns[commitIdx].Width += contentWidth - used
	}
	return columns
}

// GetTableRows converts a PipelineConnection into table rows.
func GetTableRows(pipelines gitlab.PipelineConnection) []table.Row {
	var rows []table.Row
	for _, node := range pipelines.Nodes {
		mrIID := icon.Dash
		sourceBranch := icon.Dash
		if node.MergeRequest != nil {
			mrIID = "!" + node.MergeRequest.IID
			sourceBranch = node.MergeRequest.SourceBranch
		}
		rows = append(rows, table.Row{
			table.FormatTime(node.CreatedAt),
			pipelineStatusIcon(node.Status),
			node.IID,
			node.Commit.Title,
			strconv.Itoa(node.Jobs.Count),
			node.User.Name,
			mrIID,
			sourceBranch,
			formatSource(node.Source),
			formatDuration(node.Duration),
			node.Commit.ShortId,
			node.Path,
		})
	}
	return rows
}

// GetColIndex returns the index of a column by name.
func GetColIndex(colName string) int {
	return slices.IndexFunc(Cols, func(c table.Column) bool {
		return c.Name == colName
	})
}

func pipelineStatusIcon(status string) string {
	icons := map[string]string{
		"success":  icon.CircleCheck,
		"failed":   icon.CircleCross,
		"running":  icon.CircleRunning,
		"pending":  icon.CirclePause,
		"canceled": icon.CircleCancel,
		"skipped":  icon.CircleSkip,
		"manual":   icon.Gear,
		"created":  icon.CircleDot,
	}
	if v, ok := icons[strings.ToLower(status)]; ok {
		return v
	}
	return icon.Dash
}

func formatDuration(d *int) string {
	if d == nil {
		return "-"
	}
	s := *d
	if s < 60 {
		return fmt.Sprintf("%ds", s)
	}
	m := s / 60
	rem := s % 60
	if m < 60 {
		return fmt.Sprintf("%dm %ds", m, rem)
	}
	h := m / 60
	rm := m % 60
	return fmt.Sprintf("%dh %dm", h, rm)
}

func formatSource(source string) string {
	replacer := strings.NewReplacer("_event", "", "_", " ")
	return replacer.Replace(source)
}
