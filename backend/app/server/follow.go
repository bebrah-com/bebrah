package server

import (
	"bebrah/app/db"
	"bebrah/app/middleware"
	"bebrah/app/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Description follow someone
// @Tags follows
// @Param followed_id path uint64 true "followed user id"
// @Success 200 {string} success
// @Router /follows/:followed_id [post]
func follow(c *gin.Context) {
	followedIdStr := c.Param("followed_id")
	followedId, err := strconv.ParseUint(followedIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	userId, err := middleware.GetUserIdFromGin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request in token"})
		return
	}

	var follow model.Follow
	db.Db().Where("followed_id = ? and follower_id = ? and deleted_at IS NULL", followedId, userId).First(&follow)

	db.Db().Create(&model.Follow{
		FollowedID: followedId,
		FollowerID: userId,
	})

	c.String(http.StatusOK, "success")
}

// @BasePath /api/v1

// @Description unfollow someone
// @Tags follows
// @Param followed_id path uint64 true "followed user id"
// @Success 200 {string} success
// @Router /follows/:followed_id [delete]
func unfollow(c *gin.Context) {
	followedIdStr := c.Param("followed_id")
	followedId, err := strconv.ParseUint(followedIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	userId, err := middleware.GetUserIdFromGin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request in token"})
		return
	}

	var follow model.Follow
	db.Db().Where("followed_id = ? and follower_id = ? and deleted_at IS NULL", followedId, userId).First(&follow)
	if follow.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you have not followed this user"})
		return
	}

	now := time.Now()
	follow.DeletedAt = &now
	db.Db().Save(&follow)

	c.String(http.StatusOK, "success")
}

// @BasePath /api/v1

// @Description list followed users by follower id
// @Tags follows
// @Param follower_id path uint64 true "follower id"
// @Success 200 {string} success
// @Router /follows/:follower_id [get]
func listFollows(c *gin.Context) {
	followerIdStr := c.Param("follower_id")
	followerId, err := strconv.ParseUint(followerIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var follows []model.Follow
	db.Db().Where("follower_id = ? and deleted_at IS NULL", followerId).Order("created_at DESC").Preload("Followed").Preload("Follower").Find(&follows)

	for _, follow := range follows {
		if follow.Followed != nil {
			follow.Followed.Password = ""
			follow.Followed.Token = ""
		}

		if follow.Follower != nil {
			follow.Follower.Password = ""
			follow.Follower.Token = ""
		}
	}

	c.JSON(http.StatusOK, model.ListFollowsResp{
		Follows: follows,
		Count:   int64(len(follows)),
	})
}

func setupFollow(r *gin.RouterGroup) {
	like := r.Group("/follows", middleware.JWTAuthMiddleware())
	like.GET("/:follower_id", listFollows)
	like.POST("/:followed_id", follow)
	like.DELETE("/:followed_id", unfollow)
}
