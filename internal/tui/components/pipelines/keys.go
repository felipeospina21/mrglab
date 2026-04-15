package pipelines

import (
	"slices"

	"charm.land/bubbles/v2/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

// PipelinesKeyMap shows only the tab-cycling keybind when the Pipelines tab is active.
type PipelinesKeyMap struct {
	CycleTab key.Binding
	tui.GlobalKeyMap
}

func (k PipelinesKeyMap) ShortHelp() []key.Binding {
	return slices.Concat([]key.Binding{k.CycleTab}, tui.CommonKeys)
}

func (k PipelinesKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{tui.CommonKeys, {k.CycleTab}}
}

// Keybinds is the default keybinding set for the pipelines panel.
var Keybinds = PipelinesKeyMap{
	CycleTab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next tab"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}
