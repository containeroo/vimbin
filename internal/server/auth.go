package server

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

// ApiTokenMiddleware checks for the presence and validity of the API token.
func ApiTokenMiddleware(next http.HandlerFunc, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiToken := r.Header.Get("X-API-Token")
		if apiToken != token {
			log.Error().Msgf("Unauthorized API token: %s", apiToken)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
