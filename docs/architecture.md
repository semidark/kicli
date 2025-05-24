Absolutely! Here is a **complete, expanded** `docs/architecture.md` for **kicli**, with a rich **Component Design** section added. This version is organized, clear, and ready for contributors/maintainers.

---

# kicli Architecture

## Overview

**kicli** is a fast, keyboard-driven terminal application that seamlessly integrates a PTY shell and an LLM-powered AI assistant within a modern, keyboard-centric TUI. Its mission is to provide effortless access to AI-enhanced command-line workflows—without sacrificing speed, privacy, or the classic shell user experience.

The design is guided by modularity, security, and ease of contribution. kicli leverages modern Go libraries such as [Bubbletea](https://github.com/charmbracelet/bubbletea) (for TUI) and [modernc.org/sqlite](https://github.com/cznic/sqlite) (for CGO-free local storage).


---

## Table of Contents

- [Goals & Vision](#goals--vision)
- [Success Criteria](#success-criteria)
- [Non-goals & Scope](#non-goals--scope)
- [UI Layout & Navigation](#ui-layout--navigation)
- [Core Architectural Principles](#core-architectural-principles)
- [Dependencies](#dependencies)
- [Directory & Code Structure](#directory--code-structure)
- [Component Design](#component-design)
  - [Main Bubbletea Model (KicliModel)](#main-bubbletea-model-kiclimodel)
  - [PTY Handler](#pty-handler)
  - [AI Client](#ai-client)
  - [Configuration Manager](#configuration-manager)
  - [Storage Layer](#storage-layer)
  - [TUI Views & Styles](#tui-views--styles)
  - [Utilities](#utilities)
- [Privacy & Telemetry](#privacy--telemetry)

---

## Goals & Vision

- **Powerful TUI terminal:** Classic shell experience, with modern navigation and rendering.
- **Integrated AI assistant:** LLM chat and AI command suggestions, AI context from your real terminal session.
- **Keyboard-centric workflow:** 100% operable with keyboard; user-configurable keybindings.
- **Private by default:** Zero telemetry; only the configured LLM endpoint is ever contacted.
- **Open, modular, & extensible:** Codebase and documentation designed for clarity and easy contribution.

---

## Success Criteria

kicli 1.0 is considered **done** when:

- Works on Linux and macOS (Windows support after 1.0)
- Launches a PTY shell, fully scrollable/viewable with Bubbletea
- AI chat pane, command suggestion, and command confirmation implemented
- No AI-suggested code is executed without explicit user confirmation
- All history is stored locally, CGO-free (in SQLite)
- Config files and history follow XDG
- Zero telemetry; no external calls except your chosen LLM

---

## Non-goals & Scope

- Not an IDE, SSH client, or terminal multiplexer (inspired by tmux, but not aiming to replace it)
- No full local LLM inference; only acts as a client (API interface) to OpenAI API-compatible endpoints
- No non-Go dependencies requiring CGO (SQLite and others via pure Go)

---

## UI Layout & Navigation

```
┌───────────────────────────── 65% ──────────────────────┬──────── 35% ──────┐
│                          PTY Shell (scrollable)        │      AI Chat      │
│                                                        │   (scrollable)    │
├────────────────────────────────────────────────────────┤                   │
│  Command Generation Field (1 row, text input)          │                   │
└────────────────────────────────────────────────────────┴───────────────────┘
```

- **Left 65%:** PTY shell session, scrollable, interactive.
- **Right 35%:** AI chat: rich chat interface, markdown renderings.
- **Bottom row (left):** Command Generation Field; explicit "edit then execute" for LLM-suggested commands.

**Navigation:**

- Pane Focus: `Ctrl+←` / `Ctrl+→`
- Scroll: `Ctrl+↑` / `Ctrl+↓`
- Enter: submit input in current field
- `Esc`: clear/cancel field or revert focus

See [keybindings.md](keybindings.md) for more.

---

## Core Architectural Principles

1. **State-driven UI:** A single Bubbletea model represents and synchronizes all UI state.
2. **Strict modularity:** Clear separation of shell, AI, storage, configuration, and UI rendering.
3. **Concurrent but safe:** All blocking I/O (AI and PTY) handled in goroutines, with results delivered via channels/messages and applied via the Bubbletea update loop.
4. **Centralized, type-safe theming:** All colors and alignment/styles in one place.
5. **Privacy-first:** Only connects to endpoints controlled by the user.
6. **CGO-free Go:** All dependencies support `CGO_ENABLED=0`.
---

## Dependencies

### Core Libraries

| Library                      | Purpose                                      |
|------------------------------|----------------------------------------------|
| Bubbletea                    | Terminal UI (TUI) framework                  |
| Lipgloss                     | Styling and layout for TUI                   |
| Bubbles                      | TUI components for inputs, viewports, etc.   |
| creack/pty                   | PTY spawn and shell communication           |
| modernc.org/sqlite           | CGO-free embedded SQLite                     |
| sashabaranov/go-openai       | OpenAI (and compatible) API client           |
| glamour                      | Markdown terminal rendering                  |
| gopkg.in/yaml.v3             | YAML config parsing                          |
| adrg/xdg                     | XDG-compliant config/data location (optional)|

Why: All selected for pure-Go, cross-platform capability (**CGO is never required**).

See [`go.mod`](../go.mod) for all direct and indirect dependencies.

---

## Directory & Code Structure

```
kicli/
├── cmd/kicli/main.go          # Main application entry point
├── internal/
│   ├── app/                   # Main Bubbletea model and application logic
│   ├── tui/
│   │   ├── views/             # Layout/rendering code for shells, panes, inputs
│   │   └── styles/            # Lipgloss style definitions and themes
│   ├── ptyhandler/            # Creack/pty wrapper; shell process management
│   ├── aiclient/              # OpenAI-compatible LLM API client implementation
│   ├── configmanager/         # Config file and environment variable handling
│   ├── storage/               # SQLite local history backend
│   └── util/                  # Utilities, helpers, error wrappers
├── assets/                    # Default config templates, sample files
└── docs/                      # This and other documentation
```

---

## Component Design

### Main Bubbletea Model (`KicliModel`)

**File:** `internal/app/model.go`

- Represents the **entire application state**.
- Composed of:
  - `ptyManager` — interface for PTY shell I/O
  - `aiClient` — interface for LLM chat and command suggestion
  - `config` — effective configuration
  - `db` — SQLite-backed history interface
  - UI state: viewports, text inputs (AI chat, command gen)
  - Current focus pane, error/loading state, spinners, window size
- Handles ALL events via [Bubbletea](https://github.com/charmbracelet/bubbletea) update loop:
  - User keys and navigation
  - PTY output (as messages)
  - AI responses (as messages, streaming and complete)
  - Scroll actions and resize events
  - Message types include: `PtyOutputMsg`, `PtyExitedMsg`, `AIResponseChunkMsg`, `CommandGeneratedMsg`, `HistoryLoadedMsg`, etc.
- *Design ensures concurrency and UI state are synchronized and race-free.*

---

### PTY Handler

**Directory:** `internal/ptyhandler/`

- Manages spawning, writing to, and reading from the user shell (bash, zsh, etc.) over a PTY (creack/pty)
- Maintains an internal scrollback buffer of shell output for context/AIs
- Canonical interface:
  - `Start(shellCmd, rows, cols)`
  - `Write([]byte)`
  - `ReadChan() <-chan []byte` // emits shell output lines/data
  - `SetSize(rows, cols)`
  - `GetVisibleBuffer() string` (for AI context)
  - Handles process exit/cleanup and errors
- Interacts with the Bubbletea model only via async channels and message passing

---

### AI Client

**Directory:** `internal/aiclient/`

- Speaks with any OpenAI API-compatible endpoint, using HTTP(S)
- Supports both chat-completion and command-suggestion endpoints
- Provides both *blocking* and *streaming* interfaces to the model:
  - `StreamChatCompletion(prompt, history, terminalContext)`
  - `StreamCommandSuggestion(description, chatHistory, terminalContext)`
  - `CancelRequest(requestID)`
- Streams results over Go channels, integrating with Bubbletea updates
- Handles context preparation: collecting recent chat logs from history, current shell buffer from PTY for in-context chat/command generation
- No automatic code execution; AI responses are suggestions until user-accepted

---

### Configuration Manager

**Directory:** `internal/configmanager/`

- Loads YAML config from XDG path (`$XDG_CONFIG_HOME/kicli/config.yaml`; falls back on `~/.config/kicli/config.yaml` )
- Supports override of critical options (like API key, endpoint) via environment variables (`KICLI_API_KEY`, `KICLI_API_URL`)
- Defines effective config in a structured Go type (`AppConfig`), including:
  - AI server parameters (model, endpoint, streaming)
  - Shell binary path/name
  - Keybindings, theme, UI preferences
- Provides helpers: load, save, validate config, resolve XDG paths

---

### Storage Layer

**Directory:** `internal/storage/`

- Implements unencrypted SQLite3-based storage for all:
  - Chat messages and turn history (role, text, timestamps)
  - Shell command history (command, timestamps)
  - Session management (in future versions)
- API is sync and exposed via a simple Go interface (`HistoryStore`):
  - `AddChatMessage`, `GetChatMessages`, `AddShellCommand`, `GetShellCommands`, `Init`
- Uses [modernc.org/sqlite](https://github.com/cznic/sqlite) for zero-C, CGO-free SQLite
- Database path follows XDG data specification

---

### TUI Views & Styles

**Directory:** `internal/tui/`

- `views/`  — rendering/layout of terminal panes, chat, command input using Bubbletea primitives
- `styles/` — centralized [Lipgloss](https://github.com/charmbracelet/lipgloss) color palette, spacing, alignment
- Implements 65%/35% vertical split, Markdown rendering (via [glamour](https://github.com/charmbracelet/glamour)) in AI pane, and consistent navigation/focus cues

---

### Utilities

**Directory:** `internal/util/`

- Small helpers: error wrapping, XDG path resolution, decoding, etc.
- Centralizes any utility functionality that might be shared across modules

---

## Privacy & Telemetry

- **By default, kicli sends zero telemetry.**
- No usage or error analytics are collected or transmitted.
- The *only* outbound network calls are made to the user-configured LLM endpoint for explicit chat or command requests.
- All history is stored locally in unencrypted SQLite, under your XDG data dir.
  - AI context (shell output, chat history) sent to LLM is always user-mediated, and behavior is fully documented.

See [privacy.md](privacy.md) for details.

---

_For user documentation, dev setup, configuration, and contribution, see the root [README.md](../README.md) and the [docs/](./) folder._

---