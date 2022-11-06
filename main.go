package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tcluri/microblog-backend-go/internal/database"
)

type apiConfig struct {
	dbClient database.Client
}

func main() {
	// Setup a new database
	const databasePath = "db.json"
	createDB := database.NewClient(databasePath)
	err := createDB.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig {
		dbClient: createDB,
	}

	m := http.NewServeMux()

	m.HandleFunc("/", testHandler)
	m.HandleFunc("/err", testErrHandler)
	m.HandleFunc("/users", apiCfg.endpointUsersHandler)
	m.HandleFunc("/users/", apiCfg.endpointUsersHandler)
	m.HandleFunc("/posts", apiCfg.endpointPostsHandler)
	m.HandleFunc("/posts/", apiCfg.endpointPostsHandler)
	// m.HandleFunc("/users/{EMAIL}", apiCfg.endpointUsersHandler)

	const addr = "localhost:8080"
	srv := http.Server{
		Handler: m,
		Addr: addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on ", addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

type errorBody struct {
	Error string `json:"error"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("server error"))
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("don't call respondWithError with a nil err!")
		return
	}
	log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)
			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "error marshalling",
			})
			w.Write(response)
			return
		}
		w.WriteHeader(code)
		w.Write(response)
	}
}
