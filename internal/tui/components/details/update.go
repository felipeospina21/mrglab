package details

import (
	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type (
	ClosePanelMsg    struct{}
	MergeMRMsg       struct{}
	OpenInBrowserMsg struct{}
	FullscreenMsg    struct{}
	RespondCommentMsg struct {
		DiscussionId string
		NoteableId   string
	}
)

// Init returns nil (no initialization needed).
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles key events for the details panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		match := tui.KeyMatcher(msg)
		switch {
		case match(Keybinds.ClosePanel):
			return m, func() tea.Msg { return ClosePanelMsg{} }
		case match(Keybinds.Merge):
			return m, func() tea.Msg { return MergeMRMsg{} }
		case match(Keybinds.OpenInBrowser):
			return m, func() tea.Msg { return OpenInBrowserMsg{} }
		case match(Keybinds.RespondComment):
			d := m.selectedDiscussion()
			if d == nil {
				break
			}
			return m, func() tea.Msg {
				return RespondCommentMsg{
					DiscussionId: d.Id,
					NoteableId:   m.MRId,
				}
			}
		case match(Keybinds.NextDiscussion):
			m.nextDiscussion()
			m.refreshContent()
		case match(Keybinds.PrevDiscussion):
			m.prevDiscussion()
			m.refreshContent()
		case match(Keybinds.Fullscreen):
			return m, func() tea.Msg { return FullscreenMsg{} }
		}
	case tea.WindowSizeMsg:
		frameY := PanelStyle.GetVerticalFrameSize()
		m.SetViewportViewSize(tea.WindowSizeMsg{Width: msg.Width, Height: msg.Height - frameY})
	}
	m.Viewport, cmd = m.Viewport.Update(msg)
	return m, cmd
}

// ViewContent returns the panel content as a tea.View.
func (m Model) ViewContent() tea.View {
	return tea.NewView(m.View())
}

func (m *Model) refreshContent() {
	vc := m.getViewportContent(m.MRDescription, m.MRDetails)
	m.Viewport.SetContent(vc.content)
	m.Viewport.SetYOffset(vc.selectedDiscLineOffset)
}

func (m *Model) resolvableDiscussions() []int {
	var indices []int
	for i, d := range m.Discussions {
		if d.Resolvable {
			indices = append(indices, i)
		}
	}
	return indices
}

func (m *Model) selectedDiscussion() *resolvableDiscussion {
	indices := m.resolvableDiscussions()
	if len(indices) == 0 {
		return nil
	}
	idx := m.DiscussionIdx % len(indices)
	d := m.Discussions[indices[idx]]
	return &resolvableDiscussion{DiscussionNode: d, displayIndex: idx}
}

type resolvableDiscussion struct {
	gitlab.DiscussionNode
	displayIndex int
}

func (m *Model) nextDiscussion() {
	indices := m.resolvableDiscussions()
	if len(indices) == 0 {
		return
	}
	m.DiscussionIdx = (m.DiscussionIdx + 1) % len(indices)
}

func (m *Model) prevDiscussion() {
	indices := m.resolvableDiscussions()
	if len(indices) == 0 {
		return
	}
	m.DiscussionIdx = (m.DiscussionIdx - 1 + len(indices)) % len(indices)
}
