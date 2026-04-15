package pipelines

import (
	"slices"

	"charm.land/bubbles/v2/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

// PipelinesKeyMap shows only the tab-cycling keybind when the Pipelines tab is active.
type PipelinesKeyMap struct {
	CycleTab       key.Binding
	OpenInBrowser  key.Binding
	Details        key.Binding
	RetryPipeline  key.Binding
	CancelPipeline key.Binding
	tui.GlobalKeyMap
}

func (k PipelinesKeyMap) ShortHelp() []key.Binding {
	return slices.Concat([]key.Binding{k.Details, k.CycleTab, k.OpenInBrowser, k.RetryPipeline, k.CancelPipeline}, tui.CommonKeys)
}

func (k PipelinesKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{tui.CommonKeys, {k.Details, k.CycleTab, k.OpenInBrowser, k.RetryPipeline, k.CancelPipeline}}
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
	Details: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "view details"),
	),
	RetryPipeline: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "retry failed jobs"),
	),
	CancelPipeline: key.NewBinding(
		key.WithKeys("C"),
		key.WithHelp("C", "cancel pipeline"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}
