package server

import (
	"bebrah/app/middleware"

	"github.com/gin-gonic/gin"
)

func setupWork(r *gin.RouterGroup) {
	work := r.Group("/work", middleware.JWTAuthMiddleware())
	// get work
	work.GET("/", func(c *gin.Context) {
		// TODO: implement
	})

	// add work
	work.POST("/", func(c *gin.Context) {
		// TODO: implement
	})
}
