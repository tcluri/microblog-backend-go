package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfg.handlerRetrievePosts(w, r)
	case http.MethodPost:
		apiCfg.handlerCreatePost(w, r)
	case http.MethodDelete:
	 	apiCfg.handlerDeletePost(w, r)
	default:
		respondWithError(w, 400, errors.New("method not supported"))
	}
}
