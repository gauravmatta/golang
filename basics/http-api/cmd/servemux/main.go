package main

import (
	"log"
	"log/slog"

	//"log/slog"
	"net/http"

	apphttp "github.com/gauravmatta/golang/basics/http-api/http/servemux"
	"github.com/gauravmatta/golang/basics/http-api/memstore"
	"github.com/gauravmatta/golang/basics/http-api/middleware"
	"github.com/rs/cors"
)

//"github.com/gauravmatta/golang/basics/http-api/middleware")

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
	//// Adding middleware http
	router = middleware.Apply(router,
		middleware.RateLimiter(200),
		middleware.PanicRecovery(logger),
	)
	//// CORS middleware
	router = cors.Default().Handler(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	err = server.ListenAndServe()
}

func initializeRoutes(h *apphttp.NoteHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/notes", h.GetAll)
	mux.HandleFunc("GET /api/notes/{id}", h.Get)
	mux.HandleFunc("POST /api/notes", h.Post)
	mux.HandleFunc("PUT /api/notes/{id}", h.Put)
	mux.HandleFunc("DELETE /api/notes/{id}", h.Delete)
	return mux
}
