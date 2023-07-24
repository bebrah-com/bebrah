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
	"strconv"
	"time"

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
// @Success 200 {object} model.ListWorksResp
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
		if err := db.Db().Where("user_id = ?", req.UserId).Order("created_at DESC").Offset(countOffset(req.Page, req.PageSize)).Preload("User").Find(&works).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := db.Db().Model(&model.Work{}).Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "count works failed"})
			return
		}

		// list all works
		if err := db.Db().Order("created_at DESC").Offset(countOffset(req.Page, req.PageSize)).Preload("User").Find(&works).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, model.ListWorksResp{
		Works: works,
		Count: count,
	})
}

// @BasePath /api/v1

// @Description get a work by work_id
// @Tags works
// @Param id path uint64 true "work id"
// @Success 200 {object} model.GetWorkResp
// @Router /works/:id [get]
func getWorkById(c *gin.Context) {
	id := c.Param("id")
	workId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "invalid request"})
		return
	}

	var work model.Work
	if err := db.Db().Where("id = ?", workId).Preload("User").First(&work).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	work.Viewed++
	db.Db().Model(&work).Update("viewed", work.Viewed)

	if work.User == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	work.User.Password = ""
	work.User.Token = ""
	c.JSON(http.StatusOK, model.GetWorkResp{
		Work: work,
	})
}

// @BasePath /api/v1

// @Description upload a work
// @Tags auth
// @Param file formData file true "work file"
// @Param name formData string true "work name"
// @Param description formData string false "work description"
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

	name := c.PostForm("name")
	if len(name) == 0 {
		// generate a title by create time
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		name = fileNameStr + extString
	}

	image := model.Work{
		UserID:   userId,
		Data:     encodeToBase64(fileData),
		WorkName: name,
		WorkDesc: c.PostForm("descrition"),
	}
	if err := db.Db().Create(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "Work uploaded successfully")
}

func setupWork(r *gin.RouterGroup) {
	work := r.Group("/works", middleware.JWTAuthMiddleware())
	work.GET("/", listWorks)
	work.GET("/:id", getWorkById)
	work.POST("/", uploadWork)
}

func encodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
