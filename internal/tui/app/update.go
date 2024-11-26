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
	isLeftPanelFocused := m.ctx.FocusedPanel == context.LeftPanel
	isMainPanelFocused := m.ctx.FocusedPanel == context.MainPanel
	isRightPanelFocused := m.ctx.FocusedPanel == context.RightPanel
	isModalFocused := m.ctx.FocusedPanel == context.Modal

	switch msg := msg.(type) {

	case error:
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(msg.Error())

	case tea.KeyMsg:
		match := tui.KeyMatcher(msg)
		gk := tui.GlobalKeys
		lpk := projects.Keybinds
		mpk := mergerequests.Keybinds
		rpk := details.Keybinds

		switch {
		case match(gk.MockFetch):
			if m.ctx.Task.Status == task.TaskStarted {
				m.ctx.Task.Status = task.TaskFinished
			} else if m.ctx.Task.Status == task.TaskFinished || m.ctx.Task.Status == task.TaskIdle {
				m.ctx.Task.Status = task.TaskStarted
			}
		case match(gk.ThrowError):
			cmds = append(cmds, func() tea.Msg {
				return errors.New("mocked")
			})

		case match(gk.Quit):
			return m, tea.Quit

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

		if isModalFocused {
			if match(modal.Keybinds.Close) {
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
			}
		}

		if isLeftPanelFocused {
			m.Projects.List, cmd = m.Projects.List.Update(msg)
			switch {
			case match(lpk.MRList):
				cmds = append(cmds, m.startTask(m.fetchMergeRequestsList))
			}
		}

		if isMainPanelFocused {
			m.MergeRequests.Table, cmd = m.MergeRequests.Table.Update(msg)
			switch {
			case match(mpk.Details):
				resizeCmd := m.Details.SetViewportViewSize(
					tea.WindowSizeMsg{Width: m.getViewportViewWidth(), Height: m.ctx.PanelHeight},
				)

				cmds = append(cmds,
					resizeCmd,
					m.startTask(m.fetchSingleMergeRequest),
				)

			case match(mpk.Merge):
				cmds = append(cmds, m.startTask(m.acceptMergeRequest))

			case match(mpk.OpenInBrowser):
				m.openInBrowser()
			}
		}

		if isRightPanelFocused {
			m.Details.Viewport, cmd = m.Details.Viewport.Update(msg)
			switch {
			case match(rpk.ClosePanel):
				m.toggleRightPanel()
				m.MergeRequests.SetFocus()
				m.setHelpKeys(mergerequests.Keybinds)

			case match(rpk.Merge):
				cmds = append(cmds, m.startTask(m.acceptMergeRequest))

			case match(rpk.OpenInBrowser):
				m.openInBrowser()
			}
		}

	case spinner.TickMsg:
		var spin tea.Cmd
		cmd = m.updateSpinnerViewCommand(msg)
		m.Spinner, spin = m.Spinner.Update(msg)
		cmds = append(cmds, cmd, spin)

	case tea.WindowSizeMsg:
		// Sets window in context
		m.ctx.Window = msg

		m.setLeftPanelHeight()
		m.setStatuslineWidth()

	case task.TaskMsg:
		if msg.Err != nil {
			l, f := logger.New(logger.NewLogger{})
			defer f.Close()
			l.Error(msg.Err)

		}
		// TODO: Rethink this logic
		if msg.SectionType == task.TaskSectionMR {
			if msg.TaskID == task.FetchMRs {
				t := finishTask[table.Model](
					&m,
					msg,
					mergerequests.Keybinds,
					m.getMergeRequestModel(msg),
				)

				if msg.Err == nil {
					if m.ctx.IsLeftPanelOpen {
						m.toggleLeftPanel()
						m.MergeRequests.SetFocus()
					}
					m.MergeRequests.Table = t
				}

			}

			if msg.TaskID == task.FetchDiscussions {
				mr := finishTask[details.MergeRequestDetails](
					&m,
					msg,
					details.Keybinds,
					m.getMergeRequestDetails(msg),
				)

				// get title
				titleIdx := mergerequests.GetColIndex(mergerequests.ColNames.Title)
				t := m.MergeRequests.Table.SelectedRow()[titleIdx]
				m.Details.Content.Title = t

				// get description
				idx := mergerequests.GetColIndex(mergerequests.ColNames.Description)
				d := m.MergeRequests.Table.SelectedRow()[idx]
				c := m.Details.GetViewportContent(d, details.MergeRequestDetails(mr))
				m.Details.Viewport.SetContent(c)

				if !m.ctx.IsRightPanelOpen {
					m.toggleRightPanel()
					m.Details.SetFocus()
				}

			}

			if msg.TaskID == task.MergeMR {
				finishTask[any](
					&m,
					msg,
					mergerequests.Keybinds,
					func() any {
						// TODO: implement some kind of success notification
						return nil
					},
				)

				res := msg.Msg.(message.MergeRequestMergedMsg)
				if len(res.Errors) > 0 {
					// TODO: show errors in statusline
					e := strings.Join(res.Errors, ", ")
					cmd = func() tea.Msg {
						return errors.New(e)
					}
				} else {
					if m.ctx.FocusedPanel == context.MainPanel {
						cmd = m.startTask(m.fetchMergeRequestsList)
					}

					if m.ctx.FocusedPanel == context.RightPanel {
						cmd = m.startTask(m.fetchSingleMergeRequest)
					}
				}
				cmds = append(cmds, cmd)
			}
		}
	}

	return m, tea.Batch(cmds...)
}
