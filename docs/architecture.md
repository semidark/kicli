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
- [Privacy & Telemetry](#privacy--telemetry)
- [Roadmap](#roadmap)

---

## Goals & Vision

- **Powerful TUI terminal:** Rich terminal experience on the command line, using a native shell and a fast, responsive interface.
- **Integrated AI assistant:** Embedded LLM chat and contextual command generation, fully user-configurable.
- **Keyboard-centric workflow:** Every interaction, navigation, and action can be performed via the keyboard.
- **Private by default:** Zero telemetry; only user-configured LLM endpoints receive data.
- **Open, modular, and extensible:** Easy to understand, contribute, and maintain.

---

## Success Criteria

kicli 1.0 is considered **done** when:

- Runs on Linux and macOS (Windows support after 1.0)
- Launches a user default POSIX shell in a PTY and displays it with proper rendering and scrolling
- Embeds an AI chat pane; contextually suggests shell commands and answers user questions
- AI commands are never auto-executed—user confirmation is always required
- All chat and shell history are stored (unencrypted) in a local SQLite database
- Configuration follows XDG conventions and supports environment variable overrides
- No telemetry; only the user-configured LLM endpoint is ever contacted

---

## Non-goals & Scope

- Not an IDE or terminal multiplexer (but inspired by tmux layouts)
- No SSH client, remote connections, or graphical file editors
- LLM inference is **not** performed locally; kicli is a client for OpenAI-compatible endpoints only
- All dependencies must support `CGO_ENABLED=0`; SQLite and all others are pure Go libraries

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

- **Left 65%:** PTY shell output and keyboard input, scrollable, full shell session.
- **Right 35%:** AI chat, scrollable conversation with LLM.
- **Bottom row (left):** "Command Generation Field" — editable, shows suggested shell commands from AI; only executes on user confirmation.

**Keyboard Navigation:**

- Switch focus: `Ctrl+→` / `Ctrl+←`
- Scroll panes: `Ctrl+↑` / `Ctrl+↓`
- Enter submits input in focused text field
- `Esc` often restores focus or clears input

More details: [docs/keybindings.md](keybindings.md)

---

## Core Architectural Principles

1. **State-driven UI:** All visible state is modeled in a single Bubbletea Model (`KicliModel`). Updates are handled via messages (`tea.Msg`), and the UI is redrawn accordingly.
2. **Strict modularity:** Each of the major domains (PTY I/O, AI client, configuration, storage) is implemented as its own package with clean interfaces.
3. **Concurrency:** PTY I/O and LLM API calls run in goroutines, communicating via channels and translated into Bubbletea messages for UI/reactivity.
4. **Consistent theming:** All colors, styles, and UI layout logic are centralized (Lipgloss/Styles).
5. **Privacy-first:** No analytics, background beacons, or hidden outbound connections.
6. **Portable, vendor-free Go:** Only pure-Go dependencies allowed (see [implementation plan](implementation-plan.md)).

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
├── cmd/kicli/main.go          # Application entry point
├── internal/
│   ├── app/                   # Main Bubbletea model and base logic
│   ├── tui/
│   │   ├── views/             # View rendering logic (panes, inputs, layouts)
│   │   └── styles/            # Central Lipgloss styles and themes
│   ├── ptyhandler/            # PTY session & I/O management
│   ├── aiclient/              # OpenAI (or compatible) API client logic
│   ├── configmanager/         # YAML & environment variable config loader
│   ├── storage/               # SQLite history backend
│   └── util/                  # General-purpose helpers
├── assets/                    # Default templates, such as config.yaml
└── docs/                      # Documentation (this file, others)
```

---

## Component Design

### 1. `KicliModel` (Main State Model)

Resides in `internal/app/model.go`. Holds all relevant state:

- PTY manager instance (`ptyManager`)
- AI client instance (`aiClient`)
- Config (`config`)
- Database/storage backend (`db`)
- Viewport models for shell and AI chat panes
- Text input models for AI chat and command generation
- App focus state, window size
- Loading states, spinners, error messages

#### Message-Driven Updates

All I/O (PTY output, AI responses, window resizing) are sent via custom `tea.Msg` types, ensuring state changes and side effects are handled synchronously within the Bubbletea update loop.

### 2. PTY Management (`internal/ptyhandler/`)

- Launches and supervises the user's default shell in a PTY
- Reads and writes bytes over the PTY, maintaining a virtual scrollback buffer for context
- Communicates via channels and custom messages
- Handles PTY resizing and forwarding of shell output

### 3. AI Client (`internal/aiclient/`)

- Communicates with any OpenAI API-compatible endpoint (OpenAI, LiteLLM, Ollama, etc.)
- Builds context-rich prompts from chat history, terminal buffer, and user request
- Supports streaming responses for efficient, dynamic AI chat
- Never sends data anywhere except the configured API endpoint

### 4. Configuration Management (`internal/configmanager/`)

- Loads and validates configuration from XDG paths (`$XDG_CONFIG_HOME`, or fallback)
- Supports YAML files and environment variable overrides for secrets (API keys, endpoint URLs)
- Holds all theming, keybindings, and AI model settings

### 5. History & Storage (`internal/storage/`)

- All chat messages and executed commands are stored in a **local, unencrypted SQLite DB**
- No history leaves the machine, unless sent to the LLM as prompt context
- Clean schema for easy querying and future extensibility

### 6. TUI Rendering & Styles (`internal/tui/`)

- Split between `views/` (render logic for each pane or input) and `styles/` (centralized color and layout definitions)
- Uses [Lipgloss](https://github.com/charmbracelet/lipgloss) for modern, type-safe terminal UI styling

---

## Privacy & Telemetry

- **No telemetry sent—ever.** The only outbound network calls are to the user-supplied LLM endpoint.
- All local data (chat and shell history) lives in unencrypted SQLite, under XDG data dir.
- Only information the user enters, and shell/AI context intentionally passed by user, is sent to LLM endpoints.

For more, see: [docs/privacy.md](privacy.md)

---

## Roadmap

- See [docs/implementation-plan.md](implementation-plan.md) for a detailed, phased technical plan.
- Major upcoming areas:
  - Windows (ConPTY) support
  - Extensible plugin system
  - In-TUI config editing
  - Richer theming and accessibility

---

## Diagrams

*(You can add mermaid diagrams or ASCII overviews here as the design evolves!)*

---

## References

- [Bubbletea](https://github.com/charmbracelet/bubbletea)
- [Lipgloss](https://github.com/charmbracelet/lipgloss)
- [creack/pty](https://github.com/creack/pty)
- [modernc.org/sqlite](https://github.com/cznic/sqlite)

---

_For more details on usage, development, configuration, and contribution, see the root [README.md](../README.md) and the [docs/](./) folder._
