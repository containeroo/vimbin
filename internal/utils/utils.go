package utils

import (
	"crypto/tls"
	"net"
	"net/http"
	"strconv"
)

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
