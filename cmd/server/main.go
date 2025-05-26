// main.go

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	godotenv "github.com/joho/godotenv"

	bootstrap "github.com/panudev/golang-clean-architecture/internal/bootstrap"
	"github.com/panudev/golang-clean-architecture/internal/infrastructure/mysql"
	"github.com/panudev/golang-clean-architecture/internal/interface/route"
	"github.com/panudev/golang-clean-architecture/pkg/config"
	"github.com/panudev/golang-clean-architecture/pkg/logger"
)

func main() {
	// Load ENV
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found. Using system env instead.")
	}

	// Load config
	appConfig := config.LoadConfig()

	// Init logger
	appLog := logger.NewLogger()

	// Setup DB connection
	dbConn := mysql.NewMySQLConnection(appConfig)
	userHandler := bootstrap.InitializeUserHandler(dbConn, appConfig)

	// Setup Gin router
	router := gin.Default()
	// Register routes
	route.RegisterRoutes(router, userHandler)

	appLog.Infof("🚀 Server is running on port %s", appConfig.Port)
	router.Run(":" + appConfig.Port)
}
