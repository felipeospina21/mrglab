// Package modal implements a centered overlay modal component.
package modal

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/felipeospina21/mrglab/internal/context"
	"github.com/felipeospina21/mrglab/internal/tui/components/help"
	"github.com/felipeospina21/mrglab/internal/tui/style"
)

// Model holds the state for the modal overlay.
type Model struct {
	Header    string
	Content   string
	Highlight bool
	IsError   bool
	ctx       *context.AppContext
}

// New creates a new modal model.
func New(ctx *context.AppContext) Model {
	return Model{
		ctx: ctx,
	}
}

// View renders the modal box centered over the dimmed background.
func (m Model) View(background string) string {
	w := m.ctx.Window.Width
	h := m.ctx.Window.Height

	modalW := modalSize(w)
	modalH := modalSize(h)

	header := headerStyle.Width(modalW).Render(m.Header)
	if m.IsError {
		header = headerStyle.Background(lipgloss.Color(style.Red[600])).Width(modalW).Render(m.Header)
	}
	footer := helpStyle.Render("esc close · ctrl+s submit · ctrl+y copy")
	contentH := max(modalH-lipgloss.Height(header)-lipgloss.Height(footer)-boxStyle.GetVerticalFrameSize(), 1)

	contentW := modalW - boxStyle.GetHorizontalFrameSize()
	content := m.Content
	if m.Highlight {
		content = lipgloss.NewStyle().
			Background(lipgloss.Color(style.Violet[400])).
			Foreground(lipgloss.Color(style.Black)).
			Render(content)
	}

	body := lipgloss.NewStyle().
		Width(contentW).
		Height(contentH).
		MaxHeight(contentH).
		Render(content)

	box := boxStyle.Width(modalW).Render(
		lipgloss.JoinVertical(0, header, body, footer),
	)

	dimmed := dimContent(background, w, h)

	return placeOverlay(w, h, box, dimmed)
}

// SetFocus sets the focused panel to the modal.
func (m *Model) SetFocus() {
	m.ctx.FocusedPanel = context.Modal
}

// RenderHelp renders a full help view with styles suited for the modal.
func (m Model) RenderHelp(km help.KeyMap) string {
	h := help.New()
	h.Styles.FullKey = lipgloss.NewStyle().Foreground(lipgloss.Color(style.White))
	h.Styles.FullDesc = lipgloss.NewStyle().Foreground(lipgloss.Color(style.MediumGray))
	h.Styles.FullSeparator = lipgloss.NewStyle()
	return h.FullHelpView(km.FullHelp())
}

// modalSize returns the modal dimension based on available space.
// Smaller terminals get a larger ratio so content remains usable.
func modalSize(available int) int {
	switch {
	case available < 40:
		return available - 2
	case available < 80:
		return available * 85 / 100
	default:
		return available * 3 / 4
	}
}

// dimContent strips existing colors and applies a dim foreground.
func dimContent(s string, w, h int) string {
	// Ensure background fills the full screen
	bg := lipgloss.NewStyle().Width(w).Height(h).Render(s)
	lines := strings.Split(bg, "\n")
	for i, line := range lines {
		// Strip existing ANSI sequences and re-render dimmed
		plain := stripAnsi(line)
		lines[i] = dimStyle.Render(plain)
	}
	return strings.Join(lines, "\n")
}

// placeOverlay centers fg on top of bg by replacing characters in bg.
func placeOverlay(w, h int, fg, bg string) string {
	fgLines := strings.Split(fg, "\n")
	bgLines := strings.Split(bg, "\n")

	fgW := lipgloss.Width(fg)
	fgH := len(fgLines)

	startY := (h - fgH) / 2
	startX := (w - fgW) / 2
	if startY < 0 {
		startY = 0
	}
	if startX < 0 {
		startX = 0
	}

	for i, fgLine := range fgLines {
		bgIdx := startY + i
		if bgIdx >= len(bgLines) {
			break
		}
		bgLine := bgLines[bgIdx]
		bgRunes := []rune(stripAnsi(bgLine))

		var prefix string
		if startX > 0 && startX <= len(bgRunes) {
			prefix = dimStyle.Render(string(bgRunes[:startX]))
		} else {
			prefix = strings.Repeat(" ", startX)
		}

		fgVisualW := lipgloss.Width(fgLine)
		endX := startX + fgVisualW
		var suffix string
		if endX < len(bgRunes) {
			suffix = dimStyle.Render(string(bgRunes[endX:]))
		}

		bgLines[bgIdx] = prefix + fgLine + suffix
	}

	return strings.Join(bgLines, "\n")
}

// stripAnsi removes ANSI escape sequences from a string.
func stripAnsi(s string) string {
	var out strings.Builder
	inEsc := false
	for _, r := range s {
		if r == '\x1b' {
			inEsc = true
			continue
		}
		if inEsc {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '~' {
				inEsc = false
			}
			continue
		}
		out.WriteRune(r)
	}
	return out.String()
}
