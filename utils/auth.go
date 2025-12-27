package utils

import (
	"github.com/gin-gonic/gin"
)

func GetUserId(ctx *gin.Context) (int, bool) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		return 0, false
	}

	return userId.(int), true
}
