package projects

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/mitchellh/mapstructure"
)

const ellipsis = "…"

type Model struct {
	List list.Model
	ctx  *context.AppContext
}

func New(ctx *context.AppContext) Model {
	projects := config.GlobalConfig.Filters.Projects

	var li []list.Item
	var i Item

	for _, val := range projects {
		err := mapstructure.Decode(val, &i)
		if err != nil {
			fmt.Println("Error loading projects", err)
			os.Exit(1)
		}
		li = append(li, i)

	}

	l := list.New(li, itemDelegate{ShowDescription: true, Styles: NewDefaultItemStyles()}, 30, len(li))
	l.Title = "Projects"
	l.Styles.Title = TitleStyle
	// l.Styles.PaginationStyle = PaginationStyle
	l.SetShowHelp(false)

	ctx.IsLeftPanelOpen = true

	return Model{
		List: l,
		ctx:  ctx,
	}
}

func (m *Model) SelectProject() {
	s := m.List.SelectedItem()
	i, ok := s.(Item)

	if ok {
		m.ctx.SelectedProject.ID = i.ID
		m.ctx.SelectedProject.Name = i.Name
	}
}

type Item struct {
	Name, ID string
}

func (i Item) Title() string       { return i.Name }
func (i Item) Description() string { return i.ID }
func (i Item) FilterValue() string { return i.Name }

type itemDelegate struct {
	ShowDescription bool
	Styles          DefaultItemStyles
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	var (
		title, desc  string
		matchedRunes []int
		s            = &d.Styles
	)

	if i, ok := listItem.(Item); ok {
		title = i.Title()
		desc = i.Description()
	} else {
		return
	}

	if m.Width() <= 0 {
		// short-circuit
		return
	}

	// Prevent text from exceeding list width
	textwidth := m.Width() - s.NormalTitle.GetPaddingLeft() - s.NormalTitle.GetPaddingRight()
	title = ansi.Truncate(title, textwidth, ellipsis)
	if d.ShowDescription {
		var lines []string
		for i, line := range strings.Split(desc, "\n") {
			if i >= d.Height()-1 {
				break
			}
			lines = append(lines, ansi.Truncate(line, textwidth, ellipsis))
		}
		desc = strings.Join(lines, "\n")
	}

	// Conditions
	var (
		isSelected  = index == m.Index()
		emptyFilter = m.FilterState() == list.Filtering && m.FilterValue() == ""
		isFiltered  = m.FilterState() == list.Filtering || m.FilterState() == list.FilterApplied
	)

	// if isFiltered && index < len(m.filteredItems) {
	// 	// Get indices of matched characters
	// 	matchedRunes = m.MatchesForItem(index)
	// }

	if emptyFilter {
		title = s.DimmedTitle.Render(title)
		desc = s.DimmedDesc.Render(desc)
	} else if isSelected && m.FilterState() != list.Filtering {
		if isFiltered {
			// Highlight matches
			unmatched := s.SelectedTitle.Inline(true)
			matched := unmatched.Inherit(s.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = s.SelectedTitle.Render(title)
		desc = s.SelectedDesc.Render(desc)
	} else {
		if isFiltered {
			// Highlight matches
			unmatched := s.NormalTitle.Inline(true)
			matched := unmatched.Inherit(s.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = s.NormalTitle.Render(title)
		desc = s.NormalDesc.Render(desc)
	}

	if d.ShowDescription {
		fmt.Fprintf(w, "%s\n%s", title, desc) //nolint: errcheck
		return
	}
	fmt.Fprintf(w, "%s", title) //nolint: errcheck
}

func GetFrameSize() (int, int) {
	x, y := DocStyle.GetFrameSize()
	xItem, yItem := ItemStyle.GetFrameSize()
	xTitle, yTitle := TitleStyle.GetFrameSize()

	return x + xItem + xTitle, y + yItem + yTitle
}

func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.LeftPanel
}
