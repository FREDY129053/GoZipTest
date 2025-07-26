package routers

import (
	"zip-app/internal/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func SetupRouter(handler handlers.ZipHandler) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		subscriptionRouter(api, handler)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, gin.H{"message": "service good"})
	})

	return router
}
