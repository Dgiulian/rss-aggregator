package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/giulian/rssaggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not get feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))

}

func (apiCfg *apiConfig) handlerDeleteFeed(w http.ResponseWriter, r *http.Request) {
	feedId, err := uuid.Parse(chi.URLParam(r, "feedId"))
	if err != nil {
		respondWithError(w, 400, "feed Id not found")
		return
	}

	feed, err := apiCfg.DB.DeleteFeed(r.Context(), feedId)
	if err != nil {
		respondWithError(w, 400, "feed Id not found")
		return
	}

	respondWithJSON(w, 200, feed)
}
