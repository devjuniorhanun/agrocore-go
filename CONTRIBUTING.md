# Contributing to AgroCore

Thank you for your interest in contributing to AgroCore.

This document describes the development workflow, quality requirements, commit
conventions, and architectural expectations used by the project.

## Development Philosophy

AgroCore evolves incrementally.

Changes should solve concrete requirements without introducing unnecessary
abstractions or infrastructure.

Contributors should prefer:

- simple and explicit code;
- small focused changes;
- business-oriented modules;
- automated tests;
- clear architectural boundaries;
- professional documentation;
- documented technical decisions.

Avoid introducing abstractions only because they may become useful in the
future.

## Requirements

Before contributing, install:

- Go;
- Git;
- Git Flow;
- GNU Make.

Check the installed tools:

```bash
go version
git --version
git flow version
make --version
```

## Repository Setup

Clone the repository:

```bash
git clone git@github.com:devjuniorhanun/agrocore-go.git
```

Enter the project:

```bash
cd agrocore-go
```

Fetch the latest remote state:

```bash
git fetch origin
```

Switch to the development branch:

```bash
git switch develop
```

Update it:

```bash
git pull --ff-only origin develop
```

## Git Flow

AgroCore uses Git Flow.

Permanent branches:

```text
main
develop
```

`main` represents production-ready history.

`develop` is the integration branch for ongoing development.

Feature development must normally start from `develop`.

Example:

```bash
git switch develop
git pull --ff-only origin develop
git flow feature start agricultural-year
```

This creates a branch similar to:

```text
feature/agricultural-year
```

## Pull Request Workflow

The preferred workflow is:

```text
develop
   |
   v
feature/*
   |
   v
commits
   |
   v
push
   |
   v
Pull Request
   |
   v
CI
   |
   v
Squash and Merge
   |
   v
develop
```

Once a feature is being integrated through a Pull Request, do not also use
`git flow feature finish` to merge the same feature locally.

The Pull Request is the integration mechanism.

## Commit Convention

AgroCore uses Conventional Commits.

Format:

```text
<type>: <description>
```

Examples:

```text
feat: add agricultural year creation
fix: reject invalid HTTP port
docs: document agricultural year module
test: add repository integration tests
refactor: simplify application service
ci: update quality workflow
chore: update development tooling
```

Common commit types:

| Type | Purpose |
| --- | --- |
| `feat` | New functionality |
| `fix` | Bug fix |
| `docs` | Documentation |
| `test` | Tests |
| `refactor` | Internal restructuring without behavior change |
| `perf` | Performance improvement |
| `build` | Build system or dependencies |
| `ci` | Continuous integration |
| `chore` | Maintenance |

Commit descriptions should be written in English.

## Language Convention

Source code is written in English.

This includes:

- package names;
- variables;
- functions;
- types;
- interfaces;
- error messages;
- tests;
- documentation;
- branch names;
- commit messages.

Business concepts originally expressed in Portuguese should use a consistent
English representation in source code.

Examples:

```text
Ano Agrícola -> AgriculturalYear
Cultura      -> Culture
Safra        -> Crop
Produtor     -> Producer
Fazenda      -> Farm
Talhão       -> Field
```

## Code Organization

The project is organized primarily around business capabilities.

Example:

```text
internal/
└── agricultural/
    └── agriculturalyear/
```

Avoid creating large global packages such as:

```text
controllers/
services/
repositories/
models/
```

when doing so would mix unrelated business capabilities.

Infrastructure implementations may live below the business module when that
keeps the relationship explicit.

Example:

```text
internal/
└── agricultural/
    └── agriculturalyear/
        └── memory/
```

## Domain Code

Domain code should not depend directly on:

- HTTP;
- databases;
- Redis;
- message brokers;
- external frameworks;
- infrastructure-specific configuration.

Business behavior should remain testable without external infrastructure.

## Interfaces

Do not create an interface for every type.

Interfaces should represent useful boundaries.

Prefer defining interfaces close to the code that consumes them.

For example, an application use case that requires persistence may define a
small repository interface containing only the operations it needs.

## Business Dates

Business fields that represent only a calendar date should use the project's
shared date value object.

Location:

```text
internal/shared/date
```

Business dates represent:

```text
year
month
day
```

They must not expose unnecessary time-of-day or timezone semantics.

Technical timestamps may use `time.Time` when appropriate.

See ADR 0005.

## Error Handling

Errors must be handled explicitly.

Do not silently ignore errors.

Prefer errors that allow callers and tests to determine the failure category
when that distinction is useful.

