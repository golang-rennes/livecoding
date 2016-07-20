package server

import "github.com/gin-gonic/gin"

var accounts = gin.Accounts{
	"foo":   "bar",
	"admin": "admin",
}

func basicAuth() gin.HandlerFunc {
	return gin.BasicAuth(accounts)
}
