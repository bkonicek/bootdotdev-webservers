package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func validateChirpHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Valid bool   `json:"valid,omitempty"`
		Error string `json:"error,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		log.Printf("Error decoding parameters: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	resp := returnVals{}
	// Check chirp length doesn't exceed limit
	if len(params.Body) > 140 {
		resp.Error = "Chirp is too long"
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Valid = true
		w.WriteHeader(http.StatusOK)
	}
	dat, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error encoding response: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.Write(dat)
}
