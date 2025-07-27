package routers

import (
	"github.com/gin-gonic/gin"

	"zip-app/internal/api/handlers"
)

func subscriptionRouter(router *gin.RouterGroup, handler handlers.ZipHandler) {
	subsRouter := router.Group("/zip_task")
	{
		subsRouter.POST("/", handler.CreateTask)
		subsRouter.GET("/:id/status", handler.CheckStatus)
		subsRouter.PUT("/:id", handler.UpdateTask)
		subsRouter.GET("/download/:file", handler.DownloadArchive)
		// subsRouter.POST("/", handler.CreateSubscription)
		// subsRouter.PUT("/:id", handler.FullUpdateSubscription)
		// subsRouter.PATCH("/:id", handler.PatchUpdateSubscription)
		// subsRouter.DELETE("/:id", handler.DeleteSubscription)
		// subsRouter.GET("/sub_sum", handler.GetSubscriptionSumInfo)
	}
}
