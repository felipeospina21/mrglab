package gitlab

import (
	"testing"

	"github.com/hasura/go-graphql-client"
)

func TestMergeRequestsVariables(t *testing.T) {
	vars := MergeRequestsQueryVariables{ProjectFullPath: graphql.ID("my-group/my-project")}
	got := mergeRequestsVariables(vars)

	if got["fullPath"] != graphql.ID("my-group/my-project") {
		t.Errorf("fullPath = %v, want %v", got["fullPath"], "my-group/my-project")
	}
}

func TestMergeRequestVariables(t *testing.T) {
	vars := MergeRequestQueryVariables{
		MRIID:                       "42",
		MergeRequestsQueryVariables: MergeRequestsQueryVariables{ProjectFullPath: graphql.ID("g/p")},
	}
	got := mergeRequestVariables(vars)

	if got["fullPath"] != graphql.ID("g/p") {
		t.Errorf("fullPath = %v, want g/p", got["fullPath"])
	}
	if got["mrIID"] != "42" {
		t.Errorf("mrIID = %v, want 42", got["mrIID"])
	}
}

func TestAcceptMergeRequestVariables(t *testing.T) {
	input := MergeRequestAcceptInput{
		ProjectPath:              graphql.ID("g/p"),
		IID:                      "10",
		Sha:                      "abc",
		ShouldRemoveSourceBranch: true,
		Squash:                   false,
	}
	got := acceptMergeRequestVariables(input)

	if got["sha"] != "abc" {
		t.Errorf("sha = %v, want abc", got["sha"])
	}
	if got["iid"] != "10" {
		t.Errorf("iid = %v, want 10", got["iid"])
	}
	if got["shouldRemoveSourceBranch"] != true {
		t.Errorf("shouldRemoveSourceBranch = %v, want true", got["shouldRemoveSourceBranch"])
	}
	if got["squash"] != false {
		t.Errorf("squash = %v, want false", got["squash"])
	}
}

func TestCreateNoteVariables(t *testing.T) {
	input := CreateNoteInput{
		NoteableId:   NoteableID("gid://gitlab/MergeRequest/1"),
		DiscussionId: DiscussionID("gid://gitlab/Discussion/2"),
		Body:         "LGTM",
	}
	got := createNoteVariables(input)

	if got["noteableId"] != input.NoteableId {
		t.Errorf("noteableId = %v, want %v", got["noteableId"], input.NoteableId)
	}
	if got["discussionId"] != input.DiscussionId {
		t.Errorf("discussionId = %v, want %v", got["discussionId"], input.DiscussionId)
	}
	if got["body"] != "LGTM" {
		t.Errorf("body = %v, want LGTM", got["body"])
	}
}

func TestCreateMergeRequestVariables(t *testing.T) {
	input := CreateMergeRequestInput{
		ProjectPath:  graphql.ID("g/p"),
		SourceBranch: "feature",
		TargetBranch: "main",
		Title:        "feat: new",
		Description:  "desc",
	}
	got := createMergeRequestVariables(input)

	if got["projectPath"] != graphql.ID("g/p") {
		t.Errorf("projectPath = %v, want g/p", got["projectPath"])
	}
	if got["sourceBranch"] != "feature" {
		t.Errorf("sourceBranch = %v, want feature", got["sourceBranch"])
	}
	if got["targetBranch"] != "main" {
		t.Errorf("targetBranch = %v, want main", got["targetBranch"])
	}
}

func TestNoteableIDGetGraphQLType(t *testing.T) {
	if got := NoteableID("x").GetGraphQLType(); got != "NoteableID" {
		t.Errorf("GetGraphQLType() = %q, want NoteableID", got)
	}
}

func TestDiscussionIDGetGraphQLType(t *testing.T) {
	if got := DiscussionID("x").GetGraphQLType(); got != "DiscussionID" {
		t.Errorf("GetGraphQLType() = %q, want DiscussionID", got)
	}
}
