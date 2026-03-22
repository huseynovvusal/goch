# Contributing to Goch 🤝

Thanks for your interest in contributing! This document explains how to set up your environment, make changes, and submit pull requests.

## Table of contents

- [Development setup](#development-setup)
- [Project structure](#project-structure)
- [Coding guidelines](#coding-guidelines)
- [Testing](#testing)
- [Commits and branches](#commits-and-branches)
- [Pull requests](#pull-requests)
- [Issues and feature requests](#issues-and-feature-requests)
- [Community standards](#community-standards)

## Development setup

Prerequisites:

- Go 1.20+
- Git
- `make`

Clone and build:

```sh
git clone https://github.com/huseynovvusal/goch.git
cd goch
make build
```

Run tests:

```sh
make test
```

## Project structure

```text
cmd/                 # CLI entrypoints
internal/
  chat/             # UDP chat send/receive, message model
  config/           # Ports and app configuration
  discovery/        # Presence broadcast + discovery
  tui/              # Bubble Tea TUI (model/view/update)
  utils/network/    # Network helpers, broadcast detection
```

## Coding guidelines

- Use standard Go formatting: `go fmt ./...`
- Keep functions small and clear
- Handle errors explicitly
- Add comments for exported functions and tricky logic
- Prefer dependency-free solutions where reasonable

## Testing

- Tests live next to code in the same package (`*_test.go`)
- Aim for happy paths + edge cases
- Keep tests deterministic; avoid network I/O in unit tests

Run all tests:

```sh
make test
```

## CI and automation

The repository has automated checks in GitHub Actions (`.github/workflows/ci.yml`):

- `go fmt` check (`gofmt -l`)
- `go vet` static checks
- `go test ./...` coverage collection
- `golangci-lint run ./...`
- `codecov` upload from `coverage.out`

Also, we use Dependabot to keep `go.mod` dependencies updated, configured in `.github/dependabot.yml`.

## Commits and branches

- Base new work off `main`: `git checkout -b feat/short-description`
- Commit messages: imperative mood
  - Example: `Add UDP chat listener`
  - Example: `Fix discovery channel timeout`

## Pull requests

- Target branch: `main`
- Add a clear summary + scope
- Include test plan and trouble reproduction steps
- Ensure CI passes before requesting review
- Smaller PRs are easier to review and merge

## Issues and feature requests

- Open GitHub issues for bugs and feature requests
- Include platform (OS), Go version, steps to reproduce, and expected behavior

## Community standards

Please see `CODE_OF_CONDUCT.md` for community behavior expectations.

Thank you for helping make Goch better! 🙏
