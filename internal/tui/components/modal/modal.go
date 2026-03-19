package modal

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
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
	header := headerStyle.Render(m.Header)
	h := m.ctx.PanelHeight - lipgloss.Height(header)
	content := lipgloss.JoinVertical(0,
		m.Content,
		helpStyle.Render("Press esc to close modal"),
	)
	body := lipgloss.JoinVertical(0,
		header,
		bodyStyle(h).Render(content),
	)

	return body
}

func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.Modal
}
