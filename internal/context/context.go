package context

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui/components/help"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

type AppContext struct {
	SelectedProject struct {
		Name string
		ID   string
	}

	SelectedMRID string
	Window       tea.WindowSizeMsg
	Keybinds     help.KeyMap
	TaskStatus   task.TaskStatus
}
