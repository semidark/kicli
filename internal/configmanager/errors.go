package configmanager

import "errors"

var (
	// ErrConfigNotFound indicates the configuration file was not found
	ErrConfigNotFound = errors.New("config file not found")

	// ErrInvalidConfig indicates the configuration contains invalid values
	ErrInvalidConfig = errors.New("invalid configuration")

	// ErrMissingAPIKey indicates the AI API key is required but not provided
	ErrMissingAPIKey = errors.New("AI API key is required")

	// ErrInvalidAPIURL indicates the AI API URL is malformed
	ErrInvalidAPIURL = errors.New("invalid AI API URL")
)
