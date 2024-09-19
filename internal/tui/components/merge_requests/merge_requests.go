package mergerequests

import (
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
)

type Model struct {
	Table table.Model
	ctx   *context.AppContext
}

func New(ctx *context.AppContext) Model {
	return Model{
		Table: table.Model{},
		ctx:   ctx,
	}
}

// func New() {
// 	table.InitModel(table.InitModelParams{
// 		Rows:      r,
// 		Colums:    table.GetMergeReqsColums(window.Width - 10),
// 		StyleFunc: table.StyleIconsColumns(table.Styles(table.DefaultStyle()), table.MergeReqsIconCols),
// 	})
// }
