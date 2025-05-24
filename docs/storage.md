# Storage & History in kicli

This document describes how **kicli** handles the storage of terminal and AI chat history, what is persisted and where, privacy/encryption considerations, and details for developers or advanced users.

---

## Table of Contents

- [Overview](#overview)
- [Location of History Database](#location-of-history-database)
- [What is Stored?](#what-is-stored)
- [Database Schema](#database-schema)
- [Privacy & Security](#privacy--security)
- [Schema Evolution & Backups](#schema-evolution--backups)
- [Developer Notes](#developer-notes)
- [Troubleshooting / FAQ](#troubleshooting--faq)
- [Related Links](#related-links)

---

## Overview

kicli persists both Terminal (PTY) command history and AI chat history.  
All user inputs, executed commands, and AI responses are stored in a **local SQLite database**, to provide context for AI, enable history navigation, and support session restoration.

- No cloud storage or telemetry.
- Data is stored in unencrypted SQLite files under standard XDG data directories.

---

## Location of History Database

- **Primary path:**  
  `$XDG_DATA_HOME/kicli/history.db`  
  (On most Linux systems, this is `~/.local/share/kicli/history.db`)
- **Fallback:**  
  If `XDG_DATA_HOME` is not set:  
  `~/.local/share/kicli/history.db`
- File is created automatically if missing.

The database is scheme-stable and fully accessible by any local SQLite client/tool.

---

## What is Stored?

kicli logs three main types of data:

1. **Shell Commands**  
   - Each command executed within the kicli PTY shell.
   - Timestamp of execution is recorded.
   - Session identifier used (to differentiate between sessions—future versions may allow multi-session browsing).

2. **AI Chat Messages**
   - Each message (from user or assistant) in the AI chat pane.
   - Role is stored as `"user"`, `"assistant"`, or `"system"`.
   - Timestamp and session ID are included.

3. **Sessions** *(for future multi-session support)*
   - A session is created on each new launch; session metadata may be stored for navigation/history grouping.

**Note:**  
Command suggestions made by the AI are only stored after user review/execution.

---

## Database Schema

A simplified schema is as follows:

```sql
-- Table: shell_commands
CREATE TABLE IF NOT EXISTS shell_commands (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    session_id TEXT NOT NULL,
    command TEXT NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Table: chat_history
CREATE TABLE IF NOT EXISTS chat_history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    session_id TEXT NOT NULL,
    role TEXT CHECK(role IN ('user', 'assistant', 'system')) NOT NULL,
    content TEXT NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Table: sessions
CREATE TABLE IF NOT EXISTS sessions (
    session_id TEXT PRIMARY KEY,
    start_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    description TEXT
);
```

- **`session_id`** identifies a launch of kicli; may support reconnecting/persisting in the future.
- **Forward compatibility:** Schema may evolve, but only in additive (non-breaking) ways.

---

## Privacy & Security

- **Data never leaves your device:** Unless you configure your preferred LLM endpoint, history remains local.
- **No encryption:** History is stored as *plain* SQLite. Anyone with file access can read it.
    - You are responsible for your own disk/file permissions and OS-level security.
- **AI prompt context:** Recent history from the database may be included in prompts sent to your LLM endpoint. See [Privacy Policy](privacy.md).

---

## Schema Evolution & Backups

- Schema migrations are automatic and non-destructive—new columns/tables may be added in future releases.
- You can browse, back up, or edit your history directly with any SQLite client:
    ```sh
    sqlite3 ~/.local/share/kicli/history.db
    ```
- Backup files as needed; restoring from backup simply means swapping the file.

---

## Developer Notes

- Storage backend is pure Go (`modernc.org/sqlite`), so **no CGO required**.
- Data access is exclusively through internal/query APIs for safety.
- If you add new fields or tables, use additive migrations and maintain backward compatibility.
- Write operations (including chat/command history) are performed synchronously or queued to prevent data loss; error handling/reporting is surfaced in the UI status area.

---

## Troubleshooting / FAQ

- **Can I delete my history?**
    - Yes, just delete `history.db` (kicli will recreate it on next launch, but previous sessions will be lost).
- **Does kicli sync or upload my history?**
    - Never. All data is local, unless you explicitly submit it to your LLM endpoint through context.
- **How do I clear a session’s history?**
    - In early versions, delete the DB; future versions may add in-app tools or CLI flags.
- **What if my DB gets corrupted?**
    - Shutdown kicli, move or delete the corrupted `history.db`, and restart.

---

## Related Links

- [Configuration](configuration.md)
- [Privacy](privacy.md)
- [Architecture](architecture.md)

---

**For privacy-sensitive workflows, consider regularly deleting `history.db` or running kicli on encrypted disks. You are in full control of all local data.**
