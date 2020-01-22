package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/adrenallen/leaderboarder/datasource"
	"github.com/adrenallen/leaderboarder/leaderboard"
	"github.com/gorilla/mux"
)

var h leaderboard.Handler

func main() {
	ds := datasource.File{FilePath: "./data.record"}
	h = leaderboard.Handler{Data: ds}
	router := mux.NewRouter()
	router.HandleFunc("/", retrieveLeaderboard).Methods("GET")
	router.HandleFunc("/new", newEntry).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func retrieveLeaderboard(w http.ResponseWriter, r *http.Request) {
	entries := h.GetAll()
	entriesJSON, _ := json.Marshal(entries)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(entriesJSON)
}

func newEntry(w http.ResponseWriter, r *http.Request) {
	n := r.FormValue("name")
	s := r.FormValue("score")
	m := r.FormValue("meta")
	sFloat, err := strconv.ParseFloat(s, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = h.NewEntry(n, sFloat, m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
}
