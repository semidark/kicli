# Interface & Message Specifications  
_File: docs/interfaces.md_  

This document is the single source-of-truth for **all public Go interfaces** that the kicli code-base exposes between packages, together with the custom Bubble Tea message types used for inter-goroutine/UI communication and the conventions we follow for error handling.

> IMPORTANT  
> • All cross-package interaction MUST happen through the interfaces defined here.  
> • Interfaces are considered *stable* once released in a tagged version. Breaking changes require a major version bump.

---

## Table of Contents  
1. [General Conventions](#general-conventions)  
2. [Core Runtime Interfaces](#core-runtime-interfaces)  
   ‑ [PTYManager](#ptymanager)  
   ‑ [AIClient](#aiclient)  
   ‑ [HistoryStore](#historystore)  
   ‑ [ConfigProvider](#configprovider) *(internal)*  
3. [Bubble Tea Message Types](#bubble-tea-message-types)  
4. [Error-Handling Patterns](#error-handling-patterns)  
5. [Versioning & Compatibility](#versioning--compatibility)

---

## General Conventions

| Item                           | Rule / Comment                                                                                     |
|--------------------------------|-----------------------------------------------------------------------------------------------------|
| Import paths                   | All internal packages live under `kicli/internal/<pkg>`                                             |
| Context                        | **Never** accept `context.Context` in interface methods except for *long-running* client ops (AI). |
| Blocking vs async              | Expensive I/O (PTY reads, HTTP calls) MUST be async and surface results as Bubble Tea messages.     |
| Thread-safety                  | Concrete implementations **must** document whether methods are goroutine-safe.                      |
| Timeouts / cancellation        | AIClient streaming methods MUST support cancellation via a `requestID` + `CancelRequest`.           |
| Errors                         | Always wrap with `%w`, never hide root error. Prefer sentinel errors over string matching.          |
| Paths                          | Use `internal/util/xdg.go` helpers for XDG compliance.                                              |

---

## Core Runtime Interfaces

### PTYManager

```go
// Package ptyhandler
type PTYManager interface {
    // Start spawns the user's default shell inside a new PTY.
    // rows/cols correspond to the initial terminal size.
    Start(rows, cols int) error

    // Stop terminates the PTY session and child process gracefully.
    Stop() error

    // Write writes raw bytes to the PTY (stdin of the shell).
    Write(p []byte) (int, error)

    // ReadChan returns a channel that continuously emits PTY output chunks.
    // The channel is closed on PTY exit.
    ReadChan() <-chan []byte

    // ErrorChan returns asynchronous PTY-level errors (process exit, resize failures, etc.).
    ErrorChan() <-chan error

    // SetSize resizes the PTY; MUST return os.ErrClosed if called after Stop().
    SetSize(rows, cols int) error

    // VisibleBuffer returns the *current* display content (VT-100 24x80 style)
    // used as context for AI prompts. Should be cheap; can be a snapshot.
    VisibleBuffer() string

    // ScrollbackLines fetches the last N lines of scrollback (impl may cap).
    ScrollbackLines(n int) []string

    // SessionInfo returns metadata about the PTY session.
    SessionInfo() PTYSessionInfo
}
```

Implementation notes  
• Must be **goroutine-safe** for `Write`, `SetSize`, `VisibleBuffer`, and `ScrollbackLines`.  
• Uses `github.com/creack/pty` under the hood.  
• Scrollback buffer length configurable through package constant (`DefaultScrollback = 5000`).  

---

### AIClient

```go
// Package aiclient
type AIClient interface {
    // Configure replaces the current settings at runtime.
    Configure(cfg configmanager.AIConfig) error

    // StreamChatCompletion sends <userMessage> plus history & terminalCtx to the
    // LLM and returns immediately with a requestID and a channel.
    //
    // The channel emits ChatMessageChunk values until chunk.IsFinal == true or an
    // error chunk is sent. Implementations MUST close the channel when done.
    StreamChatCompletion(
        userMessage string,
        history     []ChatMessage,
        terminalCtx string,
    ) (requestID string, stream <-chan ChatMessageChunk, err error)

    // StreamCommandSuggestion works like StreamChatCompletion but yields command-
    // only text (bash/zsh snippets).
    StreamCommandSuggestion(
        description string,
        history     []ChatMessage,
        terminalCtx string,
    ) (requestID string, stream <-chan CommandSuggestionChunk, err error)

    // CancelRequest attempts to abort an in-flight request.
    // It is idempotent; calling on an already-finished request is a no-op.
    CancelRequest(requestID string)
}
```

Supporting structs

```go
type ChatMessage struct {
    Role    string // "user", "assistant", "system"
    Content string
}

type ChatMessageChunk struct {
    Content  string
    IsFinal  bool
    Error    error
}

type CommandSuggestionChunk struct {
    Command  string
    IsFinal  bool
    Error    error
}
```

Implementation notes  
• `AIClient` **must** be goroutine-safe.  
• Long-lived HTTP connections recommended (`http.Client` w/ keep-alive).  
• Stream channel MUST send exactly one chunk where `IsFinal==true OR Error!=nil`.  

Consider adding to AIClient interface:

```go
type ContextBuilder interface {
    BuildChatContext(history []ChatMessage, terminalCtx string, maxTokens int) string
    BuildCommandContext(description, terminalCtx string, maxTokens int) string
}
```

---

### HistoryStore

```go
// Package storage
type HistoryStore interface {
    // Init opens (or creates) the SQLite DB at dbPath and runs migrations.
    Init(dbPath string) error

    // AddChatMessage persists a single chat turn.
    AddChatMessage(
        sessionID string,
        role      string,
        content   string,
    ) error

    // GetChatMessages returns the last <limit> chat messages, newest last.
    GetChatMessages(
        sessionID string,
        limit     int,
    ) ([]aiclient.ChatMessage, error)

    // AddShellCommand logs an executed command.
    AddShellCommand(sessionID, command string) error

    // GetShellCommands returns the last <limit> commands, newest last.
    GetShellCommands(
        sessionID string,
        limit     int,
    ) ([]ShellCommandEntry, error)

    // AddChatMessageBatch persists a batch of chat turns.
    AddChatMessageBatch(sessionID string, messages []ChatMessage) error
}

type ShellCommandEntry struct {
    Command   string
    Timestamp time.Time
}
```

Implementation notes  
• Uses `modernc.org/sqlite`.  
• Each method opens its own transaction unless part of a bigger batch (future).  
• All timestamps in UTC (`time.Now().UTC()`).  

---

### ConfigProvider (internal)

For completeness; mostly used inside `configmanager`.

```go
// Package configmanager
type ConfigProvider interface {
    // Load reads config from fs/env, returns concrete AppConfig.
    Load() (AppConfig, error)
    // Save writes the supplied cfg to disk (non-atomic overwrite).
    Save(cfg AppConfig) error
    // Watch returns a channel that emits AppConfig updates.
    Watch() <-chan AppConfig
}
```

---

## Bubble Tea Message Types

All messages live in package `internal/app/msg`.

| Message Type (struct)            | Emitted By             | Purpose |
|----------------------------------|------------------------|---------|
| `PtyOutputMsg{Data []byte}`      | ptyhandler read loop   | Raw PTY output for shell viewport |
| `PtyExitedMsg{Err error}`        | ptyhandler             | Shell process exited or PTY closed |
| `AIResponseChunkMsg{Chunk string}` | aiclient streaming    | Partial AI chat text; NOT final |
| `AIResponseCompleteMsg{Full string; Role string; Err error}` | aiclient | Final chat response OR terminal error |
| `CommandSuggestionChunkMsg{Chunk string}` | aiclient stream | Partial command suggestion |
| `CommandGeneratedMsg{Command string; Error error}` | app logic | Final command suggestion ready |
| `AIRequestSentMsg{RequestID string}` | app logic           | UI spinner start |
| `ExecuteCommandInPtyMsg{Command string}` | app logic -> pty | User accepted suggestion |
| `ErrorOccurredMsg{Err error}`    | any layer             | Generic surfaced error |
| `ConfigLoadedMsg{Cfg AppConfig; Err error}` | configmanager | Config load result |
| `HistoryLoadedMsg{Chat []ChatMessage; Shell []storage.ShellCommandEntry; Err error}` | storage | Initial history fetch |
| `Tea.WindowSizeMsg` *(builtin)*  | Bubble Tea runtime     | Terminal resize |

Message guarantees  
1. **Chunk vs Complete**: For streaming messages (`...ChunkMsg`) there will always be a *corresponding complete* message unless an error occurs (in which case `Err != nil`).  
2. All message structs are **plain data** – no methods, no side-effects.  
3. Messages MUST be sent on the main thread back to Bubble Tea via `tea.Cmd`.

---

## Error-Handling Patterns

| Pattern                           | Details |
|-----------------------------------|---------|
| **Immediate Return**              | For synchronous APIs (`Init`, `Write`, `SetSize`) errors are returned directly. |
| **Async Propagation**             | Long-running goroutines translate errors into `ErrorOccurredMsg` (or specific message variants). |
| **Wrapped Errors**                | Always `fmt.Errorf("context: %w", err)` so callers can `errors.Is/As`. |
| **Sentinel Errors**               | `var ErrCancelled = errors.New("cancelled")` inside `aiclient` for aborted requests. |
| **Fatal vs Non-Fatal**            | Fatal errors (config load fail, PTY start fail) bubble to `tea.Quit`. Non-fatal errors update UI status. |
| **Retry Responsibility**          | Retry logic lives in the caller (usually `KicliModel`), not in individual packages. |
| **Logging**                       | Use `util/log.go` thin wrapper; *no* direct `log.Printf` in library code. |
| **Panics**                        | Library code must not panic; recover only in main. |

---

## Versioning & Compatibility

| Interface        | Stability Level | Notes |
|------------------|-----------------|-------|
| PTYManager       | Beta            | Subject to minor additions until v1.0 |
| AIClient         | Beta            | Streaming API stable; may add extra methods |
| HistoryStore     | Stable (alpha)  | Additive changes only |
| Bubble Tea msgs  | Beta            | Field additions allowed; renames are breaking |

Once kicli reaches `v1.0.0` all "Stable" interfaces become *SemVer* frozen (only compatible additions allowed).

---

_Questions or proposed changes to these interfaces? Please open an issue or start a discussion._