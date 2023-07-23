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

// @BasePath /api/v1

// @Description list all works (by user id or not)
// @Tags works
// @Param request body model.ListWorksReq true "list works request"
// @Success 200 {object} model.ListWokrsResp
// @Router /works [get]
func listWorks(c *gin.Context) {
	var req model.ListWorksReq
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

	c.JSON(http.StatusOK, model.ListWokrsResp{
		Works: works,
		Count: count,
	})
}

// @BasePath /api/v1

// @Description upload a work
// @Tags auth
// @Param file formData file true "work file"
// @Success 200 {string} success
// @Router /works [post]
func uploadWork(c *gin.Context) {
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

	c.String(http.StatusOK, "Work uploaded successfully")
}

func setupWork(r *gin.RouterGroup) {
	work := r.Group("/works", middleware.JWTAuthMiddleware())
	// list work
	work.GET("/", listWorks)

	// get work
	work.GET("/:id", func(c *gin.Context) {
		// TODO: implement
	})

	// upload work
	work.POST("/", uploadWork)
}

func encodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
