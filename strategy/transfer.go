package strategy

import (
	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/model"
)

//Transfer struct
type Transfer struct{}

//Handle func
func (f *Transfer) Handle(evt *model.Event, db *db.Database) (model.Response, error) {

	err := db.Transfer(evt)
	if err != nil {
		return model.Response{}, err
	}

	balanceOrigin, _ := db.GetBalance(evt.Origin)
	balanceDestination, _ := db.GetBalance(evt.Destination)

	var response = model.Response{Origin: &model.ResponseBody{ID: evt.Origin, Balance: balanceOrigin}, Destination: &model.ResponseBody{ID: evt.Destination, Balance: balanceDestination}}

	return response, nil
}
