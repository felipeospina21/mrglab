package task

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TaskFinishedMsg struct {
	TaskID      TaskID
	SectionType TaskSection
	Err         error
	Msg         tea.Msg
}

type (
	TaskStatus  = uint
	TaskSection = uint
	TaskID      = uint
)

const (
	FetchMRs TaskID = iota
	FetchDiscussions
)

const (
	TaskIdle TaskStatus = iota
	TaskStarted
	TaskFinished
)

const (
	TaskSectionMR TaskSection = iota
)
