package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", //index)
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "About Page")
		})
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
