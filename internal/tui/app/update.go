package app

import (
	"errors"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/message"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/table"
	"github.com/felipeospina21/mrglab/internal/tui/task"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	isLeftPanelFocused := m.ctx.FocusedPanel == context.LeftPanel
	isMainPanelFocused := m.ctx.FocusedPanel == context.MainPanel
	isRightPanelFocused := m.ctx.FocusedPanel == context.RightPanel

	switch msg := msg.(type) {

	case error:
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(msg.Error())

	case tea.KeyMsg:
		match := keyMatcher(msg)
		gk := tui.GlobalKeys
		lpk := projects.Keybinds
		mpk := mergerequests.Keybinds
		rpk := details.Keybinds

		switch {
		case match(gk.ThrowError):
			cmds = append(cmds, func() tea.Msg {
				return errors.New("mocked")
			})

		case match(gk.Quit):
			return m, tea.Quit

		case match(gk.ToggleLeftPanel):
			m.toggleLeftPanel()
			if m.ctx.IsLeftPanelOpen {
				m.Projects.SetFocus()
			} else {
				m.MergeRequests.SetFocus()
			}
		}

		if isLeftPanelFocused {
			m.Projects.List, cmd = m.Projects.List.Update(msg)
			switch {
			case match(lpk.MRList):
				cb := func() tea.Cmd {
					m.Projects.SelectProject()
					return m.Projects.GetListCmd()
				}
				cmds = append(cmds, m.startTask(cb))
			}
		}

		if isMainPanelFocused {
			m.MergeRequests.Table, cmd = m.MergeRequests.Table.Update(msg)
			switch {
			case match(mpk.Details):
				resizeCmd := m.Details.SetViewportViewSize(
					tea.WindowSizeMsg{Width: m.getViewportViewWidth(), Height: m.ctx.PanelHeight},
				)

				mr := func() tea.Cmd {
					m.SelectMR()
					return m.MergeRequests.FetchMergeRequest()
				}

				cmds = append(cmds,
					resizeCmd,
					m.startTask(mr),
				)

			case match(mpk.Merge):
				merge := func() tea.Cmd {
					m.SelectMR()
					return m.MergeRequests.AcceptMergeRequest()
				}
				cmds = append(cmds, m.startTask(merge))
			}
		}

		if isRightPanelFocused {
			m.Details.Viewport, cmd = m.Details.Viewport.Update(msg)
			switch {
			case match(rpk.ClosePanel):
				m.toggleRightPanel()
				m.MergeRequests.SetFocus()

			case match(rpk.Merge):
				merge := func() tea.Cmd {
					m.SelectMR()
					return m.MergeRequests.AcceptMergeRequest()
				}
				cmds = append(cmds, m.startTask(merge))
			}
		}

	case spinner.TickMsg:
		cmd = m.updateSpinnerViewCommand(msg)
		cmds = append(cmds, cmd)

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
					m.GetMergeRequestModel(msg),
				)

				if msg.Err == nil {
					m.toggleLeftPanel()
					m.MergeRequests.SetFocus()
					m.MergeRequests.Table = t
				}

			}

			if msg.TaskID == task.FetchDiscussions {
				mr := finishTask[details.MergeRequestDetails](
					&m,
					msg,
					m.GetMergeRequestDetails(msg),
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

				m.toggleRightPanel()
				m.Details.SetFocus()

			}

			if msg.TaskID == task.MergeMR {
				finishTask[any](
					&m,
					msg,
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
					cmds = append(cmds, cmd)
				}
			}
		}
	}

	m.Modal, cmd = m.Modal.Update(msg)
	return m, tea.Batch(cmds...)
}

func keyMatcher(msg tea.KeyMsg) func(key.Binding) bool {
	return func(k key.Binding) bool {
		return key.Matches(msg, k)
	}
}
