package middleware

import "github.com/ofojichigozie/iphms-go-backend/utils"

type CurrentUser struct {
	UserId uint   `json:"userId"`
	Role   string `json:"role"`
}

const (
	CurrentUserKey = "currentUser"
)

func FromJWTClaims(claims *utils.JWTClaims) CurrentUser {
	return CurrentUser{
		UserId: claims.UserID,
		Role:   claims.Role,
	}
}
