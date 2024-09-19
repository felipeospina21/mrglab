package context

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppContext struct {
	SelectedProject struct {
		Name string
		ID   string
	}

	SelectedMRID string
	Window       tea.WindowSizeMsg
}
