package strategy

import (
	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/model"
)

//Withdraw struct
type Withdraw struct{}

//Handle func
func (f *Withdraw) Handle(evt *model.Event, db *db.Database) (model.Response, error) {
	balance, err := db.Withdraw(evt)

	if err != nil {
		return model.Response{}, err
	}
	var response = model.Response{Origin: &model.ResponseBody{ID: evt.Destination, Balance: balance}}
	return response, nil
}
