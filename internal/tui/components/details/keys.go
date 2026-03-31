package details

import (
	"slices"

	"charm.land/bubbles/v2/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type DetailsKeyMap struct {
	ClosePanel     key.Binding
	Merge          key.Binding
	RespondComment key.Binding
	NextDiscussion key.Binding
	PrevDiscussion key.Binding
	OpenInBrowser  key.Binding
	tui.GlobalKeyMap
}

func (k DetailsKeyMap) ShortHelp() []key.Binding {
	return slices.Concat(
		[]key.Binding{k.ClosePanel, k.OpenInBrowser, k.Merge, k.RespondComment, k.NextDiscussion, k.PrevDiscussion},
		tui.CommonKeys,
	)
}

func (k DetailsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		tui.CommonKeys,
		{k.ClosePanel, k.OpenInBrowser, k.Merge, k.RespondComment, k.NextDiscussion, k.PrevDiscussion},
	}
}

var Keybinds = DetailsKeyMap{
	ClosePanel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "close panel"),
	),
	Merge: key.NewBinding(
		key.WithKeys("M"),
		key.WithHelp("M", "merge mr"),
	),
	RespondComment: key.NewBinding(
		key.WithKeys("C"),
		key.WithHelp("C", "respond comment"),
	),
	NextDiscussion: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "next discussion"),
	),
	PrevDiscussion: key.NewBinding(
		key.WithKeys("N"),
		key.WithHelp("N", "prev discussion"),
	),
	OpenInBrowser: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "open in browser"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}
