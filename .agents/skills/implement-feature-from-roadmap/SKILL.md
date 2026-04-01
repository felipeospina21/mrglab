---
name: implement feature from ROADMAP.md
description: When implementing a new feature from the ROADMAP.md list, follow this workflow.
---

## 1. Understand the codebase structure

- `internal/gitlab/` — GitLab API client, GraphQL queries, REST calls
  - `queries.go` — GraphQL query/mutation structs and variable builders
  - `mergerequests.go` — MR-related API methods (fetch, templates, project info)
  - `client.go` — Client struct, config, dev mode flag
  - `mock.go` — Mock data for dev mode (`-dev` flag)
- `internal/tui/` — Bubble Tea TUI layer
  - `msg.go` — All message types (events between components)
  - `app/update.go` — Main update loop, message routing, keybind handling
  - `app/model.go` — App model struct, initialization
  - `app/view.go` — Main view rendering
  - `app/commands.go` — Tea commands dispatched from the app layer
  - `app/form.go` — Reusable form component (textinput + textarea fields)
  - `components/mergerequests/commands.go` — MR-specific tea commands (fetch, create, etc.)
  - `components/mergerequests/mergerequests.go` — MR list component model
  - `components/modal/modal.go` — Modal overlay component
  - `components/details/` — MR details panel
  - `components/loader/` — Loading spinner view
- `internal/exec/` — Shell command helpers (browser, clipboard, git)
- `internal/config/` — Config loading from TOML
- `internal/context/` — Shared app context (selected project, MR, window size)

## 2. Implementation checklist

**API layer (`internal/gitlab/`):**
- Add new GraphQL query/mutation structs in `queries.go` if needed
- Add new API methods in the relevant file (e.g., `mergerequests.go`)
- Add mock data in `mock.go` for dev mode — return mocks when `c.devMode` is true
- Use parallel goroutines + channels when fetching from multiple sources

**Message types (`internal/tui/msg.go`):**
- Define a new `XxxMsg` struct for each async result
- Include an `Err error` field for error handling

**Commands (`components/.../commands.go`):**
- Create tea commands that call the API and return the message type
- Keep commands in the component package that owns the data

**Update loop (`app/update.go`):**
- Handle the trigger keybind/action in the appropriate `case` block
- Handle the result message to update the model and UI
- For modal-based features: set `m.isModalOpen`, `m.Modal.Header`, `m.Modal.Content`
- Show `loader.View(m.Spinner.View())` while waiting for async data
- Use a `ready` flag to gate input handling until data arrives

**View:**
- For form-based features, create a form struct (see `form.go` as reference)
- Use `textinput.Model` for single-line fields, `textarea.Model` for multi-line
- Use `Tab`/`Shift+Tab` for field navigation
- Use `lipgloss` for layout and styles from `internal/tui/style/`

## 3. Patterns to follow

- **Dev mode**: Always add mock data so `mrglab -dev` works without API calls
- **Parallel fetching**: Use goroutines + channels when calling multiple endpoints
- **Graceful degradation**: If an API call fails or returns empty, leave fields empty rather than erroring
- **Modal flow**: Open modal → show loader → fetch data → swap to form/content → handle submit/close → reset state
- **State cleanup**: Reset all feature-specific state on both modal close and submit
- **Smart defaults**: Pre-populate fields from available context (local git, API data)

## 4. Verify

```bash
go build ./... && go vet ./...
mrglab -dev  # test with mock data first
```

## 5. Update docs

- Mark the feature as done in `ROADMAP.md`: `- [x] Feature name`
- Add feature description to README.md **Features** list
- Add new keybindings to the relevant README.md **Keybindings** table
