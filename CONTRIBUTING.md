# Contributing to mrglab

Thanks for your interest in contributing to mrglab! This document outlines the process and guidelines for contributing.

## Prerequisites

- Go 1.23 or higher
- A [Nerd Font](https://github.com/ryanoasis/nerd-fonts) installed (for icon rendering)
- Familiarity with [Bubble Tea](https://github.com/charmbracelet/bubbletea)

## Getting started

```bash
# Clone your fork
git clone https://github.com/<your-username>/mrglab.git
cd mrglab

# Build
go build ./...

# Run in dev mode (no GitLab token needed)
go run . -dev
```

## Issue-first workflow

**Every pull request must reference an existing GitHub issue.**

1. Check [existing issues](https://github.com/felipeospina21/mrglab/issues) or the [IMPROVEMENTS.md](IMPROVEMENTS.md) for known tasks
2. If no issue exists for your change, **create one first** and wait for feedback before starting work
3. In your PR description, reference the issue (e.g. `Closes #42` or `Fixes #42`)

PRs without a linked issue will not be reviewed.

## Branch and PR workflow

1. Fork the repository
2. Create a branch from `main` using the convention: `feat/short-description`, `fix/short-description`, or `refactor/short-description`
3. Make your changes (one logical change per PR)
4. Push to your fork and open a PR targeting `main`

## Commit messages

This project uses [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add keyboard shortcut for copying MR URL
fix: resolve table overflow on narrow terminals
refactor: extract layout computation into module
chore: bump bubbletea to v1.3.0
docs: update README with keybinding tables
```

## Before submitting

Make sure the following pass locally:

```bash
go build ./...
go vet ./...
```

Also:

- Test your changes in dev mode (`go run . -dev`)
- Add godoc comments to any new exported types or functions
- Format your code with `gofmt`

## Code guidelines

- Follow standard Go conventions ([Effective Go](https://go.dev/doc/effective_go), [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments))
- Keep packages focused — UI components live in `internal/tui/components/`, API logic in `internal/gitlab/`
- Avoid adding dependencies unless necessary
- Don't modify or remove existing tests without discussion

## Reporting bugs

Open an issue with:

- Steps to reproduce
- Expected vs actual behavior
- Terminal emulator and OS
- Go version (`go version`)

## Feature proposals

Open an issue describing:

- The problem or use case
- Your proposed solution
- Any alternatives you considered

Please wait for maintainer feedback before implementing.

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).
