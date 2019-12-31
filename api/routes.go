package api

import (

	"github.com/labstack/echo"

)

func InitRoutes(e *echo.Echo) {
	// https://github.com/gin-gonic/contrib/blob/master/ginrus/example/example.go
	// e.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	// e.Use(gin.Recovery())
	e.Use(AuthMiddleware)
	e.GET("/:username/accounts", AccountListHandler)
	e.POST("/:username/transactions", AddTransactionHandler)
}
