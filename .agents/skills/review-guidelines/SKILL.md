---
name: review-guidelines
description: Code review guidelines for mrglab. Use when reviewing changes on a feature branch.
---

## Workflow

1. Run `git diff main...HEAD` to see all changes in the current feature branch
2. Read the changed files for full context
3. Review against the project architecture skill and CONTRIBUTING.md conventions
4. Provide feedback in the sections below — skip sections with no findings

## Best Practices

- Go idioms, error handling, naming, package organization
- Alignment with existing project patterns (bubbletea model/update/view, GitLabAPI interface, etc.)
- Don't be nitpicky about style — focus on substance

## Convention Compliance

- Does it follow the project structure from the architecture skill?
- Are new files in the right packages?
- Are exports documented with godoc comments?
- Does it follow conventional commits if commit messages are visible?

## Security Assessment

- Hardcoded secrets, tokens, or credentials
- Unsafe use of user input (injection, path traversal)
- Insecure HTTP, disabled TLS verification
- Use of unsafe package
- Exposed sensitive data in logs or error messages

## Output

Be constructive. Suggest concrete fixes with file paths and line references.
