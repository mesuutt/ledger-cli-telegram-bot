package db

import (
	"fmt"

	// "github.com/boltdb/bolt"

	"github.com/boltdb/bolt"
)

func AddAlias(userID int, alias, account string) error {
	return  DB.Update(func(tx *bolt.Tx) error {

		// FIXME: create user bucket on conversation start
		userB, err := tx.CreateBucketIfNotExists([]byte(fmt.Sprintf("u_%d", userID)))
		if err != nil {
		    return err
		}

		b, err := userB.CreateBucketIfNotExists([]byte("aliases"))
		return b.Put([]byte(alias), []byte(account))
	})
}
