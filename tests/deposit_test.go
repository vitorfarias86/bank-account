package tests

import (
	"testing"

	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/factory"

	"github.com/vitorfarias86/bank-account/model"
)

func TestDeposit(t *testing.T) {

	t.Run("Should deposit in an acount", func(t *testing.T) {
		dao := db.Database{Data: make(map[string]int)}

		event := model.Event{Type: "deposit", Destination: "100", Amount: 10}
		factory.Command[event.Type].Handle(&event, &dao)

		actual := event.Amount
		expected, _ := dao.GetBalance(event.Destination)

		if actual != expected {
			t.Errorf("got %d want %d", actual, expected)
		}

	})

	t.Run("Should Create an account if not exist", func(t *testing.T) {
		dao := db.Database{Data: make(map[string]int)}

		actual := dao.AccountExist("1234")

		expected := dao.CreateAccount("1234")

		if !actual != expected {
			t.Errorf("got %t want %t", actual, expected)
		}

	})

	t.Run("Shouldn't Create an account if exist", func(t *testing.T) {
		dao := db.Database{Data: make(map[string]int)}
		// Given an account created
		_ = dao.CreateAccount("1234")

		// expect
		expected := dao.CreateAccount("1234")

		if expected {
			t.Errorf("Account was created ")
		}

	})
}
