package configmanager

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	mgr, err := New()
	if err != nil {
		t.Fatalf("Failed to create config manager: %v", err)
	}

	if mgr == nil {
		t.Fatal("Config manager is nil")
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	// Test that defaults are set
	if cfg.AI.ModelName == "" {
		t.Error("Default AI model name should not be empty")
	}

	if cfg.AI.APIURL == "" {
		t.Error("Default API URL should not be empty")
	}

	if cfg.Advanced.MaxScrollbackLines <= 0 {
		t.Error("Default max scrollback lines should be positive")
	}
}

func TestApplyEnvOverrides(t *testing.T) {
	// Set test environment variables
	os.Setenv("KICLI_API_KEY", "test-key")
	os.Setenv("KICLI_MODEL_NAME", "test-model")
	defer func() {
		os.Unsetenv("KICLI_API_KEY")
		os.Unsetenv("KICLI_MODEL_NAME")
	}()

	cfg := DefaultConfig()
	err := applyEnvOverrides(&cfg)
	if err != nil {
		t.Fatalf("Failed to apply env overrides: %v", err)
	}

	if cfg.AI.APIKey != "test-key" {
		t.Errorf("Expected API key 'test-key', got '%s'", cfg.AI.APIKey)
	}

	if cfg.AI.ModelName != "test-model" {
		t.Errorf("Expected model name 'test-model', got '%s'", cfg.AI.ModelName)
	}
}

func TestValidateConfig(t *testing.T) {
	// Test valid config
	cfg := DefaultConfig()
	cfg.AI.APIKey = "test-key"

	err := validateConfig(&cfg)
	if err != nil {
		t.Errorf("Valid config should not fail validation: %v", err)
	}

	// Test invalid config - missing API key
	cfg.AI.APIKey = ""
	err = validateConfig(&cfg)
	if err == nil {
		t.Error("Config with missing API key should fail validation")
	}
}
