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

## Development setup

Prerequisites:

- Go 1.20+
- Git

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
cmd/                 # CLI entrypoints (Cobra)
internal/
  chat/             # UDP chat send/receive, message model
  config/           # Ports and app configuration
  discovery/        # Presence broadcast + discovery
  tui/              # Bubble Tea TUI (model/view/update)
  utils/network/    # Network helpers, broadcast detection
```

## Coding guidelines

- Use standard Go formatting: `go fmt ./...`
- Keep functions small and focused; prefer clarity over cleverness
- Handle errors explicitly; avoid ignoring errors
- Add comments/docstrings for exported functions and tricky code
- Prefer dependency-free solutions where practical

Optional (if installed):

```sh
# Static checks
go vet ./...
```

## Testing

- Put tests next to code in the same package; files end with `_test.go`
- Aim for happy path + 1–2 edge cases
- Keep tests deterministic; avoid network calls when unit testing helpers

Run all tests:

```sh
make test
```

## Commits and branches

- Create a feature branch from `main`: `git checkout -b feat/short-description`
- Write concise commit messages in the imperative mood
  - Example: "Add UDP chat listener" / "Fix broadcast address detection"

## Pull requests

- Open PRs against `main`
- Include a clear description of the change and screenshots for TUI changes
- Ensure `make test` passes
- Keep PRs focused and reasonably small; large PRs are harder to review

## Issues and feature requests

- Use GitHub Issues to report bugs or propose features
- When reporting a bug, include steps to reproduce, expected vs actual behavior, and your OS/Go version

Thank you for helping make Goch better! 🙏
