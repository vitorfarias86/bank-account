package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorfarias86/bank-account/db"
)

func main() {

	var database db.Database = db.Database{}

	router := mux.NewRouter()
	router.HandleFunc("/reset", ResetData(&database)).Methods("POST")
	router.HandleFunc("/balance/{account_id}", GetBalance).Methods("GET")
	router.HandleFunc("/event", CreateEvent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

//ResetData reset all in memory data
func ResetData(database *db.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		database.Initialize()
	}
}

//GetBalance get balance
func GetBalance(w http.ResponseWriter, r *http.Request) {}

//CreateEvent create an event if the account doesn't exists create a new one
func CreateEvent(w http.ResponseWriter, r *http.Request) {}
