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

	d = &Database{Data: make(map[string]int)}

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

//Deposit func
func (d *Database) Deposit(evt *model.Event) {

	accountCreated := d.CreateAccount(evt.Destination)
	if accountCreated {
		fmt.Printf("Account %d created", evt.Destination)
	}
	actualBalance, _ := d.GetBalance(evt.Destination)
	d.Data[evt.Destination] = actualBalance + evt.Amount
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
