package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tcluri/microblog-backend-go/internal/database"
)

func (apiCfg apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	getPath := r.URL.Path
	getEmail := strings.TrimPrefix(getPath, "/users/")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	dataJson := database.User{}
	err = json.Unmarshal(body, &dataJson)
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshall the body into a json
	user, err := apiCfg.dbClient.UpdateUser(getEmail, dataJson.Password, dataJson.Name, dataJson.Age)
	if err != nil {
		fmt.Println(err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}
