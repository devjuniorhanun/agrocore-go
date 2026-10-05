# ADR 0003: Use the Go Standard Library for the Initial HTTP Server

## Status

Accepted

## Context

AgroCore requires an HTTP server to expose its API.

Several third-party HTTP routers and frameworks are available in the Go
ecosystem. However, the initial API has simple routing requirements that can
be satisfied by the Go standard library.

Introducing an external dependency at this stage would increase the
dependency surface without solving a current requirement.

## Decision

The initial HTTP server will use Go's standard `net/http` package.

Third-party routers or HTTP frameworks may be evaluated later if concrete
routing or middleware requirements justify their adoption.

## Consequences

### Positive

- No external HTTP dependency.
- Smaller dependency surface.
- Direct use of Go HTTP primitives.
- Easier understanding of the application's HTTP lifecycle.
- Reduced framework coupling.

### Negative

- Some higher-level routing and middleware conveniences must be implemented
  explicitly.
- Future requirements may justify migrating to or introducing a third-party
  router.