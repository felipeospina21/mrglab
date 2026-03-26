# mrglab Code Restructure Plan

## Current Pain Points

After reviewing the full codebase, these are the main issues hurting readability and maintainability:

1. ~~**God `Update()` method**~~ — ✅ Fixed. Components now own their `Update()`. `app/update.go` is a thin dispatcher.

2. **Shared mutable state via `*context.AppContext`** — Every component holds a pointer to the same `AppContext` and mutates it freely (focus, selected MR, task status, panel state). This makes data flow invisible and debugging hard.

3. **Global config singleton** — `config.GlobalConfig` is accessed directly from `api`, `statusline`, `tui`, and `app` packages. This creates hidden coupling and makes testing impossible without global state manipulation.

4. ~~**`message` package is a thin wrapper**~~ — ✅ Deleted. Replaced with typed messages in `tui/msg.go`. See `TYPED_MESSAGES.md`.

5. **API functions repeat the same boilerplate** — Every function in `api/mrapi.go` does: check dev mode → find project by ID → build variables → create client → execute. This pattern is copy-pasted 3 times.

6. ~~**Task system uses raw `uint` types**~~ — ✅ Fixed. `task/task.go` deleted. Each async operation returns its own typed message. `ctx.Task` replaced with `ctx.TaskStatus` + `ctx.TaskErr`.

7. ~~**Components don't own their own Update logic**~~ — ✅ Fixed. Each component has an `Update()` method in its own `update.go` file. Components return action messages for app-level coordination.

8. **Mixed concerns in `app/model.go`** — ~~Layout calculations (`getFrameSize`, `getEmptyTableSize`, `setLeftPanelHeight`)~~(✅ extracted to `layout.go`), state mutations (`setStatus`, `startTask`, `finishTask`), and panel toggling all live in the same file.

9. **Custom `table` and `help` packages** — These are forked/modified versions of `charmbracelet/bubbles`. They're large (482 LOC for table, 232 for help) and mixed with app-specific utilities like `FormatTime`, `StyleIconsColumns`, etc.

---

## Proposed Structure

```
mrglab/
├── main.go
├── internal/
│   ├── config/
│   │   └── config.go              # (unchanged, but remove global singleton)
│   │
│   ├── gitlab/                    # renamed from api/ + gql/ + data/
│   │   ├── client.go              # GraphQL client setup (from api.utils.go)
│   │   ├── queries.go             # All query/mutation types + variables (merged gql/)
│   │   ├── mergerequests.go       # MR API operations (from mrapi.go)
│   │   ├── notes.go               # Note API operations (from noteapi.go)
│   │   ├── types.go               # Response types (from entities.go + query types)
│   │   └── mock.go                # Dev mode mock data (from data/)
│   │
│   ├── exec/
│   │   └── browser.go             # (unchanged)
│   │
│   ├── logger/
│   │   └── logger.go              # (unchanged)
│   │
│   └── tui/
│       ├── app/
│       │   ├── app.go             # Model struct + Init + constructor
│       │   ├── update.go          # Top-level Update (thin dispatcher)
│       │   ├── view.go            # Top-level View (layout only)
│       │   ├── layout.go          # Frame/size calculations (extracted from model.go)
│       │   └── commands.go        # App-level tea.Cmds
│       │
│       ├── components/
│       │   ├── projects/
│       │   │   ├── projects.go    # Model + New + Update + View
│       │   │   ├── keys.go        # Keybindings
│       │   │   ├── styles.go      # Styles
│       │   │   └── commands.go    # tea.Cmds
│       │   │
│       │   ├── mrlist/            # renamed from mergerequests/
│       │   │   ├── mrlist.go      # Model + New + Update + View
│       │   │   ├── keys.go
│       │   │   ├── columns.go     # Column definitions + row building
│       │   │   ├── styles.go
│       │   │   └── commands.go
│       │   │
│       │   ├── details/
│       │   │   ├── details.go     # Model + New + Update + View
│       │   │   ├── keys.go
│       │   │   ├── render.go      # Pipelines, discussions, approvals rendering
│       │   │   └── styles.go
│       │   │
│       │   ├── modal/
│       │   │   ├── modal.go
│       │   │   ├── keys.go
│       │   │   └── styles.go
│       │   │
│       │   ├── statusline/
│       │   │   ├── statusline.go
│       │   │   └── styles.go
│       │   │
│       │   ├── table/             # (keep as-is, it's a custom widget)
│       │   │   ├── table.go
│       │   │   ├── styles.go
│       │   │   └── utils.go
│       │   │
│       │   └── help/              # (keep as-is)
│       │       └── help.go
│       │
│       ├── keys.go                # Global keybindings
│       ├── msg.go                 # All custom tea.Msg types (replaces message/ + task/)
│       ├── style/
│       │   ├── colors.go
│       │   └── theme.go           # renamed from mainstyles.go
│       └── util.go                # Max, Min, Clamp, Truncate
```

