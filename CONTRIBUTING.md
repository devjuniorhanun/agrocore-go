# Contributing to AgroCore

Thank you for your interest in contributing to AgroCore.

## Development Requirements

- Go 1.27+
- Git
- Git Flow

## Branching Strategy

The project uses Git Flow.

Do not develop directly on `main` or `develop`.

Start a feature with:

```bash
git flow feature start <feature-name>
```

Example:

```bash
git flow feature start product-registration
```

## Commit Messages

Commit messages follow Conventional Commits.

Common types:

- `feat` — new functionality
- `fix` — bug fix
- `docs` — documentation
- `test` — tests
- `refactor` — code changes without behavioral changes
- `perf` — performance improvements
- `build` — build system or dependency changes
- `ci` — continuous integration changes
- `chore` — maintenance tasks

Examples:

```text
feat: add product registration
fix: prevent empty product names
docs: document database decision
test: add product validation tests
refactor: extract product repository
```

## Code Quality

Before committing Go code, run:

```bash
gofmt -w .
go vet ./...
go test ./...
```

All commands must complete successfully.

## Documentation

Documentation must be updated when a change affects:

- public behavior;
- development workflows;
- architecture;
- configuration;
- external dependencies;
- deployment or operation.

Significant architectural decisions should be documented as ADRs under
`docs/decisions/`.

## Security

Never commit:

- passwords;
- API keys;
- tokens;
- private certificates;
- production environment files;
- other credentials or secrets.

Use environment variables and `.env.example` when configuration examples
are required.

## Pull Requests

Pull requests should:

- have a clear purpose;
- contain cohesive changes;
- pass automated checks;
- include or update tests when behavior changes;
- update documentation when necessary.