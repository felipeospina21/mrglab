package projects

import (
	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type FetchMRListMsg struct{}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		match := tui.KeyMatcher(msg)
		if match(Keybinds.MRList) {
			return m, func() tea.Msg { return FetchMRListMsg{} }
		}
	}
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
