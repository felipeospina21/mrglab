// Package statusline implements the bottom status bar component.
package statusline

import (
	"fmt"
	"image/color"

	"charm.land/bubbles/v2/help"
	tssl "github.com/felipeospina21/tuishell/statusline"
	"github.com/felipeospina21/tuishell/style"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
)

// Re-export from tuishell.
var (
	ModesEnum      = tssl.ModesEnum
	StatusBarStyle = tssl.StatusBarStyle()
	SpinnerStyle   = tssl.SpinnerStyle(style.DefaultTheme())
)

// Model wraps the tuishell statusline with mrglab's context.
type Model struct {
	tssl.Model
	ctx *context.AppContext
}

// New creates a new status bar model.
func New(ctx *context.AppContext, keybinds help.KeyMap) Model {
	theme := style.DefaultTheme()
	m := tssl.New(theme, ctx.DevMode, keybinds)
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
	theme := style.DefaultTheme()
	switch status {
	case ModesEnum.Loading:
		return theme.StatusLoading
	case ModesEnum.Error:
		return theme.StatusError
	case ModesEnum.Dev:
		return theme.StatusDev
	default:
		return theme.StatusNormal
	}
}
