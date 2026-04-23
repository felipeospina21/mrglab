package details

import (
	"fmt"
	"image/color"

	"github.com/felipeospina21/tuishell/style"
)

// icons holds hex color strings for icon styling.
// Populated from theme via SetTheme; defaults match the legacy palette
// so tests pass without calling SetTheme.
var icons = struct {
	Success       string
	SuccessBright string
	Danger        string
	DangerBright  string
	Info          string
	InfoBright    string
	Warning       string
	WarningBright string
	Caution       string
	Text          string
	TextInverse   string
	Border        string
	Primary       string
	PrimaryBright string
}{
	Success:       "#3ad994",
	SuccessBright: "#3ad994",
	Danger:        "#f9a8a8",
	DangerBright:  "#f47575",
	Info:          "#3ac4d9",
	InfoBright:    "#1ca7be",
	Warning:       "#ffe043",
	WarningBright: "#ffcc14",
	Caution:       "#ff8237",
	Text:          "#C4C4C4",
	TextInverse:   "#111",
	Border:        "#3f4145",
	Primary:       "#b8a6ff",
	PrimaryBright: "#9673ff",
}

// refreshIcons updates icon colors from the given theme.
func refreshIcons(t style.Theme) {
	icons.Success = colorHex(t.Success)
	icons.SuccessBright = colorHex(t.SuccessBright)
	icons.Danger = colorHex(t.Danger)
	icons.DangerBright = colorHex(t.DangerBright)
	icons.Info = colorHex(t.Info)
	icons.InfoBright = colorHex(t.InfoBright)
	icons.Warning = colorHex(t.Warning)
	icons.WarningBright = colorHex(t.WarningBright)
	icons.Caution = colorHex(t.Caution)
	icons.Text = colorHex(t.Text)
	icons.TextInverse = colorHex(t.TextInverse)
	icons.Border = colorHex(t.Border)
	icons.Primary = colorHex(t.Primary)
	icons.PrimaryBright = colorHex(t.PrimaryBright)
}

// colorHex converts a color.Color to a hex string like "#rrggbb".
func colorHex(c color.Color) string {
	if c == nil {
		return ""
	}
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}
