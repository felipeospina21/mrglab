package projects

import (
	"slices"

	"charm.land/bubbles/v2/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type ProjectsKeyMap struct {
	MRList     key.Binding
	CursorUp   key.Binding
	CursorDown key.Binding
	Filter     key.Binding
	tui.GlobalKeyMap
}

func (k ProjectsKeyMap) ShortHelp() []key.Binding {
	return slices.Concat(
		[]key.Binding{k.MRList, k.CursorUp, k.CursorDown, k.Filter},
		tui.CommonKeys,
	)
}

func (k ProjectsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		tui.CommonKeys,
		{k.MRList, k.CursorUp, k.CursorDown, k.Filter},
	}
}

var Keybinds = ProjectsKeyMap{
	MRList: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "view merge requests"),
	),
	CursorUp: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	CursorDown: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Filter: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}
