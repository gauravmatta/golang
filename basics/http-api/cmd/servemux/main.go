package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/gauravmatta/golang/basics/http-api/memstore"
)

func main() {
	repo, err := memstore.NewInmemoryRepository() // With in-memory database
	if err != nil {
		log.Fatal("Error:", err)
	}
	h := &apphttp.NoteHandler{
		Repository: repo, // Injecting dependency
	}
	router := initializeRoutes(h) // configure routes

	logger := slog.Default()
	// Adding middleware http
	router = middleware.Apply(router,
		middleware.RateLimiter(200),
		middleware.PanicRecovery(logger),
	)
	// CORS middleware
	router = cors.Default().Handler(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}
