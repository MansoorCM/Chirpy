package main

import (
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			Id:   dbChirp.ID,
			Body: dbChirp.Body,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].Id < chirps[j].Id
	})

	respondWithJson(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		respondWithError(w, 404, "query id is invalid.")
	}
	chirp, err := cfg.DB.GetChirp(idInt)

	if err != nil {
		respondWithError(w, 404, "chirp not found")
	}

	// chirpWithGivenId := []Chirp{}
	// chirpWithGivenId = append(chirpWithGivenId, Chirp{Id: chirp.ID, Body: chirp.Body})
	respondWithJson(w, http.StatusOK, chirp)
}
