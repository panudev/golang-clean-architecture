//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/panudev/golang-clean-architecture/internal/infrastructure/mysql"
	"github.com/panudev/golang-clean-architecture/internal/interface/handler"
	"github.com/panudev/golang-clean-architecture/internal/usecase"
	"github.com/panudev/golang-clean-architecture/pkg/config"
	"gorm.io/gorm"
)

func InitializeUserHandler(db *gorm.DB, cfg *config.Config) *handler.UserHandler {
	wire.Build(
		mysql.NewUserRepo,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
	)
	return nil
}
