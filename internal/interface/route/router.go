package route

import (
	"github.com/gin-gonic/gin"
	handler "github.com/panudev/golang-clean-architecture/internal/interface/handler/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// User routes
			users := v1.Group("/users")
			{
				users.GET("", userHandler.ListUsers)         // GET /api/v1/users
				users.POST("", userHandler.CreateUser)       // POST /api/v1/users
				users.GET("/:id", userHandler.GetUser)       // GET /api/v1/users/:id
				users.PUT("/:id", userHandler.UpdateUser)    // PUT /api/v1/users/:id
				users.DELETE("/:id", userHandler.DeleteUser) // DELETE /api/v1/users/:id
			}
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
