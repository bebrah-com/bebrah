package server

import (
	"bebrah/app/middleware"

	"github.com/gin-gonic/gin"
)

func setupProfile(r *gin.RouterGroup) {
	profile := r.Group("/profile", middleware.JWTAuthMiddleware())
	// get profile
	profile.GET("/", func(c *gin.Context) {
		// TODO: implement
	})
	// edit profile
	profile.POST("/", func(c *gin.Context) {
		// TODO: implement
	})
}
