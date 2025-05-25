# kicli

*A Go-powered TUI terminal with an integrated AI assistant.*

> **tmux + ChatGPT-like experience, in a single fast, keyboard-focused terminal UI.**

---

## Overview

`kicli` is a next-generation terminal application that combines a feature-rich PTY shell with an always-available LLM-powered AI chat assistant — running side-by-side in a modern, cross-platform TUI. Built from scratch in Go with [Bubbletea](https://github.com/charmbracelet/bubbletea), `kicli` aims to level-up your terminal with AI-driven command suggestions, contextual chat, and instant assistance, all within your familiar workflow.

- **No telemetry or tracking; fully offline, except your chosen LLM endpoint.**
- **Configurable, cross-language, and blazing fast.**

---

## Features

- **Integrated PTY shell.**  Use your preferred POSIX shell as usual.
- **AI command suggestions.**  Get shell commands generated/recommended based on chat and terminal history/context.
- **Context-aware AI chat.**  Interact with an OpenAI-compatible LLM; shell history and buffer are seamlessly included as context.
- **Keyboard-driven TUI.**  Fast, mouse-free, with configurable keybindings.
- **Modern UI.**  Scrollable multi-pane design using Bubbletea and Lipgloss.
- **Private and secure.**  Zero telemetry, zero outbound calls except user-specified LLM endpoint.
- **SQLite-backed history.**  All shell and chat history is stored locally (CGO-free).
- **Cross-platform.**  Linux and macOS (Windows support planned).

---

## Screenshot

*(Coming soon)*

---

## Getting Started

### 1. **Requirements**

- Go 1.18+ (pure Go dependencies, no CGO)
- Linux or macOS (Windows support planned)

## Dependencies

kicli relies on the following major libraries:

- [Bubbletea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) — Styling
- [creack/pty](https://github.com/creack/pty) — PTY shell support
- [modernc.org/sqlite](https://github.com/cznic/sqlite) — Pure-Go SQLite backend
- [sashabaranov/go-openai](https://github.com/sashabaranov/go-openai) — OpenAI API client
- [glamour](https://github.com/charmbracelet/glamour) — Markdown rendering
- [bubbles](https://github.com/charmbracelet/bubbles) — TUI components
- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) — YAML
- [adrg/xdg](https://github.com/adrg/xdg) — XDG path handling)

Full dependency details: see [`go.mod`](go.mod).

### 2. **Build and Run**

```sh
git clone https://github.com/semidark/kicli.git
cd kicli
go build ./cmd/kicli
./kicli
```

### 3. **Configuration**

- Copy the sample config to your XDG config directory:
  ```sh
  mkdir -p ~/.config/kicli
  cp assets/config.yaml ~/.config/kicli/config.yaml
  ```
- Configure your LLM API Key and endpoint in `config.yaml` or via environment variables (`KICLI_API_KEY`, `KICLI_API_URL`).
  > You must **explicitly configure** the LLM endpoint and API key you wish to use, either in `kicli`'s `config.yaml` or with the kicli-specific `KICLI_API_URL` / `KICLI_API_KEY` environment variables. We won't automatically pick up any environment LLM API related variables from your shell.
  >
  > This is a conscious design decision to give you the power and responsibility to choose **where your data goes, and who has access to it**.
- See [docs/configuration.md](docs/configuration.md) for all details.

---

- **Customizable key bindings:** See [docs/keybindings.md](docs/keybindings.md)

---

## Privacy

- No telemetry or analytics.
- Only outbound API calls are to your configured LLM endpoint.
- Data for chat and executed commands is stored _locally_ in SQLite.
- For details, see [docs/privacy.md](docs/privacy.md).

---

## Documentation

- [Architecture & Design](docs/architecture.md)
- [Configuration Reference](docs/configuration.md)
- [Implementation Roadmap & Progress](docs/implementation-plan.md)
- [Keybindings & Customization](docs/keybindings.md)
- [Privacy & Data Handling](docs/privacy.md)
- [Storage Details](docs/storage.md)
- [Contributing Guide](CONTRIBUTING.md)

---

## Roadmap

- [ ] **v0.1.0**: Core terminal & foundation (PTY shell, config, history, split TUI)
- [ ] **v0.5.0**: AI features & command suggestion (chat pane, safe command execution)
- [ ] **v0.9.0**: Streaming, error handling, polish (robust UX, documentation)
- [ ] **v1.0.0**: Production-ready release (stability, security, comprehensive docs)
- [ ] **v1.1.0+**: Advanced features (Windows support, theming, plugins, extensibility)

See [Implementation Plan](docs/implementation-plan.md) for detailed milestones and deliverables.

---

## License

MIT — see [LICENSE](LICENSE)
