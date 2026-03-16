# Contributing

Thanks for contributing to `chill-institute-contracts`.

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

This repo owns:

- protobuf source files in `proto/`
- generated Go artifacts in `gen/go/`
- generated TypeScript artifacts in `gen/ts/`
- generated OpenAPI output in `gen/openapi/`

This repo does not own:

- backend implementation
- web, mobile, or CLI application logic
- deploy/runtime configuration

## Consumer Usage

Go consumers import generated artifacts directly:

```go
import chillv4 "github.com/chill-institute/chill-institute-contracts/gen/go/chill/v4"
import chillv4connect "github.com/chill-institute/chill-institute-contracts/gen/go/chill/v4/chillv4connect"
```

TypeScript consumers install `@chill-institute/contracts`:

```bash
npm install @chill-institute/contracts
```

Example import:

```ts
import { UserService } from "@chill-institute/contracts/chill/v4/api_pb";
```

## Release Flow

Normal contract change flow:

1. Update `proto/`.
2. Run `mise run generate`.
3. Review generated output in `gen/go`, `gen/ts`, and `gen/openapi`.
4. Run `mise run verify`.
5. Commit schema and generated output together.
6. Push `main`.
7. GitHub Actions runs `semantic-release`, bumps the version automatically, tags `vX.Y.Z`, publishes the package to npm, creates the GitHub release, and commits the updated `package.json` back to `main`.

Example:

```bash
git push origin main
```

Versioning notes:

- Releases are driven by conventional commits.
- `feat:` produces a minor release.
- `fix:`, `perf:`, `refactor:`, `docs:`, `test:`, `build:`, `ci:`, and `chore:` produce a patch release.
- Breaking changes produce a major release.
- Do not manually edit `package.json` just to cut a release.

## Pull Requests

- Explain the contract change clearly.
- Include regenerated artifacts in the same PR.
- Flag any compatibility or migration risk.
