package model

//Event struct
type Event struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	Amount      int    `json:"amount"`
}
