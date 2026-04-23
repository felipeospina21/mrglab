// Package modal re-exports the tuishell modal component adapted for mrglab's context.
package modal

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/help"
	tsmodal "github.com/felipeospina21/tuishell/modal"
	"github.com/felipeospina21/tuishell/style"
	"github.com/felipeospina21/mrglab/internal/context"
)

var pkgTheme style.Theme

// SetTheme sets the theme used by the modal package.
func SetTheme(t style.Theme) { pkgTheme = t }

// Re-export message types.
type (
	CloseModalMsg     = tsmodal.CloseModalMsg
	SubmitModalMsg    = tsmodal.SubmitModalMsg
	CopyModalMsg      = tsmodal.CopyModalMsg
	ResetHighlightMsg = tsmodal.ResetHighlightMsg
)

// Keybinds re-exports the modal keybindings.
var Keybinds = tsmodal.Keybinds

// Model wraps the tuishell modal with mrglab's context.
type Model struct {
	tsmodal.Model
}

// New creates a new modal model using mrglab's AppContext.
func New(ctx *context.AppContext) Model {
	return Model{Model: tsmodal.New(&ctx.AppContext, pkgTheme)}
}

// Update wraps the inner Update to return the mrglab modal.Model type.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	inner, cmd := m.Model.Update(msg)
	m.Model = inner
	return m, cmd
}

// RenderHelp delegates to the inner model.
func (m Model) RenderHelp(km help.KeyMap) string {
	return m.Model.RenderHelp(km)
}

// ContentWidth returns the usable content width inside the modal.
func ContentWidth(windowW int) int {
	return tsmodal.ContentWidth(pkgTheme, windowW)
}

// ContentHeight returns the usable content height inside the modal.
func ContentHeight(windowH int) int {
	return tsmodal.ContentHeight(pkgTheme, windowH)
}
