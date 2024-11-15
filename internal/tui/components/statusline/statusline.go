package statusline

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/help"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
)

type Modes struct {
	Normal  string
	Insert  string
	Loading string
	Error   string
	Dev     string
}

var ModesEnum = Modes{
	Normal:  "NORMAL",
	Insert:  "INSERT",
	Loading: "LOADING",
	Error:   "ERROR",
	Dev:     "DEVELOP",
}

type Model struct {
	Status  string
	Content string
	Width   int
	Spinner spinner.Model
	Help    help.Model
	ctx     *context.AppContext
}

func New(ctx *context.AppContext) Model {
	status := ModesEnum.Normal
	if config.GlobalConfig.DevMode {
		status = ModesEnum.Dev
	}
	return Model{
		Status: status,
		Spinner: spinner.New(
			spinner.WithSpinner(spinner.Dot),
			spinner.WithStyle(spinnerStyle),
		),
		ctx:  ctx,
		Help: help.New(),
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
	statusVal := statusText.Render(tui.Truncate(m.Content, width/4))
	encoding := encodingStyle.Render("UTF-8")
	projectName := projectStyle.Render(fmt.Sprintf("%s %s", icon.Gitlab, m.ctx.SelectedProject.Name))

	help := helpText.
		Width(width - w(statusKey) - w(statusVal) - w(encoding) - w(projectName)).
		Render(" " + m.Help.View(m.ctx.Keybinds) + " ")

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		help,
		encoding,
		projectName,
	)

	return StatusBarStyle.Render(bar)
}

func GetFrameSize() (int, int) {
	x, y := StatusBarStyle.GetFrameSize()
	xNugget, yNugget := statusNugget.GetFrameSize()
	xStatus, yStatus := statusStyle.GetFrameSize()

	return x + xNugget + xStatus, y + yNugget + yStatus
}
