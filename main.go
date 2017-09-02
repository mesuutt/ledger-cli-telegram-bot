package main

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/mesuutt/ledger/api"
)

func main() {

	var conf config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
		return
	}

	router := gin.New()
	// https://github.com/gin-gonic/contrib/blob/master/ginrus/example/example.go
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	router.Use(gin.Recovery())
	router.Use(api.AuthMiddleware())
	router.GET("/:username/accounts", api.AccountListHandler)
	router.POST("/:username/transactions", api.AddTransactionHandler)

	router.Run()

}
