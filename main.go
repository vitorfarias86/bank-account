package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/factory"
	"github.com/vitorfarias86/bank-account/model"
)

func main() {

	database := db.Database{Data: make(map[string]int)}

	router := mux.NewRouter()
	router.HandleFunc("/reset", ResetData(&database)).Methods("POST")
	router.HandleFunc("/balance", GetBalance(&database)).Methods("GET")
	router.HandleFunc("/event", HandleEvent(&database)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

//ResetData reset all in memory data
func ResetData(database *db.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		database.Initialize()
		fmt.Fprintf(w, fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)))
	}
}

//GetBalance get balance
func GetBalance(database *db.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		accountID := r.URL.Query().Get("account_id")
		fmt.Printf("#v", accountID)
		balance, err := database.GetBalance(accountID)
		status := http.StatusOK
		if err != nil {
			status = http.StatusNotFound
			http.Error(w, "", http.StatusNotFound)
		}
		fmt.Fprintf(w, fmt.Sprintf("%d %d", status, balance))
	}
}

//HandleEvent handle all events
func HandleEvent(database *db.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var event model.Event

		err := json.NewDecoder(r.Body).Decode(&event)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		response, err := factory.Command[event.Type].Handle(&event, database)
		status := http.StatusCreated
		if err != nil {
			status = http.StatusNotFound
			http.Error(w, "", http.StatusNotFound)
			fmt.Fprintf(w, fmt.Sprintf("%d %s", status, "0"))
		} else {
			b, _ := json.Marshal(response)

			fmt.Fprintf(w, fmt.Sprintf("%d %s", status, string(b)))
			w.WriteHeader(status)
		}

	}
}
