package modal

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type KeyMap struct {
	Close key.Binding
	Tab   key.Binding
	tui.GlobalKeyMap
}

// ShortHelp implements the KeyMap interface.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Close, k.Quit}
}

// FullHelp implements the KeyMap interface.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Close, k.Quit},
	}
}

var Keybinds = KeyMap{
	Close: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "close modal"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", ""),
	),
	GlobalKeyMap: tui.GlobalKeys(),
}
