# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/), and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added

- Pipelines tab with table view (status, IID, commit, jobs count, author, branch, source, duration)
- Tab navigation between Merge Requests and Pipelines views (`tab` key)
- Pipeline details panel with stages and jobs breakdown
- Navigate non-success jobs in pipeline details (`n`/`N`)
- Run individual jobs: play manual jobs, retry failed/canceled/skipped (`P`)
- Cancel running jobs from pipeline details (`X`)
- Retry all failed jobs in a pipeline (`r`)
- Cancel a running pipeline (`C`)
- Open pipeline in browser from both table and details view
- Create new merge request with default templates (`N`)
- Refresh merge requests list (`R`)

### Fixed

- Help keys showing MR commands when switching projects while on pipelines tab
- Open-in-browser from pipeline details opening MR URL instead of pipeline URL
- Retried job duplicates appearing in pipeline details
- Content disappearing when navigating jobs in pipeline details

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