---

## Key Changes Explained

### 1. Components own their Update

This is the single biggest improvement. Each component should handle its own key events when focused.

**Before** (`app/update.go` handles everything):

```go
// app/update.go — 270 lines of nested switches
if isMainPanelFocused {
    m.MergeRequests.Table, cmd = m.MergeRequests.Table.Update(msg)
    switch {
    case match(mpk.Details):
        // ... 10 lines of setup
    case match(mpk.Merge):
        // ...
    case match(mpk.OpenInBrowser):
        // ...
    }
}
```

**After** (component handles its own keys):

```go
// components/mrlist/mrlist.go
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        match := keys.KeyMatcher(msg)
        switch {
        case match(Keybinds.Details):
            return m, m.fetchDetails()
        case match(Keybinds.Merge):
            return m, m.acceptMR()
        case match(Keybinds.OpenInBrowser):
            m.openInBrowser()
        }
    }
    m.Table, cmd = m.Table.Update(msg)
    return m, cmd
}

// app/update.go — now a thin dispatcher
if isMainPanelFocused {
    m.MergeRequests, cmd = m.MergeRequests.Update(msg)
    cmds = append(cmds, cmd)
}
```

This makes `app/update.go` a ~80-line dispatcher instead of a 270-line monolith.

### 2. Replace global config with dependency injection

**Before:**

```go
// api/mrapi.go — reaches into global
cfg := &config.GlobalConfig

// config/config.go
var GlobalConfig Config
```

**After:**

```go
// gitlab/client.go
type Client struct {
    gql      *graphql.Client
    devMode  bool
}

func NewClient(cfg config.Config) *Client {
    // ...
}

// main.go
cfg, err := config.Load()
client := gitlab.NewClient(cfg)
m := app.New(cfg, client)
```

Pass `*Client` into components that need API access. This eliminates the global and makes testing straightforward — just pass a mock client.

### 3. Unify message types into `tui/msg.go`

**Before** (scattered across `task/task.go` and `message/messsage.go`):

```go
// task/task.go — generic wrapper
type TaskMsg struct {
    TaskID      TaskID
    SectionType TaskSection
    Status      TaskStatus
    Err         error
    Msg         tea.Msg  // type-asserted everywhere
}

// message/messsage.go — thin wrappers
type MergeRequestsListFetchedMsg struct { Mrs gql.MergeRequestConnection }
```

**After** (`tui/msg.go` — typed messages, no type assertions):

```go
package tui

// Each message is its own type — no wrapper, no type assertions
type MRListFetchedMsg struct {
    Mrs    gql.MergeRequestConnection
    Err    error
}

type MRDetailsFetchedMsg struct {
    Discussions []gql.DiscussionNode
    Stages      []gql.CiStageNode
    Branches    [2]string
    Approvals   []gql.ApprovalRule
    Err         error
}

type MRMergedMsg struct {
    Errors []string
    Err    error
}
```

This eliminates the `TaskMsg` wrapper and all the `msg.Msg.(type)` assertions in `update.go`. Each command returns its specific message type, and the `Update` function pattern-matches on concrete types.

### 4. Merge `api/` + `gql/` + `data/` into `gitlab/`

These three packages are tightly coupled — `api` imports `gql` for types and `data` for mocks. Merging them into `gitlab/` removes circular-dependency risk and makes the API surface obvious:

```go
import "github.com/felipeospina21/mrglab/internal/gitlab"

mrs, err := gitlab.Client.GetMergeRequests(projectID, opts)
mr, err := gitlab.Client.GetMergeRequest(projectID, iid)
res, err := gitlab.Client.AcceptMergeRequest(projectID, input)
```

### 5. Extract layout calculations from `app/model.go`

`model.go` currently mixes the `Model` struct definition with 10+ layout helper methods. Move size/frame calculations to `app/layout.go`:

```go
// app/layout.go
func (m Model) frameSize() (int, int) { ... }
func (m Model) emptyTableSize() (int, int) { ... }
func (m *Model) resizeLeftPanel() { ... }
func (m *Model) resizeStatusline() { ... }
func (m Model) viewportWidth() int { ... }
```

`app/app.go` stays clean with just the struct, constructor, and `Init()`.

### 6. Extract rendering logic from `details/details.go`

`details.go` is 409 lines because it mixes the bubbletea model with markdown rendering for pipelines, discussions, approvals, and branches. Split into:

- `details.go` — Model, New, Update, View, SetFocus (~80 lines)
- `render.go` — renderPipelines, renderDiscussions, renderApprovals, renderBranches, glamourRender (~300 lines)

