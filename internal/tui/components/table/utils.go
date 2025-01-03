package table

import (
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

type InitModelParams struct {
	Rows      []Row
	Colums    []Column
	StyleFunc StyleFunc
	Height    int
}

func InitModel(params InitModelParams) Model {
	s := DefaultStyle()

	t := New(
		WithColumns(params.Colums),
		WithRows(params.Rows),
		WithFocused(true),
		WithHeight(params.Height),
		WithStyles(Styles(s)),
		WithStyleFunc(params.StyleFunc),
	)

	return t
}

func StyleIconsColumns(s Styles, iconColIdx []int) StyleFunc {
	return func(row, col int, value string) lipgloss.Style {
		type color = lipgloss.Color
		isIconCol := slices.Contains(iconColIdx, col)
		defStyle := s.Cell.Foreground

		iconStyle := map[string]lipgloss.Style{
			icon.Alert:       defStyle(color(style.Yellow[300])),
			icon.Time:        defStyle(color(style.Blue[500])),
			icon.Empty:       defStyle(color("")),
			icon.Dash:        defStyle(color(style.DarkGray)),
			icon.Check:       defStyle(color(style.Green[300])),
			icon.Clock:       defStyle(color(style.White)),
			icon.Rebase:      defStyle(color(style.Red[300])),
			icon.Cross:       defStyle(color(style.Red[300])),
			icon.Conflict:    defStyle(color(style.Yellow[400])),
			icon.Discussion:  defStyle(color(style.Green[300])),
			icon.Edit:        defStyle(color(style.Violet[400])),
			icon.CircleCheck: defStyle(color(style.Green[300])),
			icon.CircleCross: defStyle(color(style.Red[300])),
			icon.CirclePlay:  defStyle(color(style.Violet[300])),
			icon.CircleDash:  defStyle(color(style.Red[400])),
			icon.Gear:        defStyle(color(style.Yellow[300])),
			icon.Plus:        defStyle(color(style.Green[300])),
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

func ParseTimeString(d string) time.Time {
	t, _ := time.Parse(time.RFC3339, d)
	return t
}

func FormatTime(t time.Time) string {
	locale := t.Local()

	r := time.Since(locale)

	days := math.Floor(r.Hours()) / 24
	week := days / 7

	switch {
	case week > 4:
		return fmt.Sprintf("%.0fM", week/4)

	case days > 7:
		return fmt.Sprintf("%.0fw", week)

	case math.Floor(r.Hours()) > 24:
		return fmt.Sprintf("%.0fd", days)

	case math.Floor(r.Hours()) > 0:
		return fmt.Sprintf("%.0fh", r.Hours())

	case math.Floor(r.Minutes()) > 0:
		return fmt.Sprintf("%.0fm", r.Minutes())

	default:
		return fmt.Sprintf("%.0fs", r.Seconds())
	}
}

func FormatPercentage(v float32) string {
	if v == 0 {
		return ""
	}

	return fmt.Sprintf("%.2f %%", v)
}

func FormatDuration(d float32) string {
	seconds := d / 60.0

	x := time.Duration(d * float32(time.Second))

	switch {
	case seconds > 0:
		return fmt.Sprintf("%.0f m", x.Minutes())
	case seconds < 0:
		return fmt.Sprintf("%.0f m", x.Minutes())
	default:
		return ""
	}
}

func ColWidth(w int, p int) int {
	pr := float32(p) / float32(100)
	return int(float32(w) * pr)
}

func RenderIcon(b bool, i string) string {
	if b {
		return i
	}

	return icon.Empty
}

func GetColIndex(cols []Column, n string) int {
	return slices.IndexFunc(cols, func(c Column) bool {
		return c.Name == n
	})
}
