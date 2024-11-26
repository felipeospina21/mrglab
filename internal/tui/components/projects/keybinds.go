package projects

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type ProjectsKeyMap struct {
	MRList key.Binding
	tui.GlobalKeyMap
}

func (k ProjectsKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.MRList}
}

func (k ProjectsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.MRList}, // first column
		// {k.ShowFullHelp, k.Quit, k.Filter, k.ReloadConfig}, // second column
	}
}

var Keybinds = ProjectsKeyMap{
	MRList: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "view merge requests"),
	),
	GlobalKeyMap: tui.GlobalKeys(),
}
