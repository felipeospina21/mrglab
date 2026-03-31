// Package statusline implements the bottom status bar component.
package statusline

import (
	"fmt"

	"image/color"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

// Modes defines the possible status bar mode labels.
type Modes struct {
	Normal  string
	Loading string
	Error   string
	Dev     string
}

// ModesEnum contains the available status bar mode values.
var ModesEnum = Modes{
	Normal:  "NORMAL",
	Loading: "LOADING",
	Error:   "ERROR",
	Dev:     "DEVELOP",
}

// Model holds the state for the status bar.
type Model struct {
	Status   string
	Content  string
	Width    int
	Spinner  spinner.Model
	Help     help.Model
	Keybinds help.KeyMap
	ctx      *context.AppContext
}

// New creates a new status bar model.
func New(ctx *context.AppContext, keybinds help.KeyMap) Model {
	status := ModesEnum.Normal
	if ctx.DevMode {
		status = ModesEnum.Dev
	}
	return Model{
		Status:   status,
		Keybinds: keybinds,
		Spinner: spinner.New(
			spinner.WithSpinner(spinner.Dot),
			spinner.WithStyle(SpinnerStyle),
		),
		ctx:  ctx,
		Help: help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return m.Spinner.Tick
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
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

	modeColor := modeBackground(m.Status)
	statusKey := statusStyle.Background(modeColor).Render(m.Status)
	statusVal := statusText.Render(tui.Truncate(m.Content, width/4))
	encoding := encodingStyle.Render("UTF-8")
	projectName := projectStyle.Render(fmt.Sprintf("%s %s", icon.Gitlab, m.ctx.SelectedProject.Name))

	helpWidth := width - w(statusKey) - w(statusVal) - w(encoding) - w(projectName) - 2
	if helpWidth < 0 {
		helpWidth = 0
	}
	m.Help.SetWidth(helpWidth)
	help := helpText.
		Width(helpWidth + 2).
		Render(" " + m.Help.View(m.Keybinds) + " ")

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		help,
		encoding,
		projectName,
	)

	return StatusBarStyle.Render(bar)
}

func modeBackground(status string) color.Color {
	switch status {
	case ModesEnum.Loading:
		return lipgloss.Color(style.StatuslineModeLoading)
	case ModesEnum.Error:
		return lipgloss.Color(style.StatuslineModeError)
	case ModesEnum.Dev:
		return lipgloss.Color(style.StatuslineModeDev)
	default:
		return lipgloss.Color(style.StatuslineModeNormal)
	}
}

// GetFrameSize returns the total frame size of the status bar.
func GetFrameSize() (int, int) {
	x, y := StatusBarStyle.GetFrameSize()
	xNugget, yNugget := statusNugget.GetFrameSize()
	xStatus, yStatus := statusStyle.GetFrameSize()

	return x + xNugget + xStatus, y + yNugget + yStatus
}
