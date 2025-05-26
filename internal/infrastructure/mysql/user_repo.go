package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/panudev/golang-clean-architecture/internal/domain"
	"github.com/panudev/golang-clean-architecture/pkg/config"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewMySQLConnection(cfg *config.Config) *gorm.DB {
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ")/" + cfg.DBName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ failed to connect database: %v", err)
	}
	return db
}
