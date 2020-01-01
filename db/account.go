package db

import (
	"fmt"

	// "github.com/boltdb/bolt"

	"github.com/boltdb/bolt"
)

func getOrCreateUserAliasBucket(tx *bolt.Tx, userID int) (*bolt.Bucket, error) {
	userB, err := tx.CreateBucketIfNotExists([]byte(fmt.Sprintf("u_%d", userID)))
	if err != nil {
		return nil, err
	}

	return userB.CreateBucketIfNotExists([]byte("aliases"))
}

func AddAlias(userID int, alias, account string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b, err := getOrCreateUserAliasBucket(tx, userID)
		if err != nil {
			return err
		}

		return b.Put([]byte(alias), []byte(account))
	})
}

func GetUserAliases(userID int) map[string]string {
	aliases := make(map[string]string)

	DB.Update(func(tx *bolt.Tx) error {
		b, err := getOrCreateUserAliasBucket(tx, userID)
		if err != nil {
			return err
		}

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			aliases[string(k)] = string(v)
		}

		return nil
	})

	return aliases
}

func DeleteAlias(userID int, name string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b, err := getOrCreateUserAliasBucket(tx, userID)
		if err != nil {
			return err
		}

		return b.Delete([]byte(name))
	})
}

