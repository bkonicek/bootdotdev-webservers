package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	templateFilePath := templateDir + "metrics.html"
	tmpl, err := os.ReadFile(templateFilePath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(string(tmpl), cfg.fileserverHits)))
}

func (cfg *apiConfig) metricsResetHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}
