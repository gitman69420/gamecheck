package handlers

import (
	"fmt"
	"gamecheck-backend/external/rawg"
	"gamecheck-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SearchGamesHandler(ctx *gin.Context) {

	search, searchExists := ctx.GetQuery("search")

	if !searchExists {
		utils.GinBadResponse(ctx, "Missing search query param")
		return
	}

	if utils.Strlen(search) < 3 {
		utils.GinBadResponse(ctx, "Search query is too short (<3 characters)")
		return
	}

	res, _, err := rawg.GetGamesFromKeywordSearch(h.r, search)
	if err != nil {
		ctx.AbortWithError(http.StatusServiceUnavailable, fmt.Errorf("Unable to reach RAWG services"))
		return
	}

	utils.GinSuccessResponse(ctx, utils.CreateListResponse(res))

}
