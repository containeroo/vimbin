package handlers

import (
	"net/http"
	"vimbin/internal/config"
	"vimbin/internal/server"

	"github.com/rs/zerolog/log"
)

func init() {
	server.Register("/", "Home site with editor", false, Home, "GET")
}

// Home handles HTTP requests for the home page.
//
// This function logs the incoming request, retrieves content from storage,
// and renders the home page using an HTML template. It sets the page title
// and content based on the retrieved information from the storage.
//
// Parameters:
//   - w: http.ResponseWriter
//     The HTTP response writer.
//   - r: *http.Request
//     The HTTP request being processed.
func Home(w http.ResponseWriter, r *http.Request) {
	log.Trace().Msg(generateHTTPRequestLogEntry(r))

	page := Page{
		Title:      "vimbin - a pastebin with vim motion",
		Content:    config.App.Storage.Content.Get(),
		Token:      config.App.Server.Api.Token.Get(),
		Theme:      config.App.Server.Web.Theme,
		LightTheme: config.App.Server.Web.LightTheme,
		DarkTheme:  config.App.Server.Web.DarkTheme,
		Version:    config.App.Version,
	}

	if err := config.App.HtmlTemplate.Execute(w, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
