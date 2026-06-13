package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	html :=
		`<doctype html>
        <html>
	<head>
		<title>Hello Gopher</title>
	</head>
	<body>
		<b>Hello Gopher!</b>
        <p>
            <a href="/welcome">Welcome</a> |  <a href="/message">Message</a>
        </p>
	</body>
</html>`
	_, err := fmt.Fprintf(w, html)
	if err != nil {
		return
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to Go Web Programming\n")
	if err != nil {
		return
	}
}
func message(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "net/http package is used to build web apps\n")
	if err != nil {
		return
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "from Handler method\n")
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/message", message)
	mux.Handle("/test", http.HandlerFunc(test))
	fmt.Println("Listening...")
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux, // handler to invoke, http.DefaultServeMux if nil
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
