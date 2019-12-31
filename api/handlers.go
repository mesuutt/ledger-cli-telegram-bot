package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/shopspring/decimal"

	"github.com/mesuutt/teledger/ledger"
)

type AccountList struct {
	Accounts []string
	Length   int
}

func AccountListHandler(c echo.Context) error {
	// https://golang.org/doc/effective_go.html#interface_conversions
	// Type switching from interface to struct
	user := c.Get("user").(ledger.User)
	accounts := user.GetAccounts()
	response := &AccountList{
		Accounts: accounts,
		Length:   len(accounts),
	}

	json, _ := json.Marshal(response)

	return c.String(http.StatusOK, string(json))
}

func AddTransactionHandler(c echo.Context) error {
	// https://golang.org/doc/effective_go.html#interface_conversions
	// Type switching from interface to struct
	user := c.Get("user").(ledger.User)
	j := user.GetJournal()

	var payload TransactionPayload
	c.Bind(&payload)

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

	return c.String(http.StatusOK, "")
}
