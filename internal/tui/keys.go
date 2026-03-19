package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type GlobalKeyMap struct {
	Help            key.Binding
	Quit            key.Binding
	ThrowError      key.Binding
	MockFetch       key.Binding
	ToggleLeftPanel key.Binding
	OpenModal       key.Binding
}

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

func KeyMatcher(msg tea.KeyMsg) func(key.Binding) bool {
	return func(k key.Binding) bool {
		return key.Matches(msg, k)
	}
}
