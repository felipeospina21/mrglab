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

func GetMergeReqsColums(width int) []table.Column {
	id := int(float32(width) * 0.04)
	title := int(float32(width) * 0.5)
	author := int(float32(width) * 0.2)
	status := int(float32(width) * 0.1)
	i := int(float32(width) * 0.04)
	url := 0

	columns := []table.Column{
		// {
		// 	Title:  "",
		// 	Width:  &ctx.Styles.PrSection.CiCellWidth,
		// 	Grow:   new(bool),
		// 	Hidden: ciLayout.Hidden,
		// },
		// TODO: Add icons, decide columns
		{
			Title: "created at",
			Width: id,
		},
		{
			Title: "draft",
			Width: i,
		},
		{
			Title: "title",
			Width: title,
		},
		{
			Title: "author",
			Width: author,
		},
		{
			Title: "status",
			Width: status,
		},
		{
			Title: "has conflicts",
			Width: 0,
			// Hidden: true,
		},
		{
			Title: "discussions",
			Width: i,
		},
		{
			Title: "changes",
			Width: i,
		},
		{
			Title: "url",
			Width: url,
		},
		{
			Title: "description",
			Width: 0,
		},
		{
			Title: "id",
			Width: 0,
		},
	}

	return columns
}
