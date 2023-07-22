package server

import (
	"bebrah/app/db"
	"bebrah/app/middleware"
	"bebrah/app/model"
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

func setupWork(r *gin.RouterGroup) {
	work := r.Group("/work", middleware.JWTAuthMiddleware())
	// get work
	work.GET("/", func(c *gin.Context) {
		// TODO: implement
	})

	// add work
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
			Data:   string(fileData),
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
