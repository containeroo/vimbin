package config

import (
	"fmt"
	"os"

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

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// No configuration file found
		return nil
	}

	// Set the Viper config file path
	viper.SetConfigFile(configPath)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read config file: %v", err)
	}

	if err := viper.Unmarshal(c, func(d *mapstructure.DecoderConfig) {
		d.ZeroFields = true // Zero out any existing fields
		d.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			customTokenDecodeHook, // Custom decoder hook for the Token field
		)
	}); err != nil {
		return fmt.Errorf("Failed to unmarshal config file: %v", err)
	}

	return nil
}
