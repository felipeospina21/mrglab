package pipelines

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
)

type (
	CycleTabMsg      struct{}
	OpenInBrowserMsg struct{}
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
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		tableFrameX := table.DocStyle.GetHorizontalFrameSize() + 2
		tableW := msg.Width - tableFrameX
		tableH := msg.Height - 1 - 2 - 1
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
	header := fmt.Sprintf("%s - %s", m.ctx.SelectedProject.Name, "Pipelines")
	return tea.NewView(table.RenderPanel(&m.Table, m.Loading, m.SpinnerView, header))
}
