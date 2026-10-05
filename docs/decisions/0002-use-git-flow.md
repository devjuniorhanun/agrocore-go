# ADR 0002: Use Git Flow as the Branching Strategy

## Status

Accepted

## Context

AgroCore requires a predictable workflow for developing features,
preparing releases, and applying production fixes.

The repository will maintain a stable production branch while allowing
ongoing development to continue independently.

## Decision

The project will use Git Flow as its branching strategy.

The permanent branches are:

- `main`: production-ready code.
- `develop`: integration branch for ongoing development.

Supporting branches are:

- `feature/*`
- `release/*`
- `hotfix/*`
- `bugfix/*`

Version tags will use the `v` prefix.

Example:

```text
v0.1.0
v1.0.0