package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		apiCfg.handlerCreateUser(w, r)
	case http.MethodGet:
		apiCfg.handlerGetUser(w, r)
	case http.MethodPut:
	 	apiCfg.handlerUpdateUser(w, r)
	case http.MethodDelete:
	 	apiCfg.handlerDeleteUser(w, r)
	default:
		respondWithError(w, 400, errors.New("method not supported"))
	}
}
