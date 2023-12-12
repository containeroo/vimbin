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
		go Run(testPort)

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

func TestNewRouter(t *testing.T) {
	t.Run("Router configuration", func(t *testing.T) {
		// Set up a router
		Handlers = []Handler{
			{
				Path:        "/example",
				Description: "Example handler description",
				Handler: func(r http.ResponseWriter, w *http.Request) {
					r.Write([]byte("Test file content"))
					r.Header().Set("Content-Type", "text/plain")
					r.WriteHeader(http.StatusOK)
				},
				Methods: []string{"GET"},
			},
		}

		router := newRouter()

		// Create a test request
		req := httptest.NewRequest("GET", "/example", nil)
		w := httptest.NewRecorder()

		// Serve the request using the router
		router.ServeHTTP(w, req)

		// Check if the request is handled correctly
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Test file content")

		// Cleanup: clear the Handlers slice
		Handlers = nil
	})
}
