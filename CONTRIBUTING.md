Absolutely! Here’s a thoughtful starter for your `CONTRIBUTING.md` file, tailored for a project like `kicli`—modern Go, with modular structure, a focus on open collaboration, and a strong emphasis on clarity and privacy.

---

# Contributing to kicli

First off: **Thank you for taking the time to contribute!**

`kicli` welcomes issues, pull requests, and questions from anyone interested in AI-powered terminal workflows. Whether you're suggesting new features, fixing bugs, refining documentation, or improving code quality, your help makes a big difference.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
  - [Questions & Support](#questions--support)
  - [Filing Issues](#filing-issues)
  - [Suggesting Features](#suggesting-features)
  - [Pull Requests](#pull-requests)
- [Development Setup](#development-setup)
  - [Project Structure](#project-structure)
  - [Building](#building)
  - [Running](#running)
  - [Testing](#testing)
- [Coding Guidelines](#coding-guidelines)
  - [Go Style](#go-style)
  - [UI Consistency](#ui-consistency)
- [Commit Messages](#commit-messages)
- [License](#license)

---

## Code of Conduct

All contributors are expected to follow our [Code of Conduct](CODE_OF_CONDUCT.md).
Be respectful and inclusive!

---

## How to Contribute

### Questions & Support

- Try reading the [README](README.md) and [docs/](docs/) first.
- For general questions, feel free to [open a discussion](https://github.com/semidark/kicli/discussions) or use issues labeled `question`.

### Filing Issues

- Use [GitHub Issues](https://github.com/semidark/kicli/issues/new) for bugs or unexpected behavior.
- Include:
  - Steps to reproduce
  - Behavior you expected
  - kicli version, OS (and shell) info
- Add logs or screenshots when appropriate.

### Suggesting Features

- Check the [roadmap](docs/implementation-plan.md) first.
- Please describe:
  - What you'd like to see
  - Motivation or use case
  - Any relevant context/alternatives

### Pull Requests

We use feature branches and review all PRs.

1. [Fork](https://github.com/semidark/kicli/fork) and `git clone` your fork.
2. Create a new branch:
   ```sh
   git checkout -b my-feature
   ```
3. Implement your changes (see below).
4. Test locally.
5. Push to your fork and open a PR with a clear title and description.
6. Reference relevant issues (e.g., "Closes #42").

---

## Development Setup

### Project Structure

See [`docs/architecture.md`](docs/architecture.md) for details. Some key folders:
- `cmd/kicli`: main entry point
- `internal/app`: main Bubbletea model and app logic
- `internal/tui/`: view/layout code
- `internal/ptyhandler/`: PTY shell backend
- `internal/aiclient/`: LLM/OpenAI client
- `internal/configmanager/`: load configuration
- `internal/storage/`: SQLite history backend

### Building

```sh
go build ./cmd/kicli
```

### Running

Just run:

```sh
./kicli
```

(Requires a valid config at `~/.config/kicli/config.yaml` and LLM API credentials.)

### Testing

```sh
go test ./...
```

Unit tests are preferred for new logic. TUI integration tests will be added as coverage grows.

---

## Coding Guidelines

### Go Style

- Follow [Effective Go](https://go.dev/doc/effective_go) and idiomatic Go conventions.
- Use `go fmt` and organize imports.
- Prefer pure Go dependencies (`CGO_ENABLED=0`).
- Code should be modular—see component docs in [`docs/architecture.md`](docs/architecture.md).

### UI Consistency

- Use [Bubbletea](https://github.com/charmbracelet/bubbletea) models and idioms, including state-driven updates via `tea.Msg` and `tea.Cmd`.
- All styling in [`internal/tui/styles/`](internal/tui/styles/).
- Keep keybindings and layout consistent; new components should fit the 65%/35% split and keyboard navigation model.
- If in doubt, open an issue or PR draft to discuss UI ideas before major refactors.

---

## Commit Messages

- Aim for clarity: `fix`, `add`, `improve`, `doc`, or concise summary.
- Link issues with `Fixes #42` when relevant.

---

## License

By contributing, you agree your code will be released under the [MIT License](LICENSE).

---

Thank you for making `kicli` better!
  
*Questions? Suggestions? Feedback?* [Open an issue or discussion!](https://github.com/semidark/kicli/issues)

---