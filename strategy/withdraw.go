package strategy

import (
	"fmt"

	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/model"
)

//Withdraw struct
type Withdraw struct{}

//Handle func
func (f *Withdraw) Handle(evt *model.Event, db *db.Database) {
	fmt.Print("Withdraw called")
}
