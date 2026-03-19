package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

// PanelSize holds the computed width and height for a UI region.
type PanelSize struct {
	Width  int
	Height int
}

// Layout holds all computed dimensions for the current window size and panel state.
// Computed once on WindowSizeMsg or panel toggle, then consumed by View and applyLayout.
type Layout struct {
	Window     tea.WindowSizeMsg
	LeftPanel  PanelSize
	MainPanel  PanelSize
	RightPanel PanelSize
	Statusline PanelSize
	ContentH   int // usable content height (window minus main frame minus statusline)
}

const (
	leftPanelWidth       = 30 // matches projects.DocStyle.Width(30)
	mainPanelHeaderLines = 1  // the "Project - Merge Requests" header
	statuslineLines      = 1  // statusline is a single rendered line

	// table.DocStyle uses BorderStyle(RoundedBorder()) which renders borders
	// but lipgloss GetFrameSize() reports 0 for it. Account for the actual rendered border.
	tableBorderX = 2
	tableBorderY = 2

	// table.View() renders 1 line taller than SetHeight due to headersView()
	// border bottom not being counted by lipgloss.Height() inside SetHeight.
	tableViewOverhead = 1
)

func computeLayout(win tea.WindowSizeMsg, leftOpen, rightOpen bool) Layout {
	mainFrameX, mainFrameY := style.MainFrameStyle.GetFrameSize()

	innerW := win.Width - mainFrameX
	innerH := win.Height - mainFrameY

	// Statusline: full inner width, fixed height
	slFrameY := statusline.StatusBarStyle.GetVerticalFrameSize()
	slH := statuslineLines + slFrameY
	slFrameX := statusline.StatusBarStyle.GetHorizontalFrameSize()

	contentH := innerH - slH

	// Left panel
	leftW := 0
	if leftOpen {
		leftW = leftPanelWidth + projects.DocStyle.GetHorizontalFrameSize()
	}

	// Main panel gets remaining width, minus right panel if open
	mainW := innerW - leftW
	rightW := 0
	if rightOpen && !leftOpen {
		detailsFrameX := details.PanelStyle.GetHorizontalFrameSize()
		rightW = mainW/2 - detailsFrameX
		mainW = mainW - rightW - detailsFrameX
	}

	// Left panel: subtract DocStyle vertical frame (MarginBottom:2) so rendered output fits
	leftH := contentH - projects.DocStyle.GetVerticalFrameSize()

	return Layout{
		Window:     win,
		LeftPanel:  PanelSize{Width: leftW, Height: leftH},
		MainPanel:  PanelSize{Width: mainW, Height: contentH},
		RightPanel: PanelSize{Width: rightW, Height: contentH},
		Statusline: PanelSize{Width: innerW - slFrameX, Height: slH},
		ContentH:   contentH,
	}
}

// applyLayout pushes the computed dimensions to all components in-place.
func (m *Model) applyLayout() {
	l := m.layout

	// Left panel
	m.Projects.List.SetHeight(l.LeftPanel.Height)
	m.ctx.PanelHeight = l.ContentH

	// Statusline
	m.Statusline.Width = l.Statusline.Width

	// Main panel table — update in-place instead of reconstructing
	if len(m.MergeRequests.Table.Rows()) > 0 {
		tableFrameX := table.DocStyle.GetHorizontalFrameSize() + tableBorderX
		tableW := l.MainPanel.Width - tableFrameX
		m.MergeRequests.Table.SetWidth(tableW)
		m.MergeRequests.Table.SetHeight(l.ContentH - mainPanelHeaderLines - tableBorderY - tableViewOverhead)
		m.MergeRequests.Table.SetColumns(mergerequests.GetTableColums(tableW))
	}

	// Right panel (details viewport)
	if m.isRightOpen {
		detailsFrameY := details.PanelStyle.GetVerticalFrameSize()
		m.Details.SetViewportViewSize(
			tea.WindowSizeMsg{Width: l.RightPanel.Width, Height: l.ContentH - detailsFrameY - tableViewOverhead},
		)
	}
}
