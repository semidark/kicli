package app

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/semidark/kicli/internal/configmanager"
)

// KicliModel represents the main application state for the Bubbletea TUI
type KicliModel struct {
	// Core components (interfaces defined in docs/interfaces.md)
	configManager configmanager.ConfigManager
	config        configmanager.AppConfig
	// ptyManager   ptyhandler.PTYManager
	// aiClient     aiclient.AIClient
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

// NewKicliModel creates a new instance of the main application model
func NewKicliModel() (*KicliModel, error) {
	configManager, err := configmanager.New()
	if err != nil {
		return nil, err
	}

	return &KicliModel{
		configManager: configManager,
		isLoading:     true,
	}, nil
}

// Init initializes the model (required by Bubbletea)
func (m KicliModel) Init() tea.Cmd {
	// Load configuration as the first step
	return func() tea.Msg {
		cfg, err := m.configManager.Load()
		return ConfigLoadedMsg{
			Cfg: cfg,
			Err: err,
		}
	}
}

// Update handles messages and updates the model (required by Bubbletea)
func (m KicliModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		// TODO: Update viewport sizes based on layout (65%/35% split)

	case ConfigLoadedMsg:
		if msg.Err != nil {
			m.errorMessage = "Failed to load configuration: " + msg.Err.Error()
			m.isLoading = false
			return m, nil
		}
		m.config = msg.Cfg
		m.isLoading = false
		// TODO: Initialize other components with the loaded config
		return m, nil

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

	if m.isLoading {
		return "Loading configuration..."
	}

	if m.errorMessage != "" {
		return "Error: " + m.errorMessage + "\n\nPress 'q' or Ctrl+C to quit"
	}

	return "kicli TUI - Under construction\n\nConfiguration loaded successfully!\n\nPress 'q' or Ctrl+C to quit"
}
