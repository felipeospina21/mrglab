package mergerequests

import (
	"testing"
	"time"

	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
)

func TestIsMergeable(t *testing.T) {
	tests := []struct {
		name         string
		status       string
		hasConflicts bool
		expected     string
	}{
		{"mergeable no conflicts", "mergeable", false, icon.CircleCheck},
		{"mergeable with conflicts", "mergeable", true, icon.CircleDash},
		{"not mergeable no conflicts", "checking", false, icon.Dash},
		{"case insensitive", "Mergeable", false, icon.CircleCheck},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMergeable(tt.status, tt.hasConflicts); got != tt.expected {
				t.Errorf("isMergeable(%q, %v) = %q, want %q", tt.status, tt.hasConflicts, got, tt.expected)
			}
		})
	}
}

func TestApprovals(t *testing.T) {
	tests := []struct {
		name     string
		rules    []gitlab.ApprovalRule
		total    int
		expected string
	}{
		{
			"fully approved",
			[]gitlab.ApprovalRule{{ApprovalsRequired: 2, ApprovedBy: gitlab.ApprovedBy{Nodes: []gitlab.ApprovedByNode{{Name: "A"}, {Name: "B"}}}}},
			2,
			icon.Check,
		},
		{
			"partially approved",
			[]gitlab.ApprovalRule{{ApprovalsRequired: 2, ApprovedBy: gitlab.ApprovedBy{Nodes: []gitlab.ApprovedByNode{{Name: "A"}}}}},
			2,
			"1/2",
		},
		{
			"no rules zero total",
			nil,
			0,
			icon.Check,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := approvals(tt.rules, tt.total); got != tt.expected {
				t.Errorf("approvals() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	tests := []struct {
		name       string
		add, del   int
		expected   string
	}{
		{"basic", 10, 5, "+10 / -5"},
		{"zeros", 0, 0, "+0 / -0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diff(tt.add, tt.del); got != tt.expected {
				t.Errorf("diff(%d, %d) = %q, want %q", tt.add, tt.del, got, tt.expected)
			}
		})
	}
}

func TestGetStatusIcon(t *testing.T) {
	tests := []struct {
		name     string
		status   string
		expected string
	}{
		{"mergeable", "mergeable", icon.Check},
		{"need_rebase", "need_rebase", icon.Alert},
		{"conflict", "conflict", icon.Conflict},
		{"unknown", "unknown_status", ""},
		{"case insensitive", "Mergeable", icon.Check},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatusIcon(tt.status); got != tt.expected {
				t.Errorf("getStatusIcon(%q) = %q, want %q", tt.status, got, tt.expected)
			}
		})
	}
}

func TestGetTableRows(t *testing.T) {
	mrs := gitlab.MergeRequestConnection{
		Count: 1,
		Edges: []gitlab.MergeRequestEdge{
			{
				Node: gitlab.MergeRequestNode{
					IID:                 "123",
					Title:               "Test MR",
					Author:              gitlab.Author{Name: "Dev"},
					Draft:               true,
					DetailedMergeStatus: "mergeable",
					Conflicts:           false,
					ApprovalsRequired:   1,
					UserNotesCount:      2,
					WebURL:              "https://gitlab.com/mr/123",
					Description:         "desc",
					DiffHeadSha:         "abc123",
					CreatedAt:           time.Now(),
					UpdatedAt:           time.Now(),
					DiffStatsSummary:    gitlab.DiffStatsSummary{Additions: 5, Deletions: 3},
					ApprovalState: gitlab.MergeRequestApprovalState{
						Rules: []gitlab.ApprovalRule{{ApprovalsRequired: 1, Approved: true, ApprovedBy: gitlab.ApprovedBy{Nodes: []gitlab.ApprovedByNode{{Name: "A"}}}}},
					},
				},
			},
		},
	}

	rows := GetTableRows(mrs)
	if len(rows) != 1 {
		t.Fatalf("GetTableRows() returned %d rows, want 1", len(rows))
	}

	row := rows[0]
	// Title is at index 4 (after created_at, isMergeable, isDraft, IID display)
	if row[4] != "Test MR" {
		t.Errorf("row title = %q, want %q", row[4], "Test MR")
	}
	// Author at index 5
	if row[5] != "Dev" {
		t.Errorf("row author = %q, want %q", row[5], "Dev")
	}
	// IID at index 14
	if row[14] != "123" {
		t.Errorf("row IID = %q, want %q", row[14], "123")
	}
}

func TestGetTableColums(t *testing.T) {
	cols := GetTableColums(200)
	if len(cols) != len(Cols) {
		t.Fatalf("GetTableColums() returned %d cols, want %d", len(cols), len(Cols))
	}

	titleIdx := GetColIndex(ColNames.Title)
	if titleIdx < 0 {
		t.Fatal("title column not found")
	}
	// Title column should have absorbed remaining width
	if cols[titleIdx].Width <= 0 {
		t.Errorf("title column width = %d, want > 0", cols[titleIdx].Width)
	}
}

func TestGetColIndex(t *testing.T) {
	tests := []struct {
		name     string
		colName  string
		wantGte0 bool
	}{
		{"existing column", ColNames.Title, true},
		{"missing column", "nonexistent", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetColIndex(tt.colName)
			if tt.wantGte0 && got < 0 {
				t.Errorf("GetColIndex(%q) = %d, want >= 0", tt.colName, got)
			}
			if !tt.wantGte0 && got >= 0 {
				t.Errorf("GetColIndex(%q) = %d, want < 0", tt.colName, got)
			}
		})
	}
}
