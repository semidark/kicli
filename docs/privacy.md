# Privacy & Data Policy

**kicli is designed for privacy. It never collects or transmits your data except where you explicitly configure it.** This document explains what is (and is not) collected, stored, or transmitted by kicli.

---

## Table of Contents

- [Summary](#summary)
- [What Data Never Leaves Your Machine](#what-data-never-leaves-your-machine)
- [Outbound Connections](#outbound-connections)
- [History & Local Storage](#history--local-storage)
- [AI Context Sharing](#ai-context-sharing)
- [How to Control Your Data](#how-to-control-your-data)
- [Third-Party Endpoints](#third-party-endpoints)
- [Handling Sensitive Data](#handling-sensitive-data)
- [Policy Changes](#policy-changes)
- [Questions](#questions)

---

## Summary

**kicli sends no telemetry, analytics, or usage data.**  
- No crash reporter, no stats gathering, no “phone home”, and no “opt out” requirement.
- There are _no_ hidden or silent outbound network connections, except those you explicitly configure (for AI).
- All shell and chat histories are _stored locally only_, in plain SQLite.

---

## What Data Never Leaves Your Machine

These will **never be transmitted anywhere** by kicli:

- Your shell command history
- AI chat history (unless used as prompt context, see below)
- Terminal (PTY) buffer contents
- Configuration files and secrets
- Usage patterns, environment variables, or file paths

The only exception: you may opt to use these as *context* in requests to your AI endpoint (see below).

---

## Outbound Connections

kicli makes outbound internet connections **only** in two cases:

1. **To your configured LLM/AI API endpoint.**
    - _This endpoint is entirely user-controlled._ (e.g. OpenAI, ollama, or your own server)
    - No default endpoints are shipped or hard-coded.
2. **To check for LLM API reachability (on startup or when sending messages):** only as needed, never to any other server.

kicli does **not** perform any of:
- Silent updates
- Remote crash reporting
- 3rd-party metrics

---

## History & Local Storage

- **All history is stored locally** in plaintext SQLite in the standard data directory (`$XDG_DATA_HOME/kicli/history.db`, or `~/.local/share/kicli/history.db`).
- No encryption at rest.
- Deleting the database file deletes your data entirely.
- Your configuration file (`config.yaml`) is kept in your local XDG config directory.

---

## AI Context Sharing

**When you interact with the AI assistant:**
- kicli sends your chat message, and (optionally) parts of your recent shell buffer and/or chat history _as prompt context_ to your configured LLM endpoint.
- **You are in full control of which endpoint this is**—see [configuration](configuration.md).
- kicli will not auto-send your full history; it only transmits the minimum context required for your chosen feature to work (AI chat and suggestions).
- **No data is ever sent to any server or endpoint besides the one you configure.**

If you are highly privacy-sensitive, you can choose to:
- Use a local/private LLM endpoint
- Regularly clear or disable chat and command history

---

## How to Control Your Data

- **To clear history:** delete `~/.local/share/kicli/history.db`.
- **To control what’s sent to AI:**  
    - You determine the LLM provider, model, and endpoint.
    - If you do not configure an endpoint, AI features are disabled and nothing is transmitted.

- **To use environment variables only:**  
    - You don’t need to write your API key to a file; set via `KICLI_API_KEY`.

---

## Third-Party Endpoints

kicli is not responsible for the privacy practices of your chosen LLM provider.  
_You must trust and understand the endpoint you configure for AI features._  
- For example, using `api.openai.com` means data will be handled per OpenAI’s policy.
- Using a local [Ollama](https://github.com/jmorganca/ollama) instance keeps LLM context entirely on your machine.

---

## Handling Sensitive Data

- If highly sensitive data may appear in your shell or chat sessions, be cautious what is included as context in AI requests.
- Always audit your settings and environment for privacy compliance if this matters to you.

---

## Policy Changes

- kicli is strongly committed to zero-telemetry and full user data control.
- Any policy or code changes that alter these guarantees will be clearly announced in release notes and documentation.

---

## Questions

Need more info? Found something unclear?  
- [Open a privacy issue or discussion.](https://github.com/semidark/kicli/issues)
- All legitimate privacy concerns will be treated as high priority.

---

**kicli: empowering you, and protecting your privacy, by design.**
