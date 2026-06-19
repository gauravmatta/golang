package main

import (
	apphttp "customerManager/http/servemux"
	"customerManager/memstore"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	repo, err := memstore.NewInmemoryRepository()
	if err != nil {
		log.Fatal("Error:", err)
	}
	h := &apphttp.CustomerHandler{
		Repository: repo, // Injecting dependency
	}
	router := initializeRoutes(h)
	router = cors.Default().Handler(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func initializeRoutes(h *apphttp.CustomerHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/customers", h.FindAll)
	//mux.HandleFunc("GET /api/customers/{id}", h.Get)
	mux.HandleFunc("POST /api/customers", h.Post)
	//mux.HandleFunc("PUT /api/customers/{id}", h.Put)
	//mux.HandleFunc("DELETE /api/customers/{id}", h.Delete)
	return mux
}
