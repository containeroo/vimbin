package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	// CreateTempFile creates a temporary file with the given content and returns a pointer to it.
	createTempFile := func(content string) (*os.File, error) {
		tmpfile, err := os.CreateTemp("", "example.*.yaml")
		if err != nil {
			return nil, fmt.Errorf("Failed to create temp file: %v", err)
		}

		if _, err := tmpfile.Write([]byte(content)); err != nil {
			return nil, fmt.Errorf("Failed to write to temp file: %v", err)
		}

		if err := tmpfile.Close(); err != nil {
			return nil, fmt.Errorf("Failed to close temp file: %v", err)
		}

		return tmpfile, nil
	}

	t.Run("Valid configuration", func(t *testing.T) {
		content := `
server:
  api:
    skipInsecureVerify: true
`
		filePath, err := createTempFile(content)
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(filePath.Name())

		config := &Config{}
		err = config.Read(filePath.Name())
		assert.NoError(t, err)
		assert.Equal(t, true, config.Server.Api.SkipInsecureVerify)
		assert.Equal(t, "", config.Server.Api.Token.Get())
	})

	t.Run("Validate token", func(t *testing.T) {
		content := `
server:
  api:
    token: token
`
		filePath, err := createTempFile(content)
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(filePath.Name())

		config := &Config{}
		err = config.Read(filePath.Name())
		assert.NoError(t, err)
		assert.Equal(t, "token", config.Server.Api.Token.Get())
	})

	// Test reading a non-existent file
	t.Run("Invalid file path", func(t *testing.T) {
		config := &Config{}
		err := config.Read("path/to/non_existent.yaml")
		assert.EqualError(t, err, "Failed to read config file: open path/to/non_existent.yaml: no such file or directory")
	})

	t.Run("Invalid YAML", func(t *testing.T) {
		content := `
server:
  api:
    skipInsecureVerify: not_a_boolean
`
		filePath, err := createTempFile(content)
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(filePath.Name())

		// Run the test
		cfg := &Config{}
		err = cfg.Read(filePath.Name())
		assert.EqualError(t, err, "Failed to unmarshal config file: 1 error(s) decoding:\n\n* cannot parse 'server.api.skipInsecureVerify' as bool: strconv.ParseBool: parsing \"not_a_boolean\": invalid syntax")
	})
}
