package details

import (
	"strings"
	"testing"

	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

func TestShortID(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		expected string
	}{
		{"full gitlab ID", "gid://gitlab/DiffNote/1234567890", "567890"},
		{"short ID", "abc", "abc"},
		{"exactly 6", "123456", "123456"},
		{"7 chars no slash", "1234567", "234567"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortID(tt.id); got != tt.expected {
				t.Errorf("ShortID(%q) = %q, want %q", tt.id, got, tt.expected)
			}
		})
	}
}

func TestGetStageIconStatus(t *testing.T) {
	tests := []struct {
		name         string
		status       string
		expectedIcon string
		expectedColor string
	}{
		{"success", "success", icon.CircleCheck, style.Green[400]},
		{"failed", "failed", icon.CircleCross, style.Red[400]},
		{"running", "running", icon.CircleRunning, style.Blue[400]},
		{"manual", "manual", icon.Gear, style.White},
		{"canceled", "canceled", icon.CircleCancel, style.White},
		{"case insensitive", "SUCCESS", icon.CircleCheck, style.Green[400]},
		{"unknown", "unknown_status", icon.Dash, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getStageIconStatus(tt.status)
			if got.icon != tt.expectedIcon {
				t.Errorf("getStageIconStatus(%q).icon = %q, want %q", tt.status, got.icon, tt.expectedIcon)
			}
			if got.color != tt.expectedColor {
				t.Errorf("getStageIconStatus(%q).color = %q, want %q", tt.status, got.color, tt.expectedColor)
			}
		})
	}
}

func TestTimeAgo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic", "5m", "5m ago"},
		{"hours", "3h", "3h ago"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeAgo(tt.input); got != tt.expected {
				t.Errorf("timeAgo(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestRenderStatus(t *testing.T) {
	got := renderStatus("mergeable")
	if got == "" {
		t.Error("renderStatus returned empty string")
	}
	if !strings.Contains(got, "Mergeable") {
		t.Errorf("renderStatus should contain capitalized status, got %q", got)
	}
}

func TestRenderBranches(t *testing.T) {
	got := renderBranches("feature-branch", "main")
	if !strings.Contains(got, "main") {
		t.Error("renderBranches should contain target branch")
	}
	if !strings.Contains(got, "feature-branch") {
		t.Error("renderBranches should contain source branch")
	}
}

func TestRenderApprovals(t *testing.T) {
	rules := []gitlab.ApprovalRule{
		{
			Name:              "Code Review",
			ApprovalsRequired: 1,
			Approved:          true,
			ApprovedBy:        gitlab.ApprovedBy{Nodes: []gitlab.ApprovedByNode{{Name: "Alice"}}},
		},
		{
			Name:              "Security",
			ApprovalsRequired: 1,
			Approved:          false,
			ApprovedBy:        gitlab.ApprovedBy{},
		},
	}
	got := renderApprovals(rules)
	if !strings.Contains(got, "Approvals") {
		t.Error("renderApprovals should contain title")
	}
	if !strings.Contains(got, "Code Review") {
		t.Error("renderApprovals should contain rule name")
	}
	if !strings.Contains(got, "Alice") {
		t.Error("renderApprovals should contain approver name")
	}
	if !strings.Contains(got, "Security") {
		t.Error("renderApprovals should contain unapproved rule")
	}
}

func TestRenderPipelines(t *testing.T) {
	stages := []gitlab.CiStageNode{
		{
			Name:   "build",
			Status: "success",
			Jobs:   gitlab.JobsConnection{Nodes: []gitlab.JobsNode{{Name: "compile", Status: "success"}}},
		},
		{
			Name:   "test",
			Status: "failed",
			Jobs:   gitlab.JobsConnection{Nodes: []gitlab.JobsNode{{Name: "unit-tests", Status: "failed"}}},
		},
	}
	got := renderPipelines(stages, "")
	if !strings.Contains(got, "Pipeline") {
		t.Error("renderPipelines should contain title")
	}
	if !strings.Contains(got, "build") {
		t.Error("renderPipelines should contain stage name")
	}
	if !strings.Contains(got, "unit-tests") {
		t.Error("renderPipelines should contain failed job name")
	}
}

func TestRenderApprovalsEmpty(t *testing.T) {
	got := renderApprovals(nil)
	if !strings.Contains(got, "Approvals") {
		t.Error("renderApprovals with nil should still contain title")
	}
}

func TestRenderPipelinesEmpty(t *testing.T) {
	got := renderPipelines(nil, "")
	if !strings.Contains(got, "Pipeline") {
		t.Error("renderPipelines with nil should still contain title")
	}
}
