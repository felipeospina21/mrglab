package app

import (
	"errors"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/components/modal"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case error:
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(msg.Error())

	case tea.KeyMsg:
		cmd = m.handleGlobalKeys(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}

		switch m.ctx.FocusedPanel {
		case context.Modal:
			if !m.Input.Focused() {
				cmds = append(cmds, m.Input.Focus())
			}
			m.Modal, cmd = m.Modal.Update(msg)
			cmds = append(cmds, cmd)

		case context.LeftPanel:
			m.Projects, cmd = m.Projects.Update(msg)
			cmds = append(cmds, cmd)

		case context.MainPanel:
			m.MergeRequests, cmd = m.MergeRequests.Update(msg)
			cmds = append(cmds, cmd)

		case context.RightPanel:
			m.Details, cmd = m.Details.Update(msg)
			cmds = append(cmds, cmd)
		}

	// Component action messages
	case projects.FetchMRListMsg:
		cmds = append(cmds, m.startTask(m.fetchMergeRequestsList))

	case mergerequests.ViewDetailsMsg:
		cmds = append(cmds, m.startTask(m.fetchSingleMergeRequest))

	case mergerequests.MergeMRMsg, details.MergeMRMsg:
		cmds = append(cmds, m.startTask(m.acceptMergeRequest))

	case mergerequests.OpenInBrowserMsg, details.OpenInBrowserMsg:
		m.openInBrowser()

	case details.ClosePanelMsg:
		m.toggleRightPanel()
		m.MergeRequests.SetFocus()
		m.setHelpKeys(mergerequests.Keybinds)

	case details.RespondCommentMsg:
		m.ctx.IsModalOpen = true
		m.ctx.FocusedPanel = context.Modal
		m.Modal.Header = "Respond to thread"
		m.Modal.Content = m.Input.View()
		cmds = append(cmds, m.Input.Focus())

	case modal.CloseModalMsg:
		if m.ctx.Task.Err != nil {
			mode := statusline.ModesEnum.Normal
			if config.GlobalConfig.DevMode {
				mode = statusline.ModesEnum.Dev
			}
			m.setStatus(mode, "")
		}
		m.ctx.IsModalOpen = false
		if m.ctx.IsRightPanelOpen {
			m.Details.SetFocus()
			m.setHelpKeys(details.Keybinds)
		} else {
			m.MergeRequests.SetFocus()
			m.setHelpKeys(mergerequests.Keybinds)
		}

	case spinner.TickMsg:
		var spin tea.Cmd
		cmd = m.updateSpinnerViewCommand(msg)
		m.Spinner, spin = m.Spinner.Update(msg)
		cmds = append(cmds, cmd, spin)

	case tea.WindowSizeMsg:
		m.ctx.Window = msg
		m.recomputeLayout()

	case task.TaskMsg:
		cmd = m.handleTaskMsg(msg)
		cmds = append(cmds, cmd)
	}

	m.Input, cmd = m.Input.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) handleGlobalKeys(msg tea.KeyMsg) tea.Cmd {
	match := tui.KeyMatcher(msg)
	gk := tui.GlobalKeys()

	switch {
	case match(gk.MockFetch):
		if m.ctx.Task.Status == task.TaskStarted {
			m.ctx.Task.Status = task.TaskFinished
		} else if m.ctx.Task.Status == task.TaskFinished || m.ctx.Task.Status == task.TaskIdle {
			m.ctx.Task.Status = task.TaskStarted
		}

	case match(gk.ThrowError):
		return func() tea.Msg { return errors.New("mocked") }

	case match(gk.Quit):
		return tea.Quit

	case match(gk.OpenModal):
		m.ctx.IsModalOpen = true
		if m.ctx.Task.Err != nil {
			m.Modal.Header = "Error"
			m.Modal.Content = m.ctx.Task.Err.Error()
		}
		m.Modal.SetFocus()

	case match(gk.ToggleLeftPanel):
		m.toggleLeftPanel()
		if m.ctx.IsRightPanelOpen {
			m.toggleRightPanel()
		}
		if m.ctx.IsLeftPanelOpen {
			m.Projects.SetFocus()
			m.setHelpKeys(projects.Keybinds)
		} else {
			m.MergeRequests.SetFocus()
			m.setHelpKeys(mergerequests.Keybinds)
		}
	}

	return nil
}

func (m *Model) handleTaskMsg(msg task.TaskMsg) tea.Cmd {
	if msg.Err != nil {
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(msg.Err)
	}

	if msg.SectionType != task.TaskSectionMR {
		return nil
	}

	switch msg.TaskID {
	case task.FetchMRs:
		t := finishTask[table.Model](
			m, msg, mergerequests.Keybinds, m.getMergeRequestModel(msg),
		)
		if msg.Err == nil {
			if m.ctx.IsLeftPanelOpen {
				m.toggleLeftPanel()
				m.MergeRequests.SetFocus()
			}
			m.MergeRequests.Table = t
			m.recomputeLayout()
		}

	case task.FetchDiscussions:
		mr := finishTask[details.MergeRequestDetails](
			m, msg, details.Keybinds, m.getMergeRequestDetails(msg),
		)

		titleIdx := mergerequests.GetColIndex(mergerequests.ColNames.Title)
		m.Details.Content.Title = m.MergeRequests.Table.SelectedRow()[titleIdx]

		idx := mergerequests.GetColIndex(mergerequests.ColNames.Description)
		d := m.MergeRequests.Table.SelectedRow()[idx]

		rl := computeLayout(m.ctx.Window, false, true)
		m.Details.SetViewportViewSize(
			tea.WindowSizeMsg{Width: rl.RightPanel.Width, Height: rl.ContentH - details.PanelStyle.GetVerticalFrameSize() - tableViewOverhead},
		)

		c := m.Details.GetViewportContent(d, details.MergeRequestDetails(mr))
		m.Details.Viewport.SetContent(c)

		if !m.ctx.IsRightPanelOpen {
			m.toggleRightPanel()
			m.Details.SetFocus()
		}

	case task.MergeMR:
		finishTask[any](
			m, msg, mergerequests.Keybinds,
			func() any { return nil },
		)

		res := msg.Msg.(message.MergeRequestMergedMsg)
		if len(res.Errors) > 0 {
			e := strings.Join(res.Errors, ", ")
			return func() tea.Msg { return errors.New(e) }
		}
		if m.ctx.FocusedPanel == context.MainPanel {
			return m.startTask(m.fetchMergeRequestsList)
		}
		if m.ctx.FocusedPanel == context.RightPanel {
			return m.startTask(m.fetchSingleMergeRequest)
		}
	}

	return nil
}
