package utils

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
)

// Environment variable names
const (
	VIMBINDebugEnv = "VIMBIN_DEBUG"
	VIMBINTraceEnv = "VIMBIN_TRACE"
)

// GetDebugAndTrace retrieves the debug and trace flags from environment variables.
//
// The function checks the "VIMBIN_DEBUG" and "VIMBIN_TRACE" environment variables,
// parses their values, and returns the corresponding boolean flags.
//
// Returns:
//   - debug: bool
//     True if "VIMBIN_DEBUG" is set to "true", false otherwise.
//   - trace: bool
//     True if "VIMBIN_TRACE" is set to "true", false otherwise.
//   - err: error
//     An error if there was an issue parsing the environment variables or if
//     both "VIMBIN_DEBUG" and "VIMBIN_TRACE" are set simultaneously.
func GetDebugAndTrace() (debug bool, trace bool, err error) {
	// Check and parse VIMBIN_DEBUG environment variable
	if debugEnv := os.Getenv(VIMBINDebugEnv); debugEnv != "" {
		debug, err = strconv.ParseBool(debugEnv)
		if err != nil {
			return false, false, fmt.Errorf("Unable to parse '%s'. %s", VIMBINDebugEnv, err)
		}
	}

	// Check and parse VIMBIN_TRACE environment variable
	if traceEnv := os.Getenv(VIMBINTraceEnv); traceEnv != "" {
		trace, err = strconv.ParseBool(traceEnv)
		if err != nil {
			return false, false, fmt.Errorf("Unable to parse '%s'. %s", VIMBINTraceEnv, err)
		}
	}

	// Check for mutual exclusivity of debug and trace
	if debug && trace {
		return false, false, fmt.Errorf("'%s' and '%s' are mutually exclusive", VIMBINDebugEnv, VIMBINTraceEnv)
	}

	return debug, trace, nil
}

// IsInList checks if a value is in a list.
//
// Parameters:
//   - value: string
//     The value to check for in the list.
//   - list: []string
//     The list of strings to check against.
//
// Returns:
//   - bool
//     True if the value is found in the list, false otherwise.
func IsInList(value string, list []string) bool {
	for _, v := range list {
		if value == v {
			return true
		}
	}

	return false
}

// ExtractHostAndPort extracts the hostname and port from the listen address.
//
// Parameters:
//   - address: string
//     The listen address in the format "host:port".
//
// Returns:
//   - string
//     The extracted hostname.
//   - int
//     The extracted port.
//   - error
//     An error if the extraction process fails.
func ExtractHostAndPort(address string) (string, int, error) {
	host, portStr, err := net.SplitHostPort(address)
	if err != nil {
		return "", 0, err
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, err
	}

	return host, port, nil
}

// CreateHTTPClient creates an HTTP client with optional insecure skip verify.
//
// Parameters:
//   - insecureSkipVerify: bool
//     If true, the client is created with InsecureSkipVerify for TLS, allowing connections to servers
//     with self-signed certificates or certificates that cannot be verified. Use with caution.
//
// Returns:
//   - *http.Client
//     An HTTP client configured based on the insecureSkipVerify parameter.
func CreateHTTPClient(insecureSkipVerify bool) *http.Client {
	var httpClient *http.Client

	// Configure the client with InsecureSkipVerify if needed
	if insecureSkipVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: tr}
	} else {
		httpClient = &http.Client{}
	}

	return httpClient
}

// GenerateRandomToken generates a random token of the specified length.
//
// Parameters:
//   - length: int
//     The desired length of the generated token.
//
// Returns:
//   - string
//     A random token of the specified length.
//   - error
//     An error if the random token generation fails.
func GenerateRandomToken(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("Invalid token length '%d'. Must be at minimum 1", length)
	}
	// Calculate the number of bytes needed to create the token
	numBytes := (length * 3) / 4

	// Generate random bytes
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode random bytes to base64 to create the token
	token := base64.URLEncoding.EncodeToString(randomBytes)

	// Trim the padding '=' characters
	token = token[:length]

	return token, nil
}
