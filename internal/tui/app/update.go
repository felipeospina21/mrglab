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
				cmds = append(cmds, m.startCommand(cb))
			}
		}

		if isMainPanelFocused {
			m.MergeRequests.Table, cmd = m.MergeRequests.Table.Update(msg)
			switch {
			case match(mpk.Details):
				resizeCmd := m.Details.SetViewportViewSize(
					tea.WindowSizeMsg{Width: m.getViewportViewWidth(), Height: m.ctx.Window.Height},
				)

				discussions := func() tea.Cmd {
					m.SelectMRID()
					// TODO: set md header to mr title
					return m.MergeRequests.GetMRNotesCmd()
				}

				pipeline := func() tea.Cmd {
					m.SelectMRID()
					return m.MergeRequests.GetMRPipelineCmd()
				}

				cmds = append(cmds,
					resizeCmd,
					m.startCommand(discussions),
					m.startCommand(pipeline),
				)
			}
		}

		if isRightPanelFocused {
			m.Details.Viewport, cmd = m.Details.Viewport.Update(msg)
			switch {
			case match(rpk.ClosePanel):
				m.toggleRightPanel()
				m.MergeRequests.SetFocus()
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
		// TODO: Rethink this logic
		if msg.SectionType == task.TaskSectionMR {
			if msg.TaskID == task.FetchMRs {
				t := endCommand[table.Model](
					&m,
					msg,
					m.GetMergeRequestModel(msg),
				)

				m.toggleLeftPanel()
				m.MergeRequests.SetFocus()
				m.MergeRequests.Table = t
			}

			// if msg.TaskID == task.FetchDiscussions {
			// 	s := endCommand[string](
			// 		&m,
			// 		msg,
			// 		m.GetMergeRequestDiscussions(msg),
			// 	)
			//
			// 	idx := mergerequests.GetColIndex(mergerequests.ColNames.Description)
			// 	d := m.MergeRequests.Table.SelectedRow()[idx]
			//
			// 	// Merge Mr details with comments
			// 	var content strings.Builder
			// 	content.WriteString(d)
			// 	content.WriteString("\n\n")
			// 	content.WriteString(s)
			// 	m.Details.SetStyledContent(content.String())
			//
			// 	if !m.ctx.IsRightPanelOpen {
			// 		m.toggleRightPanel()
			// 		m.Details.SetFocus()
			// 	}
			//
			// }

			if msg.TaskID == task.FetchPipeline {
				s := endCommand[string](
					&m,
					msg,
					m.GetMergeRequestPipeline(msg),
				)

				var content strings.Builder
				c := m.Details.Content
				content.WriteString(string(c))
				content.WriteString("\n\n")
				content.WriteString(s)
				m.Details.SetStyledContent(content.String())
				if !m.ctx.IsRightPanelOpen {
					m.toggleRightPanel()
					m.Details.SetFocus()
				}
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func keyMatcher(msg tea.KeyMsg) func(key.Binding) bool {
	return func(k key.Binding) bool {
		return key.Matches(msg, k)
	}
}
