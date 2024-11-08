package context

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui/components/help"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

type focusedPanel uint

const (
	LeftPanel focusedPanel = iota
	MainPanel
	RightPanel
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
	Window           tea.WindowSizeMsg
	Keybinds         help.KeyMap
	Task             task.TaskMsg
	IsLeftPanelOpen  bool
	IsRightPanelOpen bool
	FocusedPanel     focusedPanel
	PanelHeight      int
}
