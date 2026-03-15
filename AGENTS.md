# Contracts

`contracts` is the public contract boundary for `chill.institute`.

## Priorities

- Treat changes here as consumer-facing API changes.
- Keep protobuf sources, generated artifacts, and package metadata in sync.
- Make compatibility decisions deliberately and call out migration risk clearly.

## Workflow

- `mise install`
- `mise run generate`
- `mise run verify`

## Boundaries

- Source schemas live in `proto/`.
- Generated outputs live in `gen/go/`, `gen/ts/`, and `gen/openapi/`.
- This repo does not own backend, web, mobile, or CLI application logic.
- See [docs/ARCHITECTURE.md](./docs/ARCHITECTURE.md) for the public-boundary model and generation flow.
