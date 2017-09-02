package api

// http://choly.ca/post/go-json-marshalling/
type TransactionPayload struct {
	From   string `json:"from" binding:"required"`
	To     string `json:"to" binding:"required"`
	Amount string `json:"amount" binding:"required"`
	Payee  string `json:"payee"`
	Date   string `json:"date"` // Payload oldugu icin string
}
