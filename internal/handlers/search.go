package handlers

import (
	"fmt"
	"gamecheck-backend/external/rawg"
	"gamecheck-backend/utils"
	"net/http"

	rawgSdk "github.com/dimuska139/rawg-sdk-go/v3"
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

	page := ctx.GetUint("page")
	size := ctx.GetUint("size")

	isPaginationParamsPresent := page != 0 && size != 0

	var res []*rawgSdk.Game
	var err error

	if isPaginationParamsPresent {
		res, _, err = rawg.GetGamesFromKeywordSearchPaginated(h.r, search, size, page)
	} else {
		res, _, err = rawg.GetGamesFromKeywordSearch(h.r, search)
	}

	if err != nil {
		ctx.AbortWithError(http.StatusServiceUnavailable, fmt.Errorf("Unable to reach RAWG services"))
		return
	}

	utils.GinSuccessResponse(ctx, utils.CreateListResponse(res))
}
