// Package style provides color palettes and shared lipgloss styles for the TUI.
package style

import "github.com/felipeospina21/tuishell/style"

type colorShades = style.ColorShades

var (
	Blue   = style.Blue
	Red    = style.Red
	Green  = style.Green
	Yellow = style.Yellow
	Violet = style.Violet
	Orange = style.Orange
)

var (
	DarkGray   = "#3f4145"
	DarkerGray = "#1e1e24"
	MediumGray = "#999999"
	White      = "#C4C4C4"
	Black      = "#111"

	StatuslineText     = "#FFFDF5"
	StatuslineEncoding = "#A550DF"
	StatuslineProject  = "#6124DF"

	StatuslineModeNormal  = Violet[600]
	StatuslineModeLoading = "#1A7A94"
	StatuslineModeError   = "#CE3060"
	StatuslineModeDev     = "#4E8212"
)
