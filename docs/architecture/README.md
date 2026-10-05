# Architecture

This directory contains the architectural documentation for AgroCore.

## Current Architecture

AgroCore is currently in its initial bootstrap phase.

The application consists of a single Go executable and does not yet
implement HTTP transport, persistence, domain modules, or external
infrastructure.

```text
main.go
   |
   v
Application startup