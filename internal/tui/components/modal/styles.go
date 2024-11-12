package modal

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var headerStyle = table.TitleStyle.
	Foreground(lipgloss.Color(style.White)).
	Background(lipgloss.Color(style.Red[600])).
	Padding(0, 1)

var bodyStyle = func(h int) lipgloss.Style {
	return lipgloss.NewStyle().Padding(1).Height(h)
}
