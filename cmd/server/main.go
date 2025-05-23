// main.go

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	godotenv "github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/panudev/golang-clean-architecture/internal/infrastructure/mysql"
	"github.com/panudev/golang-clean-architecture/internal/interface/handler"
	"github.com/panudev/golang-clean-architecture/internal/usecase"
	"github.com/panudev/golang-clean-architecture/pkg/config"
	"github.com/panudev/golang-clean-architecture/pkg/logger"
)

func main() {
	// Load ENV
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found. Using system env instead.")
	}

	// Load config
	cfg := config.LoadConfig()

	// Init logger
	appLog := logger.NewLogger()

	// Setup DB connection
	db := mysql.NewMySQLConnection(cfg)
	repo := mysql.NewUserRepo(db)
	uc := usecase.NewUserUseCase(repo)
	h := handler.NewUserHandler(uc)

	// Setup Gin router
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// API group
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/users/:id", h.GetUser)
		}
	}

	appLog.Infof("🚀 Server is running on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
