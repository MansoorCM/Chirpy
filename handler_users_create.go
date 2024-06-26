package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	userBody := &UserBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userBody)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(userBody.Email, userBody.Password)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create user")
		return
	}

	respondWithJson(w, http.StatusCreated, UserResp{Id: user.ID, Email: user.Email})
}

type UserResp struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}
