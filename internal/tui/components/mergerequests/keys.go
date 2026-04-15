package mergerequests

import (
	"slices"

	"charm.land/bubbles/v2/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type MergeReqsKeyMap struct {
	OpenInBrowser key.Binding
	Details       key.Binding
	Merge         key.Binding
	CreateMR      key.Binding
	Refetch       key.Binding
	tui.GlobalKeyMap
}

func (k MergeReqsKeyMap) ShortHelp() []key.Binding {
	return slices.Concat(
		[]key.Binding{k.Details, k.OpenInBrowser, k.Merge, k.CreateMR, k.Refetch},
		tui.CommonKeys,
	)
}

func (k MergeReqsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		tui.CommonKeys,
		{k.Details, k.OpenInBrowser, k.Merge, k.CreateMR, k.Refetch},
	}
}

var Keybinds = MergeReqsKeyMap{
	OpenInBrowser: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "open in browser"),
	),
	Details: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "view details"),
	),
	Merge: key.NewBinding(
		key.WithKeys("M"),
		key.WithHelp("M", "merge MR"),
	),
	CreateMR: key.NewBinding(
		key.WithKeys("N"),
		key.WithHelp("N", "new MR"),
	),
	Refetch: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refetch MRs"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}
