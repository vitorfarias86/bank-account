package strategy

import (
	"github.com/vitorfarias86/bank-account/db"
	"github.com/vitorfarias86/bank-account/model"
)

//EventStrategy interface
type EventStrategy interface {
	Handle(evt *model.Event, db *db.Database) (model.Response, error)
}
