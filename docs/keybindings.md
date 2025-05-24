# Keybindings & Keyboard Navigation

kicli is fully keyboard-driven. You can customize nearly all actions and navigation keys to match your workflow.

---

## Table of Contents

- [Default Keybindings](#default-keybindings)
- [Customizing Keybindings](#customizing-keybindings)
- [Available Actions](#available-actions)
- [Best Practices & Notes](#best-practices--notes)
- [Troubleshooting](#troubleshooting)

---

## Default Keybindings

| Action                       | Default Key         | Description                                      |
|------------------------------|---------------------|--------------------------------------------------|
| Focus: next pane             | `Ctrl+Right`        | Move focus to the next pane (Shell → Chat → Cmd) |
| Focus: previous pane         | `Ctrl+Left`         | Move focus to previous pane                      |
| Scroll up (current pane)     | `Ctrl+Up`           | Scroll up in the focused Shell/Chat pane         |
| Scroll down (current pane)   | `Ctrl+Down`         | Scroll down in Shell/Chat panes                  |
| Accept/Submit                | `Enter`             | Confirm input (on chat or command fields)        |
| Cancel/Clear                 | `Esc`               | Cancel current input / clear field / return focus|
| Quit                         | `Ctrl+C`            | Quit kicli                                       |
| Copy AI code block           | *(TBD)*             | Planned: Copy code blocks from chat pane         |
| Access settings              | *(TBD)*             | Planned: Open settings pane                      |

You can view or change these in your config file (see below).

---

## Customizing Keybindings

Keybindings are defined in your `config.yaml` under the `keybindings:` section.  
**Example (`~/.config/kicli/config.yaml`):**
```yaml
keybindings:
  focus_next_pane: "ctrl+right"
  focus_prev_pane: "ctrl+left"
  scroll_up: "ctrl+up"
  scroll_down: "ctrl+down"
  confirm: "enter"
  cancel: "esc"
  quit: "ctrl+c"
```

**Notes:**
- If a keybinding is not specified, an internal default is used.
- Common modifier names: `ctrl`, `alt`, and basic keys (`up`, `down`, `left`, `right`, `enter`, `esc`, etc).
- To reset to defaults: delete/comment out the key in your config file.

**Changing a keybinding:**  
Example: To use `Tab` for moving to the next pane:
```yaml
keybindings:
  focus_next_pane: "tab"
```

**After editing, restart kicli** for changes to apply.

---

## Available Actions

You may customize these actions (values on the left are config keys):

- `focus_next_pane` — Next focus (cycles: Shell → Chat → Command → Shell)
- `focus_prev_pane` — Previous focus
- `scroll_up` — Scroll shell or chat pane up
- `scroll_down` — Scroll shell or chat pane down
- `confirm` — Confirm/submit input in current field (chat send, execute command)
- `cancel` — Abort/cancel/clear current input or AI call, or return focus to Shell
- `quit` — Quit kicli
- *(future)* `copy_code` — Copy highlighted AI code block
- *(future)* `settings` — Open settings pane or mode

Actions not mapped in your config use defaults.

---

## Best Practices & Notes

- Pick keys that do not clash with your shell or terminal’s own shortcuts.
- If a key does not work, double-check spelling and format (`ctrl+key`, `alt+key`, case-insensitive).
- Avoid assigning the same key to multiple actions.
- If you want to use special keys, consult [Bubbletea key documentation](https://github.com/charmbracelet/bubbletea#keyboard-input).

---

## Troubleshooting

- **Keybinding not working?**  
  - Check spelling in config (`ctrl+right` not `ctrl+→`)
  - Ensure no other program (terminal emulator/shell) intercepts the key.
  - If unsure what key your terminal sends, run `showkey -a` or similar.

- **Can’t focus/move as expected?**  
  - Restore the default keybinding by removing or correcting your config entry.

- **Reset to defaults:**  
  - Remove the `keybindings` section and restart kicli.

**If you have suggestions for new actions or default keys, open an [issue or discussion](https://github.com/semidark/kicli/issues).**

---

## Related Docs

- [Configuration](configuration.md)
- [Architecture](architecture.md)