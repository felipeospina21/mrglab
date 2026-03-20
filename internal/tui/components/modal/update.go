package modal

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type (
	CloseModalMsg     struct{}
	SubmitModalMsg    struct{}
	CopyModalMsg      struct{}
	ResetHighlightMsg struct{}
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		match := tui.KeyMatcher(msg)
		switch {
		case match(Keybinds.Close):
			return m, func() tea.Msg { return CloseModalMsg{} }
		case match(Keybinds.Submit):
			return m, func() tea.Msg { return SubmitModalMsg{} }
		case match(Keybinds.Copy):
			m.Highlight = true
			return m, tea.Batch(
				func() tea.Msg { return CopyModalMsg{} },
				tea.Tick(100*time.Millisecond, func(time.Time) tea.Msg { return ResetHighlightMsg{} }),
			)
		}
	case ResetHighlightMsg:
		m.Highlight = false
		return m, nil
	}
	return m, nil
}
