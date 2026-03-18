package details

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type (
	ClosePanelMsg     struct{}
	MergeMRMsg        struct{}
	OpenInBrowserMsg  struct{}
	RespondCommentMsg struct{}
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		match := tui.KeyMatcher(msg)
		switch {
		case match(Keybinds.ClosePanel):
			return m, func() tea.Msg { return ClosePanelMsg{} }
		case match(Keybinds.Merge):
			return m, func() tea.Msg { return MergeMRMsg{} }
		case match(Keybinds.OpenInBrowser):
			return m, func() tea.Msg { return OpenInBrowserMsg{} }
		case match(Keybinds.RespondComment):
			return m, func() tea.Msg { return RespondCommentMsg{} }
		}
	}
	m.Viewport, cmd = m.Viewport.Update(msg)
	return m, cmd
}
