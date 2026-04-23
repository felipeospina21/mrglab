package statusline

import (
	"charm.land/lipgloss/v2"
)

var (
	statusNugget  = lipgloss.NewStyle()
	statusStyle   = lipgloss.NewStyle()
	encodingStyle = lipgloss.NewStyle()
	statusText    = lipgloss.NewStyle()
	helpText      = lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)
	projectStyle  = lipgloss.NewStyle()
)

func refreshStatuslineStyles() {
	t := pkgTheme
	statusNugget = lipgloss.NewStyle().
		Foreground(t.StatusText).
		Padding(0, 1)
	statusStyle = lipgloss.NewStyle().
		Foreground(t.StatusText).
		Padding(0, 1).
		MarginRight(1)
	encodingStyle = statusNugget.
		Background(t.StatusAccent1).
		Align(lipgloss.Right)
	statusText = lipgloss.NewStyle().Foreground(t.StatusText)
	projectStyle = statusNugget.Background(t.StatusAccent2)
}
