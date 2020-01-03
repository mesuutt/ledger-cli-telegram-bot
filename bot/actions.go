package bot

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	"github.com/mesuutt/teledger/db"
	"github.com/mesuutt/teledger/ledger"
)

func SetAlias(senderID int, name, accountName string) error {
	if db.GetAccountByAlias(senderID, name) != "" {
		return errors.New(fmt.Sprintf("alias %s already exist.", name))
	}

	err := db.AddAlias(senderID, name, accountName)
	if err != nil {
		return err
	}
	user := ledger.User{Username: strconv.Itoa(senderID)}
	err = user.AddAlias(name, accountName)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAlias(senderID int, name string) error {
	user := ledger.User{Username: strconv.Itoa(senderID)}
	err := user.DeleteAlias(name)
	if err != nil {
		return err
	}

	err = db.DeleteAlias(senderID, name)
	if err != nil {
		return err
	}

	return nil
}

func AddTransaction(senderID int, text string) ([]*ledger.Transaction, error) {
	user := ledger.User{Username: strconv.Itoa(senderID)}
	j := user.GetJournal()

	match := GetRegexSubMatch(transactionRegex, text)
	if _, ok := match["amount"]; !ok {
		return nil, errors.New("Invalid transaction syntax")
	}

	var amount decimal.Decimal
	amount, err := decimal.NewFromString(match["amount"])
	if err != nil {
		keymap := map[string]string{
			"q": "1", "w": "2", "e": "3",
			"r": "4", "t": "5", "y": "6",
			"u": "7", "i": "8", "o": "9",
			"p": "0", ".": ".",
		}

		amountStr := ""
		/*
			// we are already checking by regex above.
			expr := regexp.MustCompile(`^[qwertyuiop.]+$`)
			m := expr.Match([]byte(match["amount"]))
			if !m {
				return nil, errors.New("amount must contains only numbers or [qwertyuiop.]")
			}*/

		for _, k := range match["amount"] {
			amountStr += keymap[string(k)]
		}

		amount, err = decimal.NewFromString(amountStr)
		if err != nil {
			return nil, errors.New("amount parse error: " + err.Error())
		}
	}

	if match["month"] == "" {
		match["month"] = time.Now().Format("01")
	}
	if match["day"] == "" {
		match["day"] = time.Now().Format("02")
	}

	var transactions []*ledger.Transaction
	loopCount := 1
	if match["to2"] != "" {
		loopCount += 1
	}

	for i := 0; i < loopCount; i++ {
		from := match["from"]
		to := match[fmt.Sprintf("to%d", i + 1)]
		if i > 0 {
			from = match[fmt.Sprintf("to%d", i)]
			// to = match[fmt.Sprintf("to%d", i + 1)] // Next account
		}

		transaction := &ledger.Transaction{
			FromAccount: &ledger.Account{Name: from},
			ToAccount:   &ledger.Account{Name: to},
			Amount:      amount,
			Payee:       match["payee"],
			Date:        fmt.Sprintf("%v/%v/%v", time.Now().Year(), match["month"], match["day"]),
		}

		if err := j.AddTransaction(transaction); err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil

}


func DeleteTransaction(senderID int, id string) error {
	user := ledger.User{Username: strconv.Itoa(senderID)}
	err := user.DeleteTransaction(id)
	if err != nil {
		return err
	}

	return nil
}
