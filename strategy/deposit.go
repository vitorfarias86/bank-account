package strategy

import (
	"fmt"

	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/model"
)

//Deposit struct
type Deposit struct{}

//Handle func
func (f *Deposit) Handle(evt *model.Event, db *db.Database) {
	fmt.Print("Deposit called")
}
