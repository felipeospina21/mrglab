package modal

import (
	"time"

	tea "charm.land/bubbletea/v2"
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
	case tea.KeyPressMsg:
		match := tui.KeyMatcher(msg)
		switch {
		case match(Keybinds.Close):
			return m, func() tea.Msg { return CloseModalMsg{} }
		case m.HasSubmit && match(Keybinds.Submit):
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
