# Continuous Integration

AgroCore uses GitHub Actions to validate changes automatically.

## Purpose

Continuous Integration provides an automated quality gate for changes
submitted to the repository.

The CI workflow complements local development checks and helps prevent
invalid changes from being integrated into the main development branches.

## Triggers

The workflow runs for:

- pushes to `develop`;
- pushes to `main`;
- pull requests targeting `develop`;
- pull requests targeting `main`.

## Quality Checks

The CI pipeline performs the following checks:

1. source code formatting;
2. `go vet`;
3. automated tests;
4. Go race detector;
5. application build.

## Local Validation

Developers should run the equivalent checks before pushing changes:

```bash
gofmt -w .
go vet ./...
go test -count=1 ./...
go test -race -count=1 ./...
go build ./...
```

CI is not a replacement for local validation.

## Dependencies

The workflow uses GitHub Actions for repository checkout and Go environment
setup.

These Actions are CI infrastructure dependencies and are not runtime
dependencies of the AgroCore application.