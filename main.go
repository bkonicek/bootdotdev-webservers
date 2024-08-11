package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	const port = ":8080"
	const rootPath = "."

	apiMetrics := &apiConfig{}

	mux := http.NewServeMux()
	mux.Handle("/app/*", http.StripPrefix(
		"/app", apiMetrics.middlewareMetricsInc(
			http.FileServer(http.Dir(rootPath)),
		),
	))
	mux.HandleFunc("GET /healthz", readinessHandler)
	mux.HandleFunc("GET /metrics", apiMetrics.metricsHandler)
	mux.HandleFunc("/reset", apiMetrics.metricsResetHandler)

	server := &http.Server{
		Handler: mux,
		Addr:    port,
	}

	log.Printf("serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
