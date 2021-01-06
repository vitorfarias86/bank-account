package db

import (
	"fmt"
)

//Database struct
type Database struct {
	Data map[int]int
}

//Initialize reset the db
func (d *Database) Initialize() *Database {
	fmt.Println("Initialize called")
	return &Database{}
}

//GetBalance return the balance by accountId
func (d *Database) GetBalance(accountID int) (int, error) {
	balance := d.Data[accountID]
	// if balance == nil {

	// 	return -1, fmt.Errorf("account %d does not exist", accountID)
	// }
	return balance, nil
}
