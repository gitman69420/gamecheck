package middlewares

import (
	"gamecheck-backend/utils"

	"github.com/gin-gonic/gin"
)

var Auth gin.HandlerFunc = func(ctx *gin.Context) {
	jwt := ctx.Request.Header.Get("JWT")

	if jwt == "" {
		utils.GinUnauthorizedResponse(ctx)
		return
	}

	// TODO: check session and get user_id

	ctx.Set("user_id", 1001)
	ctx.Next()

}
