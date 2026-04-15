package app

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
)

// ProjectsPanel wraps projects.Model to implement tea.Model.
type ProjectsPanel struct {
	*projects.Model
}

func (p ProjectsPanel) Init() tea.Cmd { return p.Model.Init() }

func (p ProjectsPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m, cmd := p.Model.Update(msg)
	p.Model = &m
	return p, cmd
}

func (p ProjectsPanel) View() tea.View { return p.Model.View() }

// SelectedLabel implements tuishell.SelectionProvider.
func (p ProjectsPanel) SelectedLabel() string {
	if item := p.Model.List.SelectedItem(); item != nil {
		return item.FilterValue()
	}
	return ""
}

// MergeRequestsPanel wraps mergerequests.Model to implement tea.Model.
type MergeRequestsPanel struct {
	*mergerequests.Model
	ActiveTab int
	TabNames  []string
	width     int
	height    int
}

func (p MergeRequestsPanel) Init() tea.Cmd { return p.Model.Init() }

func (p MergeRequestsPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if wsm, ok := msg.(tea.WindowSizeMsg); ok {
		p.width = wsm.Width
		p.height = wsm.Height
		// Reserve 1 line for tab bar
		wsm.Height -= 1
		msg = wsm
	}
	// On non-MR tabs, only handle tab key and window resize
	if p.ActiveTab != 0 {
		if keyMsg, ok := msg.(tea.KeyPressMsg); ok {
			if keyMsg.String() == "tab" {
				return p, func() tea.Msg { return mergerequests.CycleTabMsg{} }
			}
			return p, nil
		}
	}
	m, cmd := p.Model.Update(msg)
	p.Model = &m
	return p, cmd
}

func (p MergeRequestsPanel) View() tea.View {
	tabBar := p.renderTabBar()
	var content string
	if p.ActiveTab == 0 {
		content = p.Model.View().Content
	} else {
		content = lipgloss.NewStyle().
			Width(p.width).
			Height(p.height-2).
			Align(lipgloss.Center, lipgloss.Center).
			Foreground(theme.TextDimmed).
			Render(p.TabNames[p.ActiveTab])
	}
	return tea.NewView(lipgloss.JoinVertical(0, tabBar, content))
}

func (p MergeRequestsPanel) renderTabBar() string {
	activeStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Primary).
		Underline(true).
		Padding(0, 1).
		MarginBottom(1)

	inactiveStyle := lipgloss.NewStyle().
		Foreground(theme.TextDimmed).
		Padding(0, 1).
		MarginBottom(1)

	var tabs []string
	for i, name := range p.TabNames {
		if i == p.ActiveTab {
			tabs = append(tabs, activeStyle.Render(name))
		} else {
			tabs = append(tabs, inactiveStyle.Render(name))
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Bottom, tabs...)
}

// DetailsPanel wraps details.Model to implement tea.Model.
type DetailsPanel struct {
	*details.Model
}

func (p DetailsPanel) Init() tea.Cmd { return p.Model.Init() }

func (p DetailsPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m, cmd := p.Model.Update(msg)
	p.Model = &m
	return p, cmd
}

func (p DetailsPanel) View() tea.View { return p.Model.ViewContent() }
