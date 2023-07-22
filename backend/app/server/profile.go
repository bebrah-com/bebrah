package server

import (
	"bebrah/app/middleware"

	"github.com/gin-gonic/gin"
)

func setupProfile(r *gin.RouterGroup) {
	profile := r.Group("/profile", middleware.JWTAuthMiddleware())
	// get profile
	profile.GET("/me", func(c *gin.Context) {
		// TODO: implement
	})
	// edit profile
	profile.POST("/me", func(c *gin.Context) {
		// TODO: implement
	})
	// get some user's profile, contain:
	// - user's info
	// - user's nfts
	// - user's works
	// - user's followers
	// - user's followings
	// - user's likes
	profile.GET("/{user_id}", func(c *gin.Context) {
		// TODO: implement
	})
}
