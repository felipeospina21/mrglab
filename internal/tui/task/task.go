package task

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TaskMsg struct {
	TaskID      TaskID
	SectionType TaskSection
	Status      TaskStatus
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
	FetchPipeline
)

const (
	TaskIdle TaskStatus = iota
	TaskStarted
	TaskFinished
)

const (
	TaskSectionMR TaskSection = iota
)
