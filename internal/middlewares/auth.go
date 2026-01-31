package middlewares

import (
	"fmt"
	"gamecheck-backend/auth"
	"gamecheck-backend/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt := ctx.GetHeader("JWT")

		if jwt == "" {
			utils.GinUnauthorizedResponse(ctx)
			return
		}

		claims, err := auth.VerifyJWT(jwt)

		if err != nil {
			utils.GinUnauthorizedResponseWithMessage(ctx, fmt.Sprintf("Failed authentication: %s", err))
			return
		}

		ctx.Set("user_id", claims.UserId)
		ctx.Set("steamid", claims.SteamId)
		ctx.Next()

	}
}
