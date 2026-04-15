package pipelines

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/tuishell/style"
)

// CycleTabMsg signals the app to cycle to the next tab.
type CycleTabMsg struct{}

// Init returns nil (no initialization needed).
func (m Model) Init() tea.Cmd { return nil }

// Update handles key events and window resize for the pipelines panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		match := tui.KeyMatcher(msg)
		if match(Keybinds.CycleTab) {
			return m, func() tea.Msg { return CycleTabMsg{} }
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

// View returns a centered placeholder.
func (m Model) View() tea.View {
	content := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center).
		Foreground(style.DefaultTheme().TextDimmed).
		Render("Pipelines")
	return tea.NewView(content)
}
