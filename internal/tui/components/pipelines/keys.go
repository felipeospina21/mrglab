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
	Refetch        key.Binding
	RetryPipeline  key.Binding
	CancelPipeline key.Binding
	FilterStatus   key.Binding
	tui.GlobalKeyMap
}

func (k PipelinesKeyMap) ShortHelp() []key.Binding {
	return slices.Concat([]key.Binding{k.Details, k.CycleTab, k.OpenInBrowser, k.Refetch, k.RetryPipeline, k.CancelPipeline, k.FilterStatus}, tui.CommonKeys)
}

func (k PipelinesKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{tui.CommonKeys, {k.Details, k.CycleTab, k.OpenInBrowser, k.Refetch, k.RetryPipeline, k.CancelPipeline, k.FilterStatus}}
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
	Refetch: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refetch pipelines"),
	),
	RetryPipeline: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "retry failed jobs"),
	),
	CancelPipeline: key.NewBinding(
		key.WithKeys("C"),
		key.WithHelp("C", "cancel pipeline"),
	),
	FilterStatus: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "filter by status"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}
