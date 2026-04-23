package gitlab

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hasura/go-graphql-client"
)

// GetProjectMergeRequests fetches the open merge requests for a project.
func (c *Client) GetProjectMergeRequests(
	projectID string,
	vars MergeRequestsQueryVariables,
) (MergeRequestConnection, error) {
	if c.demoMode {
		c.sleep(1 * time.Second)
		return mergeRequestConnectionMock, nil
	}

	var query getProjectMrs
	vars.ProjectFullPath = c.projectFullPath(projectID)
	variables := mergeRequestsVariables(vars)

	err := c.gql.Query(context.Background(), &query, variables)
	if err != nil {
		return MergeRequestConnection{}, err
	}

	return query.Project.MergeRequests, nil
}

// GetMergeRequest fetches a single merge request by IID.
func (c *Client) GetMergeRequest(
	projectID string,
	vars MergeRequestQueryVariables,
) (MergeRequestResponse, error) {
	if c.demoMode {
		c.sleep(800 * time.Millisecond)
		return mergeRequestResponseMock, nil
	}

	var query getMergeRequest
	vars.ProjectFullPath = c.projectFullPath(projectID)
	variables := mergeRequestVariables(vars)

	err := c.gql.Query(context.Background(), &query, variables)
	if err != nil {
		return MergeRequestResponse{}, err
	}

	return query.Project.MergeRequest, nil
}

// ProjectInfo holds the default branch and MR description template from project settings.
type ProjectInfo struct {
	DefaultBranch         string
	MergeRequestsTemplate string
}

// GetProjectInfo fetches the project's default branch and MR description template.
func (c *Client) GetProjectInfo(projectID string) (ProjectInfo, error) {
	if c.demoMode {
		c.sleep(200 * time.Millisecond)
		return ProjectInfo{DefaultBranch: "main", MergeRequestsTemplate: mrDescriptionTemplatesMock[0].Content}, nil
	}

	url := fmt.Sprintf("%s/api/v4/projects/%s", c.cfg.BaseURL, projectID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ProjectInfo{}, err
	}
	req.Header.Set("PRIVATE-TOKEN", c.cfg.APIToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ProjectInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ProjectInfo{}, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	var data struct {
		DefaultBranch         string `json:"default_branch"`
		MergeRequestsTemplate string `json:"merge_requests_template"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return ProjectInfo{}, err
	}

	return ProjectInfo{
		DefaultBranch:         data.DefaultBranch,
		MergeRequestsTemplate: data.MergeRequestsTemplate,
	}, nil
}

// MRDescriptionTemplate represents a resolved MR description template.
type MRDescriptionTemplate struct {
	Name    string
	Content string
}

// GetMRDescriptionTemplates fetches MR description templates from all available sources
// in parallel: project-level files, group-level templates (REST), and the project default
// description setting. Results are merged and deduplicated by name (project files win).
func (c *Client) GetMRDescriptionTemplates(projectID string) ([]MRDescriptionTemplate, error) {
	if c.demoMode {
		c.sleep(300 * time.Millisecond)
		return mrDescriptionTemplatesMock, nil
	}

	type fileResult struct {
		templates []MRDescriptionTemplate
		err       error
	}
	type restResult struct {
		templates []MRDescriptionTemplate
		err       error
	}
	type defaultResult struct {
		content string
		err     error
	}

	fileCh := make(chan fileResult, 1)
	restCh := make(chan restResult, 1)
	defaultCh := make(chan defaultResult, 1)

	fullPath := c.projectFullPath(projectID)

	// Source 1: project-level files via GraphQL
	go func() {
		templates, err := c.fetchProjectFileTemplates(fullPath)
		fileCh <- fileResult{templates, err}
	}()

	// Source 2: group-level templates via REST
	go func() {
		templates, err := c.fetchRESTTemplates(projectID)
		restCh <- restResult{templates, err}
	}()

	// Source 3: project default description
	go func() {
		info, err := c.GetProjectInfo(projectID)
		defaultCh <- defaultResult{info.MergeRequestsTemplate, err}
	}()

	fileRes := <-fileCh
	restRes := <-restCh
	defaultRes := <-defaultCh

	// Merge: project files first, then group templates (dedup by name), then default
	seen := make(map[string]bool)
	var merged []MRDescriptionTemplate

	for _, t := range fileRes.templates {
		if !seen[t.Name] {
			seen[t.Name] = true
			merged = append(merged, t)
		}
	}

	for _, t := range restRes.templates {
		if !seen[t.Name] {
			seen[t.Name] = true
			merged = append(merged, t)
		}
	}

	if len(merged) == 0 && defaultRes.content != "" {
		merged = append(merged, MRDescriptionTemplate{Name: "Default", Content: defaultRes.content})
	}

	return merged, nil
}

func (c *Client) fetchProjectFileTemplates(fullPath graphql.ID) ([]MRDescriptionTemplate, error) {
	var tree getRepositoryTree
	vars := map[string]any{
		"fullPath": fullPath,
		"path":     ".gitlab/merge_request_templates",
		"ref":      "HEAD",
	}

	err := c.gql.Query(context.Background(), &tree, vars)
	if err != nil {
		return nil, err
	}

	entries := tree.Project.Repository.Tree.Blobs.Nodes
	if len(entries) == 0 {
		return nil, nil
	}

	var paths []string
	for _, e := range entries {
		paths = append(paths, e.Path)
	}

	var blobs getRepositoryBlobs
	blobVars := map[string]any{
		"fullPath": fullPath,
		"paths":    paths,
		"ref":      "HEAD",
	}

	err = c.gql.Query(context.Background(), &blobs, blobVars)
	if err != nil {
		return nil, err
	}

	var templates []MRDescriptionTemplate
	for _, b := range blobs.Project.Repository.Blobs.Nodes {
		templates = append(templates, MRDescriptionTemplate{Name: b.Name, Content: b.RawTextBlob})
	}
	return templates, nil
}

func (c *Client) fetchRESTTemplates(projectID string) ([]MRDescriptionTemplate, error) {
	// List available templates
	listURL := fmt.Sprintf("%s/api/v4/projects/%s/templates/merge_requests", c.cfg.BaseURL, projectID)
	req, err := http.NewRequest("GET", listURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("PRIVATE-TOKEN", c.cfg.APIToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}

	var list []struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	// Fetch content for each template
	var templates []MRDescriptionTemplate
	for _, item := range list {
		content, err := c.fetchRESTTemplateContent(projectID, item.Key)
		if err != nil {
			continue
		}
		templates = append(templates, MRDescriptionTemplate{Name: item.Name, Content: content})
	}
	return templates, nil
}

func (c *Client) fetchRESTTemplateContent(projectID, key string) (string, error) {
	url := fmt.Sprintf("%s/api/v4/projects/%s/templates/merge_requests/%s", c.cfg.BaseURL, projectID, key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("PRIVATE-TOKEN", c.cfg.APIToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	var data struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	return data.Content, nil
}

// AcceptMergeRequest merges a merge request via the GitLab API.
func (c *Client) AcceptMergeRequest(
	projectID string,
	input MergeRequestAcceptInput,
) (AcceptMergeRequestResponse, error) {
	if c.demoMode {
		c.sleep(500 * time.Millisecond)
		return AcceptMergeRequestResponse{}, nil
	}

	var mutation acceptMergeRequestMutation
	input.ProjectPath = c.projectFullPath(projectID)
	variables := acceptMergeRequestVariables(input)

	err := c.gql.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		return AcceptMergeRequestResponse{}, err
	}

	return mutation.MergeRequestAccept, nil
}
