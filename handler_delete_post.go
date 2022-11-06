package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	getPath := r.URL.Path
	getRouteParam := strings.TrimPrefix(getPath, "/posts/")
	err := apiCfg.dbClient.DeletePost(getRouteParam)
	if err != nil {
		fmt.Println(err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}
