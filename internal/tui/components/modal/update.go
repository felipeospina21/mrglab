package modal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type CloseModalMsg struct{}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		match := tui.KeyMatcher(msg)
		if match(Keybinds.Close) {
			return m, func() tea.Msg { return CloseModalMsg{} }
		}
	}
	return m, nil
}
