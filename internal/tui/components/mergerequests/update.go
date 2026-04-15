package mergerequests

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
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
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Set table dimensions
		tableFrameX := table.DocStyle.GetHorizontalFrameSize() + 2 // +2 for border
		tableW := msg.Width - tableFrameX
		tableH := msg.Height - 1 - 2 - 1 // header, border, overhead
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

// View returns the panel content as a tea.View.
func (m Model) View() tea.View {
	header := fmt.Sprintf("%s - %s", m.ctx.SelectedProject.Name, "Merge Requests")
	return tea.NewView(table.RenderPanel(&m.Table, m.Loading, m.SpinnerView, header))
}
