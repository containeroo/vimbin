package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"vimbin/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Collect gathers and initializes the HTTP handlers.
//
// This function triggers the initialization functions from other files,
// populating the Handlers slice in the server package with the registered
// routes and their corresponding handler functions. It should be called
// before starting the HTTP server to ensure all handlers are registered.
func Collect() {}

// filePermission represents the default file permission used in the application.
const filePermission = 0644

// LogRequest logs the details of an HTTP request.
//
// Parameters:
//   - req: *http.Request
//     The HTTP request to log.
func LogRequest(req *http.Request) {
	if log.Logger.GetLevel() != zerolog.TraceLevel {
		return
	}

	query := strings.Map(func(r rune) rune {
		switch r {
		case '\n', '\r': // Remove newlines and carriage returns from the query string
			return -1
		case ' ': // Replace spaces with a single space
			return -1
		default: // Leave all other characters as-is
			return r
		}
	}, req.URL.RawQuery)

	log.Trace().Msgf("%s %s%s", req.Method, req.RequestURI, query)
}

// handleContentRequest handles HTTP requests for updating content.
//
// This function processes an HTTP request, decodes the JSON body, compares
// the new content to the old content, and performs the necessary actions
// based on the provided functions. It logs the request, checks for changes
// in content, writes content to a file, updates content in storage, and
// responds with appropriate JSON status messages.
//
// Parameters:
//   - w: http.ResponseWriter
//     The HTTP response writer.
//   - r: *http.Request
//     The HTTP request being processed.
//   - writeFileFunc: func(string, string) error
//     Function for writing content to a file.
//   - hasContentChangedFunc: func(string, string) bool
//     Function to check if content has changed.
//   - mergeContentFunc: func(string, string) string
//     Function to merge old and new content.
//   - saveContentFunc: func(*config.Content, string)
//     Function for updating content in storage.
//
// Note: The provided functions are injected for flexibility and can be customized
// based on specific requirements.
func handleContentRequest(
	w http.ResponseWriter,
	r *http.Request,
	writeFileFunc func(string, string) error,
	hasContentChangedFunc func(string, string) bool,
	mergeContentFunc func(string, string) string,
	saveContentFunc func(*config.Content, string),
) {
	LogRequest(r)

	// Parse JSON request body
	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		msg := fmt.Sprintf("Error decoding JSON: %v", err)
		log.Error().Msg(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	oldContent := config.App.Storage.Content.Get()

	newContent, ok := requestData["content"]
	if !ok {
		msg := "Missing 'content' field in JSON"
		log.Error().Msg(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Compare the new content to the old content
	if !hasContentChangedFunc(oldContent, newContent) {
		// Respond with JSON indicating no changes were made
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "no changes"}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			msg := fmt.Sprintf("Error marshalling response: %v", err)
			log.Error().Msg(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(jsonResponse); err != nil {
			msg := fmt.Sprintf("Error writing response: %v", err)
			log.Error().Msg(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		return
	}

	newContent = mergeContentFunc(oldContent, newContent) // Use the provided function to append or save the new content

	log.Trace().Msgf("Got new content: %s", newContent)

	// Use the provided function for writing to a file
	if err := writeFileFunc(config.App.Storage.Path, newContent); err != nil {
		msg := fmt.Sprintf("Error writing file: %v", err)
		log.Error().Msg(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// Use the provided function for updating the content in storage
	saveContentFunc(&config.App.Storage.Content, newContent)

	size := strconv.Itoa(len(newContent))
	log.Debug().Msgf("Wrote %s bytes to file '%s'", size, config.App.Storage.Path)

	// Set the X-Bytes-Written header with the number of bytes written
	w.Header().Set("X-Bytes-Written", size)

	// Respond with JSON indicating success
	response := map[string]string{"status": "success"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		msg := fmt.Sprintf("Error marshalling response: %v", err)
		log.Error().Msg(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResponse); err != nil {
		msg := fmt.Sprintf("Error writing response: %v", err)
		log.Error().Msg(msg)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
