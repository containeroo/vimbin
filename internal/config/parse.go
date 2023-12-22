package config

import (
	"fmt"
	"os"
	"path"
	"vimbin/internal/utils"

	"github.com/rs/zerolog/log"
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

	// Check if the API token was set as ENV variable
	if token := os.Getenv("VIMBIN_TOKEN"); token != "" {
		c.Server.Api.Token.Set(token)
		log.Debug().Msgf("Using API token from ENV variable: %s", token)
	}

	// Check if the API token is valid
	if c.Server.Api.Token.Get() == "" {
		if err := c.Server.Api.Token.Generate(32); err != nil {
			return fmt.Errorf("Unable to generate API token: %s", err)
		}
		log.Debug().Msgf("Generated API token: %s", c.Server.Api.Token.Get())
	}

	// Theme defaults
	c.Server.Web.LightTheme = "latte"
	if c.Server.Web.DarkTheme == "" {
		c.Server.Web.DarkTheme = "frappe"
	}

	// Check if the theme was set as ENV variable
	if theme := os.Getenv("VIMBIN_THEME"); theme != "" {
		if !utils.IsInList(theme, SupportedThemes) {
			return fmt.Errorf("Unsupported theme: %s. Supported themes are: %s", theme, SupportedThemes)
		}
		c.Server.Web.Theme = theme
		log.Debug().Msgf("Using theme from ENV variable: %s", theme)
	}

	if darkTheme := os.Getenv("VIMBIN_DARK_THEME"); darkTheme != "" {
		if !utils.IsInList(darkTheme, DarkThemes) {
			return fmt.Errorf("Unsupported dark theme: %s. Supported dark themes are: %s", darkTheme, DarkThemes)
		}
		c.Server.Web.DarkTheme = darkTheme
		log.Debug().Msgf("Using dark theme from ENV variable: %s", darkTheme)
	}

	return nil
}
