package app

import (
	"errors"
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"

	"charm.land/bubbles/v2/spinner"
	"github.com/felipeospina21/mrglab/internal/context"
	execPkg "github.com/felipeospina21/mrglab/internal/exec"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/loader"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/modal"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/mrglab/internal/tui/components/statusline"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case error:
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(msg.Error())

	case tea.KeyPressMsg:
		cmd = m.handleGlobalKeys(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}

		switch m.ctx.FocusedPanel {
		case context.Modal:
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
		if !m.isRightOpen {
			m.toggleRightPanel()
		}
		m.Details.SetFocus()
		m.Details.Ready = false
		m.Details.Viewport.SetContent("")
		titleIdx := mergerequests.GetColIndex(mergerequests.ColNames.Title)
		m.Details.Content.Title = m.MergeRequests.Table.SelectedRow()[titleIdx]
		cmds = append(cmds, m.startTask(m.fetchSingleMergeRequest))

	case mergerequests.MergeMRMsg, details.MergeMRMsg:
		cmds = append(cmds, m.startTask(m.acceptMergeRequest))

	case mergerequests.OpenInBrowserMsg, details.OpenInBrowserMsg:
		m.openInBrowser()

	case mergerequests.CreateMRMsg:
		m.isModalOpen = true
		m.pendingCreateMR = true
		m.ctx.FocusedPanel = context.Modal
		m.Modal.Header = "New Merge Request"
		m.Modal.FooterKeys = modal.CreateMRKeybinds
		m.Modal.HasSubmit = true
		m.Modal.Content = loader.View(m.Spinner.View())
		cmds = append(cmds, m.MergeRequests.FetchMRTemplates())

	case tui.MRTemplatesFetchedMsg:
		if m.pendingCreateMR {
			m.createForm.SetSize(modalContentWidth(m.ctx.Window.Width), modalContentHeight(m.ctx.Window.Height))
			if msg.Err == nil {
				if msg.SourceBranch != "" {
					m.createForm.source.SetValue(msg.SourceBranch)
				}
				if msg.DefaultBranch != "" {
					m.createForm.target.SetValue(msg.DefaultBranch)
				}
				if len(msg.Templates) > 0 {
					m.createForm.description.SetValue(msg.Templates[0].Content)
				}
			}
			m.formReady = true
			m.Modal.Content = m.createForm.View()
			cmds = append(cmds, m.createForm.Focus())
		}

	case details.ClosePanelMsg:
		m.isRightFullscreen = false
		m.toggleRightPanel()
		m.MergeRequests.SetFocus()
		m.setHelpKeys(mergerequests.Keybinds)

	case details.FullscreenMsg:
		m.isRightFullscreen = !m.isRightFullscreen
		m.recomputeLayout()

	case details.RespondCommentMsg:
		m.isModalOpen = true
		m.ctx.FocusedPanel = context.Modal
		m.Modal.Header = fmt.Sprintf("Responding to discussion (%s)", details.ShortID(msg.DiscussionId))
		m.Modal.FooterKeys = modal.MutationKeybinds
		m.Modal.HasSubmit = true
		m.Modal.Content = m.Input.View()
		m.pendingNote.DiscussionId = msg.DiscussionId
		m.pendingNote.NoteableId = msg.NoteableId
		cmds = append(cmds, m.Input.Focus())

	case modal.CloseModalMsg:
		m.Input.Blur()
		m.Input.Reset()
		m.createForm.Blur()
		m.createForm.Reset()
		m.Modal.IsError = false
		m.Modal.FooterKeys = modal.Keybinds
		m.Modal.HasSubmit = false
		m.pendingCreateMR = false
		m.formReady = false
		if m.taskErr != nil {
			mode := statusline.ModesEnum.Normal
			if m.ctx.DevMode {
				mode = statusline.ModesEnum.Dev
			}
			m.setStatus(mode, "")
		}
		m.isModalOpen = false
		switch {
		case m.isRightOpen:
			m.Details.SetFocus()
			m.setHelpKeys(details.Keybinds)
		case m.isLeftOpen:
			m.Projects.SetFocus()
			m.setHelpKeys(projects.Keybinds)
		default:
			m.MergeRequests.SetFocus()
			m.setHelpKeys(mergerequests.Keybinds)
		}

	case modal.CopyModalMsg:
		execPkg.CopyToClipboard(m.Modal.Content)

	case modal.ResetHighlightMsg:
		m.Modal.Highlight = false

	case modal.SubmitModalMsg:
		body := m.Input.Value()
		if body != "" && m.pendingNote.NoteableId != "" {
			cmds = append(cmds, m.startTask(func() tea.Cmd {
				return m.MergeRequests.CreateNote(gitlab.CreateNoteInput{
					NoteableId:   gitlab.NoteableID(m.pendingNote.NoteableId),
					DiscussionId: gitlab.DiscussionID(m.pendingNote.DiscussionId),
					Body:         body,
				})
			}))
		}
		if m.pendingCreateMR {
			createCmd := m.createMergeRequest()
			cmds = append(cmds, m.startTask(func() tea.Cmd {
				return createCmd
			}))
			m.pendingCreateMR = false
		}
		m.Input.Blur()
		m.Input.Reset()
		m.createForm.Blur()
		m.createForm.Reset()
		m.formReady = false
		m.isModalOpen = false
		if m.isRightOpen {
			m.Details.SetFocus()
			m.setHelpKeys(details.Keybinds)
		} else {
			m.MergeRequests.SetFocus()
			m.setHelpKeys(mergerequests.Keybinds)
		}

	// Typed task result messages
	case tui.MRListFetchedMsg:
		m.finishTask(msg.Err, mergerequests.Keybinds)
		if msg.Err == nil {
			t := m.getMergeRequestModel(msg)()
			if m.isLeftOpen {
				m.toggleLeftPanel()
				m.MergeRequests.SetFocus()
			}
			m.MergeRequests.Table = t
			m.recomputeLayout()
		}

	case tui.MRDetailsFetchedMsg:
		m.finishTask(msg.Err, details.Keybinds)
		if msg.Err == nil {
			mr := m.getMergeRequestDetails(msg)()

			m.Details.MRId = msg.MRId
			m.Details.Discussions = msg.Discussions
			m.Details.DiscussionIdx = 0
			m.Details.MRDetails = mr
			m.Details.MRDescription = m.MergeRequests.Table.SelectedRow()[mergerequests.GetColIndex(mergerequests.ColNames.Description)]

			titleIdx := mergerequests.GetColIndex(mergerequests.ColNames.Title)
			m.Details.Content.Title = m.MergeRequests.Table.SelectedRow()[titleIdx]

			rl := computeLayout(m.ctx.Window, false, true, m.isRightFullscreen)
			m.Details.SetViewportViewSize(
				tea.WindowSizeMsg{Width: rl.RightPanel.Width, Height: rl.ContentH - details.PanelStyle.GetVerticalFrameSize() - tableViewOverhead},
			)

			c := m.Details.GetViewportContent(m.Details.MRDescription, mr)
			m.Details.Viewport.SetContent(c)
			m.Details.Ready = true
		}

	case tui.MRMergedMsg:
		m.finishTask(msg.Err, mergerequests.Keybinds)
		if msg.Err == nil {
			if len(msg.Errors) > 0 {
				e := strings.Join(msg.Errors, ", ")
				cmds = append(cmds, func() tea.Msg { return errors.New(e) })
			} else if m.ctx.FocusedPanel == context.MainPanel {
				cmds = append(cmds, m.startTask(m.fetchMergeRequestsList))
			} else if m.ctx.FocusedPanel == context.RightPanel {
				cmds = append(cmds, m.startTask(m.fetchSingleMergeRequest))
			}
		}

	case tui.NoteCreatedMsg:
		m.finishTask(msg.Err, details.Keybinds)
		if msg.Err == nil {
			if len(msg.Errors) > 0 {
				e := strings.Join(msg.Errors, ", ")
				cmds = append(cmds, func() tea.Msg { return errors.New(e) })
			} else {
				m.Statusline.Content = "✓ Comment sent"
			}
		}

	case tui.MRCreatedMsg:
		m.finishTask(msg.Err, mergerequests.Keybinds)
		if msg.Err == nil {
			if len(msg.Errors) > 0 {
				e := strings.Join(msg.Errors, ", ")
				cmds = append(cmds, func() tea.Msg { return errors.New(e) })
			} else {
				m.Statusline.Content = "✓ MR created"
				cmds = append(cmds, m.startTask(m.fetchMergeRequestsList))
			}
		}

	case spinner.TickMsg:
		var spin tea.Cmd
		cmd = m.updateSpinnerViewCommand(msg)
		m.Spinner, spin = m.Spinner.Update(msg)
		m.Details.SpinnerView = m.Spinner.View()
		if m.pendingCreateMR && m.isModalOpen && !m.formReady {
			m.Modal.Content = loader.View(m.Spinner.View())
		}
		cmds = append(cmds, cmd, spin)

	case tea.WindowSizeMsg:
		m.ctx.Window = msg
		m.recomputeLayout()
	}

	if m.Input.Focused() {
		m.Input, cmd = m.Input.Update(msg)
		m.Modal.Content = m.Input.View()
		cmds = append(cmds, cmd)
	}

	if m.pendingCreateMR && m.isModalOpen && m.formReady {
		if keyMsg, ok := msg.(tea.KeyPressMsg); ok {
			switch keyMsg.String() {
			case "tab":
				cmds = append(cmds, m.createForm.NextField())
			case "shift+tab":
				cmds = append(cmds, m.createForm.PrevField())
			case "ctrl+d":
				m.createForm.draft = !m.createForm.draft
			default:
				cmd = m.createForm.Update(msg)
				cmds = append(cmds, cmd)
			}
		} else {
			cmd = m.createForm.Update(msg)
			cmds = append(cmds, cmd)
		}
		m.Modal.Content = m.createForm.View()
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) handleGlobalKeys(msg tea.KeyPressMsg) tea.Cmd {
	match := tui.KeyMatcher(msg)
	gk := tui.GlobalKeys(m.ctx.DevMode)

	switch {
	case match(gk.MockFetch):
		if m.taskStatus == taskStarted {
			m.taskStatus = taskFinished
		} else if m.taskStatus == taskFinished || m.taskStatus == taskIdle {
			m.taskStatus = taskStarted
		}

	case match(gk.ThrowError):
		txt := "# github.com/felipeospina21/mrglab/internal/tui/app internal/tui/app/update.go:246:12: m.Modal.IsEror undefined (type modal.Model has no field or method IsEror)"
		m.finishTask(errors.New(txt), mergerequests.Keybinds)
		return nil

	case match(gk.Quit):
		return tea.Quit

	case match(gk.OpenModal):
		if m.taskErr != nil {
			m.isModalOpen = true
			m.Modal.Header = "Error"
			m.Modal.IsError = true
			m.Modal.HasSubmit = false
			m.Modal.FooterKeys = modal.Keybinds
			m.Modal.Content = m.taskErr.Error()
			m.Modal.SetFocus()
		}

	case match(gk.Help):
		m.isModalOpen = true
		m.Modal.Header = "Keybindings"
		m.Modal.HasSubmit = false
		m.Modal.FooterKeys = modal.Keybinds
		m.Modal.Content = m.Modal.RenderHelp(m.Statusline.Keybinds)
		m.Modal.SetFocus()

	case match(gk.ToggleLeftPanel):
		m.toggleLeftPanel()
		if m.isRightOpen {
			m.toggleRightPanel()
		}
		if m.isLeftOpen {
			m.Projects.SetFocus()
			m.setHelpKeys(projects.Keybinds)
		} else {
			m.MergeRequests.SetFocus()
			m.setHelpKeys(mergerequests.Keybinds)
		}
	}

	return nil
}
