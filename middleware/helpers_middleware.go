package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) (CurrentUser, error) {
	val, exists := c.Get(CurrentUserKey)
	if !exists {
		return CurrentUser{}, errors.New("user context missing")
	}

	user, ok := val.(CurrentUser)
	if !ok {
		return CurrentUser{}, errors.New("invalid user context type")
	}

	if user.UserId == 0 {
		return CurrentUser{}, errors.New("invalid user ID")
	}

	return user, nil
}
