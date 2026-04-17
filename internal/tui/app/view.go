package app

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	v := m.Shell.RenderView()
	if m.statusFilter.IsOpen() {
		w := m.ctx.Window.Width
		h := m.ctx.Window.Height
		screen := m.statusFilter.View(v.Content, w, h)
		return tea.View{Content: screen, AltScreen: v.AltScreen}
	}
	return v
}
