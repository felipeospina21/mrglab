package statusline

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
)

type modes struct {
	Normal  string
	Insert  string
	Loading string
}

var Modes = modes{
	Normal:  "NORMAL",
	Insert:  "INSERT",
	Loading: "LOADING",
}

type Model struct {
	Status  string
	Content string
	Width   int
	ctx     *context.AppContext
	Spinner spinner.Model
	// Height  int
}

func New(ctx *context.AppContext) Model {
	return Model{
		Status: Modes.Normal,
		Spinner: spinner.New(
			spinner.WithSpinner(spinner.Dot),
			spinner.WithStyle(spinnerStyle),
		),
		ctx: ctx,
		// Content: m.Content,
	}
}

func (m Model) Init() tea.Cmd {
	return m.Spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		return m, nil
	default:
		return m, cmd
	}
}

func (m Model) View() string {
	width := m.Width
	w := lipgloss.Width

	statusKey := statusStyle.Render(m.Status)
	encoding := encodingStyle.Render("UTF-8")
	// TODO: move icon to its package
	projectName := projectStyle.Render(fmt.Sprintf(" %s", m.ctx.SelectedProject.Name))
	// TODO: refactor this logic to a separate function
	if m.Status == Modes.Loading {
		m.Content = m.Spinner.View()
		// return m.spinner.View()
	} else if m.Status == Modes.Normal {
		m.Content = ""
	}
	statusVal := statusText.
		Width(width - w(statusKey) - w(encoding) - w(projectName)).
		Render(m.Content)

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		encoding,
		projectName,
	)

	return StatusBarStyle.Render(bar)
}

// TODO: add a function to control statuses

func StartSpinner() string {
	// TODO: add spinner initialization & model (to statusline model)

	return Modes.Loading
}

func StopSpinner() string {
	// TODO: add spinner initialization & model (to statusline model)
	return Modes.Normal
}

func (m Model) startSpinner() {
	m.Content = m.Spinner.View()
}
