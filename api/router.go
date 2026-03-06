package api

import (
	"flower-shop/api/handler"
	"flower-shop/pkg/logger"
	"flower-shop/service"
	"flower-shop/storage"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           STAR FLOW PRO BACKEND
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func New(store storage.IStorage, service service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(store, service, log)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is working!")
	})

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// middleware
	r.Use(authMiddleware)

	// users
	r.POST("/api/v1/user", h.CreateUser)
	r.GET("/api/v1/users", h.GetAllUser)
	r.DELETE("/api/v1/user/:id", h.DeleteUser)

	return r
}

func authMiddleware(c *gin.Context) {
	if c.GetHeader("Authorization") != "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}
