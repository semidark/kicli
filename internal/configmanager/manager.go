package configmanager

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// ConfigManager handles loading and managing application configuration
type ConfigManager interface {
	// Load reads configuration from file and environment variables
	Load() (AppConfig, error)

	// Save writes configuration to the config file
	Save(cfg AppConfig) error

	// GetConfigPath returns the path where config should be stored
	GetConfigPath() (string, error)

	// GetDataPath returns the path where data should be stored
	GetDataPath() (string, error)
}

// manager implements ConfigManager
type manager struct {
	configPath string
	dataPath   string
}

// New creates a new ConfigManager instance
func New() (ConfigManager, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to determine config path: %w", err)
	}

	dataPath, err := getDataPath()
	if err != nil {
		return nil, fmt.Errorf("failed to determine data path: %w", err)
	}

	return &manager{
		configPath: configPath,
		dataPath:   dataPath,
	}, nil
}

// Load implements ConfigManager.Load
func (m *manager) Load() (AppConfig, error) {
	// Start with default configuration
	cfg := DefaultConfig()

	// Try to load from file
	if err := m.loadFromFile(&cfg); err != nil {
		// If file doesn't exist, that's okay - we'll use defaults + env vars
		if !os.IsNotExist(err) {
			return cfg, fmt.Errorf("failed to load config file: %w", err)
		}
	}

	// Apply environment variable overrides
	if err := applyEnvOverrides(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to apply environment overrides: %w", err)
	}

	// Validate the final configuration
	if err := validateConfig(&cfg); err != nil {
		return cfg, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}

// loadFromFile loads configuration from the YAML file
func (m *manager) loadFromFile(cfg *AppConfig) error {
	data, err := os.ReadFile(m.configPath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	return nil
}

// Save implements ConfigManager.Save
func (m *manager) Save(cfg AppConfig) error {
	// Ensure the config directory exists
	if err := os.MkdirAll(filepath.Dir(m.configPath), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Marshal to YAML
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config to YAML: %w", err)
	}

	// Write to file
	if err := os.WriteFile(m.configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// GetConfigPath returns the configuration file path
func (m *manager) GetConfigPath() (string, error) {
	return m.configPath, nil
}

// GetDataPath returns the data file path
func (m *manager) GetDataPath() (string, error) {
	return m.dataPath, nil
}
