package app

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/help"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/modal"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
)

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
}

func InitMainModel(ctx *context.AppContext) Model {
	ctx.Keybinds = projects.Keybinds
	ctx.FocusedPanel = context.LeftPanel
	ctx.TaskStatus = context.TaskIdle

	return Model{
		Projects:      projects.New(ctx),
		MergeRequests: mergerequests.New(ctx),
		Details:       details.New(ctx),
		Statusline:    statusline.New(ctx),
		Modal:         modal.New(ctx),
		Spinner: spinner.New(
			spinner.WithSpinner(spinner.Line),
			spinner.WithStyle(statusline.SpinnerStyle),
		),
		ctx: ctx,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.Statusline.Init(), m.Spinner.Tick)
}

func (m *Model) setStatus(mode string, content string) {
	switch mode {
	case statusline.ModesEnum.Normal:
		fallthrough
	case statusline.ModesEnum.Loading:
		fallthrough
	case statusline.ModesEnum.Insert:
		fallthrough
	case statusline.ModesEnum.Error:
		fallthrough
	case statusline.ModesEnum.Dev:
		m.Statusline.Status = mode
		m.Statusline.Content = content
	default:
		m.Statusline.Content = "status not supported"
	}
}

func (m *Model) startTask(cb func() tea.Cmd) tea.Cmd {
	m.ctx.TaskStatus = context.TaskStarted
	m.setStatus(statusline.ModesEnum.Loading, m.Statusline.Spinner.View())
	return cb()
}

func (m *Model) finishTask(err error, kb help.KeyMap) {
	if err != nil {
		m.setStatus(statusline.ModesEnum.Error, err.Error())
		m.ctx.TaskErr = err
	} else {
		mode := statusline.ModesEnum.Normal
		if config.GlobalConfig.DevMode {
			mode = statusline.ModesEnum.Dev
		}
		m.setStatus(mode, "")
		m.setHelpKeys(kb)
		m.ctx.TaskErr = nil
	}
	m.ctx.TaskStatus = context.TaskFinished
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
	m.ctx.IsLeftPanelOpen = !m.ctx.IsLeftPanelOpen
	m.recomputeLayout()
}

func (m *Model) toggleRightPanel() {
	m.ctx.IsRightPanelOpen = !m.ctx.IsRightPanelOpen
	m.recomputeLayout()
}

func (m *Model) recomputeLayout() {
	m.layout = computeLayout(m.ctx.Window, m.ctx.IsLeftPanelOpen, m.ctx.IsRightPanelOpen)
	m.applyLayout()
}

func (m *Model) setHelpKeys(kb help.KeyMap) {
	m.ctx.Keybinds = kb
}

func (m *Model) SelectMR() {
	idColIdx := mergerequests.GetColIndex(mergerequests.ColNames.ID)
	m.ctx.SelectedMR.IID = m.MergeRequests.Table.SelectedRow()[idColIdx]

	shaColIdx := mergerequests.GetColIndex(mergerequests.ColNames.Sha)
	m.ctx.SelectedMR.Sha = m.MergeRequests.Table.SelectedRow()[shaColIdx]

	statusColIdx := mergerequests.GetColIndex(mergerequests.ColNames.Status)
	m.ctx.SelectedMR.Status = m.MergeRequests.Table.SelectedRow()[statusColIdx]
}
