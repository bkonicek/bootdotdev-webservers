package main

import (
	"log"
	"net/http"
)

func main() {
	const port = ":8080"
	const defaultPath = "."

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(defaultPath)))

	server := &http.Server{
		Handler: mux,
		Addr:    port,
	}

	log.Printf("listening on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
