package server

import "github.com/gin-gonic/gin"

func setupNft(r *gin.RouterGroup) {
	nft := r.Group("/nft")
	// get/list nft
	nft.GET("/", func(c *gin.Context) {
	})
	// add nft
	nft.POST("/", func(c *gin.Context) {
	})
}
