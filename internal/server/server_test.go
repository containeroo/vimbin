package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("Run server and stop gracefully on SIGTERM", func(t *testing.T) {
		// Use a random port for testing
		testPort := "127.0.0.1:0"

		// Run the server in a goroutine
		go Run(testPort, "token")

		// Allow some time for the server to start
		time.Sleep(500 * time.Millisecond)

		// Send SIGTERM to stop the server gracefully
		stopServerGracefully()

		// Allow some time for the server to shut down
		time.Sleep(500 * time.Millisecond)
	})
}

// stopServerGracefully stops the server gracefully by sending SIGTERM.
func stopServerGracefully() {
	// Find the server process
	process, err := os.FindProcess(os.Getpid())
	if err != nil {
		panic(err)
	}

	// Send SIGTERM to stop the server gracefully
	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		panic(err)
	}
}

func TestNotFoundHandler(t *testing.T) {
	t.Run("Custom 404 handler returns a 404 response", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/nonexistent", nil)
		notFoundHandler(recorder, request)

		// Check if the response has a 404 status code
		assert.Equal(t, http.StatusNotFound, recorder.Code)

		// Check if the response body contains the expected content
		expectedContent := "<h1>404 Not Found</h1>"
		assert.Contains(t, recorder.Body.String(), expectedContent)
	})
}

func TestNewRouter(t *testing.T) {
	// Mock handler for testing
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	// Mock handlers for testing
	Handlers = []Handler{
		{
			Path:        "/mock",
			Handler:     mockHandler,
			Methods:     []string{"GET"},
			NeedsToken:  false,
			Description: "Mock handler without token",
		},
		{
			Path:        "/mock-with-token",
			Handler:     mockHandler,
			Methods:     []string{"GET"},
			NeedsToken:  true,
			Description: "Mock handler with token",
		},
	}

	// Set up the router
	router := newRouter("mock-token")

	// Test handler without token
	t.Run("Handler without token", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/mock", nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		// Check the status code of the response
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
	})

	// Test handler with valid token
	t.Run("Handler with valid token", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/mock-with-token", nil)
		request.Header.Set("X-API-Token", "mock-token")
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		// Check the status code of the response
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
	})

	// Test handler with invalid token
	t.Run("Handler with invalid token", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/mock-with-token", nil)
		request.Header.Set("X-API-Token", "invalid-token")
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		// Check the status code of the response
		assert.Equal(t, http.StatusUnauthorized, responseRecorder.Code)
	})

	// Test static file serving
	t.Run("Static file serving", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/static/css/vimbin.css", nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		// Check the status code of the response
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
	})

	// Test 404 handler
	t.Run("404 handler", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/non-existent", nil)
		responseRecorder := httptest.NewRecorder()

		router.ServeHTTP(responseRecorder, request)

		// Check the status code of the response
		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	})
}
