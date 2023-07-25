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

// @Description like a work
// @Tags like
// @Param work_id path uint64 true "work id"
// @Success 200 {string} success
// @Router /likes/:work_id [post]
func likeWork(c *gin.Context) {
	workIdStr := c.Param("work_id")
	workId, err := strconv.ParseUint(workIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	userId, err := middleware.GetUserIdFromGin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var work model.Work
	db.Db().Where("id = ?", workId).First(&work)
	if work.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "work not found"})
		return
	}

	var like model.Like
	db.Db().Where("user_id = ? and work_id = ? and deleted_at IS NULL", userId, workId).First(&like)
	if like.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you have liked this work"})
		return
	}

	db.Db().Create(&model.Like{
		UserID: userId,
		WorkID: workId,
	})

	work.Liked++
	db.Db().Model(&work).Update("liked", work.Liked)

	c.String(http.StatusOK, "success")
}

// @BasePath /api/v1

// @Description unlike a work
// @Tags like
// @Param work_id path uint64 true "work id"
// @Success 200 {string} success
// @Router /likes/:work_id [delete]
func unlikeWork(c *gin.Context) {
	workId := c.Param("work_id")

	userId, err := middleware.GetUserIdFromGin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var work model.Work
	db.Db().Where("id = ?", workId).First(&work)
	if work.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "work not found"})
		return
	}

	var like model.Like
	db.Db().Where("user_id = ? and work_id = ? and deleted_at IS NULL", userId, workId).First(&like)
	if like.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you have not liked this work"})
		return
	}

	now := time.Now()
	like.DeletedAt = &now
	db.Db().Save(&like)

	work.Liked--
	db.Db().Model(&work).Update("liked", work.Liked)

	c.String(http.StatusOK, "success")
}

// @BasePath /api/v1

// @Description list works by user like
// @Tags like
// @Param user_id path uint64 true "user id"
// @Success 200 {string} success
// @Router /likes/:user_id [get]
func listLikes(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var likes []model.Like
	db.Db().Where("user_id = ? and deleted_at IS NULL", userId).Order("created_at DESC").Find(&likes)

	var works []model.Work
	for _, like := range likes {
		var work model.Work
		db.Db().Where("id = ?", like.WorkID).Preload("User").First(&work)
		if work.ID != 0 {
			works = append(works, work)
		}
	}

	// TODO: pagination
	c.JSON(http.StatusOK, model.ListLikedWorksResp{
		Works: works,
		Count: int64(len(works)),
	})
}

func setupLike(r *gin.RouterGroup) {
	like := r.Group("/likes", middleware.JWTAuthMiddleware())
	like.GET("/:user_id", listLikes)
	like.POST("/:work_id", likeWork)
	like.DELETE("/:work_id", unlikeWork)
}
