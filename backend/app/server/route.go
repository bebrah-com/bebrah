package server

import (
	"bebrah/app/middleware"
	"net/http"

	_ "bebrah/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.Bearer  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func SetupRouter() *gin.Engine {
	server := gin.Default()

	server.GET("/ping", middleware.JWTAuthMiddleware(), PingExample)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")

	server.Use(cors.New(corsConfig))

	router := server.Group("/api/v1")
	setupAuth(router)
	setupWork(router)
	setupProfile(router)
	setupNft(router)
	setupLike(router)
	setupComment(router)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return server
}

// @BasePath /api/v1

// @Description do ping
// @Tags example
// @Success 200 {string} pong
// @Router /ping [get]
func PingExample(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
