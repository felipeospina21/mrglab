package mergerequests

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
)

// Messages returned to app for actions requiring app-level coordination
type (
	ViewDetailsMsg   struct{}
	MergeMRMsg       struct{}
	OpenInBrowserMsg struct{}
	CreateMRMsg      struct{}
	ReFetchMRListMsg struct{}
	CycleTabMsg      struct{}
)

// Init returns nil (no initialization needed).
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles key events for the merge requests panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		match := tui.KeyMatcher(msg)
		switch {
		case match(Keybinds.Details):
			return m, func() tea.Msg { return ViewDetailsMsg{} }
		case match(Keybinds.Merge):
			return m, func() tea.Msg { return MergeMRMsg{} }
		case match(Keybinds.OpenInBrowser):
			return m, func() tea.Msg { return OpenInBrowserMsg{} }
		case match(Keybinds.CreateMR):
			return m, func() tea.Msg { return CreateMRMsg{} }
		case match(Keybinds.Refetch):
			return m, func() tea.Msg { return ReFetchMRListMsg{} }
		case match(Keybinds.CycleTab):
			return m, func() tea.Msg { return CycleTabMsg{} }
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Set table dimensions
		tableFrameX := table.DocStyle.GetHorizontalFrameSize() + 2 // +2 for border
		tableW := msg.Width - tableFrameX
		tableH := msg.Height - 2 - 1 // border + overhead
		m.Table.W = tableW
		m.Table.H = tableH
		m.Table.SetWidth(tableW)
		m.Table.SetHeight(tableH)
		if len(m.Table.Rows()) > 0 {
			m.Table.SetColumns(GetTableColums(tableW))
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

// Header returns the panel header text.
func (m Model) Header() string {
	return fmt.Sprintf("%s", m.ctx.SelectedProject.Name)
}

// View returns the panel content as a tea.View.
func (m Model) View() tea.View {
	if m.Loading {
		content := lipgloss.NewStyle().
			Width(m.Table.W).
			Height(m.Table.H).
			Align(lipgloss.Center, lipgloss.Center).
			Render(m.SpinnerView + " Loading...")
		return tea.NewView(content)
	}
	return tea.NewView(table.DocStyle.Render(m.Table.View()))
}
