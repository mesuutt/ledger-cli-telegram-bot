package db

import (
	"log"

	"github.com/boltdb/bolt"

	"github.com/mesuutt/teledger/config"
)

var (
	DB *bolt.DB
)


func Init() {
	d, err := bolt.Open(config.Env.DBFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = d
}

func Close() {
	DB.Close()
}