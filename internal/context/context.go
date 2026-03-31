// Package context provides shared application state accessible by all TUI components.
package context

import tea "charm.land/bubbletea/v2"

type focusedPanel uint

// Panel focus constants.
const (
	LeftPanel focusedPanel = iota
	MainPanel
	RightPanel
	Modal
)

// AppContext holds shared state passed to all TUI components.
type AppContext struct {
	SelectedProject struct {
		Name string
		ID   string
	}
	SelectedMR struct {
		IID    string
		Sha    string
		Status string
	}
	Window       tea.WindowSizeMsg
	DevMode      bool
	FocusedPanel focusedPanel
	PanelHeight  int
}
