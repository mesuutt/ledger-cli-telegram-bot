package main

import (
	"flag"

	"github.com/mesuutt/teledger/bot"
	"github.com/mesuutt/teledger/config"
	"github.com/mesuutt/teledger/db"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "config.toml", "Configuration file path")
	flag.Parse()

	config.Init(configFilePath)
	config.InitLogging()
	db.Init()
	defer db.Close()

	bot.Start(config.Env.Telegram.Token)
}
