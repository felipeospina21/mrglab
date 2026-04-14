// Package context provides shared application state accessible by all TUI components.
package context

import "github.com/felipeospina21/tuishell"

// Re-export panel constants for convenience.
const (
	LeftPanel  = tuishell.LeftPanel
	MainPanel  = tuishell.MainPanel
	RightPanel = tuishell.RightPanel
	Modal      = tuishell.ModalPanel
)

// AppContext holds shared state passed to all TUI components.
type AppContext struct {
	tuishell.AppContext
	SelectedProject struct {
		Name string
		ID   string
	}
	SelectedMR struct {
		IID    string
		Sha    string
		Status string
	}
}
