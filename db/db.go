package db

import (
	"fmt"

	"github.com/vitorfarias86/bank-account/model"
)

//Database struct
type Database struct {
	Data map[string]int
}

//Initialize reset the db
func (d *Database) Initialize() *Database {

	*d = Database{Data: make(map[string]int)}

	return d
}

//GetBalance return the balance by accountId
func (d *Database) GetBalance(accountID string) (int, error) {
	balance, ok := d.Data[accountID]

	if !ok {

		return 0, fmt.Errorf("Account %s does not exist", accountID)
	}
	fmt.Printf("Balance %d - AccountID %s", balance, accountID)

	return balance, nil
}

//Deposit func - do the deposit and returns the new balance
func (d *Database) Deposit(evt *model.Event) int {

	accountCreated := d.CreateAccount(evt.Destination)
	if accountCreated {
		fmt.Printf("Account %d created", evt.Destination)
	}
	actualBalance, _ := d.GetBalance(evt.Destination)
	d.Data[evt.Destination] = actualBalance + evt.Amount

	return d.Data[evt.Destination]
}

//CreateAccount func
func (d *Database) CreateAccount(accountID string) bool {

	if !d.AccountExist(accountID) {
		d.Data[accountID] = 0
		return true
	}
	return false

}

//AccountExist func
func (d *Database) AccountExist(accountID string) bool {
	_, exist := d.Data[accountID]
	return exist
}

//Withdraw func
func (d *Database) Withdraw(evt *model.Event) (int, error) {

	accountExist := d.AccountExist(evt.Origin)
	if !accountExist {
		return 0, fmt.Errorf("Account %s does not exist", evt.Origin)
	}

	d.Data[evt.Origin] -= evt.Amount

	return d.Data[evt.Origin], nil
}

//Transfer func
func (d *Database) Transfer(evt *model.Event) error {

	balance, _ := d.GetBalance(evt.Origin)
	if evt.Amount > balance {
		return fmt.Errorf("Account %s does not have suficient amount for this transaction", evt.Origin)
	}
	d.Withdraw(evt)
	d.Deposit(evt)

	return nil
}
