package mergerequests

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipeospina21/mrglab/internal/api"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
	"github.com/xanzy/go-gitlab"
)

func (m *Model) GetMRListCmd() tea.Cmd {
	return func() tea.Msg {
		p := &gitlab.ListProjectMergeRequestsOptions{
			State: gitlab.Ptr("opened"),
		}

		mrs, err := api.GetProjectMergeRequests(m.ctx.SelectedProject.ID, p)

		return task.TaskFinishedMsg{
			TaskID:      "",
			SectionID:   0,
			SectionType: "mrs",
			Err:         err,
			Msg: MergeRequestsFetchedMsg{
				Mrs:    mrs,
				TaskId: "fetch_Mrs",
			},
		}
	}
}

func GetMRTableRows(msg task.TaskFinishedMsg) []table.Row {
	var rows []table.Row
	ml := msg.Msg.(MergeRequestsFetchedMsg)
	for _, mr := range ml.Mrs {
		r := table.Row{
			mr.CreatedAt.String(),
			strconv.FormatBool(mr.Draft),
			mr.Title,
			mr.Author.Name,
			mr.DetailedMergeStatus,
			strconv.FormatBool(mr.HasConflicts),
			strconv.Itoa(mr.UserNotesCount),
			mr.ChangesCount,
			mr.WebURL,
			mr.Description,
			strconv.Itoa(mr.IID),
		}

		rows = append(rows, r)
	}
	return rows
}
