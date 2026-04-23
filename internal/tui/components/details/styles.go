package details

import (
	"fmt"
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/tuishell/style"
)

var pkgTheme style.Theme

// themeColors holds pre-computed hex strings for use in styledIcon.
// Initialized with legacy palette defaults so tests pass without SetTheme.
var themeColors = struct {
	Info, InfoBright                         string
	Success, SuccessBright                   string
	Danger, DangerBright                     string
	Warning, WarningBright                   string
	Caution                                  string
	Text, TextInverse                        string
	Border                                   string
	Primary, PrimaryBright                   string
}{
	Info: "#3ac4d9", InfoBright: "#1ca7be",
	Success: "#6beaaf", SuccessBright: "#3ad994",
	Danger: "#f9a8a8", DangerBright: "#f47575",
	Warning: "#ffe043", WarningBright: "#ffcc14",
	Caution: "#ff8237",
	Text: "#C4C4C4", TextInverse: "#111",
	Border: "#3f4145",
	Primary: "#b8a6ff", PrimaryBright: "#9673ff",
}

// SetTheme sets the theme used by the details package and refreshes derived styles.
func SetTheme(t style.Theme) {
	pkgTheme = t
	themeColors.Info = colorHex(t.Info)
	themeColors.InfoBright = colorHex(t.InfoBright)
	themeColors.Success = colorHex(t.Success)
	themeColors.SuccessBright = colorHex(t.SuccessBright)
	themeColors.Danger = colorHex(t.Danger)
	themeColors.DangerBright = colorHex(t.DangerBright)
	themeColors.Warning = colorHex(t.Warning)
	themeColors.WarningBright = colorHex(t.WarningBright)
	themeColors.Caution = colorHex(t.Caution)
	themeColors.Text = colorHex(t.Text)
	themeColors.TextInverse = colorHex(t.TextInverse)
	themeColors.Border = colorHex(t.Border)
	themeColors.Primary = colorHex(t.Primary)
	themeColors.PrimaryBright = colorHex(t.PrimaryBright)
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

// colorHex converts a color.Color to a hex string like "#rrggbb".
func colorHex(c color.Color) string {
	if c == nil {
		return ""
	}
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
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
			Border(lipgloss.NormalBorder(), true, false, true, true).
			BorderForeground(lipgloss.Color("#3f4145"))

	iconStyle = func(c string) lipgloss.Style {
		return lipgloss.NewStyle().Foreground(lipgloss.Color(c)).MarginLeft(2)
	}

	contentStyle = lipgloss.NewStyle().MarginLeft(LeftMargin)

	sectionTextStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#C4C4C4")).MarginLeft(1)
	sectionTitleStyle        = sectionTextStyle.Bold(true).MarginLeft(0)
	sectionIndentedTextStyle = sectionTextStyle.MarginLeft(LeftMargin).Foreground(lipgloss.Color("#3f4145"))

	selectedDiscussionStyle = lipgloss.NewStyle().
				Border(lipgloss.ThickBorder(), false, false, false, true).
				BorderForeground(lipgloss.Color("#9673ff")).
				PaddingLeft(1)
)
