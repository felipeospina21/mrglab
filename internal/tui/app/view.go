package app

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	return m.Shell.RenderView()
}
