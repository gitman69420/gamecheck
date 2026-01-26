package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinUnauthorizedResponse(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
		"success": false,
	})
}

func GinBadRequest(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
		"success": false,
	})
}

func GinNotFound(ctx *gin.Context, responseBody any) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, responseBody)
}

func GinSuccessResponse(ctx *gin.Context, responseBody any) {
	ctx.JSON(http.StatusOK, responseBody)
}

func GinSuccessResponseWithMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{"message": message, "success": true})
}
