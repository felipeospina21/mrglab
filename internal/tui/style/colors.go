// Package style provides backward-compatible color constants for tests.
// Production code should use theme tokens from style.Theme instead.
package style

// ColorShades maps shade numbers (50–950) to hex color strings.
type ColorShades = map[uint]string

// Palette maps kept for test backward compatibility.
var Blue = ColorShades{
	400: "#3ac4d9", 500: "#1ca7be",
}

var Red = ColorShades{
	300: "#f9a8a8", 400: "#f47575",
}

var Green = ColorShades{
	300: "#6beaaf", 400: "#3ad994",
}

var Yellow = ColorShades{
	300: "#ffe043", 400: "#ffcc14",
}

var Violet = ColorShades{
	50: "#f2f0ff", 300: "#b8a6ff", 400: "#9673ff", 600: "#6914ff", 800: "#4c01d6",
}

var Orange = ColorShades{
	400: "#ff8237",
}

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
