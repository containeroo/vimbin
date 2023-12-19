package utils

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDebugAndTrace(t *testing.T) {
	t.Run("Both debug and trace are set", func(t *testing.T) {
		os.Setenv(VIMBINDebugEnv, "true")
		os.Setenv(VIMBINTraceEnv, "true")
		defer os.Unsetenv(VIMBINDebugEnv)
		defer os.Unsetenv(VIMBINTraceEnv)

		debug, trace, err := GetDebugAndTrace()

		assert.False(t, debug)
		assert.False(t, trace)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "'VIMBIN_DEBUG' and 'VIMBIN_TRACE' are mutually exclusive")
	})

	t.Run("Invalid debug value", func(t *testing.T) {
		os.Setenv(VIMBINDebugEnv, "invalid")
		defer os.Unsetenv(VIMBINDebugEnv)

		debug, trace, err := GetDebugAndTrace()

		assert.False(t, debug)
		assert.False(t, trace)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Unable to parse 'VIMBIN_DEBUG'")
	})

	t.Run("Invalid trace value", func(t *testing.T) {
		os.Setenv(VIMBINTraceEnv, "invalid")
		defer os.Unsetenv(VIMBINTraceEnv)

		debug, trace, err := GetDebugAndTrace()

		assert.False(t, debug)
		assert.False(t, trace)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Unable to parse 'VIMBIN_TRACE'")
	})

	t.Run("Debug is set", func(t *testing.T) {
		os.Setenv(VIMBINDebugEnv, "true")
		defer os.Unsetenv(VIMBINDebugEnv)

		debug, trace, err := GetDebugAndTrace()

		assert.True(t, debug)
		assert.False(t, trace)
		assert.NoError(t, err)
	})

	t.Run("Trace is set", func(t *testing.T) {
		os.Setenv(VIMBINTraceEnv, "true")
		defer os.Unsetenv(VIMBINTraceEnv)

		debug, trace, err := GetDebugAndTrace()

		assert.False(t, debug)
		assert.True(t, trace)
		assert.NoError(t, err)
	})

	t.Run("Neither debug nor trace is set", func(t *testing.T) {
		debug, trace, err := GetDebugAndTrace()

		assert.False(t, debug)
		assert.False(t, trace)
		assert.NoError(t, err)
	})
}

func TestIsInList(t *testing.T) {
	t.Run("Value is in the list", func(t *testing.T) {
		value := "apple"
		list := []string{"banana", "apple", "orange"}

		result := IsInList(value, list)

		assert.True(t, result)
	})

	t.Run("Value is not in the list", func(t *testing.T) {
		value := "grape"
		list := []string{"banana", "apple", "orange"}

		result := IsInList(value, list)

		assert.False(t, result)
	})

	t.Run("Empty list", func(t *testing.T) {
		value := "grape"
		var list []string

		result := IsInList(value, list)

		assert.False(t, result)
	})
}

func TestExtractHostAndPort(t *testing.T) {
	t.Run("Valid address", func(t *testing.T) {
		address := "localhost:8080"

		host, port, err := ExtractHostAndPort(address)

		assert.NoError(t, err)
		assert.Equal(t, "localhost", host)
		assert.Equal(t, 8080, port)
	})

	t.Run("Invalid address", func(t *testing.T) {
		address := "invalid"

		_, _, err := ExtractHostAndPort(address)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing port in address")
	})

	t.Run("Invalid port", func(t *testing.T) {
		address := "localhost:invalid"

		_, _, err := ExtractHostAndPort(address)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "parsing \"invalid\"")
	})
}

func TestCreateHTTPClient(t *testing.T) {
	t.Run("Create HTTP client without insecure skip verify", func(t *testing.T) {
		client := CreateHTTPClient(false)

		assert.NotNil(t, client)
		assert.Nil(t, client.Transport)
	})

	t.Run("Create HTTP client with insecure skip verify", func(t *testing.T) {
		client := CreateHTTPClient(true)

		assert.NotNil(t, client)
		assert.True(t, client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
	})
}

func TestGenerateRandomToken(t *testing.T) {
	t.Run("Generate random token with length 16", func(t *testing.T) {
		token, err := GenerateRandomToken(16)

		assert.Nil(t, err)
		assert.Equal(t, 16, len(token))
	})

	t.Run("Generate random token with length 32", func(t *testing.T) {
		token, err := GenerateRandomToken(32)

		assert.Nil(t, err)
		assert.Equal(t, 32, len(token))
	})

	t.Run("Generate random token with length 64", func(t *testing.T) {
		token, err := GenerateRandomToken(64)

		assert.Nil(t, err)
		assert.Equal(t, 64, len(token))
	})

	t.Run("Error on token generation with invalid length", func(t *testing.T) {
		_, err := GenerateRandomToken(-1)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "Invalid token length '-1'. Must be at minimum 1")
	})
}
