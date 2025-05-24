# Implementation Plan

This document outlines the phased implementation roadmap for **kicli**, including goals, deliverables, and version targets for each stage of the project.

---

## Table of Contents

- success criteria and non-goals
- [Roadmap Structure](#roadmap-structure)
- [Phase 1 (v0.1.0): Core Terminal & Foundation](#phase-1-v010-core-terminal--foundation)
- [Phase 2 (v0.5.0): AI Features & Command Suggestion](#phase-2-v050-ai-features--command-suggestion)
- [Phase 3 (v0.9.0): Streaming, Error Handling, Polish](#phase-3-v090-streaming-error-handling-polish)
- [Phase 4 (v1.0.0): Production-Ready Release](#phase-4-v100-production-ready-release)
- [Phase 5 (v1.1.0+): Advanced Features & Extensibility](#phase-5-v110-advanced-features--extensibility)
- [Contribution Flow](#contribution-flow)
- [Checklist Summary](#checklist-summary)
- [Known Non-Goals / Deferrals](#known-non-goals--deferrals)

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


## Roadmap Structure

- **Pre-v1.0** versions (~v0.x.y) provide incrementally usable milestones for contributors and alpha testers.
- **v1.0.0**: First production-ready, stable release.
- **v1.1.0+**: Ongoing improvements and community-driven expansion.

---

## Phase 1 (v0.1.0): Core Terminal & Foundation

Establish a minimal working skeleton with core shell functionality, config, and history.

- [ ] Project scaffolding: directory layout (`cmd/`, `internal/`, etc.), `go.mod`
- [ ] Configuration: XDG YAML config, environment variable overrides
- [ ] Storage: SQLite (`modernc.org/sqlite`), `HistoryStore` for shell commands
- [ ] Main Bubbletea model (`KicliModel`) with config & storage integration
- [ ] PTY handling: launch shell, I/O, scrolling; `ptyhandler` package
- [ ] TUI layout: 65%/35% split, placeholder AI/chat, command input field
- [ ] Keyboard navigation: Focus switching via `Ctrl+← / →`
- [ ] Display error/status area

**Deliverable for v0.1.0:**  
A usable TUI shell with working config file, local history, split-pane UI scaffolding, and navigation.

---

## Phase 2 (v0.5.0): AI Features & Command Suggestion

Introduce essential AI Assistant functionality and full command suggestion workflow.

- [ ] AI API client (`aiclient`): OpenAI-compatible endpoint with config, key, and model
- [ ] AI chat pane: input, output, Markdown rendering (Glamour)
- [ ] Command suggestion: context sent to LLM, response placed in Command Generation Field
- [ ] No auto-execution: explicit user confirmation required to run suggestions
- [ ] Full chat and shell command logging to SQLite per session
- [ ] Rich keyboard navigation: switch, focus, and visual cues for panes

**Deliverable for v0.5.0:**  
TUI with integrated AI chat and safe command suggestion, storing chat & command history, and robust navigation.

---

## Phase 3 (v0.9.0): Streaming, Error Handling, and UI Polish

Deliver a near-feature-complete, smooth, robust user experience before release.

- [ ] Streaming AI (chunked output, incremental display, spinner/loader UI)
- [ ] Smarter scrolling: scroll-to-bottom on update, user manual scroll preserved
- [ ] Enhanced error/status feedback and in-TUI error recovery
- [ ] Cancel/correct: abort AI calls, clear inputs with Esc, handle timeouts gracefully
- [ ] Command field enhancements: editable AI suggestions, command history browsing
- [ ] Full configuration polish: keymaps/themes/defaults, XDG path enforcement
- [ ] Documentation: config example and basic usage in `docs/`

**Deliverable for v0.9.0:**  
A "beta-quality" version with reliable streaming AI, polished interaction, self-recovering UI, and complete local docs ready for public feedback.

---

## Phase 4 (v1.0.0): Production-Ready Release

Finalize stability, security, and documentation for a robust 1.0.

- [ ] Comprehensive documentation set ([README.md](../README.md), [docs/], config/sample)
- [ ] Final bug-squashing, UI/UX refinement
- [ ] Complete test coverage of core features (unit/integration)
- [ ] Release notes and upgrade/migration instructions (if needed)
- [ ] Tag and announce v1.0.0

**Deliverable for v1.0.0:**  
First stable version, suitable for daily use, distribution, and wider adoption.

---

## Phase 5 (v1.1.0+): Advanced Features & Extensibility

Future growth: power-user improvements, ecosystem features, and community requests.

- [ ] Windows (`ConPTY`) support and cross-platform TTY quirks
- [ ] Fully configurable theming (colors, light/dark, keymaps, etc.)
- [ ] TUI-based settings editor (change config inside kicli)
- [ ] Clipboard integration: copy code blocks (if platform supports)
- [ ] Advanced session/conversation management (multiple/switchable sessions, erase)
- [ ] Early plugin/extensibility hooks (custom commands/context injectors)
- [ ] Onboarding/welcome views and in-app docs
- [ ] Enhanced logging and debug modes

**Deliverable for v1.1.0+:**  
Rich, extensible, and even more flexible kicli experience for advanced users and contributors.

---

## Contribution Flow

- Track status with checkboxes in this file.
- Open issues for features, bugs, and clarifications.
- Always branch from latest `main`, follow naming conventions.
- Submit PRs with clear, descriptive titles; link to relevant issues.
- All code must be formatted (`go fmt`) and include tests if adding non-trivial behavior.

See [CONTRIBUTING.md](../CONTRIBUTING.md).

---

## Checklist Summary

| Milestone                                  | v0.1.0 | v0.5.0 | v0.9.0 | v1.0.0 | v1.1.0+ |
|---------------------------------------------|:------:|:------:|:------:|:------:|:-------:|
| Core shell, config, split TUI, history      |   ☑️   |   -    |   -    |   -    |    -    |
| AI Chat, suggestion, safe command exec      |   -    |   ☑️   |   -    |   -    |    -    |
| Streaming, errors, polish, smart scroll     |   -    |   -    |   ☑️   |   -    |    -    |
| Final docs, bugfix, stabilization           |   -    |   -    |   -    |   ☑️   |    -    |
| Plugins, Windows, theming, clipboard, etc   |   -    |   -    |   -    |   -    |   ☑️    |

(Check ☑️ as features are delivered.)

---

## Known Non-Goals / Deferrals

- No graphical/modal GUIs or IDE code editing
- No SSH, package manager, or remote/cluster integration
- No embedded LLM inference: always relies on user-supplied endpoint
- No telemetry, analytics, or silent outbound network calls—ever

---

## Links

- [Architecture & Design](architecture.md)
- [Configuration Details](configuration.md)
- [Keybindings Reference](keybindings.md)

---

**Ready to help?**  
See contribution guidelines and join [discussions](https://github.com/semidark/kicli/discussions).