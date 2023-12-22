package handlers

import (
	"net/http"
	"vimbin/internal/config"
	"vimbin/internal/server"

	"github.com/rs/zerolog/log"
)

func init() {
	server.Register("/fetch", "Fetch content from storage file", true, Fetch, "GET")
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
	log.Trace().Msg(generateHTTPRequestLogEntry(r))

	w.Header().Set("Content-Type", "text/plain")

	content := config.App.Storage.Content.Get()
	if len(content) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if _, err := w.Write([]byte(content)); err != nil {
		log.Error().Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
