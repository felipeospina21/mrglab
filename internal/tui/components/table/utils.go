package table

import (
	"slices"

	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
)

func StyleIconsColumns(s Styles, iconColIdx []int) StyleFunc {
	t := pkgTheme
	return func(row, col int, value string) lipgloss.Style {
		isIconCol := slices.Contains(iconColIdx, col)
		defStyle := s.Cell.Foreground

		iconStyle := map[string]lipgloss.Style{
			icon.Alert:         defStyle(t.Warning),
			icon.AlertFill:     defStyle(t.WarningBright),
			icon.Time:          defStyle(t.InfoBright),
			icon.Empty:         defStyle(nil),
			icon.Dash:          defStyle(t.Border),
			icon.Check:         defStyle(t.Success),
			icon.Clock:         defStyle(t.Text),
			icon.Rebase:        defStyle(t.Danger),
			icon.Cross:         defStyle(t.Danger),
			icon.Conflict:      defStyle(t.WarningBright),
			icon.Discussion:    defStyle(t.Success),
			icon.Edit:          defStyle(t.PrimaryBright),
			icon.CircleCheck:   defStyle(t.Success),
			icon.CircleCross:   defStyle(t.Danger),
			icon.CirclePlay:    defStyle(t.Primary),
			icon.CircleDash:    defStyle(t.DangerBright),
			icon.CircleRunning: defStyle(t.Info),
			icon.CirclePause:   defStyle(t.WarningBright),
			icon.CircleCancel:  defStyle(t.Text),
			icon.CircleSkip:    defStyle(t.Caution),
			icon.CircleDot:     defStyle(t.Text),
			icon.Gear:          defStyle(t.Warning),
			icon.Plus:          defStyle(t.Success),
		}

		if isIconCol {
			v, ok := iconStyle[value]
			if ok {
				return v
			}
			return s.Cell
		}

		return s.Cell
	}
}
