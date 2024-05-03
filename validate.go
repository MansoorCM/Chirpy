package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
		resp := ValidateSuccess{getCleanedBody(chirp.Body)}
		respondWithJson(w, 200, resp)
	}

}

func getCleanedBody(body string) string {
	badWords := map[string]bool{"kerfuffle": true, "sharbert": true, "fornax": true}
	words := strings.Split(body, " ")
	for i, word := range words {
		lowercaseWord := strings.ToLower(word)
		if _, ok := badWords[lowercaseWord]; ok {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
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
	CleanedBody string `json:"cleaned_body"`
}
