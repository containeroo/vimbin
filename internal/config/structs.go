package config

import (
	"sync"
	"text/template"
)

// App is the global configuration instance.
var App Config

// Config represents the application configuration.
type Config struct {
	HtmlTemplate *template.Template `yaml:"-"`       // HtmlTemplate contains the HTML template content.
	Server       Server             `yaml:"server"`  // Server represents the server configuration.
	Storage      Storage            `yaml:"storage"` // Storage represents the storage configuration.
}

// Web represents the web configuration.
type Web struct {
	Theme   string `yaml:"server"`  // Theme is the theme to use for the web interface.
	Address string `yaml:"address"` // Address is the address to listen on for HTTP requests.
}

type Api struct {
	SkipInsecureVerify bool   `yaml:"skipInsecureVerify"` // SkipInsecureVerify skips the verification of TLS certificates.
	Address            string `yaml:"address"`            // Address is the address to push/fetch content from.
}

// Server represents the server configuration.
type Server struct {
	Web Web `yaml:"web"` // Web represents the web configuration.
	Api Api `yaml:"api"` // Api represents the api configuration.
}

// Storage represents the storage configuration.
type Storage struct {
	Name      string  `yaml:"name"`      // Name is the name of the storage file.
	Directory string  `yaml:"directory"` // Directory is the directory path for storage file.
	Path      string  `yaml:"-"`         // Path is the full path to the storage file.
	Content   Content `yaml:"-"`         // Content represents the content stored in the storage file.
}

// Content represents the content stored in the storage with thread-safe methods.
type Content struct {
	text   string       `yaml:"-"` // text is the stored content.
	mutext sync.RWMutex `yaml:"-"` // mutext is a read-write mutex for concurrent access control.
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
