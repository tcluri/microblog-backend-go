package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) handlerRetrievePosts(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("userEmail")
	if userEmail == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("no email provided to handlerRetrievePosts"))
		return
	}
	posts, err := apiCfg.dbClient.GetPosts(userEmail)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}
