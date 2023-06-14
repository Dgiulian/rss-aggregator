package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/giulian/rssaggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Using the _ we can import a module even if we're not using it in the file
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in the environment")

	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB Url not found in the environment")

	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to Database")
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	// `rssFeed` is a variable that is being assigned the result of the `urlToFeed` function, which is
	// expected to return an RSS feed parsed from the URL provided in the `feed` parameter. The contents
	// of the `rssFeed` variable are not used in the provided code snippet.
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	go startScrapping(db, 10, 1*time.Minute)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handleGetUser))

	v1Router.Get("/feed", apiCfg.handleGetFeeds)
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	v1Router.Delete("/feed/{feedId}", apiCfg.handlerDeleteFeed)

	v1Router.Get("/user/follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))
	v1Router.Post("/user/follows", apiCfg.middlewareAuth(apiCfg.handlerFollowFeed))
	v1Router.Delete("/user/follows", apiCfg.middlewareAuth(apiCfg.handlerUnfollowFeed))

	v1Router.Get("/user/posts", apiCfg.middlewareAuth(apiCfg.handleGetPosts))

	v1Router.Delete("/feed_follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFollowFeed))

	router.Mount("/v1", v1Router)

	fmt.Printf("Server started on port %v\n", portString)

	log.Fatal(srv.ListenAndServe())
}
