# ADR 0001: Use Go as the Backend Language

## Status

Accepted

## Context

AgroCore is a backend application for agricultural management.

The project requires a language suitable for building network services,
REST APIs, concurrent workloads, containerized applications, and
cloud-native services.

The project also aims to maintain a simple build and deployment process.

## Decision

Go will be used as the primary backend programming language.

## Consequences

### Positive

- Statically typed language.
- Compiled binaries.
- Strong standard library for network services.
- Built-in concurrency primitives.
- Simple dependency management through Go Modules.
- Good ecosystem for cloud-native applications.

### Negative

- The development team must maintain proficiency in Go.
- Some abstractions commonly provided by full-stack frameworks must be
  implemented or selected explicitly.