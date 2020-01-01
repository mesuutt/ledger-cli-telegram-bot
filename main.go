package main

import (
	"flag"

	"github.com/labstack/echo"

	"github.com/mesuutt/teledger/api"
	"github.com/mesuutt/teledger/bot"
	"github.com/mesuutt/teledger/config"
	"github.com/mesuutt/teledger/db"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "config.toml", "Configuration file path")
	flag.Parse()

	config.Init(configFilePath)

	// router := gin.New()
	e := echo.New()
	api.InitRoutes(e)
	db.Init()
	defer db.Close()

	go bot.Start(config.Env.Telegram.Token)

	e.Logger.Fatal(e.Start(":8080"))

}
