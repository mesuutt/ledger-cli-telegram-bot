package db

import (
	"fmt"
	"sort"
	"strings"

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

func getUserAliasBucket(tx *bolt.Tx, userID int) (*bolt.Bucket, error) {
	userB := tx.Bucket([]byte(fmt.Sprintf("u_%d", userID)))
	if userB == nil {
		return nil, &ErrBudgetNotFound{}
	}
	return userB.Bucket([]byte("aliases")), nil
}

func GetAccountByAlias(userID int, name string) string {
	var account string

	DB.View(func(tx *bolt.Tx) error {
		b, err := getUserAliasBucket(tx, userID)
		if err != nil {
		    return err
		}

		account = string(b.Get([]byte(name)))
		return nil
	})

	return account
}

func AddAlias(userID int, alias, account string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b, err := getUserAliasBucket(tx, userID)
		if err != nil {
		    return err
		}

		return b.Put([]byte(alias), []byte(account))
	})
}

func GetUserAliases(userID int) ([][]string, error) {
	aliases := make(map[string]string)

	err := DB.View(func(tx *bolt.Tx) error {
		b, err := getUserAliasBucket(tx, userID)
		if err != nil {
		    return err
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			// fmt.Printf("key=%s, value=%s\n", k, v)
			aliases[string(k)] = string(v)
		}

		return nil
	})

	if err != nil {
	    return nil, err
	}

	keys := make([]string, 0, len(aliases))
	for k := range aliases {
		keys = append(keys, k)
	}

	// Sort by case insensitive alias names
	// Maps is not ordered, so we are creating nested string array
	sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })
	sorted := [][]string{}
	for _, k := range keys {
		sorted = append(sorted, []string{k, aliases[k]})
	}

	return sorted, nil
}

func DeleteAlias(userID int, name string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b, err := getUserAliasBucket(tx, userID)
		if err != nil {
		    return err
		}
		return b.Delete([]byte(name))
	})
}
