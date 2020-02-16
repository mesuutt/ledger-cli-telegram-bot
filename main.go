package main

import (
	"github.com/mesuutt/teledger/bot"
	"github.com/mesuutt/teledger/config"
	"github.com/mesuutt/teledger/db"
)

func main() {
	config.Init()
	config.InitLogging()
	db.Init()
	defer db.Close()

	bot.Start(config.Env.Telegram.Token)
}
