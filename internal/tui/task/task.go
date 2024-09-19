package task

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TaskFinishedMsg struct {
	TaskID      string
	SectionID   int
	SectionType string
	Err         error
	Msg         tea.Msg
}
