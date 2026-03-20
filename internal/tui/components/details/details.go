package details

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui"
)

const (
	useHighPerformanceRenderer = false
	LeftMargin                 = 2
)

type MergeRequestDetails struct {
	Pipelines   []gitlab.CiStageNode
	Discussions []gitlab.DiscussionNode
	Approvals   []gitlab.ApprovalRule
	Branches    [2]string
}

type DetailsContent struct {
	Title       string
	Body        string
	Discussions string
	Pipelines   string
}

type Model struct {
	Viewport       viewport.Model
	Ready          IsDetailsResponseReady
	Content        DetailsContent
	Discussions    []gitlab.DiscussionNode
	MRDetails      MergeRequestDetails
	MRId           string
	MRDescription  string
	DiscussionIdx  int
	Err            error
	ctx            *context.AppContext
}

type (
	responseMsg            string
	contentRenderedMsg     string
	IsDetailsResponseReady bool
	errMsg                 struct{ err error }
)

func New(ctx *context.AppContext) Model {
	return Model{
		Viewport: viewport.New(10, 10),
		ctx:      ctx,
	}
}

func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.RightPanel
}

func (e errMsg) Error() string { return e.err.Error() }

func (m Model) View() string {
	return PanelStyle.Render(fmt.Sprintf("%s\n%s\n%s",
		m.HeaderView(),
		m.Viewport.View(),
		m.FooterView(),
	),
	)
}

func (m *Model) HeaderView() string {
	title := MdTitle.Render(m.Content.Title)
	line := strings.Repeat("─", tui.Max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m *Model) FooterView() string {
	info := MdInfo.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", tui.Max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m *Model) SetViewportViewSize(msg tea.WindowSizeMsg) tea.Cmd {
	w := msg.Width
	headerHeight := lipgloss.Height(m.HeaderView())
	footerHeight := lipgloss.Height(m.FooterView())
	verticalMarginHeight := headerHeight + footerHeight

	if !m.Ready {
		m.Viewport = viewport.New(w, msg.Height-verticalMarginHeight)
		m.Viewport.HighPerformanceRendering = useHighPerformanceRenderer
		m.Ready = true
		m.Viewport.YPosition = headerHeight
	} else {
		m.Viewport.Width = w
		m.Viewport.Height = msg.Height - verticalMarginHeight
	}
	if useHighPerformanceRenderer {
		return viewport.Sync(m.Viewport)
	}

	return nil
}
