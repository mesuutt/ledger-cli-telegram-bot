package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/mesuutt/teledger/ledger"
)

type config struct {
	LedgerCLI ledger.Config
}


func InitConfig(path string) {
	var conf config

	var envPath string
	if path == "" {
		if envPath = os.Getenv("TELEDGER_CONFIG_FILE_PATH"); envPath == "" {
			panic("config file path must be gived with flag or TELEDGER_CONFIG_FILE_PATH env var")
		}
	}

	if _, err := toml.DecodeFile(path, &conf); err != nil {
		panic(fmt.Errorf("config file read error : %e", err))
	}

	if conf.LedgerCLI.Journal.Dir == "" {
		panic("ledger journal dir conf cannot be empty")
	}
}