package configmanager

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

const (
	appName    = "kicli"
	configFile = "config.yaml"
	dbFile     = "history.db"
)

// getConfigPath returns the full path to the configuration file
func getConfigPath() (string, error) {
	return xdg.ConfigFile(filepath.Join(appName, configFile))
}

// getDataPath returns the full path to the data directory
func getDataPath() (string, error) {
	return xdg.DataFile(filepath.Join(appName, dbFile))
}
