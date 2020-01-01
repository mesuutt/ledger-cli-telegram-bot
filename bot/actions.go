package bot

import (
	"errors"
	"fmt"
	"strconv"

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

