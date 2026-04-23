package tui

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/tuishell/style"
)

func applyOverride(target *color.Color, val *string) {
	if val != nil {
		*target = lipgloss.Color(*val)
	}
}

// BuildTheme returns DefaultTheme with any non-nil overrides applied.
func BuildTheme(overrides config.ThemeOverrides) style.Theme {
	t := DefaultTheme()
	if overrides.Preset != nil {
		if preset, ok := presets[*overrides.Preset]; ok {
			t = preset
		}
	}
	applyOverride(&t.Primary, overrides.Primary)
	applyOverride(&t.PrimaryBright, overrides.PrimaryBright)
	applyOverride(&t.PrimaryFg, overrides.PrimaryFg)
	applyOverride(&t.PrimaryDim, overrides.PrimaryDim)
	applyOverride(&t.Info, overrides.Info)
	applyOverride(&t.InfoBright, overrides.InfoBright)
	applyOverride(&t.Success, overrides.Success)
	applyOverride(&t.SuccessBright, overrides.SuccessBright)
	applyOverride(&t.Danger, overrides.Danger)
	applyOverride(&t.DangerBright, overrides.DangerBright)
	applyOverride(&t.Warning, overrides.Warning)
	applyOverride(&t.WarningBright, overrides.WarningBright)
	applyOverride(&t.Caution, overrides.Caution)
	applyOverride(&t.Text, overrides.Text)
	applyOverride(&t.TextInverse, overrides.TextInverse)
	applyOverride(&t.TextDimmed, overrides.TextDimmed)
	applyOverride(&t.Muted, overrides.Muted)
	applyOverride(&t.Dim, overrides.Dim)
	applyOverride(&t.Border, overrides.Border)
	applyOverride(&t.ModalBorder, overrides.ModalBorder)
	applyOverride(&t.SurfaceDim, overrides.SurfaceDim)
	applyOverride(&t.SelectionBorder, overrides.SelectionBorder)
	applyOverride(&t.StatusText, overrides.StatusText)
	applyOverride(&t.StatusNormal, overrides.StatusNormal)
	applyOverride(&t.StatusLoading, overrides.StatusLoading)
	applyOverride(&t.StatusError, overrides.StatusError)
	applyOverride(&t.StatusDev, overrides.StatusDev)
	applyOverride(&t.StatusAccent1, overrides.StatusAccent1)
	applyOverride(&t.StatusAccent2, overrides.StatusAccent2)
	return t
}

// DefaultTheme returns a GitLab-inspired orange/tangerine theme.
func DefaultTheme() style.Theme {
	return style.Theme{
		Primary:       lipgloss.Color("#FC6D26"),
		PrimaryBright: lipgloss.Color("#E24329"),
		PrimaryFg:     lipgloss.Color("#FFF4ED"),
		PrimaryDim:    lipgloss.Color("#5C2900"),

		Info:          lipgloss.Color("#428FDC"),
		InfoBright:    lipgloss.Color("#6B4FBB"),
		Success:       lipgloss.Color("#6beaaf"),
		SuccessBright: lipgloss.Color("#3ad994"),
		Danger:        lipgloss.Color("#f9a8a8"),
		DangerBright:  lipgloss.Color("#f47575"),
		Warning:       lipgloss.Color("#ffe043"),
		WarningBright: lipgloss.Color("#ffcc14"),
		Caution:       lipgloss.Color("#ff8237"),

		Text:            lipgloss.Color("#C4C4C4"),
		TextInverse:     lipgloss.Color("#111"),
		TextDimmed:      lipgloss.Color("#777777"),
		Muted:           lipgloss.Color("#999999"),
		Dim:             lipgloss.Color("#444444"),
		Border:          lipgloss.Color("#3f4145"),
		ModalBorder:     lipgloss.Color("#666666"),
		SurfaceDim:      lipgloss.Color("#1A1A2E"),
		SelectionBorder: lipgloss.Color("#FC6D26"),

		StatusText:    lipgloss.Color("#FFFDF5"),
		StatusNormal:  lipgloss.Color("#5C2900"),
		StatusLoading: lipgloss.Color("#1A7A94"),
		StatusError:   lipgloss.Color("#CE3060"),
		StatusDev:     lipgloss.Color("#4E8212"),
		StatusAccent1: lipgloss.Color("#FC6D26"),
		StatusAccent2: lipgloss.Color("#E24329"),
	}
}
