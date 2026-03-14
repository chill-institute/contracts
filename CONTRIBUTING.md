# Contributing

Thanks for contributing to `chill-institute/contracts`.

## Setup

Install the toolchain:

```bash
mise install
```

Regenerate artifacts after changing protobuf definitions:

```bash
mise run generate
```

## Validation

Run the full repo checks before opening or updating a pull request:

```bash
mise run verify
```

## Development Notes

- This repo is the public contract boundary for `chill.institute`.
- Schema changes should be deliberate and consumer-aware.
- Keep generated artifacts in sync with schema changes.
- Consumers should depend on releases, not `main`.

## Pull Requests

- Explain the contract change clearly.
- Include regenerated artifacts in the same PR.
- Flag any compatibility or migration risk.
