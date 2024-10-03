package mergerequests

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type MergeReqsKeyMap struct {
	Discussions   key.Binding
	Pipelines     key.Binding
	OpenInBrowser key.Binding
	Refetch       key.Binding
	Details       key.Binding
	Merge         key.Binding
	tui.GlobalKeyMap
}

func (k MergeReqsKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Discussions, k.Pipelines, k.Details, k.OpenInBrowser, k.Merge, k.Help, k.Quit}
}

func (k MergeReqsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		tui.CommonKeys(), // first column
		{k.Discussions, k.Pipelines, k.Details, k.OpenInBrowser, k.Merge, k.Refetch}, // second column
	}
}

var Keybinds = MergeReqsKeyMap{
	Discussions: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "view comments"),
	),
	Pipelines: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "view pipelines"),
	),
	OpenInBrowser: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "open in browser"),
	),
	Refetch: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refetch"),
	),
	Details: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "view details"),
	),
	Merge: key.NewBinding(
		key.WithKeys("M"),
		key.WithHelp("M", "merge MR"),
	),
	GlobalKeyMap: tui.GlobalKeys,
}
