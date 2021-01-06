package model

// Response struct
type Response struct {
	Destination *ResponseBody `json:"destination,omitempty"`
	Origin      *ResponseBody `json:"origin,omitempty"`
}

//ResponseBody struct
type ResponseBody struct {
	ID      string `json:"id"`
	Balance int    `json:balance`
}
