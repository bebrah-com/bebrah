package server

import "github.com/gin-gonic/gin"

func setupWork(r *gin.Engine) {
	work := r.Group("/work")
	// get work
	work.GET("/", func(c *gin.Context) {
		// TODO: implement
	})

	// add work
	work.POST("/", func(c *gin.Context) {
		// TODO: implement
	})
}
