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

func (apiCfg *apiConfig) handlerFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to parse body. %v", err))
	}

	feedFollow, err := apiCfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 500, "Unable to follow that feed")
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	follows, err := apiCfg.DB.GetFollowFeed(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("errror while retrieving follows: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(follows))
}
func (apiCfg *apiConfig) handlerUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	decoder := json.NewDecoder(r.Body)
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not parse body: %v", err))
		return
	}

	follow, err := apiCfg.DB.DeleteFollowFeed(r.Context(), database.DeleteFollowFeedParams{
		UserID: user.ID,
		FeedID: params.FeedId,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("errror while retrieving follows: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedFollowToFeedFollow(follow))
}

func (apiCfg *apiConfig) handlerDeleteFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("errror while parsing the feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFollowFeedById(r.Context(), database.DeleteFollowFeedByIdParams{
		UserID: user.ID,
		ID:     feedFollowId,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("errror while retrieving follows: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{ result string }{result: "Ok"})

}
