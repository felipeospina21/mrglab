package details

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type DetailsKeyMap struct {
	ClosePanel key.Binding
	Merge      key.Binding
	tui.GlobalKeyMap
}

func (k DetailsKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.ClosePanel, k.Help, k.Quit}
}

func (k DetailsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ClosePanel, k.Merge}, // first column
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
		key.WithHelp("M", "merge MR"),
	),
	GlobalKeyMap: tui.GlobalKeys,
}
