// Package loader provides a reusable loading indicator component.
package loader

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

var (
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(style.Violet[300]))
	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(style.Violet[400]))
)

// View renders a spinner frame with a "Loading..." label.
func View(spinnerView string) string {
	return fmt.Sprintf("\n %s %s\n\n", spinnerView, textStyle.Render("Loading..."))
}
