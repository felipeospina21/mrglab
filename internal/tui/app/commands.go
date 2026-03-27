package app

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/exec"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
)

func (m Model) getMergeRequestModel(msg tui.MRListFetchedMsg) func() table.Model {
	return func() table.Model {
		rows := mergerequests.GetTableRows(msg.Mrs)
		tableW := m.layout.MainPanel.Width - table.DocStyle.GetHorizontalFrameSize() - tableBorderX
		return table.InitModel(table.InitModelParams{
			Rows:   rows,
			Colums: mergerequests.GetTableColums(tableW),
			StyleFunc: table.StyleIconsColumns(
				table.Styles(table.DefaultStyle()),
				mergerequests.IconCols(),
			),
			Height: m.layout.ContentH - mainPanelHeaderLines - tableBorderY - tableViewOverhead,
		})
	}
}

func (m Model) getMergeRequestDetails(msg tui.MRDetailsFetchedMsg) func() details.MergeRequestDetails {
	return func() details.MergeRequestDetails {
		return details.MergeRequestDetails{
			Pipelines:   msg.Stages,
			Discussions: msg.Discussions,
			Branches:    msg.Branches,
			Approvals:   msg.Approvals,
		}
	}
}

func (m Model) acceptMergeRequest() tea.Cmd {
	m.SelectMR()
	return m.MergeRequests.AcceptMergeRequest()
}

func (m Model) fetchMergeRequestsList() tea.Cmd {
	m.Projects.SelectProject()
	return m.Projects.GetListCmd()
}

func (m Model) fetchSingleMergeRequest() tea.Cmd {
	m.SelectMR()
	return m.MergeRequests.FetchMergeRequest()
}

func (m *Model) openInBrowser() {
	colIdx := mergerequests.GetColIndex(mergerequests.ColNames.URL)
	url := m.MergeRequests.Table.SelectedRow()[colIdx]
	exec.Openbrowser(url)
}

func (m Model) createMergeRequest() tea.Cmd {
	source := strings.TrimSpace(m.createForm.source.Value())
	target := strings.TrimSpace(m.createForm.target.Value())
	title := strings.TrimSpace(m.createForm.title.Value())
	description := strings.TrimSpace(m.createForm.description.Value())

	if target == "" {
		target = "main"
	}

	if source == "" || title == "" {
		return func() tea.Msg {
			return tui.MRCreatedMsg{
				Err: fmt.Errorf("source branch and title are required"),
			}
		}
	}
	return m.MergeRequests.CreateMergeRequest(gitlab.CreateMergeRequestInput{
		SourceBranch: source,
		TargetBranch: target,
		Title:        title,
		Description:  description,
	})
}

// parseBranches parses "source→target" or "source->target" format.
// If no target is given, defaults to "main".
func parseBranches(input string) (source, target string) {
	sep := "→"
	if !strings.Contains(input, sep) {
		sep = "->"
	}

	parts := strings.SplitN(input, sep, 2)
	source = strings.TrimSpace(parts[0])
	if len(parts) > 1 {
		target = strings.TrimSpace(parts[1])
	}
	if target == "" {
		target = "main"
	}
	return source, target
}
