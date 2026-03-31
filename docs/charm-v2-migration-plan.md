# Charm Ecosystem v2 Migration Plan

## Problem Statement

Migrate the entire Charm dependency stack from v1 to v2. This touches ~30 Go source files and involves import path changes, API changes (View return type, key messages, viewport constructor, adaptive colors, etc.), and behavioral changes (declarative View fields replacing imperative commands).

## Requirements

- Update import paths for bubbletea, lipgloss, bubbles, and glamour to their `charm.land` v2 paths
- Adapt all API changes per the official upgrade guides
- Maintain identical TUI behavior after migration
- Build must pass, app must run in both normal and dev mode

## Key Breaking Changes

1. Import paths: `github.com/charmbracelet/*` → `charm.land/*/v2`
2. `View() string` → `View() tea.View` (top-level model only)
3. `tea.WithAltScreen()` → `view.AltScreen = true` in View
4. `tea.KeyMsg` → `tea.KeyPressMsg` everywhere
5. `viewport.New(w, h)` → `viewport.New(viewport.WithWidth(w), viewport.WithHeight(h))`
6. Viewport fields (`Width`, `Height`, `YOffset`) → getter/setter methods
7. `help.Model.Width` → `help.Model.SetWidth()` / `help.Model.Width()`
8. `lipgloss.AdaptiveColor` → `compat.AdaptiveColor` or `lipgloss.LightDark`
9. `HighPerformanceRendering` removed from viewport
10. `charmbracelet/log` stays at v1 (no v2 exists)
11. `charmbracelet/x/ansi` — keep as-is (still used by lipgloss v2 internally)

## Reference

