package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// GlobalKeyMap defines the keybindings available in all panels.
type GlobalKeyMap struct {
	Help            key.Binding
	Quit            key.Binding
	ThrowError      key.Binding
	MockFetch       key.Binding
	ToggleLeftPanel key.Binding
	OpenModal       key.Binding
}

// CommonKeys are the keybindings shown in every panel's help.
var CommonKeys = []key.Binding{
	GlobalKeys(false).ToggleLeftPanel, GlobalKeys(false).OpenModal, GlobalKeys(false).Help, GlobalKeys(false).Quit,
}

func (k GlobalKeyMap) ShortHelp() []key.Binding {
	return CommonKeys
}

func (k GlobalKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		CommonKeys,
	}
}

// GlobalKeys returns the global keybindings, optionally including dev-mode keys.
func GlobalKeys(devMode bool) GlobalKeyMap {
	keymap := GlobalKeyMap{
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit"),
		),
		ToggleLeftPanel: key.NewBinding(
			key.WithKeys("ctrl+o"),
			key.WithHelp("ctrl+o", "toggle side panel"),
		),
		OpenModal: key.NewBinding(
			key.WithKeys("@"),
			key.WithHelp("@", "open full message modal"),
		),
	}

	if devMode {
		keymap.ThrowError = key.NewBinding(
			key.WithKeys("E"),
			key.WithHelp("E", "throw error"),
		)

		keymap.MockFetch = key.NewBinding(
			key.WithKeys("F"),
			key.WithHelp("F", "mock fetching"),
		)
	}

	return keymap
}

// KeyMatcher returns a predicate that checks if a tea.KeyMsg matches a key.Binding.
func KeyMatcher(msg tea.KeyMsg) func(key.Binding) bool {
	return func(k key.Binding) bool {
		return key.Matches(msg, k)
	}
}
