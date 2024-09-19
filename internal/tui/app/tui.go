package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
	mergerequests "github.com/felipeospina21/mrglab/internal/tui/components/merge_requests"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

type Model struct {
	Projects      projects.Model
	MergeRequests mergerequests.Model
	ctx           *context.AppContext
}

func InitMainModel(ctx *context.AppContext) Model {
	return Model{
		Projects:      projects.New(ctx),
		MergeRequests: mergerequests.New(ctx),
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
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			m.Projects.SelectProject()
			cmd = m.MergeRequests.GetMRListCmd()
			cmds = append(cmds, cmd)

		}

	case tea.WindowSizeMsg:
		m.ctx.Window = msg
		h, v := projects.ItemStyle.GetFrameSize()
		th, tv := projects.TitleStyle.GetFrameSize()
		m.Projects.List.SetSize(msg.Width-h-th, msg.Height-v-tv)

	case task.TaskFinishedMsg:
		// TODO: Update Table with msg.Msg (table rows)
	}

	m.Projects.List, cmd = m.Projects.List.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	t := table.TitleStyle.Render("Select Project")
	if m.Projects.IsOpen {
		p := projects.DocStyle.Render(m.Projects.List.View())

		return lipgloss.JoinHorizontal(0, p, t)
	}

	return t
}
