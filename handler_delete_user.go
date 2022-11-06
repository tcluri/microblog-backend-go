package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	getPath := r.URL.Path
	getEmail := strings.TrimPrefix(getPath, "/users/")
	err := apiCfg.dbClient.DeleteUser(getEmail)
	if err != nil {
		fmt.Println(err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}
