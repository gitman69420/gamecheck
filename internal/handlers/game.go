package handlers

import (
	"context"
	"errors"
	"fmt"
	"gamecheck-backend/db"
	"gamecheck-backend/utils"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PutGameHandlerOptions struct {
	IsPrivate bool `json:"is_private"`
}

func (h *Handler) PutGameHandler(ctx *gin.Context) {
	userId, userIdExists := utils.GetUserId(ctx)
	if !userIdExists {
		utils.GinUnauthorizedResponse(ctx)
	}

	game_id, ok := ctx.Params.Get("gameId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Unable to get gameId",
		})
		return
	}

	id, err := strconv.ParseInt(game_id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid value for gameId",
		})
		return
	}

	var putHandlerOptions PutGameHandlerOptions
	err = ctx.ShouldBindJSON(&putHandlerOptions)
	isRequestBodyEmpty := errors.Is(err, io.EOF)

	if err != nil && !isRequestBodyEmpty { // making this request without a body is also valid
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Errorf("Unable to resolve the request body: %s", err),
		})
		return
	}

	var isPrivate bool

	if isRequestBodyEmpty { // default fields in case request is made with an empty body
		isPrivate = false
	} else {
		isPrivate = putHandlerOptions.IsPrivate
	}

	if err = h.q.CreateUserPrivateGame(context.Background(), db.CreateUserPrivateGameParams{
		UserID:    int32(userId),
		GameID:    int32(id),
		IsPrivate: isPrivate,
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Unable to add game for user",
		})
		return
	}

	utils.GinSuccessResponseWithMessage(ctx, fmt.Sprintf("Game with ID: %d added", id))
}
