# Configuration Guide

This document covers all aspects of configuring **kicli**—from basic setup to advanced customization.

---

## Table of Contents

- [Overview](#overview)
- [Configuration File Location (XDG Paths)](#configuration-file-location-xdg-paths)
- [Quickstart: First-Time Setup](#quickstart-first-time-setup)
- [Configuration Schema & Options](#configuration-schema--options)
- [Environment Variable Overrides](#environment-variable-overrides)
- [Secrets and API Keys](#secrets-and-api-keys)
- [Theming](#theming)
- [Keybindings](#keybindings)
- [SQLite History Database Location](#sqlite-history-database-location)
- [Troubleshooting](#troubleshooting)

---

## Overview

kicli loads its configuration from a YAML file in your user config directory. Most options may also be overridden using environment variables (recommended for sensitive settings like API keys).

**Shell Note:**  
kicli always launches your *default* interactive shell based on your system. There is no setting for this—it uses your `$SHELL` (on Unix), or standard OS resolution. No shell configuration is necessary.

---

## Configuration File Location (XDG Paths)

- **Unix/Linux/macOS:**
  1. `$XDG_CONFIG_HOME/kicli/config.yaml`  
     (typically `~/.config/kicli/config.yaml`)
  2. If not found, fallback: `~/.config/kicli/config.yaml`
- **Windows:**  
  Not yet supported (planned; Windows native settings will follow XDG-like layout).

If the config file does not exist, kicli will prompt with an error.

---

## Quickstart: First-Time Setup

1. **Copy the sample configuration:**

   ```sh
   mkdir -p ~/.config/kicli
   cp assets/config.yaml ~/.config/kicli/config.yaml
   ```

2. **Edit your `config.yaml`:**
   - Set your LLM API endpoint, API key, and (optionally) adjust keybindings/themes.

3. **Start kicli:**

   ```sh
   ./kicli
   ```

4. **(Optional)**
   - Instead of editing the file, use environment variables for secrets (see below).

---

## Configuration Schema & Options

Here’s the structure of `config.yaml` (with all fields and comments):

```yaml
# ~/.config/kicli/config.yaml

ai:
  api_url: "https://api.openai.com/v1/chat/completions"   # OpenAI-compatible endpoint
  api_key: "sk-xxxxxxxxxxxxxxxxxxxx"                      # Your API key (can also use env var)
  model_name: "gpt-3.5-turbo"                             # LLM model name
  streaming_enabled: true                                 # Stream responses (default: true)

theme:
  colors:
    primary: "#00c6a8"
    secondary: "#cbf7ed"
    background: "#22223b"
    ai_assistant: "#d79921"
    user_input: "#00c6a8"
    error: "#ff0033"
    # ...add more as needed

keybindings:
  focus_next_pane: "ctrl+right"
  focus_prev_pane: "ctrl+left"
  scroll_up: "ctrl+up"
  scroll_down: "ctrl+down"
  confirm: "enter"
  cancel: "esc"
  # ...add more as supported
```

All sections/keys are **optional** (internal defaults are used), except the AI API credentials.

---

## Environment Variable Overrides

For sensitive settings and easy scripting, **any config key can be overridden with an environment variable** (just use the capitalized, `KICLI_`-prefixed version):

| Setting path       | Environment Variable         | Example Value                      |
|--------------------|-----------------------------|------------------------------------|
| `ai.api_key`       | `KICLI_API_KEY`             | `sk-xxxxxxx`                       |
| `ai.api_url`       | `KICLI_API_URL`             | `https://api.ollama.example/v1/...`|
| `ai.model_name`    | `KICLI_MODEL_NAME`          | `gpt-3.5-turbo`                    |

**Precedence:**  
1. Environment variable  
2. `config.yaml`  
3. Internal defaults

---

## Secrets and API Keys

**It’s strongly recommended to provide your API keys via environment variables** rather than storing them in `config.yaml`.  
For example:

```sh
export KICLI_API_KEY=sk-xxxxxx
./kicli
```

If both the environment and the YAML file specify a value, the **environment variable takes precedence**.

---

## Theming

- UI colors are defined under `theme.colors`.
- Use hex strings (`#rrggbb`) or Xterm color names.
- Adjust colors in `config.yaml` and restart kicli for changes to take effect.

---

## Keybindings

- All major actions, navigation, and field focus can be customized in the `keybindings` section.
- Key descriptions follow Bubbletea conventions (`ctrl+left`, `enter`, `esc`).
- To configure your keymap, edit `config.yaml` or set corresponding environment variables.

See [docs/keybindings.md](keybindings.md) for all configurable actions.

---

## SQLite History Database Location

All chat and shell history is stored unencrypted in SQLite:

- **Path:**  
  `$XDG_DATA_HOME/kicli/history.db`  
  (typically `~/.local/share/kicli/history.db`)
- If `XDG_DATA_HOME` is not set, fallback: `~/.local/share/kicli/history.db`

No data is ever sent anywhere else, except when included as LLM prompt context for AI operations.

---

## Troubleshooting

- **Missing config:**  
  kicli will report a missing `config.yaml` and exit.
- **Malformed YAML:**  
  kicli will show the parse error at startup.
- **Invalid API key/endpoint:**  
  AI features will not work; check logs/TUI error area for details.
- **Invalid keybindings/theme:**  
  kicli will report invalid keys/colors and fallback to defaults.

If problems persist, please [open an issue](https://github.com/semidark/kicli/issues).

---

## Further Reading

- [Architecture](architecture.md)
- [Keybindings](keybindings.md)
- [Privacy Statement](privacy.md)
- [Implementation Plan](implementation-plan.md)

---

**kicli always uses your system’s default shell.**  
No extra configuration is required for your shell.
