package db

import (
	"fmt"

	// "github.com/boltdb/bolt"

	"github.com/boltdb/bolt"
)

func CreateUser(userID int) error {
	return DB.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucket([]byte(fmt.Sprintf("u_%d", userID)))
		if err != nil {
			return err
		}
		_, err = root.CreateBucket([]byte("aliases"))
		return err
	})
}

func getUserAliasBucket(tx *bolt.Tx, userID int) *bolt.Bucket {
	userB := tx.Bucket([]byte(fmt.Sprintf("u_%d", userID)))
	return userB.Bucket([]byte("aliases"))
}

func GetAccountByAlias(userID int, name string) string {
	var account string

	DB.View(func(tx *bolt.Tx) error {
		 account = string(getUserAliasBucket(tx, userID).Get([]byte(name)))
		 return nil
	})

	return account
}

func AddAlias(userID int, alias, account string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := getUserAliasBucket(tx, userID)
		return b.Put([]byte(alias), []byte(account))
	})
}

func GetUserAliases(userID int) map[string]string {
	aliases := make(map[string]string)

	DB.View(func(tx *bolt.Tx) error {
		b := getUserAliasBucket(tx, userID)
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
		b := getUserAliasBucket(tx, userID)
		return b.Delete([]byte(name))
	})
}
