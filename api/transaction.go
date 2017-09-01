package api

type TransactionPayload struct {
	From   string `json:"from" binding:"required"`
	To     string `json:"to" binding:"required"`
	Amount string `json:"amount" binding:"required"`
	Payee  string `json:"payee"`
	Date   string `json:"date"`
}
