// Package details implements the merge request details side panel component.
package details

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/loader"
)

// LeftMargin is the horizontal margin used in the details panel.
const (
	LeftMargin = 2
)

// MergeRequestDetails holds the fetched details for a single merge request.
type MergeRequestDetails struct {
	Pipelines   []gitlab.CiStageNode
	Discussions []gitlab.DiscussionNode
	Approvals   []gitlab.ApprovalRule
	Branches    [2]string
}

// DetailsContent holds the rendered content sections for the details viewport.
type DetailsContent struct {
	Title       string
	Body        string
	Discussions string
	Pipelines   string
}

// Model holds the state for the details side panel.
type Model struct {
	Viewport       viewport.Model
	Ready          IsDetailsResponseReady
	Content        DetailsContent
	Discussions    []gitlab.DiscussionNode
	MRDetails      MergeRequestDetails
	MRId           string
	MRDescription  string
	SpinnerView    string
	DiscussionIdx  int
	Err            error
	ctx            *context.AppContext
}

// IsDetailsResponseReady indicates whether the details data has been loaded.
type (
	responseMsg            string
	contentRenderedMsg     string
	IsDetailsResponseReady bool
	errMsg                 struct{ err error }
)

// New creates a new details panel model.
func New(ctx *context.AppContext) Model {
	return Model{
		Viewport: viewport.New(viewport.WithWidth(10), viewport.WithHeight(10)),
		ctx:      ctx,
	}
}

// SetFocus sets the focused panel to the details (right) panel.
func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.RightPanel
}

func (e errMsg) Error() string { return e.err.Error() }

func (m Model) View() string {
	if !m.Ready {
		return PanelStyle.Render(fmt.Sprintf("%s\n%s",
			m.HeaderView(),
			loader.View(m.SpinnerView),
		))
	}
	return PanelStyle.Render(fmt.Sprintf("%s\n%s\n%s",
		m.HeaderView(),
		m.Viewport.View(),
		m.FooterView(),
	),
	)
}

func (m *Model) HeaderView() string {
	title := MdTitle.Render(m.Content.Title)
	line := strings.Repeat("─", tui.Max(0, m.Viewport.Width()-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m *Model) FooterView() string {
	info := MdInfo.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", tui.Max(0, m.Viewport.Width()-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

// SetViewportViewSize initializes or resizes the viewport to fit the given window.
func (m *Model) SetViewportViewSize(msg tea.WindowSizeMsg) tea.Cmd {
	w := msg.Width
	headerHeight := lipgloss.Height(m.HeaderView())
	footerHeight := lipgloss.Height(m.FooterView())
	verticalMarginHeight := headerHeight + footerHeight

	if !m.Ready {
		m.Viewport = viewport.New(viewport.WithWidth(w), viewport.WithHeight(msg.Height-verticalMarginHeight))
		m.Ready = true
	} else {
		m.Viewport.SetWidth(w)
		m.Viewport.SetHeight(msg.Height - verticalMarginHeight)
	}

	return nil
}
