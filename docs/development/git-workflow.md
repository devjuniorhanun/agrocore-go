# Git Workflow

AgroCore uses Git Flow for branch management and Conventional Commits
for commit messages.

## Permanent Branches

### `main`

Contains production-ready versions of the application.

Direct development on this branch is not allowed.

### `develop`

Integration branch for completed features that will be included in a
future release.

Feature development must not occur directly on this branch.

## Feature Branches

Feature branches are created from `develop`.

Create a feature:

```bash
git flow feature start <feature-name>
```

Example:

```bash
git flow feature start http-server
```

Finish a feature:

```bash
git flow feature finish <feature-name>
```

## Release Branches

Create a release:

```bash
git flow release start 0.1.0
```

Finish a release:

```bash
git flow release finish 0.1.0
```

## Hotfix Branches

Hotfix branches are created from `main`.

Create a hotfix:

```bash
git flow hotfix start 0.1.1
```

Finish a hotfix:

```bash
git flow hotfix finish 0.1.1
```

## Conventional Commits

Commit messages follow the Conventional Commits convention.

Common types:

- `feat`: new functionality
- `fix`: bug fix
- `docs`: documentation changes
- `test`: test changes
- `refactor`: code changes without changing behavior
- `perf`: performance improvements
- `build`: build system or dependency changes
- `ci`: continuous integration changes
- `chore`: maintenance tasks

Examples:

```text
feat: add health check endpoint
fix: validate empty product name
docs: document Git Flow strategy
test: add product service tests
refactor: extract product repository
chore: configure development environment
```

## Development Rules

- Do not develop directly on `main`.
- Do not develop features directly on `develop`.
- Create feature branches from `develop`.
- Keep commits small and cohesive.
- Never commit credentials or environment secrets.
- Update tests when behavior changes.
- Update documentation when architecture or behavior changes.
- Use English for code, branches, commits, and technical documentation.

## Quality Checks

Before committing Go code, run:

```bash
gofmt -w .
go vet ./...
go test ./...
```

The commands must complete successfully before the change is committed.