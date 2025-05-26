package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userdto "github.com/panudev/golang-clean-architecture/internal/interface/dto/user"
	"github.com/panudev/golang-clean-architecture/internal/usecase"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

// GetUser handles GET /users/:id
func (userHandler *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")

	// Validate ID
	userID, err := strconv.Atoi(idParam)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call UseCase
	user, err := userHandler.userUseCase.GetUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Map domain to DTO response
	response := userdto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, response)
}
