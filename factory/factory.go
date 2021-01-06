package factory

import (
	"github.com/vitorfarias86/bank-account/model"
	"github.com/vitorfarias86/bank-account/strategy"
)

//Command variable
var Command map[string]strategy.EventStrategy = map[string]strategy.EventStrategy{"deposit": &strategy.Deposit{}, "withdraw": &strategy.Withdraw{}, "transfer": &strategy.Transfer{}}

//Factory func
func Factory(event model.Event) strategy.EventStrategy {

	return Command[event.Type]
}
