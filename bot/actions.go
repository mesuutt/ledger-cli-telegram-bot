package bot

import (
	"strconv"

	"github.com/mesuutt/teledger/db"
	"github.com/mesuutt/teledger/ledger"
)

func SetAlias(senderID int, name, accountName string) error {
	err := db.AddAlias(senderID, name, accountName)
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

