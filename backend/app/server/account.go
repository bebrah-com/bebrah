package server

import (
	"github.com/gin-gonic/gin"
)

func setupAccount(r *gin.Engine) {
	account := r.Group("/account")
	account.POST("/sign-up", func(c *gin.Context) {
		// TODO: implement
	})
	account.POST("/sign-in", func(c *gin.Context) {
		// TODO: implement
	})
	account.POST("/sign-out", func(c *gin.Context) {
		// TODO: implement
	})
}
