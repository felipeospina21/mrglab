package app

import (
	"log"
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

type Model struct {
	Projects      projects.Model
	MergeRequests mergerequests.Model
	Statusline    statusline.Model
	ctx           *context.AppContext
}

func InitMainModel(ctx *context.AppContext) Model {
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case error:
		logger.Debug("err", func() {
			log.Println(msg)
		})
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			cb := func() tea.Cmd {
				m.Projects.SelectProject()
				return m.MergeRequests.GetMRListCmd()
			}
			cmd = m.startCommand(cb)
			cmds = append(cmds, cmd)

		}
	case spinner.TickMsg:
		m.Statusline.Spinner, cmd = m.Statusline.Spinner.Update(msg)
		cmds = append(cmds, cmd)

	case tea.WindowSizeMsg:
		m.ctx.Window = msg
		h, v := projects.ItemStyle.GetFrameSize()
		th, tv := projects.TitleStyle.GetFrameSize()
		m.Projects.List.SetSize(msg.Width-h-th, msg.Height-v-tv)
		m.Statusline.Width = msg.Width - statusline.StatusBarStyle.GetHorizontalFrameSize() - table.DocStyle.GetHorizontalFrameSize()

	case task.TaskFinishedMsg:
		// TODO: Rethink this logic
		if msg.SectionType == "mrs" {
			ml := msg.Msg.(mergerequests.MergeRequestsFetchedMsg)
			var rows []table.Row
			for _, mr := range ml.Mrs {
				r := table.Row{
					mr.CreatedAt.String(),
					strconv.FormatBool(mr.Draft),
					mr.Title,
					mr.Author.Name,
					mr.DetailedMergeStatus,
					strconv.FormatBool(mr.HasConflicts),
					strconv.Itoa(mr.UserNotesCount),
					mr.ChangesCount,
					mr.WebURL,
					mr.Description,
					strconv.Itoa(mr.IID),
				}

				rows = append(rows, r)
			}
			m.MergeRequests.Table = table.InitModel(table.InitModelParams{
				Rows:   rows,
				Colums: mergerequests.GetMergeReqsColums(m.ctx.Window.Width - 10),
				// StyleFunc: mergerequests.StyleIconsColumns(table.Styles(table.DefaultStyle()), table.MergeReqsIconCols),
			})
		}
		m.Statusline.Status = statusline.StopSpinner()
	}

	m.Projects.List, cmd = m.Projects.List.Update(msg)
	m.MergeRequests.Table, cmd = m.MergeRequests.Table.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.Projects.IsOpen {
		t := table.TitleStyle.Render("Select Project")
		p := projects.DocStyle.Render(m.Projects.List.View())

		return lipgloss.JoinHorizontal(0, p, t)
	}
	table := table.DocStyle.Render(m.MergeRequests.Table.View())
	h := m.ctx.Window.Height - lipgloss.Height(table)
	sl := lipgloss.PlaceVertical(h, lipgloss.Bottom, m.Statusline.View())
	return lipgloss.JoinVertical(0, table, sl)
}
