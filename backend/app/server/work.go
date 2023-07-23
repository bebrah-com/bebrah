package server

import (
	"bebrah/app/db"
	"bebrah/app/middleware"
	"bebrah/app/model"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type AddWorkReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListWorksReq struct {
	UserId   uint64 `json:"user_id"`
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"page_size"`
	Count    int64  `json:"count"`
}

func countOffset(page, pageSize uint64) int {
	var offset int

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	offset = int((page - 1) * pageSize)

	return offset
}

func setupWork(r *gin.RouterGroup) {
	work := r.Group("/work", middleware.JWTAuthMiddleware())
	// list work
	work.GET("/", func(c *gin.Context) {
		var req ListWorksReq
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "invalid request"})
			return
		}

		var works []model.Work
		var count int64

		if req.UserId != 0 {
			if err := db.Db().Model(&model.Work{}).Where("user_id = ?", req.UserId).Count(&count).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "count works by user id failed"})
				return
			}

			// list works by user id
			if err := db.Db().Where("user_id = ?", req.UserId).Order("created_at DESC").Offset(countOffset(req.Page, req.PageSize)).Find(&works).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			if err := db.Db().Model(&model.Work{}).Count(&count).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "count works failed"})
				return
			}

			// list all works
			if err := db.Db().Order("created_at DESC").Offset(countOffset(req.Page, req.PageSize)).Find(&works).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"works": works, "count": count})
	})

	// get work
	work.GET("/:id", func(c *gin.Context) {
		// TODO: implement
	})

	// upload work
	work.POST("/", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "file not found"})
			return
		}

		extString := path.Ext(file.Filename)

		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extString]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file type not allowed"})
			return
		}

		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "file open failed"})
			return
		}
		defer f.Close()

		fileData, err := ioutil.ReadAll(f)
		if err != nil && err != io.EOF {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "file read failed"})
			return
		}

		userId, err := middleware.GetUserIdFromGin(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "invalid request"})
			return
		}

		// var req AddWorkReq
		// if err := c.ShouldBindJSON(&req); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "invalid request"})
		// 	return
		// }

		// if len(req.Title) == 0 {
		// 	// generate a title by create time
		// 	fileNameInt := time.Now().Unix()
		// 	fileNameStr := strconv.FormatInt(fileNameInt, 10)
		// 	req.Title = fileNameStr + extString
		// }

		image := model.Work{
			UserID: userId,
			Data:   encodeToBase64(fileData),
			// WorkName: req.Title,
			// WorkDesc: req.Description,
		}
		if err := db.Db().Create(&image).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
	})
}

func encodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
