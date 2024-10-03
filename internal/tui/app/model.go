package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/help"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

type Model struct {
	Projects      projects.Model
	MergeRequests mergerequests.Model
	Details       details.Model
	Statusline    statusline.Model
	ctx           *context.AppContext
}

func InitMainModel(ctx *context.AppContext) Model {
	// Sets global keybinds by default
	ctx.Keybinds = tui.GlobalKeys
	ctx.FocusedPanel = context.LeftPanel
	ctx.TaskStatus = task.TaskIdle

	return Model{
		Projects:      projects.New(ctx),
		MergeRequests: mergerequests.New(ctx),
		Details:       details.New(ctx),
		Statusline:    statusline.New(ctx),
		ctx:           ctx,
		// 	isSidePanelOpen: false,
		// 	CurrView:      HomeView,
		// 	Help:          components.Help{Model: help.New()},
		// 	MergeRequests: MergeRequestsModel{},
		// 	Toast: toast.New(toast.Model{
		// 		Progress: progress.New(
		// 			progress.WithDefaultGradient(),
		// 			progress.WithFillCharacters('-', ' '),
		// 			progress.WithoutPercentage(),
		// 		),
		// 		Interval: 10,
		// 		// Type:     toast.Info,
		// 		// Show:     true,
		// 		// Message:  "Info msg",
		// 	}),
		// 	Tabs: tabs.Model{
		// 		Tabs: []string{"Merge Requests", "Issues", "Pipelines"},
		// 	},
		// 	Statusline:      statusline.Model{Status: statusline.Modes.Normal},
		// 	Paginator:       p,
	}
}

func (m Model) Init() tea.Cmd {
	return m.Statusline.Spinner.Tick
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

// command wraper that takes care of initializing spinner
// & setting corresponding status
func (m *Model) startCommand(cb func() tea.Cmd) tea.Cmd {
	m.setStatus(statusline.ModesEnum.Loading, m.Statusline.Spinner.View())
	m.startTask()
	return cb()
}

func endCommand[T any](m *Model, msg task.TaskFinishedMsg, cb func() T) T {
	if msg.Err != nil {
		m.setStatus(statusline.ModesEnum.Error, msg.Err.Error())
	} else {
		mode := statusline.ModesEnum.Normal
		if m.ctx.IsDevMode {
			mode = statusline.ModesEnum.Dev
		}
		m.setStatus(mode, "")
		m.setHelpKeys(mergerequests.Keybinds)
		m.finishTask()
	}
	return cb()
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
	m.MergeRequests.Table.SetWidth(lipgloss.Width(m.MergeRequests.Table.View()))
	m.MergeRequests.Table.UpdateViewport()
}

func (m *Model) setHelpKeys(kb help.KeyMap) {
	m.ctx.Keybinds = kb
}

func (m *Model) startTask() {
	m.ctx.TaskStatus = task.TaskStarted
}

func (m *Model) finishTask() {
	m.ctx.TaskStatus = task.TaskFinished
}

func getFrameSize() (int, int) {
	xMain, yMain := MainFrameStyle.GetFrameSize()
	xProjects, yProjects := projects.GetFrameSize()
	xStatus, yStatus := statusline.GetFrameSize()

	return xMain + xProjects + xStatus, yMain + yProjects + yStatus
}

func (m Model) getEmptyTableSize() (int, int) {
	w, h := m.ctx.Window.Width, m.ctx.Window.Height
	leftPanX, leftPanY := projects.GetFrameSize()
	leftPanW := m.Projects.List.Width()
	tableX := table.TitleStyle.GetHorizontalFrameSize()
	statusHeight := lipgloss.Height(m.Statusline.View())

	width := w - leftPanX - leftPanW - tableX
	height := h - leftPanY - statusHeight - MainFrameStyle.GetVerticalFrameSize()

	return width, height
}

func (m *Model) setLeftPanelHeight() {
	_, y := getFrameSize()
	yStatus := lipgloss.Height(m.Statusline.View())
	height := m.ctx.Window.Height - y - yStatus - 3 // FIX: find how to replace this magic num

	m.Projects.List.SetHeight(height)
	m.ctx.PanelHeight = height
}

func (m *Model) setStatuslineWidth() {
	windowW := m.ctx.Window.Width
	xStatus, _ := statusline.GetFrameSize()
	m.Statusline.Width = windowW - xStatus
}
