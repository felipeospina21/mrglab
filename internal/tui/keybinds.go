package tui

import "github.com/charmbracelet/bubbles/key"

type GlobalKeyMap struct {
	Help       key.Binding
	Quit       key.Binding
	ThrowError key.Binding
	// NextTab         key.Binding
	// PrevTab         key.Binding
	// NextPage        key.Binding
	// PrevPage        key.Binding
	// NavigateBack    key.Binding
	ToggleLeftPanel key.Binding
}

func CommonKeys() []key.Binding {
	var k GlobalKeyMap
	return []key.Binding{
		k.Help, k.ToggleLeftPanel, k.Quit,
		// k.NextTab, k.PrevTab, k.NextPage, k.PrevPage
	}
}

func (k GlobalKeyMap) ShortHelp() []key.Binding {
	// return CommonKeys()
	return []key.Binding{
		k.Help, k.ToggleLeftPanel, k.Quit,
		// k.NextTab, k.PrevTab, k.NextPage, k.PrevPage
	}
}

func (k GlobalKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		CommonKeys(),
	}
}

var GlobalKeys = GlobalKeyMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	// NextTab: key.NewBinding(
	// 	key.WithKeys("tab"),
	// 	key.WithHelp("tab", "next tab"),
	// ),
	// PrevTab: key.NewBinding(
	// 	key.WithKeys("shift+tab"),
	// 	key.WithHelp("shift+tab", "prev tab"),
	// ),
	// NextPage: key.NewBinding(
	// 	key.WithKeys("right"),
	// 	key.WithHelp("->", "next page"),
	// ),
	// PrevPage: key.NewBinding(
	// 	key.WithKeys("left"),
	// 	key.WithHelp("<-", "prev page"),
	// ),
	ToggleLeftPanel: key.NewBinding(
		key.WithKeys("ctrl+o"),
		key.WithHelp("ctrl+o", "toggle side panel"),
	),

	// TODO: make this available only when program is run whith certain cmd
	ThrowError: key.NewBinding(
		key.WithKeys("E"),
		key.WithHelp("E", "throw error"),
	),
}
