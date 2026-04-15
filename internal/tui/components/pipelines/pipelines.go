// Package pipelines implements the pipelines panel component.
package pipelines

import "github.com/felipeospina21/mrglab/internal/context"

// Model holds the state for the pipelines panel.
type Model struct {
	ctx    *context.AppContext
	width  int
	height int
}

// New creates a new pipelines model.
func New(ctx *context.AppContext) Model {
	return Model{ctx: ctx}
}

// SetFocus sets the focused panel to the main panel.
func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.MainPanel
}
