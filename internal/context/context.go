package context

import tea "github.com/charmbracelet/bubbletea"

type focusedPanel uint

const (
	LeftPanel focusedPanel = iota
	MainPanel
	RightPanel
	Modal
)

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
