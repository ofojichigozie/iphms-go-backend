package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ofojichigozie/iphms-go-backend/dtos"
	"github.com/ofojichigozie/iphms-go-backend/middleware"
	"github.com/ofojichigozie/iphms-go-backend/responses"
	"github.com/ofojichigozie/iphms-go-backend/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var body dtos.CreateUserInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	user, err := c.userService.CreateUser(body)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Failed to create user", err.Error())
		return
	}

	responses.Success(ctx, http.StatusCreated, "User created successfully", user)
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.userService.GetUsers()
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Failed to fetch users", err.Error())
		return
	}

	responses.Success(ctx, http.StatusOK, "Users retrieved successfully", users)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid user ID", "User ID must be a positive integer")
		return
	}

	user, err := c.userService.GetUserByID(uint(userID))
	if err != nil {
		responses.Error(ctx, http.StatusNotFound, "User not found", nil)
		return
	}

	responses.Success(ctx, http.StatusOK, "User retrieved successfully", user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Authentication required", err.Error())
		return
	}

	id := ctx.Param("id")
	userId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest,
			"Invalid user ID", "User ID must be a positive integer")
		return
	}

	if currentUser.Role != "admin" && currentUser.UserId != uint(userId) {
		responses.Error(ctx, http.StatusForbidden,
			"Not authorized", "You can only update your own profile")
		return
	}

	var body dtos.UpdateUserInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if body.DateOfBirth != nil {
		_, err = time.Parse("2006-01-02", *body.DateOfBirth)
		if err != nil {
			responses.Error(ctx, http.StatusBadRequest, "Invalid date format", "Use YYYY-MM-DD format")
			return
		}
	}

	user, err := c.userService.UpdateUser(uint(userId), body)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Failed to update user", err.Error())
		return
	}

	responses.Success(ctx, http.StatusOK, "User updated successfully", user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Authentication required", err.Error())
		return
	}

	id := ctx.Param("id")
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid user ID", "User ID must be a positive integer")
		return
	}

	if currentUser.Role != "admin" && currentUser.UserId != uint(userID) {
		responses.Error(ctx, http.StatusForbidden,
			"Not authorized", "You can only delete your own profile")
		return
	}

	err = c.userService.DeleteUser(uint(userID))
	if err != nil {
		responses.Error(ctx, http.StatusNotFound, "User not found", nil)
		return
	}

	responses.Success(ctx, http.StatusOK, "User deleted successfully", nil)
}
