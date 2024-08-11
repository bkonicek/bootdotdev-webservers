package main

import (
	"log"
	"net/http"
)

func main() {
	const port = ":8080"

	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    port,
	}

	log.Printf("listening on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
