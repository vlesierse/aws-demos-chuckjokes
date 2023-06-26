package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func jsonResult(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func HandleListJokes(w http.ResponseWriter, r *http.Request) {
	jsonResult(w, ListJokes())
}

func HandleGetJokeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jokeId, _ := strconv.Atoi(vars["jokeId"])
	jsonResult(w, GetJokeById(jokeId))
}

func HandleGetRandomJoke(w http.ResponseWriter, r *http.Request) {
	jsonResult(w, GetRandomJoke())
}
