# Changes Summary

## 1. Cross-Tool AI Agent Compatibility

### Problem
Kiro skills in `.kiro/skills/` weren't available to other AI coding tools (Gemini CLI, Codex, Claude Code, Copilot, Cursor).

### Solution
- Created `.agents/skills/` as the **single source of truth** for skills, using the same `<skill-name>/SKILL.md` directory structure that Kiro, Gemini CLI, and Codex all support.
- Symlinked `.kiro/skills/` → `.agents/skills/` so Kiro reads the same files.
- Added Kiro-compatible YAML frontmatter (`---name/description---`) to each `SKILL.md` — harmless to Gemini/Codex, required by Kiro.
- Created a slim `AGENTS.md` at the repo root as an index for tools that don't support `.agents/` (Claude Code, Copilot, Cursor).
- Created `scripts/sync-agents.sh` to copy `AGENTS.md` into tool-specific formats (`CLAUDE.md`, `.github/copilot-instructions.md`, `.cursor/rules/mrglab.mdc`).

### File structure
```
.agents/skills/                                    ← Source of truth
  project-architecture/SKILL.md
  implement-feature-from-roadmap/SKILL.md
  review-guidelines/SKILL.md

.kiro/skills/                                      ← Symlinks
  project-architecture/SKILL.md  →  .agents/...
  implement-feature-from-roadmap/SKILL.md  →  .agents/...
  review-guidelines/SKILL.md     →  .agents/...

AGENTS.md                                          ← Slim index for other tools
scripts/sync-agents.sh                             ← Syncs AGENTS.md to tool-specific files
```

### Token impact
- Monolithic `AGENTS.md` loaded every request: ~3K tokens
- Slim `AGENTS.md` + selective skill loading: ~200 tokens base, skills loaded only when relevant

---

## 2. Extract Reusable TUI Layout Package (`tuishell`)

### Problem
mrglab's layout, table, modal, statusline, and loader components were tightly coupled to GitLab types, making them impossible to reuse in a future Jira TUI.

### Solution
Created `tuishell` (`../tuishell/`) as a standalone Go module containing all reusable TUI infrastructure. Refactored mrglab to import from it.

### What moved to `tuishell`
| Package | Contents |
|---------|----------|
| `tuishell` (root) | `AppContext`, `Layout`, `ComputeLayout()`, `GlobalKeyMap`, `KeyMatcher`, `Max/Min/Clamp/Truncate` |
| `tuishell/style` | `Theme` struct (30 semantic color tokens), `DefaultTheme()`, `MainFrameStyle()`, color palettes |
| `tuishell/table` | Table widget, `ThemedStyles()`, `TitleStyle()`, `DocStyle()`, `FormatTime`, `ColWidth`, `GetColIndex` |
| `tuishell/modal` | Overlay modal, themed styles, keybindings, update logic |
| `tuishell/statusline` | Status bar with `ProjectLabel` slot, mode colors, spinner style |
| `tuishell/loader` | Themed loading spinner view |

### What stayed in mrglab
- `internal/gitlab/` — API client, types, queries, mock data
- `internal/config/` — TOML config
- `internal/exec/` — browser, clipboard, git helpers
- `internal/tui/msg.go` — GitLab-specific message types
- `internal/tui/icon/` — Nerd Font icons
- `internal/tui/components/projects/` — GitLab project list
- `internal/tui/components/details/` — MR details rendering
- `internal/tui/components/mergerequests/` — MR table columns, commands
- `internal/tui/components/table/utils.go` — `StyleIconsColumns` (GitLab-specific icon coloring)
- `internal/tui/app/` — mrglab-specific Model, Update, View

### How mrglab imports tuishell
Each mrglab component became a thin wrapper that re-exports from tuishell:
- `internal/context/` — embeds `tuishell.AppContext`, adds `SelectedProject`, `SelectedMR`
- `internal/tui/components/modal/` — wraps `tsmodal.Model`, adds `Update` returning correct type
- `internal/tui/components/statusline/` — wraps `tssl.Model`, injects `icon.Gitlab + project name`
- `internal/tui/components/table/` — re-exports types, keeps `StyleIconsColumns`
- `internal/tui/components/loader/` — delegates to tuishell with default theme
- `internal/tui/app/layout.go` — calls `tuishell.ComputeLayout()` with mrglab's panel styles

### Theme struct
30 semantic color tokens grouped into:
- **Primary** (4) — accent family for titles, selections, highlights
- **Semantic** (9) — info, success, danger, warning, caution with bright variants
- **Neutral** (9) — text, borders, dimmed surfaces, modal overlay
- **Statusline** (7) — mode colors and segment backgrounds

Components receive the theme at construction time. `DefaultTheme()` matches the original mrglab colors. A Jira TUI would pass its own theme:
```go
theme := style.Theme{
    Primary:      lipgloss.Color("#0052CC"),
    PrimaryBright: lipgloss.Color("#2684FF"),
    // ...
}
```

### Local development setup
`go.mod` uses a `replace` directive for local development:
```
replace github.com/felipeospina21/tuishell => ../tuishell
```
