package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
	chirpBody := &ChirpBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&chirpBody)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters")
		return
	}

	cleaned, err := chirpsValidate(chirpBody.Body)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	chirp, err := cfg.DB.CreateChirp(cleaned)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp")
		return
	}

	respondWithJson(w, http.StatusCreated, Chirp{
		Id:   chirp.ID,
		Body: chirp.Body})
}
