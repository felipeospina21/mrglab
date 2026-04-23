// Package statusline implements the bottom status bar component.
package statusline

import (
	"fmt"
	"image/color"

	"charm.land/bubbles/v2/help"
	"charm.land/lipgloss/v2"
	tssl "github.com/felipeospina21/tuishell/statusline"
	"github.com/felipeospina21/tuishell/style"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
)

// pkgTheme is initialized with sensible defaults so that code (including tests)
// that runs before SetTheme is called still gets non-nil color values.
var pkgTheme = style.Theme{
	StatusNormal:  lipgloss.Color("#6914ff"),
	StatusLoading: lipgloss.Color("#1A7A94"),
	StatusError:   lipgloss.Color("#CE3060"),
	StatusDev:     lipgloss.Color("#4E8212"),
}

// SetTheme sets the theme used by the statusline package and refreshes derived styles.
func SetTheme(t style.Theme) {
	pkgTheme = t
	StatusBarStyle = tssl.StatusBarStyle()
	SpinnerStyle = tssl.SpinnerStyle(t)
	refreshStatuslineStyles()
}

// Re-export from tuishell.
var (
	ModesEnum      = tssl.ModesEnum
	StatusBarStyle = tssl.StatusBarStyle()
	SpinnerStyle   = tssl.SpinnerStyle(pkgTheme)
)

// Model wraps the tuishell statusline with mrglab's context.
type Model struct {
	tssl.Model
	ctx *context.AppContext
}

// New creates a new status bar model.
func New(ctx *context.AppContext, keybinds help.KeyMap) Model {
	m := tssl.New(pkgTheme, ctx.DevMode, keybinds)
	return Model{Model: m, ctx: ctx}
}

// View renders the status bar, injecting the project label from mrglab context.
func (m Model) View() string {
	m.Model.ProjectLabel = fmt.Sprintf("%s %s", icon.Gitlab, m.ctx.SelectedProject.Name)
	return m.Model.View()
}

// GetFrameSize returns the total frame size of the status bar.
func GetFrameSize() (int, int) {
	return StatusBarStyle.GetFrameSize()
}

// modeBackground returns the background color for a given status mode.
// Kept for backward compatibility with tests.
func modeBackground(status string) color.Color {
	switch status {
	case ModesEnum.Loading:
		return pkgTheme.StatusLoading
	case ModesEnum.Error:
		return pkgTheme.StatusError
	case ModesEnum.Dev:
		return pkgTheme.StatusDev
	default:
		return pkgTheme.StatusNormal
	}
}