- [Bubbletea v2 Upgrade Guide](https://github.com/charmbracelet/bubbletea/blob/main/UPGRADE_GUIDE_V2.md)
- [Bubbles v2 Upgrade Guide](https://github.com/charmbracelet/bubbles/blob/main/UPGRADE_GUIDE_V2.md)
- [Lipgloss v2 Upgrade Guide](https://github.com/charmbracelet/lipgloss/blob/main/UPGRADE_GUIDE_V2.md)

---

## Task Breakdown

### Task 1: Install v2 modules and update go.mod

**Objective:** Add the v2 charm modules to go.mod

```bash
go get charm.land/bubbletea/v2@latest charm.land/lipgloss/v2@latest charm.land/bubbles/v2@latest charm.land/glamour/v2@latest
```

Run `go mod tidy` after all source changes are done (will error until imports are updated).

**Demo:** `go.mod` and `go.sum` updated with v2 module paths.

---

### Task 2: Update all import paths across the codebase

**Objective:** Search-and-replace all import paths in every `.go` file (~20 files)

| v1 | v2 |
|---|---|
| `"github.com/charmbracelet/bubbletea"` | `"charm.land/bubbletea/v2"` |
| `"github.com/charmbracelet/lipgloss"` | `"charm.land/lipgloss/v2"` |
| `"github.com/charmbracelet/bubbles/key"` | `"charm.land/bubbles/v2/key"` |
| `"github.com/charmbracelet/bubbles/spinner"` | `"charm.land/bubbles/v2/spinner"` |
| `"github.com/charmbracelet/bubbles/textarea"` | `"charm.land/bubbles/v2/textarea"` |
| `"github.com/charmbracelet/bubbles/textinput"` | `"charm.land/bubbles/v2/textinput"` |
| `"github.com/charmbracelet/bubbles/viewport"` | `"charm.land/bubbles/v2/viewport"` |
| `"github.com/charmbracelet/bubbles/list"` | `"charm.land/bubbles/v2/list"` |
| `"github.com/charmbracelet/bubbles/table"` | `"charm.land/bubbles/v2/table"` |
| `"github.com/charmbracelet/glamour"` | `"charm.land/glamour/v2"` |

**Unchanged:** `github.com/charmbracelet/log`, `github.com/charmbracelet/x/ansi`

**Demo:** All imports updated. Won't compile yet due to API changes.

---

### Task 3: Migrate `tea.KeyMsg` → `tea.KeyPressMsg` globally

**Objective:** Update all `case tea.KeyMsg:` type switches and the `KeyMatcher` helper

**Files:**
- `internal/tui/keys.go` — Change `KeyMatcher` parameter from `tea.KeyMsg` to `tea.KeyPressMsg`
- `internal/tui/app/update.go` — `case tea.KeyMsg:` → `case tea.KeyPressMsg:`; also form key handling `keyMsg, ok := msg.(tea.KeyMsg)` → `tea.KeyPressMsg`
- `internal/tui/components/table/table.go` — `case tea.KeyMsg:` → `case tea.KeyPressMsg:`
- `internal/tui/components/details/update.go` — `case tea.KeyMsg:` → `case tea.KeyPressMsg:`
- `internal/tui/components/modal/update.go` — `case tea.KeyMsg:` → `case tea.KeyPressMsg:`
- `internal/tui/components/mergerequests/update.go` — `case tea.KeyMsg:` → `case tea.KeyPressMsg:`
- `internal/tui/components/projects/update.go` — `case tea.KeyMsg:` → `case tea.KeyPressMsg:`

**Demo:** All key event handling uses v2 types.

---

### Task 4: Migrate top-level `View()` to return `tea.View` and move `WithAltScreen` to declarative

**Objective:** Change `main.go` and `app/view.go` for the new View API

**Changes in `main.go`:**
- Remove `tea.WithAltScreen()` from `tea.NewProgram(m, tea.WithAltScreen())`
- Becomes: `tea.NewProgram(m)`

**Changes in `app/view.go`:**
- Change `func (m Model) View() string` → `func (m Model) View() tea.View`
- Wrap return values with `tea.NewView(screen)` and set `v.AltScreen = true`

```go
// Before
func (m Model) View() string {
    // ...
    return screen
}

// After
func (m Model) View() tea.View {
    // ...
    v := tea.NewView(screen)
    v.AltScreen = true
    return v
}
```

**Demo:** App compiles with new View signature and uses declarative alt screen.

---

### Task 5: Migrate viewport to v2 API

**Objective:** Update viewport constructor and field access patterns

**Files:**

**`internal/tui/components/details/details.go`:**
- `viewport.New(10, 10)` → `viewport.New(viewport.WithWidth(10), viewport.WithHeight(10))`
- `m.Viewport.Width` (read) → `m.Viewport.Width()`
- `m.Viewport.Width = w` → `m.Viewport.SetWidth(w)`
- `m.Viewport.Height` (read) → `m.Viewport.Height()`
- `m.Viewport.Height = msg.Height - ...` → `m.Viewport.SetHeight(...)`
- `m.Viewport.YPosition = headerHeight` → remove (not needed in v2)
- `m.Viewport.HighPerformanceRendering = ...` → remove entirely
- Remove `viewport.Sync` call
- `m.Viewport.YOffset` reads → `m.Viewport.YOffset()`

**`internal/tui/components/table/table.go`:**
- `viewport.New(0, 20)` → `viewport.New(viewport.WithWidth(0), viewport.WithHeight(20))`
- `m.viewport.Height` (read) → `m.viewport.Height()`
- `m.viewport.Height = h - ...` → `m.viewport.SetHeight(h - ...)`
- `m.viewport.Width = w` → `m.viewport.SetWidth(w)`
- `m.viewport.Width` (read) → `m.viewport.Width()`
- `m.viewport.YOffset` (read) → `m.viewport.YOffset()`
- `m.viewport.SetYOffset(...)` stays the same

**Demo:** Viewport usage compiles with v2 API.

---

### Task 6: Migrate `lipgloss.AdaptiveColor` usage

**Objective:** Replace `lipgloss.AdaptiveColor` with `compat.AdaptiveColor` or hardcode dark theme values

Since this is a dark-themed TUI (always uses alt screen), the simplest approach is to hardcode the dark color values directly. Alternatively, use `compat.AdaptiveColor`.

**Files:**
- `internal/tui/components/help/help.go` — Uses `lipgloss.AdaptiveColor{Light: "...", Dark: "..."}`
- `internal/tui/components/statusline/styles.go` — Uses `lipgloss.AdaptiveColor`
- `internal/tui/components/projects/styles.go` — Uses `lipgloss.AdaptiveColor` in `NewDefaultItemStyles`

**Option A (simplest):** Replace with just the dark value:
```go
// Before
lipgloss.AdaptiveColor{Light: "#909090", Dark: "#626262"}
// After
lipgloss.Color("#626262")
```

**Option B (preserve adaptivity):**
```go
import "charm.land/lipgloss/v2/compat"
compat.AdaptiveColor{Light: lipgloss.Color("#909090"), Dark: lipgloss.Color("#626262")}
```

**Demo:** All adaptive color usage compiles.

---

### Task 7: Verify custom help component

**Objective:** The project has its own `help` package (`internal/tui/components/help`), not using `bubbles/help`. The custom `help.Model` has an exported `Width int` field — this is our own code, so no migration needed.

**Demo:** No-op. Custom help component works unchanged.

---

### Task 8: Migrate `bubbles/table` styles import

**Objective:** Update import in `internal/tui/components/table/styles.go`

- `"github.com/charmbracelet/bubbles/table"` → `"charm.land/bubbles/v2/table"` (already covered in Task 2)
- Verify `table.DefaultStyles()` API is unchanged in v2

**Demo:** Table styles compile.

---

### Task 9: Verify glamour v2 API compatibility

**Objective:** Check `internal/tui/components/details/render.go`

The glamour API appears unchanged in v2:
- `glamour.NewTermRenderer(...)` — same signature
- `glamour.WithStandardStyle("dark")` — same
- `glamour.WithWordWrap(width)` — same
- `glamour.WithEmoji()` — same
- `glamour.WithPreservedNewLines()` — same

**Demo:** Markdown rendering compiles.

---

### Task 10: Verify `tea.WindowSizeMsg` in context

**Objective:** `tea.WindowSizeMsg` is still a struct in v2 with `Width` and `Height` fields.

File: `internal/context/context.go` — no changes beyond import path (Task 2).

**Demo:** No-op beyond import.

---

### Task 11: Verify `charmbracelet/x/ansi` import

**Objective:** `internal/tui/components/projects/projects.go` uses `ansi.Truncate`

Lipgloss v2 still depends on `charmbracelet/x/ansi` internally. The `ansi.Truncate` function is still available. Keep the import as-is.

**Demo:** Projects component compiles.

---

### Task 12: Build and verify

**Objective:** Clean up and verify everything works

```bash
go mod tidy
go build ./...
mrglab -dev
```

Fix any remaining compilation errors discovered during build.

**Demo:** App builds and runs in dev mode with identical behavior.

---

### Task 13: Verify clean go.mod

**Objective:** Ensure no v1 charm packages remain as direct dependencies

After `go mod tidy`, verify `go.mod` no longer contains these as **direct** dependencies:
- `github.com/charmbracelet/bubbletea`
- `github.com/charmbracelet/lipgloss`
- `github.com/charmbracelet/bubbles`
- `github.com/charmbracelet/glamour`

These may remain as **indirect** (transitive deps of `charmbracelet/log`):
- `github.com/charmbracelet/log` (stays at v1, no v2)
- `github.com/charmbracelet/x/ansi` (still used directly)

**Demo:** Clean go.mod with only v2 charm direct dependencies.
