package dto

import "github.com/panudev/golang-clean-architecture/internal/domain/entity"

// UserResponse represents the user response structure
type UserResponse struct {
	ID    uint   `json:"id" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john@example.com"`
}

func ToUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
