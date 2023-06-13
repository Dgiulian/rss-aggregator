package main

import (
	"fmt"
	"net/http"

	"github.com/giulian/rssaggregator/internal/auth"
	"github.com/giulian/rssaggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Auth error: %v", err))
			return
		}
		handler(w, r, user)
	}
}
