package context

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui/components/help"
)

type focusedPanel uint

const (
	LeftPanel focusedPanel = iota
	MainPanel
	RightPanel
	Modal
)

type TaskStatus uint

const (
	TaskIdle TaskStatus = iota
	TaskStarted
	TaskFinished
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
	TaskStatus       TaskStatus
	TaskErr          error
	IsLeftPanelOpen  bool
	IsRightPanelOpen bool
	IsModalOpen      bool
	FocusedPanel     focusedPanel
	PanelHeight      int
}
