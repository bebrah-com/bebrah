package server

import (
	"bebrah/app/db"
	"bebrah/app/middleware"
	"bebrah/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/log"
	"go.uber.org/zap"
)

// @BasePath /api/v1

// @Description get my profile
// @Tags profile
// @Success 200 {object} model.GetProfileResp
// @Router /profiles/me [get]
func getMyProfile(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var user model.User
	db.Db().Where("id = ?", userId).First(&user)
	user.Password = ""
	user.Token = ""
	c.JSON(http.StatusOK, model.GetProfileResp{
		User: user,
	})
}

// @BasePath /api/v1

// @Description edit my profile, note avatar and banner should be a base64 string
// @Tags profile
// @Param request body model.EditMyProfileReq true "Edit profile request"
// @Success 200 {string} success
// @Router /profiles/me [post]
func editMyProfile(c *gin.Context) {
	var req model.EditMyProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	userId, err := middleware.GetUserIdFromGin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if userId != req.User.ID {
		c.JSON(http.StatusBadRequest, gin.H{"message": "your id is not equal to the user id you want to edit"})
		return
	}

	var user model.User
	db.Db().Where("id = ?", userId).First(&user)
	user.UserName = req.User.UserName
	user.Info = req.User.Info
	user.Avatar = req.User.Avatar
	user.Banner = req.User.Banner

	log.Info("get user", zap.Any("user", user))
	db.Db().Model(&user).Updates(&model.User{
		UserName: req.User.UserName,
		Info:     req.User.Info,
		Avatar:   req.User.Avatar,
		Banner:   req.User.Banner,
	})

	c.String(http.StatusOK, "success")
}

// @BasePath /api/v1

// @Description get profile
// @Tags profile
// @Param user_id path uint64 true "user id"
// @Success 200 {object} model.GetProfileResp
// @Router /profiles/:user_id [get]
func getProfile(c *gin.Context) {
	userId := c.Param("user_id")

	var user model.User
	db.Db().Where("id = ?", userId).First(&user)
	user.Password = ""
	user.Token = ""
	c.JSON(http.StatusOK, model.GetProfileResp{
		User: user,
	})
}

func setupProfile(r *gin.RouterGroup) {
	profile := r.Group("/profiles", middleware.JWTAuthMiddleware())
	profile.GET("/me", getMyProfile)
	profile.POST("/me", editMyProfile)
	profile.GET("/:user_id", getProfile)
}
