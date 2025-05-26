package route

import (
	"github.com/gin-gonic/gin"
	"github.com/panudev/golang-clean-architecture/internal/interface/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine, handler *handler.UserHandler) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/users/:id", handler.GetUser)
		}
	}
}
