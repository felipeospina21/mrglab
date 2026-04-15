package table

import (
	"slices"

	"charm.land/lipgloss/v2"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

func StyleIconsColumns(s Styles, iconColIdx []int) StyleFunc {
	return func(row, col int, value string) lipgloss.Style {
		isIconCol := slices.Contains(iconColIdx, col)
		defStyle := s.Cell.Foreground

		iconStyle := map[string]lipgloss.Style{
			icon.Alert:         defStyle(lipgloss.Color(style.Yellow[300])),
			icon.AlertFill:     defStyle(lipgloss.Color(style.Yellow[400])),
			icon.Time:          defStyle(lipgloss.Color(style.Blue[500])),
			icon.Empty:         defStyle(lipgloss.Color("")),
			icon.Dash:          defStyle(lipgloss.Color(style.DarkGray)),
			icon.Check:         defStyle(lipgloss.Color(style.Green[300])),
			icon.Clock:         defStyle(lipgloss.Color(style.White)),
			icon.Rebase:        defStyle(lipgloss.Color(style.Red[300])),
			icon.Cross:         defStyle(lipgloss.Color(style.Red[300])),
			icon.Conflict:      defStyle(lipgloss.Color(style.Yellow[400])),
			icon.Discussion:    defStyle(lipgloss.Color(style.Green[300])),
			icon.Edit:          defStyle(lipgloss.Color(style.Violet[400])),
			icon.CircleCheck:   defStyle(lipgloss.Color(style.Green[300])),
			icon.CircleCross:   defStyle(lipgloss.Color(style.Red[300])),
			icon.CirclePlay:    defStyle(lipgloss.Color(style.Violet[300])),
			icon.CircleDash:    defStyle(lipgloss.Color(style.Red[400])),
			icon.CircleRunning: defStyle(lipgloss.Color(style.Blue[400])),
			icon.CirclePause:   defStyle(lipgloss.Color(style.Yellow[400])),
			icon.CircleCancel:  defStyle(lipgloss.Color(style.White)),
			icon.CircleSkip:    defStyle(lipgloss.Color(style.Orange[400])),
			icon.CircleDot:     defStyle(lipgloss.Color(style.White)),
			icon.Gear:          defStyle(lipgloss.Color(style.Yellow[300])),
			icon.Plus:          defStyle(lipgloss.Color(style.Green[300])),
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
