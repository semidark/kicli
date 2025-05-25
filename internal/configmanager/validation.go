package configmanager

import (
	"fmt"
	"net/url"
	"strings"
)

// validateConfig validates the loaded configuration
func validateConfig(cfg *AppConfig) error {
	// Validate AI configuration
	if err := validateAIConfig(&cfg.AI); err != nil {
		return fmt.Errorf("AI config validation failed: %w", err)
	}

	// Validate advanced configuration
	if err := validateAdvancedConfig(&cfg.Advanced); err != nil {
		return fmt.Errorf("advanced config validation failed: %w", err)
	}

	// Validate theme configuration
	if err := validateThemeConfig(&cfg.Theme); err != nil {
		return fmt.Errorf("theme config validation failed: %w", err)
	}

	return nil
}

func validateAIConfig(cfg *AIConfig) error {
	// API Key is required
	if strings.TrimSpace(cfg.APIKey) == "" {
		return ErrMissingAPIKey
	}

	// Validate API URL
	if _, err := url.Parse(cfg.APIURL); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidAPIURL, err)
	}

	// Model name should not be empty
	if strings.TrimSpace(cfg.ModelName) == "" {
		return fmt.Errorf("%w: model name cannot be empty", ErrInvalidConfig)
	}

	return nil
}

func validateAdvancedConfig(cfg *AdvancedConfig) error {
	if cfg.MaxScrollbackLines <= 0 {
		return fmt.Errorf("%w: max_scrollback_lines must be positive", ErrInvalidConfig)
	}
	if cfg.MaxContextMessages <= 0 {
		return fmt.Errorf("%w: max_context_messages must be positive", ErrInvalidConfig)
	}
	if cfg.AITimeoutSeconds <= 0 {
		return fmt.Errorf("%w: ai_timeout_seconds must be positive", ErrInvalidConfig)
	}
	return nil
}

func validateThemeConfig(cfg *ThemeConfig) error {
	// Basic color validation - ensure they're not empty
	colors := map[string]string{
		"primary":      cfg.Colors.Primary,
		"secondary":    cfg.Colors.Secondary,
		"background":   cfg.Colors.Background,
		"ai_assistant": cfg.Colors.AIAssistant,
		"user_input":   cfg.Colors.UserInput,
		"error":        cfg.Colors.Error,
		"success":      cfg.Colors.Success,
		"warning":      cfg.Colors.Warning,
	}

	for name, color := range colors {
		if strings.TrimSpace(color) == "" {
			return fmt.Errorf("%w: color %s cannot be empty", ErrInvalidConfig, name)
		}
		// Basic hex color validation
		if strings.HasPrefix(color, "#") && len(color) != 7 {
			return fmt.Errorf("%w: color %s must be a valid hex color", ErrInvalidConfig, name)
		}
	}

	return nil
}
