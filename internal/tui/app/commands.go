package app

import (
	"fmt"
	"slices"
	"strings"

	"github.com/felipeospina21/mrglab/internal/gql"
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
		separator := strings.Repeat("-", 5)

		content.WriteString(fmt.Sprintf("**%s Discussions**", icon.Discussion))
		content.WriteString("\n\n")

		hasDiscussions := slices.ContainsFunc(notesMsg.Discussions, func(d gql.DiscussionNode) bool {
			return d.Resolvable
		})

		if !hasDiscussions {
			content.WriteString("... *No Discussions*")
		} else {
			for _, discussion := range notesMsg.Discussions {
				if !discussion.Resolvable {
					continue
				}

				content.WriteString(separator)
				if discussion.Resolved {
					resolvedAt := table.FormatTime(discussion.ResolvedAt)
					content.WriteString(fmt.Sprintf(" **%s %s** ", icon.Check, timeAgo(resolvedAt)))
				} else {
					content.WriteString(fmt.Sprintf(" %s ", icon.Dash))
				}
				content.WriteString(separator)
				content.WriteString("\n\n")

				for _, note := range discussion.Notes.Nodes {
					author := note.Author.Name
					body := note.Body
					createdAt := table.FormatTime(note.CreatedAt)

					if !note.Resolvable {
						before, _, found := strings.Cut(body, "(")
						if found {
							content.WriteString(
								fmt.Sprintf(
									"*%s %s %s %s* ",
									icon.Dot,
									author,
									before,
									timeAgo(createdAt),
								),
							)
							content.WriteString("\n\n")
						}
						continue
					}

					content.WriteString(fmt.Sprintf("`%s` ", author))
					content.WriteString(timeAgo(createdAt))
					content.WriteString("\n")

					content.WriteString(body)
					content.WriteString("\n\n")

				}
				content.WriteString("\n\n")

			}
		}

		return content.String()
	}
}

func timeAgo(time string) string {
	return fmt.Sprintf("_%s ago_", time)
}
