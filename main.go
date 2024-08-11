package main

import (
	"log"
	"net/http"
)

func main() {
	const port = ":8080"
	const rootPath = "."

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(rootPath)))

	server := &http.Server{
		Handler: mux,
		Addr:    port,
	}

	log.Printf("serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
