package main

import (
	"fmt"
	"log"
	"net/http"
)

func middlewareFirst(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareFirst - Before Handler")
		next.ServeHTTP(w, r)
		log.Println("MiddlewareFirst - After Handler")
	})
}

func middlewareSecond(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareSecond - Before Handler")
		if r.URL.Path == "/message" {
			if r.URL.Query().Get("password") == "pass123" {
				log.Println("Authorized to the system")
				next.ServeHTTP(w, r)
			} else {
				log.Println("Failed to authorize to the system")
				return
			}
		} else {
			next.ServeHTTP(w, r)
		}

		log.Println("MiddlewareSecond - After Handler")
	})
}

func index(w http.ResponseWriter, _ *http.Request) {
	log.Println("Executing index Handler")
	_, err := fmt.Fprintf(w, "Welcome")
	if err != nil {
		return
	}
}
func message(w http.ResponseWriter, _ *http.Request) {
	log.Println("Executing message Handler")
	_, err := fmt.Fprintf(w, "HTTP Middleware is awesome")
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/message", message)
	server := &http.Server{
		Addr:    ":8080",
		Handler: middlewareFirst(middlewareSecond(mux)),
	}
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
