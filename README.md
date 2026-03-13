# `chill-institute/contracts`

Canonical protobuf API contracts and generated client artifacts for Chill.

This repository is the source of truth for the public RPC surface used by official Chill clients. It contains the protobuf schema plus generated artifacts for Go and TypeScript consumers.

This repository does not include the private backend implementation.

## What Is In This Repo

- `proto/` contains the protobuf source files.
- `gen/go/` contains generated Go protobuf and Connect RPC artifacts.
- `gen/ts/` contains generated JavaScript and TypeScript artifacts for web and mobile clients.
- `gen/openapi/` contains generated OpenAPI output derived from the protobuf schema.

## Versioning

- The API contract version is encoded in the protobuf package path, for example `chill.v4`.
- Repository releases use normal semver tags such as `v0.1.0`.
- Consumers should depend on tagged releases, not `main`.

## Go Usage

Generated Go artifacts live in this module and are intended to be imported directly:

```go
import chillv4 "github.com/chill-institute/contracts/gen/go/chill/v4"
import chillv4connect "github.com/chill-institute/contracts/gen/go/chill/v4/chillv4connect"
```

## TypeScript Usage

The TypeScript package is published as `@chill-institute/contracts`.

```bash
npm install @chill-institute/contracts
```

Then import generated artifacts directly:

```ts
import { UserService } from "@chill-institute/contracts/chill/v4/api_pb";
```

## Development

This repository uses [`mise`](https://mise.jdx.dev/) for local tool management and task execution.

```bash
mise install
mise run generate
mise run verify
```

Generation is centralized in this repository. Downstream clients should consume the generated artifacts from tagged releases instead of regenerating from copied schema files.

## Release Flow

1. Update `proto/` sources.
2. Run `mise run generate`.
3. Commit both the schema changes and generated artifacts.
4. Tag a release.
5. Publish the TypeScript package from the release workflow.

## Contributing

- Do not edit generated files by hand.
- Keep schema changes backward compatible unless you are intentionally shipping a breaking change.
- Run `mise run verify` before opening a pull request.

## License

Apache-2.0. See [LICENSE](./LICENSE).
