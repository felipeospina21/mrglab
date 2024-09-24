package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/help"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
)

type Model struct {
	Projects      projects.Model
	MergeRequests mergerequests.Model
	Statusline    statusline.Model
	ctx           *context.AppContext
}

func InitMainModel(ctx *context.AppContext) Model {
	// Sets global keybinds by default
	ctx.Keybinds = tui.GlobalKeys
	return Model{
		Projects:      projects.New(ctx),
		MergeRequests: mergerequests.New(ctx),
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
	return cb()
}

type endCommandStatus struct {
	isError   bool
	error     error
	isSuccess bool
}

func endCommand[T any](m *Model, status endCommandStatus, cb func() T) T {
	if status.isError {
		m.setStatus(statusline.ModesEnum.Error, status.error.Error())
	} else {
		m.setStatus(statusline.ModesEnum.Normal, "")
		m.SetHelpKeys(mergerequests.Keybinds)
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

func (m *Model) SetHelpKeys(kb help.KeyMap) {
	m.ctx.Keybinds = kb
}
