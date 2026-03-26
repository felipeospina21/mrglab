# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/), and this project adheres to [Semantic Versioning](https://semver.org/).

## [0.1.0] - 2026-03-26

Initial public release.

### Features

- Browse open merge requests for configured GitLab projects
- View MR details: description, pipeline stages, approvals, branches, and discussions
- Merge a merge request directly from the TUI
- Open any MR in the default browser
- Respond to discussion threads (post comments)
- Navigate between resolvable discussions with `n`/`N`
- Copy modal content to clipboard
- Toggle project list and details panels
- Full-screen help modal with all keybindings
- Loading spinner when fetching MR list and details
- Dev mode with mocked data (`mrglab -dev`)

### Documentation

- README with features, keybindings, and configuration reference
- Godoc comments on all exported types and functions
- CONTRIBUTING.md with issue-first workflow
- MIT License

### CI

- Build and vet checks on pull requests
- Conventional commit validation
- PR template with checklist
