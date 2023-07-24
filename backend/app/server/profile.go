package server

import (
	"bebrah/app/db"
	"bebrah/app/middleware"
	"bebrah/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupProfile(r *gin.RouterGroup) {
	profile := r.Group("/profile", middleware.JWTAuthMiddleware())
	// get profile
	profile.GET("/me", func(c *gin.Context) {
		// TODO: implement
		userId, err := middleware.GetUserIdFromGin(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		var user model.User
		db.Db().Where("email = ?", userId).First(&user)
		c.JSON(http.StatusOK, &user)
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
