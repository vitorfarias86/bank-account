package model

import "github.com/vitorfarias86/bank-account/db"

//Event struct
type Event struct {
	Type        string `json:"type"`
	Destination int    `json:"destination"`
	Amount      int    `json:"amount"`
}

//Deposit function
func Deposit(event *Event, db *db.Database) {

}
