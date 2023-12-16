package server

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

// ApiTokenMiddleware checks for the presence and validity of the API token.
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
