package details

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/logger"
)

const useHighPerformanceRenderer = false

type DetailsContent struct {
	Title       string
	Body        string
	Discussions string
	Pipelines   string
}

type Model struct {
	Viewport viewport.Model
	Ready    IsDetailsResponseReady
	Content  DetailsContent
	Err      error
	ctx      *context.AppContext
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
	var content strings.Builder

	content.WriteString(m.renderBody())
	content.WriteString("\n\n")
	content.WriteString(m.Content.Pipelines)
	content.WriteString(m.renderDiscussions())

	return fmt.Sprintf("%s\n%s\n%s",
		m.HeaderView(),
		content.String(),
		m.FooterView(),
	)
}

func (m *Model) HeaderView() string {
	title := MdTitle.Render(m.Content.Title)
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m *Model) FooterView() string {
	info := MdInfo.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m Model) renderDiscussions() string {
	r, err := getMdRenderer(m)
	if err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)

		return ""
	}

	d, err := r.Render(m.Content.Discussions)
	if err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)

		return ""
	}

	return d
}

func (m Model) renderBody() string {
	r, err := getMdRenderer(m)
	if err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)

		return ""
	}

	b, err := r.Render(m.Content.Body)
	if err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)

		return ""
	}

	return b
}

func getMdRenderer(m Model) (*glamour.TermRenderer, error) {
	magicnumber := 8 // FIX: find where this comes from
	width := m.Viewport.Width - magicnumber
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(width),
		glamour.WithEmoji(),
		glamour.WithPreservedNewLines(),
	)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (m *Model) SetViewportViewSize(msg tea.WindowSizeMsg) tea.Cmd {
	magicnumber := 8 // FIX: find where this comes from

	w := msg.Width - magicnumber
	headerHeight := lipgloss.Height(m.HeaderView())
	footerHeight := lipgloss.Height(m.FooterView())
	verticalMarginHeight := headerHeight + footerHeight + magicnumber

	if !m.Ready {
		// Since this program is using the full size of the viewport we
		// need to wait until we've received the window dimensions before
		// we can initialize the viewport. The initial dimensions come in
		// quickly, though asynchronously, which is why we wait for them
		// here.
		m.Viewport = viewport.New(w, msg.Height-verticalMarginHeight)
		m.Viewport.YPosition = headerHeight
		m.Viewport.HighPerformanceRendering = useHighPerformanceRenderer

		// m.SetResponseContent(string(m.Content))
		m.Ready = true

		// This is only necessary for high performance rendering, which in
		// most cases you won't need.
		//
		// Render the viewport one line below the header.
		m.Viewport.YPosition = headerHeight + 1
	} else {
		m.Viewport.Width = w
		m.Viewport.Height = msg.Height - verticalMarginHeight
	}
	if useHighPerformanceRenderer {
		// Render (or re-render) the whole viewport. Necessary both to
		// initialize the viewport and when the window is resized.
		//
		// This is needed for high-performance rendering only.
		// cmds = append(cmds, viewport.Sync(m.viewport.mod))
		return viewport.Sync(m.Viewport)
	}

	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
