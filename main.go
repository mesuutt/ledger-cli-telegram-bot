package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mesuutt/ledger/api"
)

/*var userHandlers = map[string]http.HandlerFunc{
	"/{username}/accounts":     api.AccountListHandler,
	"/{username}/transactions": api.AddTransactionHandler,
}
*/
func main() {
	// router := mux.NewRouter().StrictSlash(true)
	router := gin.Default()
	router.Use(api.AuthMiddleware())
	router.GET("/:username/accounts", api.AccountListHandler)
	router.POST("/:username/transactions", api.AddTransactionHandler)

	router.Run()
}
