// Module demo - example Go module
package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	_, _ = sql.Open("postgres", "user=postgres password=postgres dbname=golang sslmode=disable")
	fmt.Println("Module demo")
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		return
	}
}
