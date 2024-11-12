package modal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

type Model struct {
	Header  string
	Content string
	ctx     *context.AppContext
}

func New(ctx *context.AppContext) Model {
	return Model{
		ctx: ctx,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// TODO: create keybinds and help model
		if msg.String() == "esc" {
			// TODO: return cmd to be captured in main modal with finishTask to reset statusline
			m.ctx.Task.Err = nil
		}
		// if msg.String() == "ctrl+c" {
		// 	return m, tea.Quit
		// }
	}
	return m, nil
}

func (m Model) View() string {
	h := m.ctx.PanelHeight + projects.DocStyle.GetVerticalFrameSize() + style.MainFrameStyle.GetVerticalFrameSize()
	body := lipgloss.JoinVertical(0,
		headerStyle.Render(m.Header),
		bodyStyle(h).Render(m.ctx.Task.Err.Error()))
	return body
}
