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

func GetAccounts(senderID int64) []string {
	user := ledger.User{Username: strconv.Itoa(int(senderID))}
	return user.GetAccounts()
}

func SetAlias(senderID int64, name, accountName string) error {
	if db.GetAccountByAlias(senderID, name) != "" {
		return errors.New(fmt.Sprintf("alias %s already exist.", name))
	}

	err := db.AddAlias(senderID, name, accountName)
	if err != nil {
		return err
	}
	user := ledger.User{Username: strconv.Itoa(int(senderID))}
	err = user.AddAlias(name, accountName)
	if err != nil {
		return err
	}

	return nil
}
func DeleteAlias(senderID int64, name string) error {
	user := ledger.User{Username: strconv.Itoa(int(senderID))}
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

func AddTransaction(senderID int64, text string) ([]*ledger.Transaction, error) {
	user := ledger.User{Username: strconv.Itoa(int(senderID))}
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
		to := match[fmt.Sprintf("to%d", i+1)]
		if i > 0 {
			from = match[fmt.Sprintf("to%d", i)]
		}

		fromAcc := db.GetAccountByAlias(senderID, from)
		toAcc := db.GetAccountByAlias(senderID, to)

		if fromAcc != "" {
			from = fromAcc
		}
		if toAcc != "" {
			to = toAcc
		}

		if from == to {
			return transactions, errors.New("you cannot use same account in same transaction")
		}

		transaction := &ledger.Transaction{
			FromAccount: &ledger.Account{Name: from},
			ToAccount:   &ledger.Account{Name: to},
			Amount:      amount,
			Payee:       match["payee"],
			Date:        fmt.Sprintf("%v/%v/%v", time.Now().Year(), match["month"], match["day"]),
		}

		transactions = append(transactions, transaction)
	}

	for i, _ := range transactions {
		if err := j.AddTransaction(transactions[i]); err != nil {
			return nil, err
		}
	}

	return transactions, nil

}
func DeleteTransaction(senderID int64, id string) error {
	user := ledger.User{Username: strconv.Itoa(int(senderID))}
	err := user.DeleteTransaction(id)
	if err != nil {
		return err
	}

	return nil
}

func GetAccountBalance(senderID int64, name string) string {
	user := ledger.User{Username: strconv.Itoa(int(senderID))}
	return user.GetAccountBalance(name)
}
