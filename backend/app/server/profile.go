package server

import "github.com/gin-gonic/gin"

func setupProfile(r *gin.RouterGroup) {
	profile := r.Group("/profile")
	// get profile
	profile.GET("/", func(c *gin.Context) {
		// TODO: implement
	})
	// edit profile
	profile.POST("/", func(c *gin.Context) {
		// TODO: implement
	})
}
