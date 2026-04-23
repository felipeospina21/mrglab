package table

import (
	"charm.land/lipgloss/v2"
	tstable "github.com/felipeospina21/tuishell/table"
	"github.com/felipeospina21/tuishell/style"
)

var pkgTheme style.Theme

// SetTheme sets the theme used by the table package and refreshes derived styles.
func SetTheme(t style.Theme) {
	pkgTheme = t
	TitleStyle = tstable.TitleStyle(t)
	DocStyle = tstable.DocStyle(t)
}

var (
	DefaultStyle = func() Styles {
		return tstable.ThemedStyles(pkgTheme)
	}

	TitleStyle = tstable.TitleStyle(pkgTheme)
	EmptyMsg   = tstable.EmptyMsg
	DocStyle   = tstable.DocStyle(pkgTheme)
)

// ThemedDocStyle returns a DocStyle for a given theme.
func ThemedDocStyle(t style.Theme) lipgloss.Style {
	return tstable.DocStyle(t)
}

// RenderPanel renders a table panel with loading state and optional header.
func RenderPanel(tbl *Model, loading bool, spinnerView, header string) string {
	return tstable.RenderPanel(pkgTheme, tbl, loading, spinnerView, header)
}
