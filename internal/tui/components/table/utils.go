package table

import (
	"fmt"
	"math"
	"time"
)

type TableColIndex uint

type TableCol struct {
	Name string
	Idx  TableColIndex
}

type InitModelParams struct {
	Rows      []Row
	Colums    []Column
	StyleFunc StyleFunc
}

func InitModel(params InitModelParams) Model {
	s := DefaultStyle()

	t := New(
		WithColumns(params.Colums),
		WithRows(params.Rows),
		WithFocused(true),
		WithHeight(len(params.Rows)+1),
		WithStyles(Styles(s)),
		WithStyleFunc(params.StyleFunc),
	)

	return t
}

// func StyleIconsColumns(s Styles, iconColIdx []int) StyleFunc {
// 	return func(row, col int, value string) lipgloss.Style {
// 		isIconCol := slices.Contains(iconColIdx, col)
//
// 		if isIconCol {
// 			switch value {
// 			case icon.Check:
// 				return s.Cell.Foreground(lipgloss.Color(style.Green[300]))
// 			case icon.Clock:
// 				return s.Cell.Foreground(lipgloss.Color(style.Yellow[300]))
// 			case icon.CircleCheck:
// 				return s.Cell.Foreground(lipgloss.Color(style.Green[300]))
// 			case icon.CircleCross:
// 				return s.Cell.Foreground(lipgloss.Color(style.Red[300]))
// 			case icon.CirclePlay:
// 				return s.Cell.Foreground(lipgloss.Color(style.Violet[400]))
// 			case icon.Gear:
// 				return s.Cell.Foreground(lipgloss.Color(style.Yellow[100]))
//
// 			}
// 		}
//
// 		return s.Cell
// 	}
// }

func FormatTime(d string) string {
	t, _ := time.Parse(time.RFC3339, d)

	locale := t.Local()

	r := time.Since(locale)

	days := math.Floor(r.Hours()) / 24
	week := days / 7

	switch {
	case week > 4:
		return fmt.Sprintf("%.0f M", week/4)

	case days > 7:
		return fmt.Sprintf("%.0f w", week)

	case math.Floor(r.Hours()) > 24:
		return fmt.Sprintf("%.0f d", days)

	case math.Floor(r.Hours()) > 0:
		return fmt.Sprintf("%.0f h", r.Hours())

	case math.Floor(r.Minutes()) > 0:
		return fmt.Sprintf("%.0f m", r.Minutes())

	default:
		return fmt.Sprintf("%.0f s", r.Seconds())
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
