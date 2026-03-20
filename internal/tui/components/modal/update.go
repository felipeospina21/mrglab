package modal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type CloseModalMsg struct{}
type SubmitModalMsg struct{}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		match := tui.KeyMatcher(msg)
		switch {
		case match(Keybinds.Close):
			return m, func() tea.Msg { return CloseModalMsg{} }
		case match(Keybinds.Submit):
			return m, func() tea.Msg { return SubmitModalMsg{} }
		}
	}
	return m, nil
}
