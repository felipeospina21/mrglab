package app

import (
	"errors"
	"fmt"
	"strings"

	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"

	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/felipeospina21/mrglab/internal/exec"
	execPkg "github.com/felipeospina21/mrglab/internal/exec"
	"github.com/felipeospina21/mrglab/internal/gitlab"
	"github.com/felipeospina21/mrglab/internal/logger"
	"github.com/felipeospina21/mrglab/internal/tui"
	"github.com/felipeospina21/mrglab/internal/tui/components/details"
	"github.com/felipeospina21/mrglab/internal/tui/components/loader"
	"github.com/felipeospina21/mrglab/internal/tui/components/mergerequests"
	"github.com/felipeospina21/mrglab/internal/tui/components/pipelines"
	"github.com/felipeospina21/mrglab/internal/tui/components/projects"
	"github.com/felipeospina21/tuishell"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case error:
		l, f := logger.New(logger.NewLogger{})
		defer f.Close()
		l.Error(msg.Error())

	// Component action messages from panels
	case projects.FetchMRListMsg, mergerequests.ReFetchMRListMsg:
		m.MergeRequests.Loading = true
		m.Pipelines.Loading = true
		cmds = append(cmds, func() tea.Msg {
			return tuishell.StartTaskMsg{Cmd: m.fetchMergeRequestsList()}
		})
		cmds = append(cmds, m.fetchPipelinesList())

	case mergerequests.ViewDetailsMsg:
		if !m.Shell.IsRightOpen() {
			cmds = append(cmds, func() tea.Msg { return tuishell.OpenRightPanelMsg{} })
		}
		m.Details.SetFocus()
		m.Details.Ready = false
		m.Details.Viewport.SetContent("")
		titleIdx := mergerequests.GetColIndex(mergerequests.ColNames.Title)
		m.Details.Content.Title = m.MergeRequests.Table.SelectedRow()[titleIdx]
		cmds = append(cmds, func() tea.Msg {
			return tuishell.StartTaskMsg{Cmd: m.fetchSingleMergeRequest()}
		})

	case mergerequests.MergeMRMsg, details.MergeMRMsg:
		cmds = append(cmds, func() tea.Msg {
			return tuishell.StartTaskMsg{Cmd: m.acceptMergeRequest()}
		})

	case mergerequests.OpenInBrowserMsg:
		m.openInBrowser()

	case details.OpenInBrowserMsg:
		if m.ActiveTab == 1 {
			pp := pipelines.GetColIndex(pipelines.ColNames.Path)
			url := m.Pipelines.Table.SelectedRow()[pp]
			exec.Openbrowser(fmt.Sprintf("%s%s", config.GlobalConfig.BaseURL, url))
		} else {
			m.openInBrowser()
		}

	case pipelines.ViewDetailsMsg:
		if !m.Shell.IsRightOpen() {
			cmds = append(cmds, func() tea.Msg { return tuishell.OpenRightPanelMsg{} })
		}
		m.Details.SetFocus()
		m.setHelpKeys(details.PipelineKeybinds)

		iidIdx := pipelines.GetColIndex(pipelines.ColNames.IID)
		iid := m.Pipelines.Table.SelectedRow()[iidIdx]
		m.Details.Content.Title = fmt.Sprintf("Pipeline #%s", iid)

		node := m.findPipelineByIID(iid)
		if node != nil {
			m.Details.PipelineNode = node
			m.Details.ActionableJobs = nil
			m.Details.ActionableJobIdx = 0
			for _, j := range node.Jobs.Nodes {
				if !j.Retried && strings.ToLower(j.Status) != "success" {
					m.Details.ActionableJobs = append(m.Details.ActionableJobs, j)
				}
			}
			selectedJob := ""
			if len(m.Details.ActionableJobs) > 0 {
				j := m.Details.ActionableJobs[0]
				selectedJob = j.Stage.Name + "/" + j.Name
			}
			c := details.RenderPipelineDetailsWithSelection(*node, selectedJob)
			m.Details.Viewport.SetContent(c)
			m.Details.Ready = true
		}

	case mergerequests.CycleTabMsg, pipelines.CycleTabMsg:
		m.ActiveTab = (m.ActiveTab + 1) % len(m.TabNames)
		if m.ActiveTab == 0 {
			m.setHelpKeys(mergerequests.Keybinds)
			m.Shell.Main = MergeRequestsPanel{Model: m.MergeRequests, ActiveTab: m.ActiveTab, TabNames: m.TabNames, ProjectName: m.ctx.SelectedProject.Name}
		} else {
			m.setHelpKeys(pipelines.Keybinds)
			m.Shell.Main = PipelinesPanel{Model: m.Pipelines, ActiveTab: m.ActiveTab, TabNames: m.TabNames, ProjectName: m.ctx.SelectedProject.Name}
		}
		l := m.Shell.Layout
		m.Shell.Main, cmd = m.Shell.Main.Update(tea.WindowSizeMsg{Width: l.MainPanel.Width, Height: l.MainPanel.Height})
		cmds = append(cmds, cmd)

	case mergerequests.CreateMRMsg:
		m.pendingCreateMR = true
		content := loader.View(m.Shell.Spinner.View())
		cmds = append(cmds,
			func() tea.Msg { return tuishell.OpenModalMsg{Header: "New Merge Request", Content: content} },
			m.MergeRequests.FetchMRTemplates(),
		)

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
			m.Shell.Modal.Content = m.createForm.View()
			cmds = append(cmds, m.createForm.Focus())
		}

	case details.ClosePanelMsg:
		m.MergeRequests.SetFocus()
		m.setHelpKeys(mergerequests.Keybinds)
		cmds = append(cmds, func() tea.Msg { return tuishell.CloseRightPanelMsg{} })

	case details.FullscreenMsg:
		cmds = append(cmds, func() tea.Msg { return tuishell.ToggleFullscreenMsg{} })

	case details.RespondCommentMsg:
		m.pendingNote.DiscussionId = msg.DiscussionId
		m.pendingNote.NoteableId = msg.NoteableId
		cmds = append(cmds,
			func() tea.Msg {
				return tuishell.OpenModalMsg{
					Header:  fmt.Sprintf("Responding to discussion (%s)", details.ShortID(msg.DiscussionId)),
					Content: m.Input.View(),
				}
			},
			m.Input.Focus(),
		)

	case tuishell.CloseModalMsg:
		m.Input.Blur()
		m.Input.Reset()
		m.createForm.Blur()
		m.createForm.Reset()
		m.pendingCreateMR = false
		m.pendingConfirm = false
		m.formReady = false
		if m.Shell.TaskErr() != nil {
			cmds = append(cmds, func() tea.Msg {
				return tuishell.SetStatusMsg{Content: ""}
			})
		}

	case tuishell.CopyModalMsg:
		execPkg.CopyToClipboard(m.Shell.Modal.Content)

	case tuishell.SubmitModalMsg, tuishell.ShellSubmitMsg:
		body := m.Input.Value()
		if body != "" && m.pendingNote.NoteableId != "" {
			cmds = append(cmds, func() tea.Msg {
				return tuishell.StartTaskMsg{Cmd: m.MergeRequests.CreateNote(gitlab.CreateNoteInput{
					NoteableId:   gitlab.NoteableID(m.pendingNote.NoteableId),
					DiscussionId: gitlab.DiscussionID(m.pendingNote.DiscussionId),
					Body:         body,
				})}
			})
		}
		if m.pendingCreateMR {
			cmds = append(cmds, func() tea.Msg {
				return tuishell.StartTaskMsg{Cmd: m.createMergeRequest()}
			})
			m.pendingCreateMR = false
		}
		m.Input.Blur()
		m.Input.Reset()
		m.createForm.Blur()
		m.createForm.Reset()
		m.formReady = false
		if m.Shell.IsRightOpen() {
			m.Details.SetFocus()
			m.setHelpKeys(details.Keybinds)
		} else {
			m.MergeRequests.SetFocus()
			m.setHelpKeys(mergerequests.Keybinds)
		}

	// Typed task result messages
	case tui.MRListFetchedMsg:
		m.MergeRequests.Loading = false
		var kb help.KeyMap = mergerequests.Keybinds
		if m.ActiveTab == 1 {
			kb = pipelines.Keybinds
		}
		m.setHelpKeys(kb)
		cmds = append(cmds, finishTaskCmd(msg.Err, kb))
		if msg.Err == nil {
			t := m.getMergeRequestModel(msg)()
			m.MergeRequests.Table = t
			m.MergeRequests.SetFocus()
			cmds = append(cmds, func() tea.Msg { return tuishell.CloseLeftPanelMsg{} })
			l := m.Shell.Layout
			m.Shell.Main, cmd = m.Shell.Main.Update(tea.WindowSizeMsg{Width: l.MainPanel.Width, Height: l.MainPanel.Height})
			cmds = append(cmds, cmd)
		}

	case tui.PipelineListFetchedMsg:
		m.Pipelines.Loading = false
		if msg.Err == nil {
			t := m.getPipelineModel(msg)()
			m.Pipelines.Table = t
			m.Pipelines.Nodes = msg.Pipelines.Nodes
		}

	case tui.MRDetailsFetchedMsg:
		cmds = append(cmds, finishTaskCmd(msg.Err, details.Keybinds))
		if msg.Err == nil {
			mr := m.getMergeRequestDetails(msg)()

			m.Details.MRId = msg.MRId
			m.Details.Discussions = msg.Discussions
			m.Details.DiscussionIdx = 0
			m.Details.MRDetails = mr
			m.Details.MRDescription = m.MergeRequests.Table.SelectedRow()[mergerequests.GetColIndex(mergerequests.ColNames.Description)]

			titleIdx := mergerequests.GetColIndex(mergerequests.ColNames.Title)
			m.Details.Content.Title = m.MergeRequests.Table.SelectedRow()[titleIdx]

			c := m.Details.GetViewportContent(m.Details.MRDescription, mr)
			m.Details.Viewport.SetContent(c)
			m.Details.Ready = true
		}

	case tui.MRMergedMsg:
		cmds = append(cmds, finishTaskCmd(msg.Err, mergerequests.Keybinds))
		if msg.Err == nil {
			if len(msg.Errors) > 0 {
				e := strings.Join(msg.Errors, ", ")
				cmds = append(cmds, func() tea.Msg { return errors.New(e) })
			} else if m.Shell.Ctx.FocusedPanel == tuishell.MainPanel {
				cmds = append(cmds, func() tea.Msg {
					return tuishell.StartTaskMsg{Cmd: m.fetchMergeRequestsList()}
				})
			} else if m.Shell.Ctx.FocusedPanel == tuishell.RightPanel {
				cmds = append(cmds, func() tea.Msg {
					return tuishell.StartTaskMsg{Cmd: m.fetchSingleMergeRequest()}
				})
			}
		}

	case tui.NoteCreatedMsg:
		cmds = append(cmds, finishTaskCmd(msg.Err, details.Keybinds))
		if msg.Err == nil {
			if len(msg.Errors) > 0 {
				e := strings.Join(msg.Errors, ", ")
				cmds = append(cmds, func() tea.Msg { return errors.New(e) })
			} else {
				cmds = append(cmds, func() tea.Msg {
					return tuishell.SetStatusMsg{Content: "✓ Comment sent"}
				})
			}
		}

	case tui.MRCreatedMsg:
		cmds = append(cmds, finishTaskCmd(msg.Err, mergerequests.Keybinds))
		if msg.Err == nil {
			if len(msg.Errors) > 0 {
				e := strings.Join(msg.Errors, ", ")
				cmds = append(cmds, func() tea.Msg { return errors.New(e) })
			} else {
				cmds = append(cmds,
					func() tea.Msg { return tuishell.SetStatusMsg{Content: "✓ MR created"} },
					func() tea.Msg { return tuishell.StartTaskMsg{Cmd: m.fetchMergeRequestsList()} },
				)
			}
		}

	case pipelines.OpenInBrowserMsg:
		pp := pipelines.GetColIndex(pipelines.ColNames.Path)
		url := m.Pipelines.Table.SelectedRow()[pp]
		exec.Openbrowser(fmt.Sprintf("%s%s", config.GlobalConfig.BaseURL, url))

	case pipelines.RetryPipelineMsg:
		iidIdx := pipelines.GetColIndex(pipelines.ColNames.IID)
		iid := m.Pipelines.Table.SelectedRow()[iidIdx]
		node := m.findPipelineByIID(iid)
		if node != nil {
			cmds = append(cmds, func() tea.Msg {
				return tuishell.StartTaskMsg{Cmd: m.Pipelines.RetryPipeline(node.ID)}
			})
		}

	case tui.PipelineRetryMsg:
		cmds = append(cmds, finishTaskCmd(msg.Err, pipelines.Keybinds))
		if msg.Err == nil {
			if len(msg.Errors) > 0 {
				e := strings.Join(msg.Errors, ", ")
				cmds = append(cmds, func() tea.Msg { return errors.New(e) })
			} else {
				cmds = append(cmds,
					func() tea.Msg { return tuishell.SetStatusMsg{Content: "✓ Pipeline retry triggered"} },
					m.fetchPipelinesList(),
				)
			}
		}

	case details.PlayJobMsg:
		if strings.ToLower(msg.Status) == "manual" {
			cmds = append(cmds, func() tea.Msg {
				return tuishell.StartTaskMsg{Cmd: m.Pipelines.PlayJob(msg.JobID)}
			})
		} else {
			cmds = append(cmds, func() tea.Msg {
				return tuishell.StartTaskMsg{Cmd: m.Pipelines.RetryJob(msg.JobID)}
			})
		}

	case tui.JobPlayMsg:
		cmds = append(cmds, finishTaskCmd(msg.Err, details.PipelineKeybinds))
		if msg.Err == nil {
			if len(msg.Errors) > 0 {
				e := strings.Join(msg.Errors, ", ")
				cmds = append(cmds, func() tea.Msg { return errors.New(e) })
			} else {
				cmds = append(cmds,
					func() tea.Msg { return tuishell.SetStatusMsg{Content: "✓ Job triggered"} },
					m.fetchPipelinesList(),
				)
			}
		}

	case tui.JobRetryMsg:
		cmds = append(cmds, finishTaskCmd(msg.Err, details.PipelineKeybinds))
		if msg.Err == nil {
			if len(msg.Errors) > 0 {
				e := strings.Join(msg.Errors, ", ")
				cmds = append(cmds, func() tea.Msg { return errors.New(e) })
			} else {
				cmds = append(cmds,
					func() tea.Msg { return tuishell.SetStatusMsg{Content: "✓ Job retriggered"} },
					m.fetchPipelinesList(),
				)
			}
		}
	}

	// Handle input focus
	if m.Input.Focused() {
		m.Input, cmd = m.Input.Update(msg)
		m.Shell.Modal.Content = m.Input.View()
		cmds = append(cmds, cmd)
	}

	// Handle form focus
	if m.pendingCreateMR && m.Shell.IsModalOpen() && m.formReady {
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
		m.Shell.Modal.Content = m.createForm.View()
	}

	// Propagate focus changes from components to shell before shell.Update
	m.Shell.Ctx.FocusedPanel = m.ctx.FocusedPanel

	// Delegate to shell for everything else
	m.Shell, cmd = m.Shell.Update(msg)
	cmds = append(cmds, cmd)

	// Sync shell context back to mrglab context after shell update
	m.ctx.AppContext = m.Shell.Ctx

	// Sync panel pointers after shell update
	if p, ok := m.Shell.Left.(ProjectsPanel); ok {
		m.Projects = p.Model
	}
	if mr, ok := m.Shell.Main.(MergeRequestsPanel); ok {
		m.MergeRequests = mr.Model
		mr.ActiveTab = m.ActiveTab
		mr.ProjectName = m.ctx.SelectedProject.Name
		m.Shell.Main = mr
	}
	if pip, ok := m.Shell.Main.(PipelinesPanel); ok {
		m.Pipelines = pip.Model
		pip.ActiveTab = m.ActiveTab
		pip.ProjectName = m.ctx.SelectedProject.Name
		m.Shell.Main = pip
	}
	if d, ok := m.Shell.Right.(DetailsPanel); ok {
		m.Details = d.Model
	}

	// Sync spinner view to components that render loading states
	sv := m.Shell.Spinner.View()
	m.Details.SpinnerView = sv
	m.MergeRequests.SpinnerView = sv
	m.Pipelines.SpinnerView = sv
	if m.pendingCreateMR && m.Shell.IsModalOpen() && !m.formReady {
		m.Shell.Modal.Content = loader.View(sv)
	}

	if m.pendingCreateMR && m.Shell.IsModalOpen() && m.formReady && m.pendingConfirm {
		m.Shell.Modal.Content = "Discard changes? (y/n)"
	}

	return m, tea.Batch(cmds...)
}

// finishTaskCmd returns a command that sends FinishTaskMsg to the shell.
func finishTaskCmd(err error, kb help.KeyMap) tea.Cmd {
	return func() tea.Msg {
		return tuishell.FinishTaskMsg{Err: err, Keybinds: kb}
	}
}
