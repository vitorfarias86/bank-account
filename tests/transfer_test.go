package tests

import (
	"testing"

	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/factory"
	"github.com/vitorfarias86/bank-account/model"
)

func TestTransfer(t *testing.T) {

	t.Run("Should transfer to an existing account", func(t *testing.T) {
		dao := db.Database{Data: make(map[string]int)}
		dao.CreateAccount("100")
		dao.CreateAccount("200")

		eventDeposit := model.Event{Type: "deposit", Destination: "100", Amount: 10}
		factory.Command[eventDeposit.Type].Handle(&eventDeposit, &dao)

		eventTransfer := model.Event{Type: "transfer", Origin: "100", Destination: "200", Amount: 5}
		factory.Command[eventTransfer.Type].Handle(&eventTransfer, &dao)

		balanceAccount100, _ := dao.GetBalance("100")
		expected := 5
		if balanceAccount100 != expected {
			t.Errorf("got %d want %d", balanceAccount100, expected)
		}

		balanceAccount200, _ := dao.GetBalance("200")
		expected = 5
		if balanceAccount200 != expected {
			t.Errorf("got %d want %d", balanceAccount100, expected)
		}
	})

}
