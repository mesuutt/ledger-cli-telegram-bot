package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mesuutt/ledger/ledger"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := ledger.User{Username: c.Param("username")}
		c.Set("user", user)
		c.Next()
	}
}
