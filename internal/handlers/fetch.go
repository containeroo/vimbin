package handlers

import (
	"net/http"
	"vimbin/internal/config"
	"vimbin/internal/server"

	"github.com/rs/zerolog/log"
)

func init() {
	server.Register("/fetch", Fetch, "Fetch content from storage file", "GET")
}

// Fetch handles HTTP requests for fetching content.
//
// This function logs the incoming request and retrieves content from storage.
// If content is present, it sets the appropriate HTTP headers and writes
// the content to the response. If no content is found, it returns a
// HTTP status code of No Content.
//
// Parameters:
//   - w: http.ResponseWriter
//     The HTTP response writer.
//   - r: *http.Request
//     The HTTP request being processed.
func Fetch(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)

	w.Header().Set("Content-Type", "application/text")

	content := config.App.Storage.Content.Get()
	if len(content) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	contentBytes := []byte(content)
	if _, err := w.Write(contentBytes); err != nil {
		log.Error().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
