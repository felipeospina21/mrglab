package modal

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
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

type DiscussionForm struct {
	Number      textinput.Model
	Response    textarea.Model
	Discussions string
}

type Model struct {
	Content        ModalContent
	DiscussionForm DiscussionForm
	Editable       bool
	ctx            *context.AppContext
}

func New(ctx *context.AppContext) Model {
	ti := textinput.New()
	ti.Focus()

	ta := textarea.New()
	// TODO: set this height dinamic
	ta.SetHeight(10)
	ta.SetWidth(50)

	return Model{
		ctx: ctx,
		DiscussionForm: DiscussionForm{
			Number:   ti,
			Response: ta,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, textinput.Blink)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var updateInput, updateTextArea tea.Cmd
	m.DiscussionForm.Response, updateInput = m.DiscussionForm.Response.Update(msg)
	m.DiscussionForm.Number, updateTextArea = m.DiscussionForm.Number.Update(msg)
	return m, tea.Batch(updateInput, updateTextArea)
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
		// TODO: Set dinamic width & height for both panels
		// Make discussions a scrollable viewport (render only open discussions)
		formStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder())
		form := formStyle.Render(lipgloss.JoinVertical(0,
			"Discussion Number",
			m.DiscussionForm.Number.View(),
			"\n\n",
			"Comment",
			m.DiscussionForm.Response.View(),
		))
		body := lipgloss.JoinHorizontal(0, form, m.DiscussionForm.Discussions)
		content = lipgloss.JoinVertical(0,
			headerStyle.Render(m.Content.Header),
			bodyStyle(h).Render(body),
		)
	}

	return content
}

func (m *Model) CycleFocus() {
	if m.DiscussionForm.Number.Focused() {
		m.DiscussionForm.Number.Blur()
		m.DiscussionForm.Response.Focus()
	} else {
		m.DiscussionForm.Response.Blur()
		m.DiscussionForm.Number.Focus()
	}
}

func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.Modal
}

func (m *Model) ResetContent() {
	m.Content.Body = ""
	m.Content.Header = ""
}
