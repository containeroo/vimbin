package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// Read reads the configuration settings from a specified config file.
//
// This method uses Viper to read and unmarshal the configuration from the provided file path.
//
// Parameters:
//   - configPath: string
//     The file path to the configuration file.
//
// Returns:
//   - error
//     An error if reading or unmarshalling the configuration fails.
func (c *Config) Read(configPath string) error {
	// Set the Viper config file path
	viper.SetConfigFile(configPath)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read config file: %v", err)
	}

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(c, func(d *mapstructure.DecoderConfig) { d.ZeroFields = true }); err != nil {
		return fmt.Errorf("Failed to unmarshal config file: %v", err)
	}

	return nil
}
