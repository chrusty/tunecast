package configuration

import (
	"fmt"

	env "github.com/caarlos0/env"
)

// Config for all of TuneCast's dependencies:
type Config struct {
	Chromecast ChromecastConfig
	Database   DatabaseConfig
	HTTP       HTTPConfig
	Library    LibraryConfig
	Logging    LoggingConfig
}

// Load prepares a new config and populates it from environment variables:
func Load() (*Config, error) {
	newConfig := &Config{}

	for _, configSection := range []interface{}{
		newConfig,
		&newConfig.Chromecast,
		&newConfig.Database,
		&newConfig.Library,
		&newConfig.HTTP,
		&newConfig.Logging,
	} {
		if err := env.Parse(configSection); err != nil {
			return nil, fmt.Errorf("Unable to load the config: %v", err)
		}
	}

	return newConfig, nil
}