### 7. Reduce `AppContext` surface

Currently `AppContext` holds everything: window size, focus state, selected items, task state, keybindings, panel visibility, and panel height. Trim it to only what genuinely needs to be shared:

```go
type AppContext struct {
    Window       tea.WindowSizeMsg
    FocusedPanel FocusedPanel
    PanelHeight  int

    SelectedProject struct {
        Name string
        ID   string
    }
    SelectedMR struct {
        IID    string
        Sha    string
        Status string
    }
}
```

Move out of context:

- `Keybinds` → each component already knows its own keybinds; the statusline help view can receive them as a parameter
- `Task` → replaced by typed messages (see point 3)
- `IsLeftPanelOpen`, `IsRightPanelOpen`, `IsModalOpen` → these are app-level state, keep them in `app.Model` directly

---

## Layout & Resize System — ✅ COMPLETED

See `LAYOUT_CHANGES.md` for full details on the implementation and bugs fixed.

### Summary

Replaced all scattered sizing logic with a centralized `layout.go` containing:

- `computeLayout(win, leftOpen, rightOpen) Layout` — single top-down calculation
- `applyLayout()` — pushes sizes to all components in-place
- Constants documenting lipgloss quirks: `tableBorderX/Y`, `tableViewOverhead`

### Bugs Fixed During Implementation

- Project list header clipping (delegate `Height()` returned 1, rendered 2)
- Table top border missing (lipgloss `BorderStyle` doesn't register in `GetFrameSize()`)
- Details panel pushed off-screen (statusline stretching left column width)
- Details panel empty on open (viewport not sized before content set)
- Table not filling available width (cell padding not accounted for in column width calculation)
- Table columns not resizing on first render (`recomputeLayout` not called after table creation)

### Column Width Fix

- `GetTableColums` now subtracts cell padding (2 chars × visible columns) before computing percentage-based widths
- Title column absorbs integer-rounding remainder so columns fill exact pixel width
- Title column bumped from 25% to 64% so visible columns sum to 100%

---

## Migration Order

Do these as independent, reviewable PRs:

1. ~~**Rename files**~~ — ✅ Done. `messsage.go` → `message.go`, `mainstyles.go` → `theme.go`, `keybinds.go` → `keys.go` (5 files).

2. ~~**Extract `details/render.go`**~~ — ✅ Done. Rendering functions split out of `details.go` (~300 LOC) into `render.go`.

3. ~~**Extract `app/layout.go`**~~ — ✅ Done. Layout calculations extracted from `model.go` into `layout.go`.

4. ~~**Rewrite layout system**~~ — ✅ Done. Replaced all scattered sizing functions with `computeLayout` + `applyLayout`. Removed magic numbers. See `LAYOUT_CHANGES.md`.

5. ~~**Add `Update` to components**~~ — ✅ Done. Each component (`mergerequests`, `projects`, `details`, `modal`) has its own `Update()` in `update.go`. Returns action messages for app-level coordination. `app/update.go` is now a thin dispatcher.

6. ~~**Replace `TaskMsg` with typed messages**~~ — ✅ Done. Created `tui/msg.go` with `MRListFetchedMsg`, `MRDetailsFetchedMsg`, `MRMergedMsg`. Deleted `task/` and `message/` packages. See `TYPED_MESSAGES.md`.

7. ~~**Merge `api/` + `gql/` + `data/` → `gitlab/`**~~ — ✅ Done. Consolidated into one package with `Client` struct. Components receive `*gitlab.Client` via constructors. Eliminated repeated boilerplate (dev mode check, project lookup, client creation).

8. ~~**Remove global config**~~ — ✅ Done. `config.GlobalConfig` is now only referenced in `main.go`. All other packages receive what they need through constructors: `*config.Config` to `InitMainModel`, `[]config.Project` to `projects.New()`, `devMode` via `AppContext`.

9. ~~**Trim `AppContext`**~~ — ✅ Done. Moved `IsLeftPanelOpen`, `IsRightPanelOpen`, `IsModalOpen`, `TaskStatus`, `TaskErr` to `app.Model`. Moved `Keybinds` to `statusline.Model`. `AppContext` now only holds genuinely shared state: `Window`, `SelectedProject`, `SelectedMR`, `FocusedPanel`, `PanelHeight`, `DevMode`.

---

## Notes

- All 9 steps are complete. The codebase now has centralized layout, component-owned updates, typed messages, a unified `gitlab` package with dependency-injected `Client`, no global config access outside `main.go`, and a minimal `AppContext`.
- The custom `table` and `help` packages are fine to keep as-is. They're stable and serve a specific purpose. Just avoid mixing app-specific utilities (like `FormatTime`) into the generic table package — those belong in `mrlist/columns.go`.
