# Roadmap

Items are ordered by priority.

## 1. Update Go Version & Dependencies — `high`

- [x] Bump `go.mod` to latest Go version (1.26.1)
- [x] Update all dependencies and verify build

## 2. Tests & Coverage — `high`

- [ ] Add unit tests for `internal/gitlab` (client methods, variable builders)
- [ ] Add unit tests for `internal/config` (loading, defaults, dev mode fallback)
- [ ] Add unit tests for `internal/tui` (utility functions, message types)
- [ ] Add unit tests for component logic (table row/column builders, details rendering)
- [ ] Add coverage threshold to CI (`go test -coverprofile`)
- [ ] Add coverage badge to README

## 3. Address Linter Warnings — `high`

- [ ] Run `golangci-lint run` and fix all warnings
- [ ] Add linter to CI workflow

## 4. Fix TODO/FIX Comments — `medium`

- [ ] `mergerequests.go:236` — Refactor Icon + Status rendering in details view
- [ ] `projects/styles.go:16` — Set width from config instead of hardcoded 30
- [ ] `statusline/styles.go:9` — Update colors with design tokens
- [ ] `table/styles.go:36` — Update border color with design tokens
- [ ] `details/render.go:286` — Replace magic number 4 with proper calculation

## 5. Documentation — `medium`

- [ ] Record demo GIF with VHS and add to README
- [ ] Add GitHub issue templates (bug report, feature request)

## 6. Refactoring — `medium`

- [ ] Move `FormatTime` to a shared `tui/format.go`
- [ ] Move `StyleIconsColumns` to `mergerequests/columns.go` or `app/commands.go`
- [ ] Make `table` package a pure reusable widget
- [ ] Rename `internal/tui/components/mergerequests/` → `internal/tui/components/mrlist/`
- [ ] Update all imports across `app/`, `tui/`, and other components
- [ ] Move `setStatus`, `startTask`, `finishTask`, `toggleLeftPanel`, `toggleRightPanel`, `SelectMR` to `app/helpers.go`

## 7. New Features — `low`

- [x] Create new MR
- [ ] Visualize pipelines/jobs
- [ ] Interact with pipelines/jobs
- [ ] Visualize issues
- [ ] Create issues
- [ ] Filter MR
- [ ] Filter pipelines/jobs
- [ ] Filter issues
- [ ] Issues board
- [ ] Review MR
- [ ] Search projects
- [ ] View project README
- [ ] Publish to Homebrew (tap via GoReleaser)
