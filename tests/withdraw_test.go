package tests

import (
	"testing"

	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/factory"
	"github.com/vitorfarias86/bank-account/model"
)

func TestWithdraw(t *testing.T) {
	t.Run("Withdraw from non-existing account should return error", func(t *testing.T) {
		dao := db.Database{Data: make(map[string]int)}
		event := model.Event{Type: "withdraw", Destination: "a1b", Amount: 10}
		_, err := factory.Command[event.Type].Handle(&event, &dao)

		if err == nil {
			t.Errorf("Withdraw from non-existing account should return error")
		}
	})

	t.Run("Withdraw from existing account ", func(t *testing.T) {
		dao := db.Database{Data: make(map[string]int)}
		eventDeposit := model.Event{Type: "deposit", Destination: "100", Amount: 10}
		factory.Command[eventDeposit.Type].Handle(&eventDeposit, &dao)

		event := model.Event{Type: "withdraw", Destination: "100", Amount: 5}
		result, _ := factory.Command[event.Type].Handle(&event, &dao)

		actualOrigin := result.Origin.ID
		actualBalance := result.Origin.Balance

		expectedOrigin := "100"
		expectedBalance := 5
		if actualOrigin != expectedOrigin && actualBalance != expectedBalance {
			t.Errorf("Withdraw from existing account error")
		}
	})
}
