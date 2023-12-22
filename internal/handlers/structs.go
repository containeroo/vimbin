package handlers

// Page represents the data structure that will be passed to the HTML template.
//
// This struct holds information about the title and content of a page, which can be
// utilized by the HTML template to render dynamic content.
type Page struct {
	Title      string // Title is the title of the page.
	Content    string // Content is the content of the page.
	Token      string // Token is the API token.
	Theme      string // Theme is the theme of the page.
	LightTheme string // LightTheme is the light theme of the page.
	DarkTheme  string // DarkTheme is the dark theme of the page.
	Version    string // Version is the version of the application.
}
