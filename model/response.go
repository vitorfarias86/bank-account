package model

// Response struct
type Response struct {
	Origin      *ResponseBody `json:"origin,omitempty"`
	Destination *ResponseBody `json:"destination,omitempty"`
}

//ResponseBody struct
type ResponseBody struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}
