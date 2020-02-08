package db

import (
	"fmt"
	"sort"
	"strings"

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
			// fmt.Printf("key=%s, value=%s\n", k, v)
			aliases[string(k)] = string(v)
		}

		return nil
	})

	keys := make([]string, 0, len(aliases))
	for k := range aliases {
		keys = append(keys, k)
	}

	// Sort by case insensitive alias names
	sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })
	sorted := make(map[string]string)
	for _, k := range keys {
		sorted[k] = aliases[k]
	}

	return sorted
}

func DeleteAlias(userID int, name string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := getUserAliasBucket(tx, userID)
		return b.Delete([]byte(name))
	})
}
