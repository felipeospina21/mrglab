package pipelines

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
)

type (
	CycleTabMsg            struct{}
	OpenInBrowserMsg       struct{}
	ViewDetailsMsg         struct{}
	RetryPipelineMsg       struct{}
	CancelPipelineMsg      struct{}
	ReFetchPipelineListMsg struct{}
)

// Init returns nil (no initialization needed).
func (m Model) Init() tea.Cmd { return nil }

// Update handles key events and window resize for the pipelines panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		match := tui.KeyMatcher(msg)
		switch {
		case match(Keybinds.CycleTab):
			return m, func() tea.Msg { return CycleTabMsg{} }

		case match(Keybinds.OpenInBrowser):
			return m, func() tea.Msg { return OpenInBrowserMsg{} }

		case match(Keybinds.Details):
			return m, func() tea.Msg { return ViewDetailsMsg{} }

		case match(Keybinds.RetryPipeline):
			return m, func() tea.Msg { return RetryPipelineMsg{} }

		case match(Keybinds.CancelPipeline):
			return m, func() tea.Msg { return CancelPipelineMsg{} }

		case match(Keybinds.Refetch):
			return m, func() tea.Msg { return ReFetchPipelineListMsg{} }
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		tableFrameX := table.DocStyle.GetHorizontalFrameSize() + 2
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
