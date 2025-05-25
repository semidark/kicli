package app

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// KicliModel represents the main application state for the Bubbletea TUI
type KicliModel struct {
	// Core components (interfaces defined in docs/interfaces.md)
	// ptyManager   ptyhandler.PTYManager
	// aiClient     aiclient.AIClient
	// config       configmanager.AppConfig
	// db           storage.HistoryStore

	// UI components
	shellViewport   viewport.Model
	chatViewport    viewport.Model
	chatInput       textinput.Model
	commandGenInput textinput.Model

	// Application state
	currentFocus FocusPane
	windowWidth  int
	windowHeight int
	isLoading    bool
	errorMessage string
	sessionID    string

	// TODO: Add other state fields as needed
}

// FocusPane represents which pane currently has focus
type FocusPane int

const (
	FocusShell FocusPane = iota
	FocusChat
	FocusCommandGen
)

// Init initializes the model (required by Bubbletea)
func (m KicliModel) Init() tea.Cmd {
	// TODO: Initialize components and return initial commands
	return nil
}

// Update handles messages and updates the model (required by Bubbletea)
func (m KicliModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		// TODO: Update viewport sizes based on layout (65%/35% split)

	case tea.KeyMsg:
		// TODO: Handle keyboard navigation and input
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		// TODO: Handle custom message types defined in interfaces.md
		// case PtyOutputMsg:
		// case AIResponseChunkMsg:
		// case CommandGeneratedMsg:
		// etc.
	}

	return m, nil
}

// View renders the TUI (required by Bubbletea)
func (m KicliModel) View() string {
	// TODO: Implement the 65%/35% split layout
	// Left: PTY shell viewport
	// Right: AI chat viewport
	// Bottom left: Command generation input field

	if m.windowWidth == 0 {
		return "Initializing kicli..."
	}

	return "kicli TUI - Under construction\n\nPress 'q' or Ctrl+C to quit"
}
