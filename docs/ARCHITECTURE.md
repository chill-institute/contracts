# Architecture

This document describes how `chill-institute-contracts` is built.

## System Context

```mermaid
graph LR
  Proto["protobuf schemas"] --> Generated["generated artifacts"]
  Generated --> Web["web"]
  Generated --> CLI["cli"]
  Generated --> Backend["backend"]
  Generated --> Future["future clients"]
```

## Components

| Component | Responsibility | Talks to |
|-----------|----------------|----------|
| `proto/` | Canonical public API schemas | Buf, generators |
| `gen/go/` | Generated Go contract types | backend consumers |
| `gen/ts/` | Generated TypeScript contract types | `web`, future JS clients |
| `gen/openapi/` | Generated OpenAPI output | docs and tooling |
| package metadata | Publish the TypeScript package and release artifacts | npm, GitHub releases |

## Generation Model

```mermaid
graph TD
  Author["edit proto"] --> Generate["mise run generate"]
  Generate --> Go["gen/go"]
  Generate --> TS["gen/ts"]
  Generate --> OpenAPI["gen/openapi"]
  Go --> Backend["backend consumers"]
  TS --> Web["web consumers"]
```

## Boundaries

- This repo owns public contracts only.
- Backend behavior, web UX, and CLI command surfaces do not live here.
- Compatibility decisions here are consumer-facing and should be treated like API changes.

## Release Model

- `main` is the release branch.
- CI verifies the repo on push.
- release automation publishes the npm package and GitHub release from `main`
