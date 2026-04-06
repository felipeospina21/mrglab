package app

import (
	tea "charm.land/bubbletea/v2"
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
}

func (p MergeRequestsPanel) Init() tea.Cmd { return p.Model.Init() }

func (p MergeRequestsPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m, cmd := p.Model.Update(msg)
	p.Model = &m
	return p, cmd
}

func (p MergeRequestsPanel) View() tea.View { return p.Model.View() }

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
