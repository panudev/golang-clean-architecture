package mysql

import (
	"errors"

	"github.com/panudev/golang-clean-architecture/internal/domain/entity"
	"github.com/panudev/golang-clean-architecture/internal/domain/repository"
	"github.com/panudev/golang-clean-architecture/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/panudev/golang-clean-architecture/pkg/config"
)

type userRepository struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewUserRepo(db *gorm.DB, logger logger.Logger) repository.UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

// GetByID retrieves a user by ID
func (repo *userRepository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := repo.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrUserNotFound
		}
		repo.logger.Error("failed to get user by ID", "error", err, "id", id)
		return nil, err
	}
	return &user, nil
}

// Create creates a new user
func (repo *userRepository) Create(user *entity.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		repo.logger.Error("failed to create user", "error", err)
		return err
	}
	return nil
}

// Update updates an existing user
func (repo *userRepository) Update(user *entity.User) error {
	if err := repo.db.Save(user).Error; err != nil {
		repo.logger.Error("failed to update user", "error", err, "id", user.ID)
		return err
	}
	return nil
}

// Delete deletes a user by ID
func (repo *userRepository) Delete(id uint) error {
	if err := repo.db.Delete(&entity.User{}, id).Error; err != nil {
		repo.logger.Error("failed to delete user", "error", err, "id", id)
		return err
	}
	return nil
}

// List retrieves all users with pagination
func (repo *userRepository) List(offset, limit int) ([]*entity.User, error) {
	var users []*entity.User
	if err := repo.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		repo.logger.Error("failed to list users", "error", err)
		return nil, err
	}
	return users, nil
}

// GetByEmail retrieves a user by email
func (repo *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrUserNotFound
		}
		repo.logger.Error("failed to get user by email", "error", err, "email", email)
		return nil, err
	}
	return &user, nil
}

// NewMySQLConnection creates a new MySQL connection
func NewMySQLConnection(cfg *config.Config, logger logger.Logger) *gorm.DB {
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect to database", "error", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("failed to get database instance", "error", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}
