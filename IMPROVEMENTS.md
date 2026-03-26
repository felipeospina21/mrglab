# Improvements Roadmap

Items are ordered by priority. Checked items are complete.

## 1. Versioning & Releases

- [ ] Tag initial release as `v0.1.0`
- [ ] Add `CHANGELOG.md`
- [ ] Add GitHub Actions release workflow (create GitHub Release on tag push)
- [ ] Set up `.goreleaser.yml` for cross-platform binary builds

## 2. Update Go Version & Dependencies

- [ ] Bump `go.mod` to latest Go version (currently 1.23.1)
- [ ] Update all dependencies
- [ ] Verify build and fix any breaking changes

## 3. Address Linter Warnings

- [ ] Run `golangci-lint run` and fix all warnings
- [ ] Add linter to CI workflow

## 4. Fix TODO/FIX Comments

- [ ] `mergerequests.go:236` — Refactor function to render Icon + Status in details view
- [ ] `projects/styles.go:16` — Set width from config instead of hardcoded 30
- [ ] `statusline/styles.go:9` — Update colors with design tokens
- [ ] `table/styles.go:36` — Update border color with design tokens
- [ ] `details/render.go:286` — Investigate magic number 4 and replace with proper calculation

## 5. Documentation & README

- [x] Add godoc comments to exported types and functions
- [x] Update README with new features (respond to discussions, keybindings)
- [x] Document configuration options
- [ ] Record demo GIF with VHS and add to README
- [ ] Add GitHub issue templates (bug report, feature request)

## 6. Loading State on MR Details

- [x] Show loading spinner when fetching MR details
- [x] Dismiss loading state when data arrives or on error
- [x] Unify loader into reusable component

## 7. Improve .gitignore

- [ ] Add `debug.log`, `c.out`, `*.exe`

## 8. Move App-Specific Utilities Out of `table` Package

- [ ] Move `FormatTime` to a shared `tui/format.go`
- [ ] Move `StyleIconsColumns` to `mergerequests/columns.go` or `app/commands.go`
- [ ] `table` package becomes a pure reusable widget

## 9. Rename `mergerequests/` → `mrlist/`

- [ ] Rename `internal/tui/components/mergerequests/` → `internal/tui/components/mrlist/`
- [ ] Update all imports across `app/`, `tui/`, and other components

## 10. Extract `app/model.go` Helpers

- [ ] Move `setStatus`, `startTask`, `finishTask`, `toggleLeftPanel`, `toggleRightPanel`, `SelectMR` to `app/helpers.go`

## 11. Reduce `AppContext` Mutations from Components

No action needed unless the codebase grows. Current `AppContext` surface (6 fields) is reasonable. Documented here for future reference:

- **Focus management** — 4 components set `ctx.FocusedPanel` directly. Could use messages instead, but adds indirection for minimal gain.
- **Selection state** — `projects.SelectProject()` and `app.SelectMR()` are natural owners.
- **PanelHeight** — set by `app/layout.go`, read by `modal`. Legitimate shared layout state.
