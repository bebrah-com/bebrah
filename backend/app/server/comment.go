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

// @Description comment a work
// @Tags comment
// @Param request body model.CommentWorkReq true "comment work request"
// @Success 200 {string} success
// @Router /comments [post]
func commentWork(c *gin.Context) {
	var req model.CommentWorkReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	userId, err := middleware.GetUserIdFromGin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request in token"})
		return
	}

	var work model.Work
	db.Db().Where("id = ?", req.WorkId).First(&work)
	if work.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "work not found"})
		return
	}

	var comment model.Comment
	db.Db().Where("user_id = ? and work_id = ? and deleted_at IS NULL", userId, req.WorkId).First(&comment)
	if comment.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you have commented this work"})
		return
	}

	db.Db().Create(&model.Comment{
		UserID:  userId,
		WorkID:  req.WorkId,
		Content: req.Content,
	})

	c.String(http.StatusOK, "success")
}

// @BasePath /api/v1

// @Description list comments by work id
// @Tags comment
// @Param work_id path uint64 true "work id"
// @Success 200 {string} success
// @Router /comments/:work_id [get]
func listCommentsByWorkId(c *gin.Context) {
	workIdStr := c.Param("work_id")
	workId, err := strconv.ParseUint(workIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var comments []model.Comment
	db.Db().Where("work_id = ? and deleted_at IS NULL", workId).Order("created_at DESC").Preload("User").Find(&comments)

	c.JSON(http.StatusOK, model.ListCommentsByWorkIdResp{
		Comments: comments,
	})
}

// @BasePath /api/v1

// @Description delete a comment
// @Tags comment
// @Param commentId path uint64 true "comment id"
// @Success 200 {string} success
// @Router /comments/:comment_id [delete]
func deleteComment(c *gin.Context) {
	commentId := c.Param("comment_id")

	userId, err := middleware.GetUserIdFromGin(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var comment model.Comment
	db.Db().Where("id = ? and deleted_at IS NULL", commentId).First(&comment)
	if comment.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "comment not found"})
		return
	}

	if comment.UserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you can only delete your own comment"})
		return
	}

	now := time.Now()
	db.Db().Model(&comment).Update("deleted_at", &now)

	c.String(http.StatusOK, "success")
}

func setupComment(r *gin.RouterGroup) {
	comment := r.Group("/comments", middleware.JWTAuthMiddleware())
	comment.POST("", commentWork)
	comment.GET("/:work_id", listCommentsByWorkId)
	comment.DELETE("/:comment_id", deleteComment)
}
