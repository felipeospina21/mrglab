package details

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/icon"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

type styledIcon struct {
	icon  string
	color string
}

func (m Model) GetViewportContent(b string, mr MergeRequestDetails) string {
	var content strings.Builder
	mergeStatus := strings.ToLower(m.ctx.SelectedMR.Status)

	content.WriteString(renderBranches(mr.Branches[0], mr.Branches[1]))
	content.WriteString("\n\n")

	if mergeStatus == "mergeable" {
		content.WriteString(renderStatus(mergeStatus))
		content.WriteString("\n\n")
	}

	content.WriteString(m.renderWithStyle(b))
	content.WriteString("\n\n")
	content.WriteString(renderPipelines(mr.Pipelines))
	content.WriteString("\n\n")
	content.WriteString(renderApprovals(mr.Approvals))
	content.WriteString("\n\n")
	content.WriteString(renderDiscussions(mr.Discussions, m))

	return content.String()
}

func renderIndentedText(content *strings.Builder, i styledIcon, text string) {
	indentStyle := sectionIndentedTextStyle.Render
	iconStyle := iconStyle(i.color).MarginLeft(0).Render
	content.WriteString(indentStyle("└ "))
	content.WriteString(iconStyle(i.icon))
	content.WriteString(sectionTextStyle.Render(text))
	content.WriteString("\n")
}

func renderStatus(status string) string {
	s := contentStyle.
		Background(lipgloss.Color(style.Green[400])).
		Foreground(lipgloss.Color("#111")).
		Padding(0, 1).
		Bold(true)

	var content strings.Builder
	content.WriteString(icon.Mergeable)
	content.WriteString(strings.ToUpper(status[:1]) + status[1:])

	return s.Render(content.String())
}

func renderBranches(source, target string) string {
	s := contentStyle.Foreground(lipgloss.Color(style.DarkGray))
	var content strings.Builder
	content.WriteString(icon.Rebase)
	content.WriteString(target)
	content.WriteString(" <- ")
	content.WriteString(source)

	return s.Render(content.String())
}

func renderApprovals(approvals []gitlab.ApprovalRule) string {
	var content strings.Builder
	content.WriteString(sectionTitleStyle.Render(fmt.Sprintf("%s Approvals", icon.Approval)))
	content.WriteString("\n\n")

	rendeRule := func(content *strings.Builder, i styledIcon, rule string) {
		content.WriteString(iconStyle(i.color).Render(i.icon))
		content.WriteString(sectionTextStyle.Render(rule))
		content.WriteString("\n")
	}
	for _, rule := range approvals {
		if len(rule.ApprovedBy.Nodes) > 0 {
			i := styledIcon{icon: icon.CircleCross, color: style.Red[400]}
			if rule.Approved {
				i = styledIcon{icon: icon.CircleCheck, color: style.Green[400]}
			}
			rendeRule(
				&content,
				i,
				rule.Name,
			)
			for _, approver := range rule.ApprovedBy.Nodes {
				renderIndentedText(
					&content,
					styledIcon{icon: icon.User, color: style.White},
					approver.Name,
				)
			}
		} else {
			rendeRule(
				&content,
				styledIcon{icon: icon.CircleCross, color: style.Red[400]},
				rule.Name,
			)
		}
	}

	return contentStyle.Render(content.String())
}

func renderPipelines(stages []gitlab.CiStageNode) string {
	var content strings.Builder

	content.WriteString(sectionTitleStyle.Render(fmt.Sprintf("%s Pipeline", icon.Pipeline)))
	content.WriteString("\n\n")

	for _, stage := range stages {

		stageStatus := getStageIconStatus(stage.Status)
		content.WriteString(iconStyle(stageStatus.color).Render(stageStatus.icon))
		content.WriteString(sectionTextStyle.Render(stage.Name))
		content.WriteString("\n")

		lower := strings.ToLower
		if lower(stage.Status) != "success" || lower(stage.Status) != "manual" {
			for _, node := range stage.Jobs.Nodes {
				if lower(node.Status) != "success" {
					nodeStatus := getStageIconStatus(node.Status)
					renderIndentedText(
						&content,
						styledIcon{icon: nodeStatus.icon, color: nodeStatus.color},
						node.Name,
					)
				}
			}
		}
	}
	return contentStyle.Render(content.String())
}

func renderDiscussions(discussions []gitlab.DiscussionNode, m Model) string {
	var bdy, title, content strings.Builder
	separator := strings.Repeat("-", 5)

	title.WriteString(fmt.Sprintf("%s Discussions", icon.Discussion))
	title.WriteString("\n\n")

	hasDiscussions := slices.ContainsFunc(discussions, func(d gitlab.DiscussionNode) bool {
		return d.Resolvable
	})

	if !hasDiscussions {
		bdy.WriteString("... *No Discussions*")
	} else {
		for _, discussion := range discussions {
			if !discussion.Resolvable {
				continue
			}

			bdy.WriteString(separator)
			if discussion.Resolved {
				resolvedAt := table.FormatTime(discussion.ResolvedAt)
				bdy.WriteString(fmt.Sprintf(" **%s %s** ", icon.Check, timeAgo(resolvedAt)))
			} else {
				bdy.WriteString(fmt.Sprintf(" %s ", icon.Dash))
			}
			bdy.WriteString(separator)
			bdy.WriteString("\n\n")

			for _, note := range discussion.Notes.Nodes {
				author := note.Author.Name
				body := note.Body
				createdAt := table.FormatTime(note.CreatedAt)

				if !note.Resolvable {
					before, _, found := strings.Cut(body, "(")
					if found {
						bdy.WriteString(
							fmt.Sprintf(
								"*%s %s %s %s* ",
								icon.Dot,
								author,
								before,
								timeAgo(createdAt),
							),
						)
						bdy.WriteString("\n\n")
					}
					continue
				}

				bdy.WriteString(fmt.Sprintf("`%s` ", author))
				bdy.WriteString(timeAgo(createdAt))
				bdy.WriteString("\n")

				bdy.WriteString(body)
				bdy.WriteString("\n\n")

			}
			bdy.WriteString("\n\n")

		}
	}

	content.WriteString(sectionTitleStyle.Render(title.String()))
	content.WriteString(sectionTextStyle.Render(m.renderWithStyle(bdy.String())))

	return contentStyle.Render(content.String())
}

func (m Model) renderWithStyle(s string) string {
	d, err := glamourRender(m, s)
	if err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)

		return ""
	}
	return d
}

func getMdRenderer(m Model) *glamour.TermRenderer {
	magicnumber := 4 // FIX: find where this comes from
	width := m.Viewport.Width - magicnumber - LeftMargin
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(width),
		glamour.WithEmoji(),
		glamour.WithPreservedNewLines(),
	)
	if err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)

		return nil
	}

	return r
}

func glamourRender(m Model, markdown string) (string, error) {
	r := getMdRenderer(m)
	out, err := r.Render(markdown)
	if err != nil {
		return "", err
	}

	// trim lines
	lines := strings.Split(out, "\n")

	var content string
	for i, s := range lines {
		content += strings.TrimSpace(s)

		// don't add an artificial newline after the last split
		if i+1 < len(lines) {
			content += "\n"
		}
	}

	return content, nil
}

func getStageIconStatus(s string) styledIcon {
	icons := map[string]styledIcon{
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

	return styledIcon{icon: icon.Dash}
}

func timeAgo(time string) string {
	return fmt.Sprintf("_%s ago_", time)
}
