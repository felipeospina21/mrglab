package details

import (
	"slices"

	"charm.land/bubbles/v2/key"
	"github.com/felipeospina21/mrglab/internal/tui"
)

type DetailsKeyMap struct {
	Merge          key.Binding
	RespondComment key.Binding
	NextDiscussion key.Binding
	PrevDiscussion key.Binding
	OpenInBrowser  key.Binding
	Fullscreen     key.Binding
	tui.GlobalKeyMap
}

func (k DetailsKeyMap) ShortHelp() []key.Binding {
	return slices.Concat(
		[]key.Binding{k.OpenInBrowser, k.Merge, k.RespondComment, k.NextDiscussion, k.PrevDiscussion, k.Fullscreen},
		tui.CommonKeys,
	)
}

func (k DetailsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		tui.CommonKeys,
		{k.OpenInBrowser, k.Merge, k.RespondComment, k.NextDiscussion, k.PrevDiscussion, k.Fullscreen},
	}
}

type PipelineDetailsKeyMap struct {
	OpenInBrowser key.Binding
	NextJob       key.Binding
	PrevJob       key.Binding
	PlayJob       key.Binding
	CancelJob     key.Binding
	Fullscreen    key.Binding
	tui.GlobalKeyMap
}

func (k PipelineDetailsKeyMap) ShortHelp() []key.Binding {
	return slices.Concat(
		[]key.Binding{k.OpenInBrowser, k.NextJob, k.PrevJob, k.PlayJob, k.CancelJob, k.Fullscreen},
		tui.CommonKeys,
	)
}

func (k PipelineDetailsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		tui.CommonKeys,
		{k.OpenInBrowser, k.NextJob, k.PrevJob, k.PlayJob, k.CancelJob, k.Fullscreen},
	}
}

var PipelineKeybinds = PipelineDetailsKeyMap{
	OpenInBrowser: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "open in browser"),
	),
	NextJob: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "next job"),
	),
	PrevJob: key.NewBinding(
		key.WithKeys("N"),
		key.WithHelp("N", "prev job"),
	),
	PlayJob: key.NewBinding(
		key.WithKeys("P"),
		key.WithHelp("P", "run job"),
	),
	CancelJob: key.NewBinding(
		key.WithKeys("X"),
		key.WithHelp("X", "cancel job"),
	),
	Fullscreen: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "fullscreen"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}

var Keybinds = DetailsKeyMap{
	Merge: key.NewBinding(
		key.WithKeys("M"),
		key.WithHelp("M", "merge mr"),
	),
	RespondComment: key.NewBinding(
		key.WithKeys("C"),
		key.WithHelp("C", "respond comment"),
	),
	NextDiscussion: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "next discussion"),
	),
	PrevDiscussion: key.NewBinding(
		key.WithKeys("N"),
		key.WithHelp("N", "prev discussion"),
	),
	OpenInBrowser: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "open in browser"),
	),
	Fullscreen: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "fullscreen"),
	),
	GlobalKeyMap: tui.GlobalKeys(false),
}
