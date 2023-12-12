package config

import (
	"fmt"
	"os"
	"path"
	"vimbin/internal/utils"
)

// Parse reads and processes the configuration settings.
//
// This method handles various configuration-related tasks, such as setting the working directory,
// expanding environment variables, checking the storage file, reading the content file,
// and validating the hostname and port.
//
// Returns:
//   - err: error
//     An error if any of the configuration tasks fail.
func (c *Config) Parse() (err error) {
	// Get the working directory
	if c.Storage.Directory == "" || c.Storage.Directory == "$(pwd)" {
		c.Storage.Directory, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("Unable to get working directory: %s", err)
		}
	}

	// Expand environment variables in the storage directory
	c.Storage.Directory = os.ExpandEnv(c.Storage.Directory)

	// Set the full path to the storage file
	c.Storage.Path = path.Join(c.Storage.Directory, c.Storage.Name)

	// Check if the storage file is valid
	if err := checkStorageFile(c.Storage.Path); err != nil {
		return fmt.Errorf("Unable to check storage file: %s", err)
	}

	// Read the content file
	content, err := os.ReadFile(c.Storage.Path)
	if err != nil {
		return fmt.Errorf("Cannot read storage file. %s", err)
	}
	c.Storage.Content.Set(string(content))

	// Check if Hostname and Port are valid
	if _, _, err := utils.ExtractHostAndPort(c.Server.Web.Address); err != nil {
		return fmt.Errorf("Unable to extract hostname and port: %s", err)
	}

	return nil
}
