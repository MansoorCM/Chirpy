package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerUsersLogin(w http.ResponseWriter, r *http.Request) {
	userBody := &UserBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(userBody)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters.")
	}

	user, err := cfg.DB.UserLogin(userBody.Email, userBody.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "passwords do not match.")
	}

	respondWithJson(w, http.StatusOK, UserResp{Id: user.ID, Email: user.Email})
}
