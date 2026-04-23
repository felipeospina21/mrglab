package statusline

import (
	"charm.land/lipgloss/v2"
)

var (
	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.
			Background(lipgloss.Color("#A550DF")).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5"))

	helpText = lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center)

	projectStyle = statusNugget.Background(lipgloss.Color("#6124DF"))
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
