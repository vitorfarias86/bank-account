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
		fmt.Fprintf(w, fmt.Sprintf("%s", http.StatusText(http.StatusOK)))
	}
}

//GetBalance get balance
func GetBalance(database *db.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		accountID := r.URL.Query().Get("account_id")
		fmt.Printf("#v", accountID)
		balance, err := database.GetBalance(accountID)
		if err != nil {
			http.Error(w, "", http.StatusNotFound)
		}
		fmt.Fprintf(w, fmt.Sprintf("%d", balance))
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

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, fmt.Sprintf("%d", 0))
		} else {
			b, _ := json.Marshal(response)
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, fmt.Sprintf("%s", string(b)))

		}
	}
}
