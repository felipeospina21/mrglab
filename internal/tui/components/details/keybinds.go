package details

import (
	"slices"

	"github.com/charmbracelet/bubbles/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type DetailsKeyMap struct {
	ClosePanel     key.Binding
	Merge          key.Binding
	RespondComment key.Binding
	OpenInBrowser  key.Binding
	tui.GlobalKeyMap
}

func (k DetailsKeyMap) ShortHelp() []key.Binding {
	return slices.Concat(
		[]key.Binding{k.ClosePanel, k.OpenInBrowser, k.Merge},
		tui.CommonKeys,
	)
}

func (k DetailsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ClosePanel, k.OpenInBrowser, k.Merge}, // first column
		// {k.ShowFullHelp, k.Quit, k.Filter, k.ReloadConfig}, // second column
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
	OpenInBrowser: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "open in browser"),
	),
	GlobalKeyMap: tui.GlobalKeys,
}