Example:

```go
if !errors.Is(err, agriculturalyear.ErrNameAlreadyExists) {
    // handle unexpected error
}
```

Do not use panic for normal application errors.

## Context

Operations that may cross application or infrastructure boundaries should
accept `context.Context` when cancellation, deadlines, or request lifetime are
relevant.

Do not use `context.Context` as a generic container for business parameters.

## Testing

New behavior should normally include automated tests.

Current test categories include:

- domain unit tests;
- application unit tests;
- infrastructure tests;
- repository tests;
- internal integration tests.

Run all tests:

```bash
make test
```

Run the race detector:

```bash
make test-race
```

Run all project checks:

```bash
make check
```

## Formatting

Go source code must be formatted with `gofmt`.

Use:

```bash
make fmt
```

Verify formatting with:

```bash
make fmt-check
```

Do not manually format Go source to imitate `gofmt`.

## Static Analysis

Run:

```bash
make vet
```

before submitting changes.

## Build

Verify that the application builds:

```bash
make build
```

Generated binaries must not be committed.

The build output directory is:

```text
bin/
```

and must remain ignored by Git.

## Quality Check

Before committing a completed change, run:

```bash
make check
```

Then:

```bash
git diff --check
```

Both commands must succeed.

## Review Before Staging

Do not automatically stage everything without reviewing it first.

Inspect:

```bash
git status --short
git diff --stat
git diff
```

Verify that no unrelated files, credentials, local configuration, generated
binaries, or editor files are included.

## Staging

Prefer staging the files that belong to the change intentionally.

Example:

```bash
git add internal/agricultural/agriculturalyear
git add internal/shared/date
git add docs
```

Then inspect:

```bash
git status
git diff --cached
```

Only commit after reviewing the staged diff.

## Definition of Done

A change is considered complete when applicable requirements are satisfied:

- implementation is complete;
- code is formatted;
- `go vet` passes;
- automated tests pass;
- race detector passes;
- application builds;
- relevant documentation is updated;
- architectural decisions are documented when necessary;
- changelog is updated when appropriate;
- `git diff --check` passes;
- the final diff has been reviewed;
- CI passes.

## Architecture Decision Records

Important architectural decisions must be documented under:

```text
docs/decisions/
```

Create an ADR when a decision:

- affects multiple modules;
- introduces an important architectural constraint;
- selects a major technology;
- establishes a project-wide convention;
- would otherwise require future developers to guess why the decision was
  made.

Do not create ADRs for trivial implementation details.

## Documentation

Documentation should explain responsibilities, architecture, behavior, and
important decisions.

Avoid comments that merely translate obvious Go syntax into English.

Prefer useful GoDoc for exported APIs.

For example, prefer:

```go
// NewRepository creates an empty in-memory agricultural year repository.
func NewRepository() *Repository
```

over comments such as:

```go
// Increment nextID by one.
nextID++
```

Documentation should explain intent rather than repeat implementation.

## Security

Never commit:

- passwords;
- API keys;
- access tokens;
- private keys;
- production credentials;
- secrets;
- real `.env` files.

Use environment variables or an appropriate secrets-management mechanism.

Only safe examples belong in:

```text
.env.example
```

## Pull Request Checklist

Before opening a Pull Request, verify:

```text
[ ] The change belongs to the intended feature.
[ ] The code is formatted.
[ ] go vet passes.
[ ] Tests pass.
[ ] Race detector passes.
[ ] The application builds.
[ ] Documentation is updated.
[ ] ADRs are updated when required.
[ ] CHANGELOG is updated when appropriate.
[ ] git diff --check passes.
[ ] The final diff has been reviewed.
[ ] No credentials or local files are included.
```

## Continuous Integration

Pull Requests must pass the configured GitHub Actions quality checks before
integration.

CI currently executes:

```bash
make check
```

A successful local `make check` does not replace CI, but it reduces avoidable
failures after pushing.

## Architectural Evolution

AgroCore currently uses a modular monolith.

Do not introduce microservices merely to distribute code.

A service boundary should be introduced only when supported by concrete
business or operational requirements.

Likewise, infrastructure such as Redis, messaging, Kubernetes, and other
components should be introduced when the application has a concrete reason to
use them.

## Questions About Architecture

Before making a significant architectural change:

1. identify the concrete problem;
2. evaluate the simplest viable solution;
3. determine which modules are affected;
4. evaluate alternatives;
5. document the decision with an ADR when appropriate;
6. implement the smallest coherent change;
7. validate it with automated tests.