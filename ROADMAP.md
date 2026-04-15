# Roadmap

Items are ordered by priority.

## 1. Fix Shell Integration Issues — `critical`

These violate tuishell's documented message-based contract and cause state desync bugs.

- [x] **Dual FocusedPanel state**: mrglab passes its own `*context.AppContext` to components, but `shell.Model` creates a separate `tuishell.AppContext`. Components call `m.ctx.FocusedPanel = context.RightPanel` which updates mrglab's copy, while the shell reads/writes `m.Shell.Ctx.FocusedPanel`. The two diverge after any shell-initiated focus change (e.g. `CloseModalMsg` restores `prevFocus` only in the shell's copy). The `MRMergedMsg` handler checks `m.ctx.FocusedPanel` which may be stale. **Fix**: make mrglab's `AppContext` point to (or embed a pointer to) the shell's `Ctx`, or sync after every shell update.
- [x] **Duplicated message handling**: `CloseModalMsg` and `ResetHighlightMsg` are handled by both mrglab and the shell. The shell already closes the modal and restores focus on `CloseModalMsg`, but mrglab also manually calls `SetFocus()` (writing to the wrong context) and directly mutates `m.Shell.Statusline.Status`/`Content`. mrglab should only handle its own cleanup (form reset, input blur) and delegate shell state to messages.
- [x] **Direct statusline mutation in CloseModalMsg**: `m.Shell.Statusline.Status = mode` and `m.Shell.Statusline.Content = ""` bypass the message bus. Use `SetStatusMsg` instead so the shell owns its statusline state.
- [x] **Remove redundant `ResetHighlightMsg` handler**: the shell already sets `m.Modal.Highlight = false` — mrglab's handler is a no-op duplicate.

## 2. Update Go Version & Dependencies — `high`

- [x] Bump `go.mod` to latest Go version (1.26.1)
- [x] Update all dependencies and verify build

## 3. Tests & Coverage — `high`

- [x] Add unit tests for `internal/gitlab` (client methods, variable builders)
- [x] Add unit tests for `internal/config` (loading, defaults, dev mode fallback)
- [x] Add unit tests for `internal/tui` (utility functions, message types)
- [x] Add unit tests for component logic (table row/column builders, details rendering)
- [x] Extract `GitLabAPI` interface for testable consumers
- [ ] Add coverage threshold to CI (`go test -coverprofile`) ([#32](https://github.com/felipeospina21/mrglab/issues/32))
- [x] Add coverage badge to README

## 4. Address Linter Warnings — `high`

- [ ] Run `golangci-lint run` and fix all warnings ([#33](https://github.com/felipeospina21/mrglab/issues/33))
- [ ] Add linter to CI workflow ([#33](https://github.com/felipeospina21/mrglab/issues/33))

## 5. Fix TODO/FIX Comments — `medium`

- [ ] `mergerequests.go:236` — Refactor Icon + Status rendering in details view ([#31](https://github.com/felipeospina21/mrglab/issues/31))
- [ ] `projects/styles.go:16` — Set width from config instead of hardcoded 30 ([#30](https://github.com/felipeospina21/mrglab/issues/30))
- [x] `statusline/styles.go:9` — Update colors with design tokens
- [ ] `table/styles.go:36` — Update border color with design tokens ([#34](https://github.com/felipeospina21/mrglab/issues/34))
- [ ] `details/render.go:286` — Replace magic number 4 with proper calculation ([#30](https://github.com/felipeospina21/mrglab/issues/30))
- [ ] Details panel content doesn't reflow to use full width when toggling fullscreen with "f" ([#29](https://github.com/felipeospina21/mrglab/issues/29))

## 6. Documentation — `medium`

- [x] Record demo GIF with VHS and add to README
- [ ] Add GitHub issue templates (bug report, feature request) ([#35](https://github.com/felipeospina21/mrglab/issues/35))

## 7. Refactoring — `medium`

- [ ] Move `FormatTime` to a shared `tui/format.go` ([#37](https://github.com/felipeospina21/mrglab/issues/37))
- [ ] Move `StyleIconsColumns` to `mergerequests/columns.go` or `app/commands.go` ([#37](https://github.com/felipeospina21/mrglab/issues/37))
- [ ] Make `table` package a pure reusable widget (extract to standalone Go module) ([#38](https://github.com/felipeospina21/mrglab/issues/38))
- [ ] Rename `internal/tui/components/mergerequests/` → `internal/tui/components/mrlist/` ([#36](https://github.com/felipeospina21/mrglab/issues/36))
- [ ] Update all imports across `app/`, `tui/`, and other components ([#36](https://github.com/felipeospina21/mrglab/issues/36))
- [ ] Move `setStatus`, `SelectMR` to `app/helpers.go` ([#37](https://github.com/felipeospina21/mrglab/issues/37))

## 8. New Features — `low`

- [x] Create new MR
- [x] Dynamic statusline colors per status mode (normal, loading, error, etc.)
- [ ] Visualize pipelines/jobs ([#41](https://github.com/felipeospina21/mrglab/issues/41))
- [ ] Interact with pipelines/jobs ([#41](https://github.com/felipeospina21/mrglab/issues/41))
- [ ] Visualize issues ([#40](https://github.com/felipeospina21/mrglab/issues/40))
- [ ] Create issues ([#40](https://github.com/felipeospina21/mrglab/issues/40))
- [ ] Filter MR ([#39](https://github.com/felipeospina21/mrglab/issues/39))
- [ ] Filter pipelines/jobs ([#39](https://github.com/felipeospina21/mrglab/issues/39))
- [ ] Filter issues ([#39](https://github.com/felipeospina21/mrglab/issues/39))
- [ ] Issues board ([#40](https://github.com/felipeospina21/mrglab/issues/40))
- [ ] Review MR ([#44](https://github.com/felipeospina21/mrglab/issues/44))
- [ ] Search projects ([#42](https://github.com/felipeospina21/mrglab/issues/42))
- [ ] View project README ([#42](https://github.com/felipeospina21/mrglab/issues/42))
- [ ] Publish to Homebrew (tap via GoReleaser) ([#43](https://github.com/felipeospina21/mrglab/issues/43))
- [ ] refetch mrs after merging ([#44](https://github.com/felipeospina21/mrglab/issues/44))
- [ ] refresh command for mrs and updated timestamp in status bar ([#44](https://github.com/felipeospina21/mrglab/issues/44))
- [x] Add missing fields to MR table columns (pipeline status, MR number, merge details/commits behind)
- [x] Fix approvals column to use rule-based logic instead of flat count
