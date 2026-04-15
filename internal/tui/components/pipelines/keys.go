package pipelines

import (
	"slices"

	"charm.land/bubbles/v2/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

// PipelinesKeyMap shows only the tab-cycling keybind when the Pipelines tab is active.
type PipelinesKeyMap struct {
	CycleTab      key.Binding
	OpenInBrowser key.Binding
	tui.GlobalKeyMap
}

func (k PipelinesKeyMap) ShortHelp() []key.Binding {
	return slices.Concat([]key.Binding{k.CycleTab, k.OpenInBrowser}, tui.CommonKeys)
}

func (k PipelinesKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{tui.CommonKeys, {k.CycleTab, k.OpenInBrowser}}
}

// Keybinds is the default keybinding set for the pipelines panel.
var Keybinds = PipelinesKeyMap{
	CycleTab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next tab"),
	),
	OpenInBrowser: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "open in browser"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}
