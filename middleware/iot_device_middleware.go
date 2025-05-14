package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ofojichigozie/iphms-go-backend/responses"
	"github.com/ofojichigozie/iphms-go-backend/services"
)

const (
	iotSecretConstant = "someiotdevicesecret"
)

func IoTDeviceMiddleware(userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		deviceID := ctx.GetHeader("X-Device-ID")
		secret := ctx.GetHeader("X-Device-Secret")

		if secret != iotSecretConstant {
			responses.Error(ctx, http.StatusUnauthorized, "Unauthorized", "Invalid device credentials")
			ctx.Abort()
			return
		}

		user, err := userService.GetUserByDeviceID(deviceID)
		if err != nil {
			responses.Error(ctx, http.StatusUnauthorized, "Unauthorized", "Device not registered")
			ctx.Abort()
			return
		}

		currentUser := CurrentUser{
			UserId: user.ID,
			Role:   user.Role,
		}

		ctx.Set(CurrentUserKey, currentUser)
		ctx.Next()
	}
}
