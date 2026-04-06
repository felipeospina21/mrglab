package modal

import (
	"charm.land/bubbles/v2/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type KeyMap struct {
	Close  key.Binding
	Submit key.Binding
	Copy   key.Binding
	tui.GlobalKeyMap
}

// ShortHelp implements the KeyMap interface.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Close, k.Copy, k.Quit}
}

// FullHelp implements the KeyMap interface.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Close, k.Copy, k.Quit},
	}
}

var Keybinds = KeyMap{
	Close: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "close"),
	),
	Submit: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "submit"),
	),
	Copy: key.NewBinding(
		key.WithKeys("ctrl+y"),
		key.WithHelp("ctrl+y", "copy"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}

// MutationKeyMap is used for modals that perform mutations (respond to comment).
type MutationKeyMap struct {
	KeyMap
}

func (k MutationKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Close, k.Submit, k.Copy}
}

func (k MutationKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Close, k.Submit, k.Copy, k.Quit},
	}
}

// MutationKeybinds are the keybindings shown in modals with submit actions.
var MutationKeybinds = MutationKeyMap{KeyMap: Keybinds}

// CreateMRKeyMap extends the modal keybindings with create-MR-specific keys.
type CreateMRKeyMap struct {
	KeyMap
	Draft key.Binding
}

func (k CreateMRKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Close, k.Submit, k.Draft, k.Copy}
}

func (k CreateMRKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Close, k.Submit, k.Draft, k.Copy, k.Quit},
	}
}

// CreateMRKeybinds are the keybindings shown in the create MR modal footer.
var CreateMRKeybinds = CreateMRKeyMap{
	KeyMap: Keybinds,
	Draft: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("ctrl+d", "draft"),
	),
}
