package api

import (
	"github.com/labstack/echo"

	"github.com/mesuutt/teledger/ledger"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := ledger.User{Username: c.Param("username")}
		c.Set("user", user)
		return next(c)
	}
}
