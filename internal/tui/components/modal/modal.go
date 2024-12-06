package modal

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

type ModalContent struct {
	Header string
	Body   string
}

type Model struct {
	Content  ModalContent
	Input    textarea.Model
	Editable bool
	ctx      *context.AppContext
}

func New(ctx *context.AppContext) Model {
	ti := textarea.New()
	ti.Focus()
	return Model{
		ctx:   ctx,
		Input: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	h := m.ctx.PanelHeight + projects.DocStyle.GetVerticalFrameSize() + style.MainFrameStyle.GetVerticalFrameSize()
	body := lipgloss.JoinVertical(0,
		m.Content.Body,
		helpStyle.Render("Press esc to close modal"),
	)
	content := lipgloss.JoinVertical(0,
		headerStyle.Render(m.Content.Header),
		bodyStyle(h).Render(body),
	)

	if m.Editable {
		content = lipgloss.JoinVertical(0,
			headerStyle.Render(m.Content.Header),
			m.Input.View(),
		)
	}

	return content
}

func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.Modal
}

func (m *Model) ResetContent() {
	m.Content.Body = ""
	m.Content.Header = ""
}
