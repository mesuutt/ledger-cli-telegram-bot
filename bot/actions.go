package bot

import (
	"github.com/mesuutt/teledger/db"
)

func SetAlias(senderID int, name, accountName string) error {
	err := db.AddAlias(senderID, name, accountName)
	if err != nil {
		return err
	}

	return nil
}
