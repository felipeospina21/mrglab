package app

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/gql"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/style"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) GetMergeRequestModel(msg task.TaskMsg) func() table.Model {
	return func() table.Model {
		mrMsg := msg.Msg.(message.MergeRequestsListFetchedMsg)
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

type MergeRequestDetails struct {
	Pipelines   string
	Discussions string
}

func (m Model) GetMergeRequestDetails(msg task.TaskMsg) func() MergeRequestDetails {
	return func() MergeRequestDetails {
		mrMsg := msg.Msg.(message.MergeRequestFetchedMsg)
		pStyle := lipgloss.NewStyle().MarginLeft(2).Render
		var discussions, pipelines strings.Builder

		processPipeline(&pipelines, mrMsg.Stages)
		processComments(&discussions, mrMsg.Discussions)

		return MergeRequestDetails{
			Pipelines:   pStyle(pipelines.String()),
			Discussions: discussions.String(),
		}
	}
}

func processPipeline(content *strings.Builder, stages []gql.CiStageNode) {
	baseStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(style.White))
	headStyle := baseStyle.Bold(true)

	content.WriteString(headStyle.Render(fmt.Sprintf("%s Pipeline", icon.Pipeline)))
	content.WriteString("\n\n")

	for _, stage := range stages {
		styledIcon := func(i stageIcon) lipgloss.Style {
			return lipgloss.NewStyle().Foreground(lipgloss.Color(i.color))
		}

		stageStatus := getStageIconStatus(stage.Status)
		content.WriteString(styledIcon(stageStatus).Render(stageStatus.icon))
		content.WriteString(" ")
		content.WriteString(baseStyle.Render(stage.Name))
		content.WriteString("\n")

		lower := strings.ToLower
		if lower(stage.Status) != "success" || lower(stage.Status) != "manual" {
			for _, node := range stage.Jobs.Nodes {
				if lower(node.Status) != "success" {
					nodeStatus := getStageIconStatus(node.Status)
					indentStyle := baseStyle.MarginLeft(2).Foreground(lipgloss.Color(style.DarkGray)).Render
					iconStyle := styledIcon(nodeStatus).Render
					textStyle := baseStyle.Render

					content.WriteString(indentStyle("└ "))
					content.WriteString(iconStyle(nodeStatus.icon))
					content.WriteString(" ")
					content.WriteString(textStyle(node.Name))
					content.WriteString("\n")
				}
			}
		}
		// TODO: iterate stage.Jobs.Nodes to render nested jobs
	}
	content.WriteString("\n\n")
}

type stageIcon struct {
	icon  string
	color string
}

func getStageIconStatus(s string) stageIcon {
	icons := map[string]stageIcon{
		"running":              {icon: icon.CircleRunning, color: style.Blue[400]},
		"preparing":            {icon: icon.CirclePause, color: style.Yellow[400]},
		"success":              {icon: icon.CircleCheck, color: style.Green[400]},
		"failed":               {icon: icon.CircleCross, color: style.Red[400]},
		"skipped":              {icon: icon.CircleSkip, color: style.Orange[400]},
		"manual":               {icon: icon.Gear, color: style.White},
		"created":              {icon: icon.CircleDot, color: style.White},
		"waiting_for_resource": {icon: icon.CircleQuestion, color: style.White},
		"scheduled":            {icon: icon.Time, color: style.White},
		"pending":              {icon: icon.CirclePause, color: style.White},
		"canceled":             {icon: icon.CircleCancel, color: style.White},
	}

	v, ok := icons[strings.ToLower(s)]
	if ok {
		return v
	}

	return stageIcon{icon: icon.Dash}
}

func processComments(content *strings.Builder, discussions []gql.DiscussionNode) {
	separator := strings.Repeat("-", 5)

	content.WriteString(fmt.Sprintf("**%s Discussions**", icon.Discussion))
	content.WriteString("\n\n")

	hasDiscussions := slices.ContainsFunc(discussions, func(d gql.DiscussionNode) bool {
		return d.Resolvable
	})

	if !hasDiscussions {
		content.WriteString("... *No Discussions*")
	} else {
		for _, discussion := range discussions {
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
}

func timeAgo(time string) string {
	return fmt.Sprintf("_%s ago_", time)
}
