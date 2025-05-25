package configmanager

// DefaultConfig returns a configuration with sensible defaults
func DefaultConfig() AppConfig {
	return AppConfig{
		AI: AIConfig{
			APIURL:           "https://api.openai.com/v1/chat/completions",
			APIKey:           "", // Must be provided by user
			ModelName:        "gpt-3.5-turbo",
			StreamingEnabled: true,
		},
		Theme: ThemeConfig{
			Colors: ColorConfig{
				Primary:     "#00c6a8",
				Secondary:   "#cbf7ed",
				Background:  "#22223b",
				AIAssistant: "#d79921",
				UserInput:   "#00c6a8",
				Error:       "#ff0033",
				Success:     "#00ff00",
				Warning:     "#ffaa00",
			},
		},
		Keybindings: KeybindingsConfig{
			FocusNextPane: "ctrl+right",
			FocusPrevPane: "ctrl+left",
			ScrollUp:      "ctrl+up",
			ScrollDown:    "ctrl+down",
			Confirm:       "enter",
			Cancel:        "esc",
			Quit:          "ctrl+c",
		},
		Advanced: AdvancedConfig{
			MaxScrollbackLines: 5000,
			MaxContextMessages: 20,
			AITimeoutSeconds:   30,
		},
	}
}
