package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ofojichigozie/iphms-go-backend/dtos"
	"github.com/ofojichigozie/iphms-go-backend/responses"
	"github.com/ofojichigozie/iphms-go-backend/services"
	"github.com/ofojichigozie/iphms-go-backend/utils"
)

type AuthController struct {
	userService services.UserService
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{userService}
}

func (c AuthController) Register(ctx *gin.Context) {
	var body dtos.CreateUserInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	user, err := c.userService.CreateUser(body)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Failed to create user", err.Error())
		return
	}

	tokens, err := utils.GenerateTokenPair(user.ID, user.Role)
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Token generation failed", err.Error())
		return
	}

	responses.Success(ctx, http.StatusCreated, "Registration successful", gin.H{
		"accessToken":  tokens["accessToken"],
		"refreshToken": tokens["refreshToken"],
		"user":         user,
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var body dtos.LoginInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	user, err := c.userService.GetUserByEmail(body.Email)
	if err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Authentication failed", "Invalid email or password")
		return
	}

	if err := utils.VerifyPassword(user.Password, body.Password); err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Authentication failed", "Invalid email or password")
		return
	}

	tokens, err := utils.GenerateTokenPair(user.ID, user.Role)
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Token generation failed", err.Error())
		return
	}

	responses.Success(ctx, http.StatusOK, "Login successful", gin.H{
		"accessToken":  tokens["accessToken"],
		"refreshToken": tokens["refreshToken"],
		"user":         user,
	})
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		responses.Error(ctx, http.StatusUnauthorized, "Authorization header required", nil)
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		responses.Error(ctx, http.StatusUnauthorized, "Invalid authorization format", nil)
		return
	}

	newTokens, err := utils.RefreshToken(tokenParts[1])
	if err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Token refresh failed", err.Error())
		return
	}

	responses.Success(ctx, http.StatusOK, "Token refreshed successfully", gin.H{
		"accessToken":  newTokens["accessToken"],
		"refreshToken": newTokens["refreshToken"],
	})
}
