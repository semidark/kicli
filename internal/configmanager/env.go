package configmanager

import (
	"os"
	"strconv"
)

// applyEnvOverrides applies environment variable overrides to the configuration
func applyEnvOverrides(cfg *AppConfig) error {
	// AI configuration overrides
	if val := os.Getenv("KICLI_API_KEY"); val != "" {
		cfg.AI.APIKey = val
	}
	if val := os.Getenv("KICLI_API_URL"); val != "" {
		cfg.AI.APIURL = val
	}
	if val := os.Getenv("KICLI_MODEL_NAME"); val != "" {
		cfg.AI.ModelName = val
	}
	if val := os.Getenv("KICLI_STREAMING_ENABLED"); val != "" {
		if parsed, err := strconv.ParseBool(val); err == nil {
			cfg.AI.StreamingEnabled = parsed
		}
	}

	// Advanced configuration overrides
	if val := os.Getenv("KICLI_MAX_SCROLLBACK_LINES"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			cfg.Advanced.MaxScrollbackLines = parsed
		}
	}
	if val := os.Getenv("KICLI_MAX_CONTEXT_MESSAGES"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			cfg.Advanced.MaxContextMessages = parsed
		}
	}
	if val := os.Getenv("KICLI_AI_TIMEOUT_SECONDS"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			cfg.Advanced.AITimeoutSeconds = parsed
		}
	}

	return nil
}
