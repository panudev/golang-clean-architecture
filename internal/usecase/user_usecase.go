package usecase

import "github.com/panudev/golang-clean-architecture/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (userUseCase *UserUseCase) GetUser(id uint) (*domain.User, error) {
	return userUseCase.repo.GetByID(id)
}
