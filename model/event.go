package model

//Event struct
type Event struct {
	Type        string `json:"type"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Amount      int    `json:"amount"`
}
