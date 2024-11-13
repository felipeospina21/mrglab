package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type GlobalKeyMap struct {
	Help            key.Binding
	Quit            key.Binding
	ThrowError      key.Binding
	ToggleLeftPanel key.Binding
	OpenModal       key.Binding
}

var CommonKeys = []key.Binding{
	GlobalKeys.ToggleLeftPanel, GlobalKeys.OpenModal, GlobalKeys.Help, GlobalKeys.Quit,
}

func (k GlobalKeyMap) ShortHelp() []key.Binding {
	return CommonKeys
}

func (k GlobalKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		CommonKeys,
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
	ToggleLeftPanel: key.NewBinding(
		key.WithKeys("ctrl+o"),
		key.WithHelp("ctrl+o", "toggle side panel"),
	),
	OpenModal: key.NewBinding(
		key.WithKeys("@"),
		key.WithHelp("@", "open full message modal"),
	),

	// TODO: make this available only when program is run whith certain cmd
	ThrowError: key.NewBinding(
		key.WithKeys("E"),
		key.WithHelp("E", "throw error"),
	),
}

func KeyMatcher(msg tea.KeyMsg) func(key.Binding) bool {
	return func(k key.Binding) bool {
		return key.Matches(msg, k)
	}
}
