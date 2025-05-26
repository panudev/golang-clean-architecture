// main.go

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	godotenv "github.com/joho/godotenv"

	_ "github.com/panudev/golang-clean-architecture/api/docs" // swagger docs
	bootstrap "github.com/panudev/golang-clean-architecture/internal/bootstrap"
	"github.com/panudev/golang-clean-architecture/internal/infrastructure/mysql"
	"github.com/panudev/golang-clean-architecture/internal/interface/route"
	"github.com/panudev/golang-clean-architecture/pkg/config"
	"github.com/panudev/golang-clean-architecture/pkg/logger"
)

// @title           Go Clean Architecture API
// @version         1.0
// @description     A RESTful API built with Go using Clean Architecture principles.
// @host            localhost:8080
// @BasePath        /api/v1
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
	dbConn := mysql.NewMySQLConnection(appConfig, appLog)
	userHandler := bootstrap.InitializeUserHandler(dbConn, appConfig, appLog)

	// Setup Gin router
	router := gin.Default()
	// Register routes
	route.RegisterRoutes(router, userHandler)

	appLog.Info("🚀 Server is running on port", "port", appConfig.Port)
	router.Run(":" + appConfig.Port)
}
