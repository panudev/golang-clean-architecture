package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/panudev/golang-clean-architecture/internal/domain"
	"github.com/panudev/golang-clean-architecture/pkg/config"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewMySQLConnection(cfg *config.Config) *gorm.DB {
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ")/" + cfg.DBName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
