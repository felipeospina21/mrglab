package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
)

// command wraper that takes care of initializing spinner
// & setting corresponding status
func (m *Model) startCommand(cb func() tea.Cmd) tea.Cmd {
	m.Statusline.Status = statusline.StartSpinner()
	return cb()
}
