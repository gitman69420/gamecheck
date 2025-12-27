package handlers

import (
	"gamecheck-backend/models"
	"gamecheck-backend/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListGamesHandler(ctx *gin.Context) {
	// userId, userIdExists := utils.GetUserId(ctx)
	// if !userIdExists {
	// 	utils.GinUnauthorizedResponse(ctx)
	// 	return
	// }

	// games, err := h.q.GetUserPublicGames(context.Background(), int32(userId))
	games := models.ListGames()

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"success": false,
	// 		"message": err,
	// 	})
	// 	return
	// }

	utils.GinSuccessResponse(ctx, utils.CreateListResponse(games))
}
