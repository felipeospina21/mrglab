package app

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/pipelines"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
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
	ActiveTab   int
	TabNames    []string
	ProjectName string
}

func (p MergeRequestsPanel) Init() tea.Cmd { return p.Model.Init() }

func (p MergeRequestsPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if wsm, ok := msg.(tea.WindowSizeMsg); ok {
		wsm.Height -= 3 // header + tab bar + margin
		msg = wsm
	}
	m, cmd := p.Model.Update(msg)
	p.Model = &m
	return p, cmd
}

func (p MergeRequestsPanel) View() tea.View {
	header := renderHeader(p.ProjectName)
	tabBar := renderTabBar(p.ActiveTab, p.TabNames)
	return tea.NewView(lipgloss.JoinVertical(0, header, tabBar, p.Model.View().Content))
}

// PipelinesPanel wraps pipelines.Model to implement tea.Model.
type PipelinesPanel struct {
	*pipelines.Model
	ActiveTab   int
	TabNames    []string
	ProjectName string
}

func (p PipelinesPanel) Init() tea.Cmd { return p.Model.Init() }

func (p PipelinesPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if wsm, ok := msg.(tea.WindowSizeMsg); ok {
		wsm.Height -= 3 // header + tab bar + margin
		msg = wsm
	}
	m, cmd := p.Model.Update(msg)
	p.Model = &m
	return p, cmd
}

func (p PipelinesPanel) View() tea.View {
	header := renderHeader(p.ProjectName)
	tabBar := renderTabBar(p.ActiveTab, p.TabNames)
	return tea.NewView(lipgloss.JoinVertical(0, header, tabBar, p.Model.View().Content))
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

// renderHeader renders the project name header above the tab bar.
func renderHeader(projectName string) string {
	return table.TitleStyle.Render(projectName)
}

// renderTabBar renders the tab bar above the main panel content.
func renderTabBar(activeTab int, tabNames []string) string {
	activeStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Primary).
		Underline(true).
		Padding(0, 1).
		MarginTop(1)

	inactiveStyle := lipgloss.NewStyle().
		Foreground(theme.TextDimmed).
		Padding(0, 1).
		MarginTop(1)

	var tabs []string
	for i, name := range tabNames {
		if i == activeTab {
			tabs = append(tabs, activeStyle.Render(name))
		} else {
			tabs = append(tabs, inactiveStyle.Render(name))
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Bottom, tabs...)
}
