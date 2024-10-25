package app

import (
	"fmt"
	"strings"

	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) GetMergeRequestModel(msg task.TaskMsg) func() table.Model {
	return func() table.Model {
		mrMsg := msg.Msg.(message.MergeRequestsFetchedMsg)
		rows := mergerequests.GetTableRows(mrMsg)
		mainPanelHeaderHeight := 1
		return table.InitModel(table.InitModelParams{
			Rows:   rows,
			Colums: mergerequests.GetTableColums(m.ctx.Window.Width),
			StyleFunc: table.StyleIconsColumns(
				table.Styles(table.DefaultStyle()),
				mergerequests.IconCols(),
			),
			Height: m.ctx.PanelHeight - mainPanelHeaderHeight,
		})
	}
}

func (m Model) GetMergeRequestDiscussions(msg task.TaskMsg) func() string {
	return func() string {
		notesMsg := msg.Msg.(message.MergeRequestNotesFetchedMsg)

		var content strings.Builder
		// var resolved strings.Builder
		separator := strings.Repeat("-", 5)
		for _, discussion := range notesMsg.Notes {
			for _, note := range discussion {
				if note.Resolved {
					// TODO: check how can this be styled better
					content.WriteString(icon.Check + " ")
				}
				createdAt := table.FormatTime(note.CreatedAt)
				author := note.Author.Name
				body := note.Body

				content.WriteString(fmt.Sprintf("`%s` ", author))
				content.WriteString(fmt.Sprintf("*%s ago*", createdAt))
				content.WriteString("\n")

				content.WriteString(body)
				content.WriteString("\n\n")

			}
			content.WriteString(separator)
			content.WriteString("\n\n")
		}

		// content.WriteString(resolved.String())
		return content.String()
	}
}

type BuildStringArgs struct {
	builder    *strings.Builder
	s          string
	addNewLine bool
	withSpace  bool
}

func BuildString(args BuildStringArgs) *strings.Builder {
	if args.withSpace {
		args.builder.WriteString(args.s + " ")
	} else {
		args.builder.WriteString(args.s)
	}

	if !args.addNewLine {
		args.builder.WriteString("\n\n")
	}
	return args.builder
}
