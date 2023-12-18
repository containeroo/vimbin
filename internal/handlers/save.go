package handlers

import (
	"net/http"
	"os"
	"vimbin/internal/config"
	"vimbin/internal/server"

	"github.com/rs/zerolog/log"
)

func init() {
	server.Register("/save", "Save content to storage file", true, Save, "POST")
}

// Save handles HTTP requests for saving content to a file.
//
// This function creates and injects specific functions for saving content,
// checking if content has changed, merging old and new content, and updating
// content in storage. It then calls the handleContentRequest function to process
// the HTTP request.
//
// Parameters:
//   - w: http.ResponseWriter
//     The HTTP response writer.
//   - r: *http.Request
//     The HTTP request be processed.
func Save(w http.ResponseWriter, r *http.Request) {
	// Define a function for saving content to a file
	writeFileFunc := func(filePath, content string) error {
		log.Trace().Msgf("Writing content to file '%s': %s", filePath, content)

		return os.WriteFile(filePath, []byte(content), filePermission)
	}

	// Define a function to check if content has changed
	hasContentChangedFunc := func(oldContent, newContent string) bool {
		hasContentChanged := oldContent != newContent
		log.Trace().Msgf("Has content changed? %t", hasContentChanged)

		return hasContentChanged
	}

	// Define a function to merge old and new content
	mergeContentFunc := func(oldContent, newContent string) string {
		mergedContent := newContent
		log.Trace().Msgf("Merging content: %s", mergedContent)

		return mergedContent
	}

	// Define a function for updating content in storage
	saveContentFunc := (*config.Content).Set

	// Process the HTTP request using the defined functions
	handleContentRequest(w, r, writeFileFunc, hasContentChangedFunc, mergeContentFunc, saveContentFunc)
}
