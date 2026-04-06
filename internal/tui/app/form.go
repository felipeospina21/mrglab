package app

import (
	"fmt"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/components/modal"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

const (
	formFieldSource = iota
	formFieldTarget
	formFieldTitle
	formFieldDescription
	formFieldCount
)

type createMRForm struct {
	source      textinput.Model
	target      textinput.Model
	title       textinput.Model
	description textarea.Model
	draft       bool
	focused     int
	width       int
}

var (
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(style.Violet[400])).
			Bold(true).
			MarginTop(1)

	arrowStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(style.MediumGray)).
			Bold(true).
			Padding(0, 0, 0, 4)
)

func newCreateMRForm() createMRForm {
	source := textinput.New()
	source.Placeholder = "source branch"
	source.CharLimit = 200

	target := textinput.New()
	target.Placeholder = "target branch"
	target.CharLimit = 200

	title := textinput.New()
	title.Placeholder = "feat: add new feature"
	title.CharLimit = 200

	desc := textarea.New()
	desc.Placeholder = "Description (populated from template if available)"
	desc.CharLimit = 0
	desc.ShowLineNumbers = false

	return createMRForm{
		source:      source,
		target:      target,
		title:       title,
		description: desc,
	}
}

func (f *createMRForm) Focus() tea.Cmd {
	f.focused = formFieldSource
	return f.source.Focus()
}

func (f *createMRForm) Blur() {
	f.source.Blur()
	f.target.Blur()
	f.title.Blur()
	f.description.Blur()
	f.focused = formFieldSource
}

func (f *createMRForm) Reset() {
	f.source.Reset()
	f.target.Reset()
	f.title.Reset()
	f.description.Reset()
	f.draft = false
	f.focused = formFieldSource
}

func (f *createMRForm) SetSize(w, h int) {
	f.width = w
	f.source.SetWidth(w)
	f.target.SetWidth(w)
	f.title.SetWidth(w)

	// labels(3 * 2) + source(1) + arrow(1+1 padding) + target(1) + title(1) + desc label(2)
	overhead := 13
	descH := h - overhead
	if descH < 3 {
		descH = 3
	}
	f.description.SetWidth(w)
	f.description.SetHeight(descH)
}

func (f *createMRForm) NextField() tea.Cmd {
	f.blurCurrent()
	f.focused = (f.focused + 1) % formFieldCount
	return f.focusCurrent()
}

func (f *createMRForm) PrevField() tea.Cmd {
	f.blurCurrent()
	f.focused = (f.focused - 1 + formFieldCount) % formFieldCount
	return f.focusCurrent()
}

func (f *createMRForm) blurCurrent() {
	switch f.focused {
	case formFieldSource:
		f.source.Blur()
	case formFieldTarget:
		f.target.Blur()
	case formFieldTitle:
		f.title.Blur()
	case formFieldDescription:
		f.description.Blur()
	}
}

func (f *createMRForm) focusCurrent() tea.Cmd {
	switch f.focused {
	case formFieldSource:
		return f.source.Focus()
	case formFieldTarget:
		return f.target.Focus()
	case formFieldTitle:
		return f.title.Focus()
	case formFieldDescription:
		return f.description.Focus()
	}
	return nil
}

func (f *createMRForm) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	switch f.focused {
	case formFieldSource:
		f.source, cmd = f.source.Update(msg)
	case formFieldTarget:
		f.target, cmd = f.target.Update(msg)
	case formFieldTitle:
		f.title, cmd = f.title.Update(msg)
	case formFieldDescription:
		f.description, cmd = f.description.Update(msg)
	}
	return cmd
}

func (f *createMRForm) View() string {
	draftIcon := icon.Empty
	if f.draft {
		draftIcon = icon.Check
	}
	draftStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(style.MediumGray))
	if f.draft {
		draftStyle = draftStyle.Foreground(lipgloss.Color(style.Violet[400]))
	}

	return lipgloss.JoinVertical(0,
		labelStyle.Render("Branches"),
		f.source.View(),
		arrowStyle.Render(fmt.Sprintf("  %s", icon.ArrowDown)),
		f.target.View(),
		labelStyle.Render("Title"),
		f.title.View(),
		draftStyle.Render(fmt.Sprintf("[%s] Draft (ctrl+d)", draftIcon)),
		labelStyle.Render("Description"),
		f.description.View(),
	)
}

func modalContentWidth(windowW int) int {
	return modal.ContentWidth(windowW)
}

func modalContentHeight(windowH int) int {
	return modal.ContentHeight(windowH)
}
