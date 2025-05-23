package usecase

import "github.com/panudev/golang-clean-architecture/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) GetUser(id uint) (*domain.User, error) {
	return u.repo.GetByID(id)
}
