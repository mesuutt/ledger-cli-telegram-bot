package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mesuutt/ledger/ledger"
)

func AccountListHandler(c *gin.Context) {
	// https://golang.org/doc/effective_go.html#interface_conversions
	// Type switching from interface to struct
	user := c.MustGet("user").(ledger.User)
	accounts := user.GetAccounts()
	json, _ := json.Marshal(accounts)
	c.String(http.StatusOK, string(json))
}

func AddTransactionHandler(c *gin.Context) {
	// https://golang.org/doc/effective_go.html#interface_conversions
	// Type switching from interface to struct
	user := c.MustGet("user").(ledger.User)
	j := user.GetJournal()

	var t TransactionPayload
	c.BindJSON(&t)

	j.AddTransaction(&ledger.Transaction{
		FromAccount: &ledger.Account{Name: t.From},
		ToAccount:   &ledger.Account{Name: t.To},
		Amount:      t.Amount,
		Payee:       t.Payee,
		Date:        t.Date,
	})

	c.String(http.StatusOK, "")
}
