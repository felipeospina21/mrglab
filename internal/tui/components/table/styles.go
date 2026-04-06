package table

import (
	"charm.land/lipgloss/v2"
	tstable "github.com/felipeospina21/tuishell/table"
	"github.com/felipeospina21/tuishell/style"
	mrgstyle "github.com/felipeospina21/mrglab/internal/tui/style"
)

var theme = mrgstyle.DefaultTheme()

var (
	DefaultStyle = func() Styles {
		return tstable.ThemedStyles(theme)
	}

	TitleStyle = tstable.TitleStyle(theme)
	EmptyMsg   = tstable.EmptyMsg
	DocStyle   = tstable.DocStyle(theme)
)

// ThemedDocStyle returns a DocStyle for a given theme.
func ThemedDocStyle(t style.Theme) lipgloss.Style {
	return tstable.DocStyle(t)
}

// RenderPanel renders a table panel with loading state and optional header.
func RenderPanel(tbl *Model, loading bool, spinnerView, header string) string {
	return tstable.RenderPanel(theme, tbl, loading, spinnerView, header)
}
