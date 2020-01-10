package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", retrieveLeaderboard).Methods("GET")
	router.HandleFunc("/new", newEntry).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func retrieveLeaderboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working!")
}

func newEntry(w http.ResponseWriter, r *http.Request) {
	return
}