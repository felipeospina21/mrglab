package details

import (
	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/tuishell/style"
)

var pkgTheme style.Theme

// SetTheme sets the theme used by the details package and refreshes derived styles.
func SetTheme(t style.Theme) {
	pkgTheme = t
	refreshIcons(t)
	PanelStyle = lipgloss.NewStyle().
		MarginTop(1).
		Border(lipgloss.NormalBorder(), true, false, true, true).
		BorderForeground(t.Border)
	sectionTextStyle = lipgloss.NewStyle().Foreground(t.Text).MarginLeft(1)
	sectionTitleStyle = sectionTextStyle.Bold(true).MarginLeft(0)
	sectionIndentedTextStyle = sectionTextStyle.MarginLeft(LeftMargin).Foreground(t.Border)
	selectedDiscussionStyle = lipgloss.NewStyle().
		Border(lipgloss.ThickBorder(), false, false, false, true).
		BorderForeground(t.PrimaryBright).
		PaddingLeft(1)
}

var (
	MdTitle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0).Margin(0)
	}()

	MdInfo = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return MdTitle.BorderStyle(b)
	}()

	PanelStyle = lipgloss.NewStyle().
			MarginTop(1).
			Border(lipgloss.NormalBorder(), true, false, true, true)

	iconStyle = func(c string) lipgloss.Style {
		return lipgloss.NewStyle().Foreground(lipgloss.Color(c)).MarginLeft(2)
	}

	contentStyle = lipgloss.NewStyle().MarginLeft(LeftMargin)

	sectionTextStyle         = lipgloss.NewStyle().MarginLeft(1)
	sectionTitleStyle        = sectionTextStyle.Bold(true).MarginLeft(0)
	sectionIndentedTextStyle = sectionTextStyle.MarginLeft(LeftMargin)

	selectedDiscussionStyle = lipgloss.NewStyle().
				Border(lipgloss.ThickBorder(), false, false, false, true).
				PaddingLeft(1)
)
