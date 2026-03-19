package mergerequests

import (
	"slices"

	"github.com/charmbracelet/bubbles/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type MergeReqsKeyMap struct {
	OpenInBrowser key.Binding
	Details       key.Binding
	Merge         key.Binding
	tui.GlobalKeyMap
}

func (k MergeReqsKeyMap) ShortHelp() []key.Binding {
	return slices.Concat(
		[]key.Binding{k.Details, k.OpenInBrowser, k.Merge},
		tui.CommonKeys,
	)
}

func (k MergeReqsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		tui.CommonKeys,
		{k.Details, k.OpenInBrowser, k.Merge},
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
	GlobalKeyMap: tui.GlobalKeys(false),
}
