package strategy

import (
	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/model"
)

//Deposit struct
type Deposit struct{}

//Handle func
func (f *Deposit) Handle(evt *model.Event, db *db.Database) (model.Response, error) {

	balance := db.Deposit(evt)
	var response = model.Response{Destination: &model.ResponseBody{ID: evt.Destination, Balance: balance}}

	return response, nil
}
