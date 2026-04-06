package projects

import (
	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/tui"
)

// FetchMRListMsg signals that the user selected a project and wants to fetch MRs.
type FetchMRListMsg struct{}

// Init returns nil (no initialization needed).
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles key events for the projects panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		match := tui.KeyMatcher(msg)
		if match(Keybinds.MRList) {
			m.SelectProject()
			return m, func() tea.Msg { return FetchMRListMsg{} }
		}
	case tea.WindowSizeMsg:
		// Account for list title (2 lines) and some padding
		m.List.SetHeight(msg.Height - 3)
		m.List.SetWidth(msg.Width)
	}
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

// View returns the panel content as a tea.View.
func (m Model) View() tea.View {
	return tea.NewView(m.List.View())
}
