# Improvements Roadmap

## 1. Fix TODO/FIX Comments

- [ ] `mergerequests.go:236` — Refactor function to render Icon + Status in details view
- [ ] `projects/styles.go:16` — Set width from config instead of hardcoded 30
- [ ] `statusline/styles.go:9` — Update colors with design tokens
- [ ] `table/styles.go:36` — Update border color with design tokens
- [ ] `details/render.go:286` — Investigate magic number 4 and replace with proper calculation

## 2. Update Go Version & Dependencies

- [ ] Bump `go.mod` to latest Go version (currently 1.23.1)
- [ ] Update all dependencies
- [ ] Verify build and fix any breaking changes

## 3. Loading State on MR Details

- [ ] Show loading modal/spinner when fetching MR details
- [ ] Dismiss loading state when data arrives or on error

## 4. Address Linter Warnings

- [ ] Run `golangci-lint run` and fix all warnings
- [ ] Consider adding linter to CI workflow

## 5. Documentation & README

- [x] Add godoc comments to exported types and functions
- [x] Update README with new features (respond to discussions, keybindings)
- [x] Document configuration options
- [ ] Add usage examples / screenshots

## Remaining Improvements

These are lower-priority cleanups that weren't part of the original 9-step migration but would further improve the codebase.

### 10. Move app-specific utilities out of `table` package

`table/utils.go` contains `FormatTime` and `StyleIconsColumns` — these are app-specific formatting helpers that don't belong in a generic table widget.

- Move `FormatTime` to a shared `tui/format.go` (used by `mergerequests/mergerequests.go` and `details/render.go`)
- Move `StyleIconsColumns` to `mergerequests/columns.go` or `app/commands.go` (only used in `app/commands.go`)
- `table` package becomes a pure reusable widget with no app-specific knowledge

### 11. Rename `mergerequests/` → `mrlist/`

The proposed structure uses `mrlist/` which is shorter and consistent with the component's role (it's a list view, not the domain concept). This is a straightforward rename:

- `internal/tui/components/mergerequests/` → `internal/tui/components/mrlist/`
- Update all imports across `app/`, `tui/`, and other components

### 12. Reduce `AppContext` mutations from components

Components currently mutate `ctx` directly for two concerns:

**a) Focus management** — 4 components set `m.ctx.FocusedPanel`:

- `projects.SetFocus()`, `mergerequests.SetFocus()`, `details.SetFocus()`, `modal.SetFocus()`

These could return a message instead (e.g., `FocusPanelMsg{Panel}`) and let `app` be the single writer. But this adds indirection for minimal gain since focus is inherently shared state. **Keep as-is unless it causes bugs.**

**b) Selection state** — `projects.SelectProject()` writes `ctx.SelectedProject`, and `app.SelectMR()` writes `ctx.SelectedMR`. These are fine — they're the natural owners of that data.

**c) `PanelHeight`** — set by `app/layout.go`, read by `modal`. This is legitimate shared layout state. **Keep as-is.**

No action needed here unless the codebase grows. The current `AppContext` surface (6 fields) is reasonable.

### 13. Extract `app/model.go` helpers

`model.go` still mixes the struct definition with state mutation helpers (`setStatus`, `startTask`, `finishTask`, `toggleLeftPanel`, `toggleRightPanel`, `SelectMR`). These could move to a separate `app/helpers.go` to keep `model.go` focused on the struct + constructor. Low priority — the file is ~150 lines, not a pain point yet.
