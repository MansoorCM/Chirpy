package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func chirpsValidate(w http.ResponseWriter, r *http.Request) {
	chirp := ChirpBody{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chirp)
	if err != nil {
		log.Printf("error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	if len(chirp.Body) > 140 {
		resp := ValidateError{Error: "chirp is too long"}
		respondWithJson(w, 400, resp)
	} else {
		resp := ValidateSuccess{true}
		respondWithJson(w, 200, resp)
	}

}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error marshalling json: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

type ChirpBody struct {
	Body string `json:"body"`
}

type ValidateError struct {
	Error string `json:"error"`
}

type ValidateSuccess struct {
	Valid bool `json:"valid"`
}
