package configmanager

// AppConfig represents the complete application configuration
type AppConfig struct {
	AI          AIConfig          `yaml:"ai"`
	Theme       ThemeConfig       `yaml:"theme"`
	Keybindings KeybindingsConfig `yaml:"keybindings"`
	Advanced    AdvancedConfig    `yaml:"advanced"`
}

// AIConfig holds AI-related configuration
type AIConfig struct {
	APIURL           string `yaml:"api_url"`
	APIKey           string `yaml:"api_key"`
	ModelName        string `yaml:"model_name"`
	StreamingEnabled bool   `yaml:"streaming_enabled"`
}

// ThemeConfig holds theme and color configuration
type ThemeConfig struct {
	Colors ColorConfig `yaml:"colors"`
}

// ColorConfig defines the color scheme
type ColorConfig struct {
	Primary     string `yaml:"primary"`
	Secondary   string `yaml:"secondary"`
	Background  string `yaml:"background"`
	AIAssistant string `yaml:"ai_assistant"`
	UserInput   string `yaml:"user_input"`
	Error       string `yaml:"error"`
	Success     string `yaml:"success"`
	Warning     string `yaml:"warning"`
}

// KeybindingsConfig holds all keybinding configuration
type KeybindingsConfig struct {
	FocusNextPane string `yaml:"focus_next_pane"`
	FocusPrevPane string `yaml:"focus_prev_pane"`
	ScrollUp      string `yaml:"scroll_up"`
	ScrollDown    string `yaml:"scroll_down"`
	Confirm       string `yaml:"confirm"`
	Cancel        string `yaml:"cancel"`
	Quit          string `yaml:"quit"`
}

// AdvancedConfig holds advanced/optional settings
type AdvancedConfig struct {
	MaxScrollbackLines int `yaml:"max_scrollback_lines"`
	MaxContextMessages int `yaml:"max_context_messages"`
	AITimeoutSeconds   int `yaml:"ai_timeout_seconds"`
}
