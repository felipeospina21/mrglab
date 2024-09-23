package statusline

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
)

type Modes struct {
	Normal  string
	Insert  string
	Loading string
	Error   string
}

var ModesEnum = Modes{
	Normal:  "NORMAL",
	Insert:  "INSERT",
	Loading: "LOADING",
	Error:   "ERROR",
}

type Model struct {
	Status  string
	Content string
	Width   int
	ctx     *context.AppContext
	Spinner spinner.Model
}

func New(ctx *context.AppContext) Model {
	return Model{
		Status: ModesEnum.Normal,
		Spinner: spinner.New(
			spinner.WithSpinner(spinner.Dot),
			spinner.WithStyle(spinnerStyle),
		),
		ctx: ctx,
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
	projectName := projectStyle.Render(fmt.Sprintf("%s %s", icon.Gitlab, m.ctx.SelectedProject.Name))

	// FIX: without this the spinner is not being updated
	if m.Status == ModesEnum.Loading {
		m.Content = m.Spinner.View()
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
