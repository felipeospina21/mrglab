package details

import (
	"fmt"
	"slices"
	"strings"

	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
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

type viewportContent struct {
	content                string
	selectedDiscLineOffset int
}

func (m Model) getViewportContent(b string, mr MergeRequestDetails) viewportContent {
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
	content.WriteString(renderPipelines(mr.Pipelines, ""))
	content.WriteString("\n\n")
	content.WriteString(renderApprovals(mr.Approvals))
	content.WriteString("\n\n")

	linesBefore := strings.Count(content.String(), "\n")
	disc := renderDiscussions(mr.Discussions, m)
	content.WriteString(disc.content)

	return viewportContent{
		content:                content.String(),
		selectedDiscLineOffset: linesBefore + disc.selectedLineOffset,
	}
}

// GetViewportContent renders the full details viewport content string.
func (m Model) GetViewportContent(b string, mr MergeRequestDetails) string {
	return m.getViewportContent(b, mr).content
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
	content.WriteString(fmt.Sprintf(" %s ", icon.ArrowLeft))
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

func renderPipelines(stages []gitlab.CiStageNode, selectedJob string) string {
	var content strings.Builder

	content.WriteString(sectionTitleStyle.Render(fmt.Sprintf("%s Pipeline Jobs", icon.Pipeline)))
	content.WriteString("\n\n")

	for _, stage := range stages {

		stageStatus := getStageIconStatus(stage.Status)
		content.WriteString(iconStyle(stageStatus.color).Render(stageStatus.icon))
		content.WriteString(sectionTextStyle.Render(stage.Name))
		content.WriteString("\n")

		lower := strings.ToLower
		if lower(stage.Status) != "success" || lower(stage.Status) != "manual" {
			for _, node := range stage.Jobs.Nodes {
				jobKey := stage.Name + "/" + node.Name
				isSelected := selectedJob != "" && jobKey == selectedJob
				if lower(node.Status) != "success" {
					nodeStatus := getStageIconStatus(node.Status)
					if isSelected {
						text := iconStyle(nodeStatus.color).MarginLeft(0).Render(nodeStatus.icon) + sectionTextStyle.Render(node.Name)
						content.WriteString(sectionIndentedTextStyle.Render("└ "))
						content.WriteString(selectedDiscussionStyle.Render(text))
						content.WriteString("\n")
					} else {
						renderIndentedText(
							&content,
							styledIcon{icon: nodeStatus.icon, color: nodeStatus.color},
							node.Name,
						)
					}
				}
			}
		}
	}
	return contentStyle.Render(content.String())
}

type discussionsResult struct {
	content            string
	selectedLineOffset int
}

func renderDiscussions(discussions []gitlab.DiscussionNode, m Model) discussionsResult {
	var title, content strings.Builder
	separator := strings.Repeat("─", 5)
	selectedLineOffset := 0

	title.WriteString(fmt.Sprintf("%s Discussions", icon.Discussion))
	title.WriteString("\n\n")

	hasDiscussions := slices.ContainsFunc(discussions, func(d gitlab.DiscussionNode) bool {
		return d.Resolvable
	})

	if !hasDiscussions {
		content.WriteString(sectionTitleStyle.Render(title.String()))
		content.WriteString(sectionTextStyle.Render(m.renderWithStyle("... *No Discussions*")))
		return discussionsResult{content: contentStyle.Render(content.String())}
	}

	content.WriteString(sectionTitleStyle.Render(title.String()))
	content.WriteString("\n")

	borderWidth := selectedDiscussionStyle.GetHorizontalFrameSize()

	resolvableIdx := 0
	for _, discussion := range discussions {
		if !discussion.Resolvable {
			continue
		}

		selected := resolvableIdx == m.DiscussionIdx

		if selected {
			selectedLineOffset = strings.Count(content.String(), "\n")
		}

		var bdy strings.Builder

		var header strings.Builder
		header.WriteString(separator)
		if discussion.Resolved {
			resolvedAt := table.FormatTime(discussion.ResolvedAt)
			header.WriteString(fmt.Sprintf(" %s Closed %s (%s)", icon.Check, timeAgo(resolvedAt), ShortID(discussion.Id)))
		} else {
			header.WriteString(fmt.Sprintf(" %s Open (%s)", icon.CircleDot, ShortID(discussion.Id)))
		}
		header.WriteString(separator)

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

		headerLine := sectionTextStyle.MarginLeft(0).Render(header.String())

		var rendered string
		if selected {
			body := sectionTextStyle.Render(m.renderWithWidth(bdy.String(), m.mdWidth()-borderWidth))
			rendered = headerLine + "\n" + selectedDiscussionStyle.Render(body)
		} else {
			body := sectionTextStyle.Render(m.renderWithStyle(bdy.String()))
			rendered = headerLine + "\n" + body
		}

		content.WriteString(rendered)
		content.WriteString("\n")

		resolvableIdx++
	}

	return discussionsResult{
		content:            contentStyle.Render(content.String()),
		selectedLineOffset: selectedLineOffset,
	}
}

func (m Model) renderWithStyle(s string) string {
	return m.renderWithWidth(s, m.mdWidth())
}

func (m Model) renderWithWidth(s string, width int) string {
	d, err := glamourRender(s, width)
	if err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(err)

		return ""
	}
	return d
}

func (m Model) mdWidth() int {
	magicnumber := 4 // FIX: find where this comes from
	return m.Viewport.Width() - magicnumber - LeftMargin
}

func getMdRenderer(width int) *glamour.TermRenderer {
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

func glamourRender(markdown string, width int) (string, error) {
	r := getMdRenderer(width)
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

// RenderPipelineDetails renders the details view for a pipeline, composing sections.
func RenderPipelineDetails(pipeline gitlab.PipelineNode) string {
	return RenderPipelineDetailsWithSelection(pipeline, "")
}

// RenderPipelineDetailsWithSelection renders pipeline details with an optional selected manual job highlight.
func RenderPipelineDetailsWithSelection(pipeline gitlab.PipelineNode, selectedJob string) string {
	var content strings.Builder

	content.WriteString(renderPipelineInfo(pipeline))
	content.WriteString("\n\n")
	content.WriteString(renderPipelines(pipelineJobsToStages(pipeline.Jobs.Nodes), selectedJob))

	return content.String()
}

func renderPipelineInfo(pipeline gitlab.PipelineNode) string {
	var content strings.Builder

	statusIcon := getStageIconStatus(pipeline.Status)
	content.WriteString(iconStyle(statusIcon.color).Render(statusIcon.icon))
	content.WriteString(sectionTextStyle.Render(fmt.Sprintf("Status: %s", pipeline.Status)))
	content.WriteString("\n")

	content.WriteString(iconStyle(style.White).Render(icon.Start))
	content.WriteString(sectionTextStyle.Render(fmt.Sprintf("Source: %s", pipeline.Source)))
	content.WriteString("\n")

	content.WriteString(iconStyle(style.White).Render(icon.PR))
	content.WriteString(sectionTextStyle.Render(fmt.Sprintf("%s %s", pipeline.Commit.ShortId, pipeline.Commit.Title)))
	content.WriteString("\n")

	if pipeline.MergeRequest != nil {
		content.WriteString(iconStyle(style.White).Render(icon.SourceBranch))
		content.WriteString(sectionTextStyle.Render(pipeline.MergeRequest.SourceBranch))
		content.WriteString("\n")
	}

	if pipeline.Duration != nil {
		content.WriteString(iconStyle(style.White).Render(icon.StopWatch))
		content.WriteString(sectionTextStyle.Render(fmt.Sprintf("Duration: %ds", *pipeline.Duration)))
		content.WriteString("\n")
	}

	return contentStyle.Render(content.String())
}

// pipelineJobsToStages converts flat PipelineJobNodes into CiStageNodes for reuse with renderPipelines.
func pipelineJobsToStages(jobs []gitlab.PipelineJobNode) []gitlab.CiStageNode {
	var stages []gitlab.CiStageNode
	idx := map[string]int{}
	for _, job := range jobs {
		sn := job.Stage.Name
		jn := gitlab.JobsNode{Name: job.Name, Status: job.Status}
		if i, ok := idx[sn]; ok {
			stages[i].Jobs.Nodes = append(stages[i].Jobs.Nodes, jn)
		} else {
			idx[sn] = len(stages)
			stages = append(stages, gitlab.CiStageNode{
				Name:   sn,
				Status: deriveStageStatus(jobs, sn),
				Jobs:   gitlab.JobsConnection{Nodes: []gitlab.JobsNode{jn}},
			})
		}
	}
	return stages
}

// deriveStageStatus determines the overall stage status from jobs in a given stage.
func deriveStageStatus(jobs []gitlab.PipelineJobNode, stageName string) string {
	hasRunning, hasFailed, hasPending := false, false, false
	for _, j := range jobs {
		if j.Stage.Name != stageName {
			continue
		}
		switch strings.ToLower(j.Status) {
		case "failed":
			hasFailed = true
		case "running":
			hasRunning = true
		case "pending", "created":
			hasPending = true
		}
	}
	switch {
	case hasFailed:
		return "failed"
	case hasRunning:
		return "running"
	case hasPending:
		return "pending"
	default:
		return "success"
	}
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
	return fmt.Sprintf("%s ago", time)
}

// ShortID extracts the last 6 characters of a GitLab resource ID.
func ShortID(id string) string {
	if i := strings.LastIndex(id, "/"); i >= 0 {
		id = id[i+1:]
	}
	if len(id) <= 6 {
		return id
	}
	return id[len(id)-6:]
}
