// Package app wires together all TUI components into the main Bubble Tea model.
package app

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/pipelines"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/style"
	"github.com/felipeospina21/tuishell"
	"github.com/felipeospina21/tuishell/popover"
	"github.com/felipeospina21/tuishell/shell"
	tsstyle "github.com/felipeospina21/tuishell/style"
)

// Model wraps shell.Model with mrglab-specific domain logic.
type Model struct {
	Shell         shell.Model
	Projects      *projects.Model
	MergeRequests *mergerequests.Model
	Pipelines     *pipelines.Model
	Details       *details.Model
	Input         textarea.Model
	ctx           *context.AppContext
	pendingNote   struct {
		DiscussionId string
		NoteableId   string
	}
	pendingCreateMR    bool
	pendingConfirm     bool
	statusFilter       popover.ListModel
	confirmPopover     popover.ConfirmModel
	pendingAction      tea.Msg
	formReady          bool
	createForm      createMRForm
	ActiveTab       int
	TabNames        []string
}

var theme = tsstyle.Theme{
	Primary:         lipgloss.Color(style.Violet[300]),
	PrimaryBright:   lipgloss.Color(style.Violet[400]),
	PrimaryFg:       lipgloss.Color(style.Violet[50]),
	PrimaryDim:      lipgloss.Color(style.Violet[800]),
	Info:            lipgloss.Color(style.Blue[400]),
	InfoBright:      lipgloss.Color(style.Blue[500]),
	Success:         lipgloss.Color(style.Green[300]),
	SuccessBright:   lipgloss.Color(style.Green[400]),
	Danger:          lipgloss.Color(style.Red[300]),
	DangerBright:    lipgloss.Color(style.Red[400]),
	Warning:         lipgloss.Color(style.Yellow[300]),
	WarningBright:   lipgloss.Color(style.Yellow[400]),
	Caution:         lipgloss.Color(style.Orange[400]),
	Text:            lipgloss.Color("#C4C4C4"),
	TextInverse:     lipgloss.Color("#111"),
	TextDimmed:      lipgloss.Color("#777777"),
	Muted:           lipgloss.Color("#999999"),
	Dim:             lipgloss.Color("#444444"),
	Border:          lipgloss.Color("#3f4145"),
	ModalBorder:     lipgloss.Color("#666666"),
	SurfaceDim:      lipgloss.Color("#1e1e24"),
	SelectionBorder: lipgloss.Color("#AD58B4"),
	StatusText:      lipgloss.Color("#FFFDF5"),
	StatusNormal:    lipgloss.Color(style.Violet[600]),
	StatusLoading:   lipgloss.Color("#1A7A94"),
	StatusError:     lipgloss.Color("#CE3060"),
	StatusDev:       lipgloss.Color("#4E8212"),
	StatusAccent1:   lipgloss.Color("#A550DF"),
	StatusAccent2:   lipgloss.Color("#6124DF"),
}

var leftPanelStyle = lipgloss.NewStyle().
	PaddingRight(4).
	Foreground(theme.Primary).
	Border(lipgloss.NormalBorder(), false, true, false, false).
	BorderForeground(theme.Border).
	Width(30)

var rightPanelStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), true, false, true, true).
	BorderForeground(theme.Border)



// InitMainModel creates and returns the initial application model.
func InitMainModel(ctx *context.AppContext, cfg *config.Config, client *gitlab.Client) Model {
	ctx.DevMode = cfg.DevMode

	ti := textarea.New()
	ti.Placeholder = "Write your reply..."
	ti.CharLimit = 0

	proj := projects.New(ctx, client, cfg.Filters.Projects)
	mrs := mergerequests.New(ctx, client)
	det := details.New(ctx)
	pip := pipelines.New(ctx, client)

	tabNames := []string{"Merge Requests", "Pipelines"}

	s := shell.New(shell.Config{
		Theme:           theme,
		LeftPanel:       ProjectsPanel{&proj},
		MainPanel:       MergeRequestsPanel{Model: &mrs, TabNames: tabNames, ProjectName: ""},
		RightPanel:      DetailsPanel{&det},
		AppIcon:         icon.Gitlab,
		Keybinds:        projects.Keybinds,
		DevMode:         cfg.DevMode,
		LeftPanelWidth:  30,
		LeftPanelStyle:  leftPanelStyle,
		RightPanelStyle: rightPanelStyle,
	})

	// Sync shell's initial context to mrglab's context
	ctx.AppContext = s.Ctx

	return Model{
		Shell:         s,
		Projects:      &proj,
		MergeRequests: &mrs,
		Pipelines:     &pip,
		Details:       &det,
		Input:         ti,
		ctx:           ctx,
		createForm:    newCreateMRForm(),
		statusFilter:   popover.NewList(tsstyle.Theme(theme)),
		confirmPopover: popover.NewConfirm(tsstyle.Theme(theme)),
		TabNames:      tabNames,
	}
}

func (m Model) Init() tea.Cmd {
	return m.Shell.Init()
}

func (m *Model) setHelpKeys(kb help.KeyMap) {
	m.Shell.Statusline.Keybinds = kb
}

// mainPanelKeybinds returns the correct keybinds for the currently active main panel tab.
func (m *Model) mainPanelKeybinds() help.KeyMap {
	if m.ActiveTab == 1 {
		return pipelines.Keybinds
	}
	return mergerequests.Keybinds
}

// detailsPanelKeybinds returns the correct keybinds for the details panel based on the active tab.
func (m *Model) detailsPanelKeybinds() help.KeyMap {
	if m.ActiveTab == 1 {
		return details.PipelineKeybinds
	}
	return details.Keybinds
}

// focusMainPanel sets focus and keybinds for the currently active main panel tab.
func (m *Model) focusMainPanel() {
	if m.ActiveTab == 1 {
		m.Pipelines.SetFocus()
	} else {
		m.MergeRequests.SetFocus()
	}
	m.setHelpKeys(m.mainPanelKeybinds())
}

// syncKeybinds updates the statusline keybinds to match the currently focused panel.
func (m *Model) syncKeybinds() {
	switch m.Shell.Ctx.FocusedPanel {
	case tuishell.LeftPanel:
		m.setHelpKeys(projects.Keybinds)
	case tuishell.RightPanel:
		m.setHelpKeys(m.detailsPanelKeybinds())
	default:
		m.setHelpKeys(m.mainPanelKeybinds())
	}
}

// SelectMR stores the currently selected merge request's IID, SHA, and status in the app context.
func (m *Model) SelectMR() {
	idColIdx := mergerequests.GetColIndex(mergerequests.ColNames.ID)
	m.ctx.SelectedMR.IID = m.MergeRequests.Table.SelectedRow()[idColIdx]

	shaColIdx := mergerequests.GetColIndex(mergerequests.ColNames.Sha)
	m.ctx.SelectedMR.Sha = m.MergeRequests.Table.SelectedRow()[shaColIdx]

	statusColIdx := mergerequests.GetColIndex(mergerequests.ColNames.Status)
	m.ctx.SelectedMR.Status = m.MergeRequests.Table.SelectedRow()[statusColIdx]
}
