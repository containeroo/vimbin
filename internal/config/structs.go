package config

import (
	"sync"
	"text/template"
	"vimbin/internal/utils"
)

// App is the global configuration instance.
var App Config

// Config represents the application configuration.
type Config struct {
	Version      string             `mapstructure:"-"`       // Version is the version of the application.
	HtmlTemplate *template.Template `mapstructure:"-"`       // HtmlTemplate contains the HTML template content.
	Server       Server             `mapstructure:"server"`  // Server represents the server configuration.
	Storage      Storage            `mapstructure:"storage"` // Storage represents the storage configuration.
}

// Web represents the web configuration.
type Web struct {
	Theme   string `mapstructure:"server"`  // Theme is the theme to use for the web interface.
	Address string `mapstructure:"address"` // Address is the address to listen on for HTTP requests.
}

// Token represents the API token.
type Token struct {
	value string
}

// Get retrieves the current token value.
//
// Returns:
//   - string
//     The current token value.
func (t *Token) Get() string {
	return t.value
}

// Set sets the token to the specified value.
//
// Parameters:
//   - token: string
//     The value to set as the token.
func (t *Token) Set(token string) {
	t.value = token
}

// Generate generates a new random token of the specified length and updates the token value.
//
// Parameters:
//   - len: int
//     The length of the new token.
//
// Returns:
//   - error
//     An error, if any, encountered during token generation.
func (t *Token) Generate(len int) error {
	tokenString, err := utils.GenerateRandomToken(len)
	if err != nil {
		return err
	}
	t.value = tokenString

	return nil
}

// Api represents the api configuration.
type Api struct {
	Token              Token  `mapstructure:"token"`              // Token is the API token.
	SkipInsecureVerify bool   `mapstructure:"skipInsecureVerify"` // SkipInsecureVerify skips the verification of TLS certificates.
	Address            string `mapstructure:"address"`            // Address is the address to push/fetch content from.
}

// Server represents the server configuration.
type Server struct {
	Web Web `mapstructure:"web"` // Web represents the web configuration.
	Api Api `mapstructure:"api"` // Api represents the api configuration.
}

// Storage represents the storage configuration.
type Storage struct {
	Name      string  `mapstructure:"name"`      // Name is the name of the storage file.
	Directory string  `mapstructure:"directory"` // Directory is the directory path for storage file.
	Path      string  `mapstructure:"-"`         // Path is the full path to the storage file.
	Content   Content `mapstructure:"-"`         // Content represents the content stored in the storage file.
}

// Content represents the content stored in the storage with thread-safe methods.
type Content struct {
	text   string       `mapstructure:"-"` // text is the stored content.
	mutext sync.RWMutex `mapstructure:"-"` // mutext is a read-write mutex for concurrent access control.
}

// Set sets the content to the specified text.
//
// Parameters:
//   - text: string
//     The text to set as the content.
func (c *Content) Set(text string) {
	c.mutext.Lock()
	defer c.mutext.Unlock()
	c.text = text
}

// Get retrieves the current content.
//
// Returns:
//   - string
//     The current content.
func (c *Content) Get() string {
	c.mutext.RLock()
	defer c.mutext.RUnlock()
	return c.text
}

// Append appends the specified text to the current content.
//
// Parameters:
//   - text: string
//     The text to append to the current content.
func (c *Content) Append(text string) {
	c.mutext.Lock()
	defer c.mutext.Unlock()
	c.text += text
}
