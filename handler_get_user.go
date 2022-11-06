package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	getPath := r.URL.Path
	getEmail := strings.TrimPrefix(getPath, "/users/")
	user, err := apiCfg.dbClient.GetUser(getEmail)
	if err != nil {
		fmt.Println(err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}
