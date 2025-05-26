package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panudev/golang-clean-architecture/internal/domain/entity"
	"github.com/panudev/golang-clean-architecture/internal/domain/repository"
	userdto "github.com/panudev/golang-clean-architecture/internal/interface/dto/user"
	usecase "github.com/panudev/golang-clean-architecture/internal/usecase/user"
	"github.com/panudev/golang-clean-architecture/pkg/logger"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
	logger      logger.Logger
}

func NewUserHandler(userUseCase *usecase.UserUseCase, logger logger.Logger) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
		logger:      logger,
	}
}

// @Summary      Get user by ID
// @Description  Get user details by user ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  userdto.UserResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userUseCase.GetUser(uint(userID))
	if err != nil {
		if err == repository.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		h.logger.Error("failed to get user", "error", err, "id", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, userdto.ToUserResponse(user))
}

// @Summary      Create new user
// @Description  Create a new user with the provided details
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      userdto.CreateUserRequest  true  "User object"
// @Success      201  {object}  userdto.UserResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var request userdto.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password, // Note: Should be hashed in usecase
	}

	if err := h.userUseCase.CreateUser(user); err != nil {
		h.logger.Error("failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, userdto.ToUserResponse(user))
}

// @Summary      Update user
// @Description  Update user details by user ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int  true  "User ID"
// @Param        user  body      userdto.UpdateUserRequest  true  "User object"
// @Success      200  {object}  userdto.UserResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request userdto.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &entity.User{
		ID:    uint(userID),
		Name:  request.Name,
		Email: request.Email,
	}

	if err := h.userUseCase.UpdateUser(user); err != nil {
		if err == repository.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		h.logger.Error("failed to update user", "error", err, "id", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, userdto.ToUserResponse(user))
}

// @Summary      Delete user
// @Description  Delete user by user ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  "No Content"
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.userUseCase.DeleteUser(uint(userID)); err != nil {
		if err == repository.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		h.logger.Error("failed to delete user", "error", err, "id", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary      List users
// @Description  Get list of users with pagination
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        offset  query     int  false  "Offset for pagination"  default(0)
// @Param        limit   query     int  false  "Limit for pagination"   default(10)
// @Success      200    {array}   userdto.UserResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	users, err := h.userUseCase.ListUsers(offset, limit)
	if err != nil {
		h.logger.Error("failed to list users", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	response := make([]userdto.UserResponse, len(users))
	for i, user := range users {
		response[i] = *userdto.ToUserResponse(user)
	}

	c.JSON(http.StatusOK, response)
}
