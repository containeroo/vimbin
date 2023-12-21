package server

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

// ApiTokenMiddleware is a middleware function that checks for the presence and validity of the API token.
//
// Parameters:
//   - next: http.HandlerFunc
//     The next HTTP handler in the chain.
//   - token: string
//     The expected API token for authentication.
//
// Behavior:
//
//	The middleware checks the 'X-API-Token' header in the incoming request against the provided token.
//	If the header is missing or the token is invalid, it responds with an HTTP 401 Unauthorized status.
//	If the token is valid, it calls the next handler in the chain.
func ApiTokenMiddleware(next http.HandlerFunc, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiToken := r.Header.Get("X-API-Token")

		if apiToken == "" {
			msg := "Missing API token. You must provide the API token in the X-API-Token header."
			log.Error().Msg(msg)
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		if apiToken != token {
			log.Error().Msgf("Unauthorized API token: %s", apiToken)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
