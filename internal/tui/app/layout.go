package app

import (
	"github.com/felipeospina21/tuishell"
)

// PanelSize holds the computed width and height for a UI region.
type PanelSize = tuishell.PanelSize

// Layout holds all computed dimensions for the current window size and panel state.
type Layout = tuishell.Layout

const (
	mainPanelHeaderLines = 1 // the "Project - Merge Requests" header

	// table.DocStyle uses BorderStyle(RoundedBorder()) which renders borders
	// but lipgloss GetFrameSize() reports 0 for it. Account for the actual rendered border.
	tableBorderX = 2
	tableBorderY = 2

	// table.View() renders 1 line taller than SetHeight due to headersView()
	// border bottom not being counted by lipgloss.Height() inside SetHeight.
	tableViewOverhead = 1
)
