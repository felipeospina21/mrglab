// Package export exposes the mrglab TUI app for embedding in launchers.
package export

import (
	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/tui/app"
)

// NewApp returns an initialized mrglab tea.Model ready to run.
func NewApp() tea.Model {
	return app.NewApp()
}
