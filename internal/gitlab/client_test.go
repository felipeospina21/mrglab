package gitlab

import (
	"testing"

	"github.com/felipeospina21/mrglab/internal/config"
)

func TestNewClientDevMode(t *testing.T) {
	cfg := &config.Config{DemoMode: true}
	c := NewClient(cfg)

	if !c.demoMode {
		t.Error("NewClient with DemoMode=true should set demoMode=true")
	}
	if c.gql != nil {
		t.Error("NewClient with DemoMode=true should not create graphql client")
	}
}

func TestNewClientNonDevMode(t *testing.T) {
	cfg := &config.Config{
		DemoMode: false,
		BaseURL:  "https://gitlab.example.com",
		APIToken: "test-token",
	}
	c := NewClient(cfg)

	if c.demoMode {
		t.Error("NewClient with DemoMode=false should set demoMode=false")
	}
	if c.gql == nil {
		t.Error("NewClient with DemoMode=false should create graphql client")
	}
}

func TestProjectFullPath(t *testing.T) {
	cfg := &config.Config{
		Filters: config.Filter{
			Projects: []config.Project{
				{ID: "123", FullPath: "my-group/my-project"},
				{ID: "456", FullPath: "other-group/other-project"},
			},
		},
	}
	c := &Client{cfg: cfg}

	got := c.projectFullPath("456")
	if string(got) != "other-group/other-project" {
		t.Errorf("projectFullPath(456) = %v, want other-group/other-project", got)
	}
}

func devClient() *Client {
	c := NewClient(&config.Config{DemoMode: true})
	c.noSleep = true
	return c
}

func TestGetProjectMergeRequestsDevMode(t *testing.T) {
	c := devClient()
	mrs, err := c.GetProjectMergeRequests("any", MergeRequestsQueryVariables{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if mrs.Count == 0 {
		t.Error("expected mock MRs, got 0")
	}
}

func TestGetMergeRequestDevMode(t *testing.T) {
	c := devClient()
	mr, err := c.GetMergeRequest("any", MergeRequestQueryVariables{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if mr.SourceBranch == "" {
		t.Error("expected mock MR with SourceBranch, got empty")
	}
}

func TestGetProjectInfoDevMode(t *testing.T) {
	c := devClient()
	info, err := c.GetProjectInfo("any")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if info.DefaultBranch != "main" {
		t.Errorf("DefaultBranch = %q, want main", info.DefaultBranch)
	}
}

func TestGetMRDescriptionTemplatesDevMode(t *testing.T) {
	c := devClient()
	templates, err := c.GetMRDescriptionTemplates("any")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(templates) == 0 {
		t.Error("expected mock templates, got 0")
	}
}

func TestAcceptMergeRequestDevMode(t *testing.T) {
	c := devClient()
	resp, err := c.AcceptMergeRequest("any", MergeRequestAcceptInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Errors) != 0 {
		t.Errorf("expected no errors, got %v", resp.Errors)
	}
}

func TestCreateNoteDevMode(t *testing.T) {
	c := devClient()
	resp, err := c.CreateNote(CreateNoteInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Errors) != 0 {
		t.Errorf("expected no errors, got %v", resp.Errors)
	}
}

func TestCreateMergeRequestDevMode(t *testing.T) {
	c := devClient()
	resp, err := c.CreateMergeRequest("any", CreateMergeRequestInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Errors) != 0 {
		t.Errorf("expected no errors, got %v", resp.Errors)
	}
}
