// Package app wires together all TUI components into the main Bubble Tea model.
package app

import (
	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"charm.land/bubbles/v2/help"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/modal"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
)

type taskStatus uint

const (
	taskIdle taskStatus = iota
	taskStarted
	taskFinished
)

// Model is the top-level Bubble Tea model that composes all TUI components.
type Model struct {
	Projects      projects.Model
	MergeRequests mergerequests.Model
	Details       details.Model
	Statusline    statusline.Model
	Modal         modal.Model
	Spinner       spinner.Model
	Input         textarea.Model
	layout        Layout
	ctx           *context.AppContext
	taskStatus    taskStatus
	taskErr       error
	isLeftOpen        bool
	isRightOpen       bool
	isRightFullscreen bool
	isModalOpen       bool
	pendingNote   struct {
		DiscussionId string
		NoteableId   string
	}
	pendingCreateMR bool
	pendingConfirm  bool
	formReady       bool
	createForm      createMRForm
}

// InitMainModel creates and returns the initial application model.
func InitMainModel(ctx *context.AppContext, cfg *config.Config, client *gitlab.Client) Model {
	ctx.FocusedPanel = context.LeftPanel
	ctx.DevMode = cfg.DevMode

	ti := textarea.New()
	ti.Placeholder = "Write your reply..."
	ti.CharLimit = 0

	return Model{
		Projects:      projects.New(ctx, client, cfg.Filters.Projects),
		MergeRequests: mergerequests.New(ctx, client),
		Details:       details.New(ctx),
		Statusline:    statusline.New(ctx, projects.Keybinds),
		Modal:         modal.New(ctx),
		Input:         ti,
		Spinner: spinner.New(
			spinner.WithSpinner(spinner.Line),
			spinner.WithStyle(statusline.SpinnerStyle),
		),
		ctx:        ctx,
		taskStatus: taskIdle,
		isLeftOpen: true,
		createForm: newCreateMRForm(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.Statusline.Init(), m.Spinner.Tick)
}

func (m *Model) setStatus(mode string, content string) {
	switch mode {
	case statusline.ModesEnum.Normal,
		statusline.ModesEnum.Loading,
		statusline.ModesEnum.Error,
		statusline.ModesEnum.Dev:
		m.Statusline.Status = mode
		m.Statusline.Content = content
	default:
		m.Statusline.Content = "status not supported"
	}
}

func (m *Model) startTask(cb func() tea.Cmd) tea.Cmd {
	m.taskStatus = taskStarted
	m.setStatus(statusline.ModesEnum.Loading, m.Statusline.Spinner.View())
	return cb()
}

func (m *Model) finishTask(err error, kb help.KeyMap) {
	if err != nil {
		m.setStatus(statusline.ModesEnum.Error, err.Error())
		m.taskErr = err
	} else {
		mode := statusline.ModesEnum.Normal
		if m.ctx.DevMode {
			mode = statusline.ModesEnum.Dev
		}
		m.setStatus(mode, "")
		m.setHelpKeys(kb)
		m.taskErr = nil
	}
	m.taskStatus = taskFinished
}

func (m *Model) updateSpinnerViewCommand(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	if m.Statusline.Status == statusline.ModesEnum.Loading {
		m.Statusline.Content = m.Statusline.Spinner.View()
	}
	m.Statusline.Spinner, cmd = m.Statusline.Spinner.Update(msg)

	return cmd
}

func (m *Model) toggleLeftPanel() {
	m.isLeftOpen = !m.isLeftOpen
	m.recomputeLayout()
}

func (m *Model) toggleRightPanel() {
	m.isRightOpen = !m.isRightOpen
	m.recomputeLayout()
}

func (m *Model) recomputeLayout() {
	m.layout = computeLayout(m.ctx.Window, m.isLeftOpen, m.isRightOpen, m.isRightFullscreen)
	m.applyLayout()
}

func (m *Model) setHelpKeys(kb help.KeyMap) {
	m.Statusline.Keybinds = kb
}

// SelectMR stores the currently selected merge request's IID, SHA, and status in the app context.
func (m *Model) SelectMR() {
	idColIdx := mergerequests.GetColIndex(mergerequests.ColNames.ID)
	m.ctx.SelectedMR.IID = m.MergeRequests.Table.SelectedRow()[idColIdx]

	shaColIdx := mergerequests.GetColIndex(mergerequests.ColNames.Sha)
	m.ctx.SelectedMR.Sha = m.MergeRequests.Table.SelectedRow()[shaColIdx]

	statusColIdx := mergerequests.GetColIndex(mergerequests.ColNames.Status)
	m.ctx.SelectedMR.Status = m.MergeRequests.Table.SelectedRow()[statusColIdx]
}
