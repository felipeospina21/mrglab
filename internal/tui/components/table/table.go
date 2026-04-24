package table

import (
	tstable "github.com/felipeospina21/tuishell/table"
)

// Re-export tuishell table types so existing mrglab code compiles unchanged.
type (
	Model           = tstable.Model
	Row             = tstable.Row
	Column          = tstable.Column
	KeyMap          = tstable.KeyMap
	Styles          = tstable.Styles
	StyleFunc       = tstable.StyleFunc
	Option          = tstable.Option
	InitModelParams = tstable.InitModelParams
)

var (
	New              = tstable.New
	InitModel        = tstable.InitModel
	DefaultKeyMap    = tstable.DefaultKeyMap
	DefaultStyles    = tstable.DefaultStyles
	WithColumns      = tstable.WithColumns
	WithRows         = tstable.WithRows
	WithHeight       = tstable.WithHeight
	WithWidth        = tstable.WithWidth
	WithFocused      = tstable.WithFocused
	WithStyles       = tstable.WithStyles
	WithStyleFunc    = tstable.WithStyleFunc
	WithKeyMap       = tstable.WithKeyMap
	FormatTime       = tstable.FormatTime
	FormatPercentage = tstable.FormatPercentage
	FormatDuration   = tstable.FormatDuration
	ColWidth         = tstable.ColWidth
	RenderIcon       = tstable.RenderIcon
	GetColIndex      = tstable.GetColIndex
	ParseTimeString  = tstable.ParseTimeString
)
