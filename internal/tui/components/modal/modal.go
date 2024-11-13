package modal

import (
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

func (m Model) View() string {
	h := m.ctx.PanelHeight + projects.DocStyle.GetVerticalFrameSize() + style.MainFrameStyle.GetVerticalFrameSize()
	content := lipgloss.JoinVertical(0,
		m.Content,
		helpStyle.Render("Press esc to close modal"),
	)
	body := lipgloss.JoinVertical(0,
		headerStyle.Render(m.Header),
		bodyStyle(h).Render(content),
	)

	return body
}

func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.Modal
}
