---
name: mrglab-architecture
description: Project structure, patterns, and conventions for mrglab. Use when planning features or understanding the codebase.
---

# Project Architecture

- `internal/tui/app/` — Main bubbletea Model, Update, View
- `internal/tui/components/` — UI components (mergerequests, details, projects, table, statusline, modal)
- `internal/tui/style/` — Shared lipgloss styles
- `internal/tui/icon/` — Nerd Font icon constants
- `internal/tui/` — Shared keys, messages, utils
- `internal/gitlab/` — GraphQL API client, queries, types, mock data
- `internal/config/` — TOML config loading from ~/.config/mrglab/mrglab.toml
- `internal/exec/` — Browser, clipboard, git helpers
- `internal/context/` — App-wide shared context
- `internal/logger/` — Structured logging

## Key Patterns

- Bubbletea v2 + lipgloss v2 (charm.land/bubbletea/v2)
- GraphQL via hasura/go-graphql-client for GitLab API
- `GitLabAPI` interface in internal/gitlab/ enables mocking for tests
- Dev mode (`-dev` flag) uses mock data, no API calls
- Config via viper + TOML
- Conventional commits, issue-first workflow, one logical change per PR
