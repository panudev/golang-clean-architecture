package repository

import (
	"errors"

	"github.com/panudev/golang-clean-architecture/internal/domain/entity"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	GetByID(id uint) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
	List(offset, limit int) ([]*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}
