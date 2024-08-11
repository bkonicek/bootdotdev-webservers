package main

import (
	"log"
	"net/http"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func main() {
	const port = ":8080"
	const rootPath = "."

	mux := http.NewServeMux()
	mux.Handle("/app/*", http.StripPrefix("/app", http.FileServer(http.Dir(rootPath))))
	mux.Handle("/healthz", http.HandlerFunc(readinessHandler))

	server := &http.Server{
		Handler: mux,
		Addr:    port,
	}

	log.Printf("serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
