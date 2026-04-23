package tui

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/felipeospina21/tuishell"
)

// GlobalKeyMap defines the keybindings available in all panels.
type GlobalKeyMap = tuishell.GlobalKeyMap

// CommonKeys are the keybindings shown in every panel's help.
var CommonKeys = tuishell.CommonKeys

// GlobalKeys returns the global keybindings, optionally including dev-mode keys.
func GlobalKeys(demoMode bool) GlobalKeyMap { return tuishell.GlobalKeys(demoMode) }

// KeyMatcher returns a predicate that checks if a tea.KeyPressMsg matches a key.Binding.
func KeyMatcher(msg tea.KeyPressMsg) func(key.Binding) bool { return tuishell.KeyMatcher(msg) }
