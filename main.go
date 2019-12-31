package main

import (
	"flag"

	"github.com/labstack/echo"

	"github.com/mesuutt/teledger/api"
)

func main() {
	var configFilePath string

	flag.StringVar(&configFilePath, "config", "config.toml", "Configuration file path")
	flag.Parse()

	InitConfig(configFilePath)

	// router := gin.New()
	e := echo.New()
	api.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
