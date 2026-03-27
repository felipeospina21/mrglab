package mergerequests

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui"
)

// Messages returned to app for actions requiring app-level coordination
type (
	ViewDetailsMsg    struct{}
	MergeMRMsg        struct{}
	OpenInBrowserMsg  struct{}
	CreateMRMsg       struct{}
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		match := tui.KeyMatcher(msg)
		switch {
		case match(Keybinds.Details):
			return m, func() tea.Msg { return ViewDetailsMsg{} }
		case match(Keybinds.Merge):
			return m, func() tea.Msg { return MergeMRMsg{} }
		case match(Keybinds.OpenInBrowser):
			return m, func() tea.Msg { return OpenInBrowserMsg{} }
		case match(Keybinds.CreateMR):
			return m, func() tea.Msg { return CreateMRMsg{} }
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}
