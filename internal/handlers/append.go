package handlers

import (
	"io"
	"net/http"
	"os"
	"vimbin/internal/config"
	"vimbin/internal/server"

	"github.com/rs/zerolog/log"
)

func init() {
	server.Register("/append", "Append content to storage file", true, Append, "POST")
}

// Append handles HTTP requests for appending content to a file.
//
// This function creates and injects specific functions for appending content,
// checking if content has changed, merging old and new content, and updating
// content in storage. It then calls the handleContentRequest function to process
// the HTTP request.
//
// Parameters:
//   - w: http.ResponseWriter
//     The HTTP response writer.
//   - r: *http.Request
//     The HTTP request being processed.
func Append(w http.ResponseWriter, r *http.Request) {
	// Define a function for appending content to a file
	writeFileFunc := func(filePath, content string) error {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePermission)
		if err != nil {
			return err
		}
		defer file.Close()

		log.Trace().Msgf("Writing content to file '%s': %s", filePath, content)

		_, err = io.WriteString(file, content+"\n")
		return err
	}

	// Define a function to check if content has changed
	hasContentChangedFunc := func(oldContent, newContent string) bool {
		hasContentChanged := (oldContent + newContent) != oldContent
		log.Trace().Msgf("Has content changed? %t", hasContentChanged)

		return hasContentChanged
	}

	// Define a function to merge old and new content
	mergeContentFunc := func(oldContent, newContent string) string {
		mergedContent := oldContent + newContent
		log.Trace().Msgf("Merging content: %s", mergedContent)

		return mergedContent
	}

	// Define a function for updating content in storage
	saveContentFunc := (*config.Content).Append

	// Process the HTTP request using the defined functions
	handleContentRequest(w, r, writeFileFunc, hasContentChangedFunc, mergeContentFunc, saveContentFunc)
}
