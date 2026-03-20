package modal

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type KeyMap struct {
	Close  key.Binding
	Submit key.Binding
	tui.GlobalKeyMap
}

// ShortHelp implements the KeyMap interface.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Close, k.Submit, k.Quit}
}

// FullHelp implements the KeyMap interface.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Close, k.Submit, k.Quit},
	}
}

var Keybinds = KeyMap{
	Close: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "close modal"),
	),
	Submit: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "submit"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}
