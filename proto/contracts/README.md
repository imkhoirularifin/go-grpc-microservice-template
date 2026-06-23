# Proto Contracts

This directory is intended to be a **git submodule** pointing to a centralized protobuf repository.

## Submodule setup

```bash
# Remove the example contracts (first-time setup only)
rm -rf proto/contracts

# Add your centralized proto repo pinned to a release tag
git submodule add -b v1.0.0 https://github.com/your-org/proto-contracts proto/contracts
git submodule update --init --recursive
```

## Versioning

- Tag releases in the proto repository using [Semantic Versioning](https://semver.org/) (`v1.0.0`, `v1.1.0`, …).
- Pin consuming services to a specific tag via the submodule branch (`-b v1.0.0`).
- Run `buf breaking` against the previous tag before publishing a new proto release.

## Layout

```
proto/contracts/
├── buf.yaml
├── common/v1/
│   └── pagination.proto
└── user/v1/
    └── user.proto
```

## Generate Go code

From the repository root:

```bash
make proto
```
