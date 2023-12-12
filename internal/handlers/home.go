package handlers

import (
	"html/template"
	"net/http"
	"vimbin/internal/config"
	"vimbin/internal/server"
)

func init() {
	server.Register("/", Home, "Home site with editor", "GET")
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
	LogRequest(r)

	page := Page{
		Title:   "vimbin - a pastebin with vim motion",
		Content: config.App.Storage.Content.Get(),
	}

	tmpl, err := template.New("index").Parse(string(config.App.HtmlTemplate))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
