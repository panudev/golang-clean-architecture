package usecase

import (
	"github.com/panudev/golang-clean-architecture/internal/domain/entity"
	"github.com/panudev/golang-clean-architecture/internal/domain/repository"
	"github.com/panudev/golang-clean-architecture/pkg/logger"
)

type UserUseCase struct {
	repo   repository.UserRepository
	logger logger.Logger
}

func NewUserUseCase(repo repository.UserRepository, logger logger.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		logger: logger,
	}
}

// GetUser retrieves a user by ID
func (uc *UserUseCase) GetUser(id uint) (*entity.User, error) {
	user, err := uc.repo.GetByID(id)
	if err != nil {
		uc.logger.Error("failed to get user", "error", err, "id", id)
		return nil, err
	}
	return user, nil
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(user *entity.User) error {
	if err := uc.repo.Create(user); err != nil {
		uc.logger.Error("failed to create user", "error", err)
		return err
	}
	return nil
}

// UpdateUser updates an existing user
func (uc *UserUseCase) UpdateUser(user *entity.User) error {
	if err := uc.repo.Update(user); err != nil {
		uc.logger.Error("failed to update user", "error", err, "id", user.ID)
		return err
	}
	return nil
}

// DeleteUser deletes a user by ID
func (uc *UserUseCase) DeleteUser(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		uc.logger.Error("failed to delete user", "error", err, "id", id)
		return err
	}
	return nil
}

// ListUsers retrieves users with pagination
func (uc *UserUseCase) ListUsers(offset, limit int) ([]*entity.User, error) {
	users, err := uc.repo.List(offset, limit)
	if err != nil {
		uc.logger.Error("failed to list users", "error", err)
		return nil, err
	}
	return users, nil
}

// GetUserByEmail retrieves a user by email
func (uc *UserUseCase) GetUserByEmail(email string) (*entity.User, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		uc.logger.Error("failed to get user by email", "error", err, "email", email)
		return nil, err
	}
	return user, nil
}
