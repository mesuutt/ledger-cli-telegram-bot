package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mesuutt/ledger/ledger"
	"github.com/shopspring/decimal"
)

type AccountList struct {
	Accounts []string
	Length   int
}

func AccountListHandler(c *gin.Context) {
	// https://golang.org/doc/effective_go.html#interface_conversions
	// Type switching from interface to struct
	user := c.MustGet("user").(ledger.User)
	accounts := user.GetAccounts()
	response := &AccountList{
		Accounts: accounts,
		Length:   len(accounts),
	}

	json, _ := json.Marshal(response)

	c.String(http.StatusOK, string(json))
}

func AddTransactionHandler(c *gin.Context) {
	// https://golang.org/doc/effective_go.html#interface_conversions
	// Type switching from interface to struct
	user := c.MustGet("user").(ledger.User)
	j := user.GetJournal()

	var payload TransactionPayload
	c.BindJSON(&payload)

	amount, _ := decimal.NewFromString(payload.Amount)
	if payload.Date == "" {
		payload.Date = time.Now().Format("2006/01/02")
	}

	transaction := &ledger.Transaction{
		FromAccount: &ledger.Account{Name: payload.From},
		ToAccount:   &ledger.Account{Name: payload.To},
		Amount:      amount,
		Payee:       payload.Payee,
		Date:        payload.Date,
	}

	j.AddTransaction(transaction)

	c.String(http.StatusOK, "")
}
