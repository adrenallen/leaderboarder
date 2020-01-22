package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/adrenallen/leaderboarder/datasource"
	"github.com/adrenallen/leaderboarder/leaderboard"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var h leaderboard.Handler

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("%s\nFailed to load .env - using defaults", err.Error())
	}

	port := getEnvOrDefault("PORT", "9001")

	ds := getDatasource()

	h = leaderboard.Handler{Data: ds}
	router := mux.NewRouter()
	router.HandleFunc("/", retrieveLeaderboard).Methods("GET")
	router.HandleFunc("/submit", newEntry).Methods("POST")

	log.Fatal(http.ListenAndServe(":" + port, router))
}

func getDatasource() datasource.DataSource {
	//TODO - update this to check for the type of datasource
	filePath := getEnvOrDefault("DATA_FILE", "./data.record")
	return &datasource.File{FilePath: filePath}
}

func getEnvOrDefault(name string, defaultVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	return val
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
