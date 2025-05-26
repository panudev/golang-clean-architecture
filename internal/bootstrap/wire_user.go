//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/panudev/golang-clean-architecture/internal/infrastructure/mysql"
	handler "github.com/panudev/golang-clean-architecture/internal/interface/handler/user"
	usecase "github.com/panudev/golang-clean-architecture/internal/usecase/user"
	"github.com/panudev/golang-clean-architecture/pkg/config"
	"github.com/panudev/golang-clean-architecture/pkg/logger"
	"gorm.io/gorm"
)

func InitializeUserHandler(db *gorm.DB, cfg *config.Config, logger logger.Logger) *handler.UserHandler {
	wire.Build(
		mysql.NewUserRepo,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
	)
	return nil
}

// ProvideLogger provides a logger instance
func ProvideLogger() logger.Logger {
	return logger.NewLogger()
}

// ProvideDB provides a database connection
func ProvideDB(cfg *config.Config) *gorm.DB {
	return mysql.NewMySQLConnection(cfg, ProvideLogger())
}

// InitializeApp initializes the entire application
func InitializeApp(cfg *config.Config) (*handler.UserHandler, error) {
	wire.Build(
		ProvideLogger,
		ProvideDB,
		InitializeUserHandler,
	)
	return nil, nil
}
