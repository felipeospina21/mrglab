package app

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
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
		var resolved strings.Builder
		separator := strings.Repeat("-", 5)
		for _, discussion := range notesMsg.Notes {
			for idx, note := range discussion {
				createdAt := table.FormatTime(note.CreatedAt)
				author := note.Author.Name
				body := note.Body

				if idx == 0 {
					content.WriteString(icon.User + " ")
					content.WriteString(author + " ")
					content.WriteString("-" + " ")
					content.WriteString(icon.Clock + " ")
					content.WriteString(createdAt + " ")
					content.WriteString("\n\n")

					content.WriteString(icon.Discussion + " ")
					content.WriteString(body)
					content.WriteString("\n\n")
				} else {
					// FIX: make comments responses to wrap
					l := lipgloss.NewStyle().MarginLeft(4).Render
					// w := lipgloss.NewStyle().Width(viewportWidth - 40).Render
					content.WriteString(l("--> "))
					content.WriteString(icon.User + " ")
					content.WriteString(author + " ")
					content.WriteString("-" + " ")
					content.WriteString(icon.Clock + " ")
					content.WriteString(createdAt + " ")
					content.WriteString("\n\n")

					content.WriteString(l("    "))
					content.WriteString(icon.Discussion + " ")
					content.WriteString(body)
					content.WriteString("\n\n")

				}

			}
			content.WriteString("\t\t" + separator)
			content.WriteString("\n\n")
		}

		content.WriteString(resolved.String())
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
