package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ofojichigozie/iphms-go-backend/responses"
	"github.com/ofojichigozie/iphms-go-backend/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		cookie, err := ctx.Cookie("accessToken")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			responses.Error(ctx, http.StatusUnauthorized, "Unauthorized", "An access token is required")
			ctx.Abort()
			return
		}

		user, err := utils.VerifyToken(accessToken)
		if err != nil {
			responses.Error(ctx, http.StatusUnauthorized, "Unauthorized", err.Error())
			ctx.Abort()
			return
		}

		ctx.Set(CurrentUserKey, FromJWTClaims(user))
		ctx.Next()
	}
}
