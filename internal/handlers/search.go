package handlers

import (
	"encoding/json"
	"fmt"
	"gamecheck-backend/cacher"
	"gamecheck-backend/external/rawg"
	"gamecheck-backend/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	rawgSdk "github.com/dimuska139/rawg-sdk-go/v3"
	"github.com/gin-gonic/gin"
	"github.com/valkey-io/valkey-go"
)

func (h *Handler) getCurrentValueOfSearchMaxIndex(ctx *gin.Context, keyword string) (int, error) {
	searchMaxIndexCacheKey := cacher.CacheKeyForSearchMaxIndex(keyword)

	if searchMaximumIndex := ctx.GetInt(searchMaxIndexCacheKey); searchMaximumIndex != 0 {
		return searchMaximumIndex, nil
	}

	getSearchMaxIndexCmd := h.c.B().Get().Key(searchMaxIndexCacheKey).Build()

	res := h.c.Do(ctx, getSearchMaxIndexCmd)

	if res.Error() != nil {
		return 0, res.Error()
	}

	searchMaxIndex, err := res.AsInt64()
	if err != nil {
		return 0, err
	}

	return int(searchMaxIndex), nil
}

func (h *Handler) setCurrentValueOfSearchMaxIndex(ctx *gin.Context, keyword string, value int) error {
	searchMaxIndexCacheKey := cacher.CacheKeyForSearchMaxIndex(keyword)
	valueStr := strconv.Itoa(value)
	setSearchMaxIndexCmd := h.c.B().Set().Key(searchMaxIndexCacheKey).Value(valueStr).Build()

	res := h.c.Do(ctx, setSearchMaxIndexCmd)

	if res.Error() != nil {
		return res.Error()
	}

	ctx.Set(searchMaxIndexCacheKey, value)

	return nil
}

func (h *Handler) getGamesSearchResultFromCache(ctx *gin.Context, keyword string, pageNumber uint, pageSize uint) ([]*rawgSdk.Game, error) {
	startIndex := utils.GetPaginationStartIndex(pageNumber, pageSize)
	endIndex := startIndex + pageSize

	keywordCacheKey := cacher.CacheKeyForSearch(keyword)

	start := strconv.Itoa(int(startIndex))
	end := strconv.Itoa(int(endIndex))

	readKeywordSearchCmd := h.c.B().Zrange().Key(keywordCacheKey).Min(start).Max(end).Build()

	gamesSetResult := h.c.Do(ctx, readKeywordSearchCmd)

	if valkey.IsValkeyNil(gamesSetResult.Error()) {
		return []*rawgSdk.Game{}, nil
	}

	gameCacheKeys, err := gamesSetResult.AsStrSlice()
	if err != nil {
		return nil, err
	}

	if uint(len(gameCacheKeys)) < pageSize {
		return nil, fmt.Errorf("Cache for requested page and size is either empty or incomplete")
	}

	// get the games for the respective cache keys
	readGamesCmd := h.c.B().Mget().Key(gameCacheKeys...).Build()
	gamesCacheRes := h.c.Do(ctx, readGamesCmd)

	if valkey.IsValkeyNil(gamesCacheRes.Error()) {
		return []*rawgSdk.Game{}, nil
	}

	gamesCache, err := gamesCacheRes.AsStrSlice()
	if err != nil {
		return nil, err
	}

	finalGames := []*rawgSdk.Game{}

	for _, stringifiedGameJson := range gamesCache {
		game := rawgSdk.Game{}
		err := json.Unmarshal([]byte(stringifiedGameJson), &game)
		if err != nil {
			return nil, err
		}
		finalGames = append(finalGames, &game)
	}

	return finalGames, nil
}

// updateGameSearchMaxIndex updates the Search Maximum Index based on the pagination parameters and the result set
func (h *Handler) updateGameSearchMaxIndex(ctx *gin.Context, resultsCount int, keyword string, pageNumber uint, pageSize uint) {
	startIndex := (int)(utils.GetPaginationStartIndex(pageNumber, pageSize))
	searchMaxIndex, err := h.getCurrentValueOfSearchMaxIndex(ctx, keyword)

	isSearchMaxIndexAlreadySet := !valkey.IsValkeyNil(err)

	// unable to get the current Search Maximum Index value
	if err != nil && !valkey.IsValkeyNil(err) {
		// do nothing, who cares, we can always set it later
		return
	}

	// happy path case:
	// if the result size is a full page size (we're targetting a page somewhere in the middle of the total collection)
	if resultsCount == int(pageSize) {
		// don't worry about it
		return
	}

	// earlier final page case:
	// if fetched results are 0
	if resultsCount == 0 {
		// if it's not set, we can set it for the first time
		// or if we've queried an earlier page to get 0 results
		if !isSearchMaxIndexAlreadySet || startIndex < searchMaxIndex {
			h.setCurrentValueOfSearchMaxIndex(ctx, keyword, startIndex)
		}

		return
	}

	// final case:
	// if the results size is less than a full page size (we successfully found the last page relative to the page size)
	newSearchMaxIndex := startIndex + resultsCount
	h.setCurrentValueOfSearchMaxIndex(ctx, keyword, newSearchMaxIndex)
}

func (h *Handler) saveGamesSearchResultToCache(ctx *gin.Context, results []*rawgSdk.Game, keyword string, pageNumber uint, pageSize uint) error {
	h.updateGameSearchMaxIndex(ctx, len(results), keyword, pageNumber, pageSize)
	if len(results) == 0 {
		return nil
	}

	// save all the games to the cache
	var gameCacheKeys []string

	h.c.Dedicated(func(dedicatedClient valkey.DedicatedClient) error {
		for _, game := range results {
			gameCacheKey := cacher.CacheKeyForGame(game.ID)

			stringifiedJsonGame, err := json.Marshal(game)
			if err != nil {
				h.logger.Fatalf("`rawg.Game` cannot be JSON Marshalled: %s\n", err)
				continue
			}

			writeGameCmd := dedicatedClient.B().Set().Key(gameCacheKey).Value(string(stringifiedJsonGame)).Ex(time.Hour * 24 * 7).Build()
			if err := dedicatedClient.Do(ctx, writeGameCmd).Error(); err != nil {
				h.logger.Printf("Cannot write key %s to cache\n", gameCacheKey)
				continue
			}

			h.logger.Printf("Written key %s to cache successfuly", gameCacheKey)

			gameCacheKeys = append(gameCacheKeys, gameCacheKey)
		}
		return nil
	})

	startIndex := utils.GetPaginationStartIndex(pageNumber, pageSize)

	// save the keywords search to the cache
	keywordCacheKey := cacher.CacheKeyForSearch(keyword)

	writeKeywordSearchCmdBuilder := h.c.B().Zadd().Key(keywordCacheKey).Nx().ScoreMember()
	for idx, gameId := range gameCacheKeys {
		score := float64(int(startIndex) + idx)
		writeKeywordSearchCmdBuilder.ScoreMember(score, gameId)
	}

	writeKeywordSearchCmd := writeKeywordSearchCmdBuilder.Build()
	if err := h.c.Do(ctx, writeKeywordSearchCmd).Error(); err != nil {
		return fmt.Errorf("Unable to write search to cache: %s", keywordCacheKey)
	}

	setSearchExpireCmd := h.c.B().Expire().Key(keywordCacheKey).Seconds(int64(time.Hour.Seconds() * 24)).Build()
	if err := h.c.Do(ctx, setSearchExpireCmd).Error(); err != nil {
		return fmt.Errorf("Unable to set expiry for search cache: %s", keywordCacheKey)
	}

	h.logger.Printf("Written key %s to cache successfuly", keywordCacheKey)
	return nil
}

func (h *Handler) SearchGamesHandler(ctx *gin.Context) {

	searchQuery, searchExists := ctx.GetQuery("search")

	if !searchExists {
		utils.GinBadRequest(ctx, "Missing search query param")
		return
	}

	search := strings.ToLower(strings.Trim(searchQuery, " "))

	if utils.Strlen(search) < 3 {
		utils.GinBadRequest(ctx, "Search query is too short (<3 characters)")
		return
	}

	page := ctx.GetUint("page")
	size := ctx.GetUint("size")

	isPaginationParamsPresent := page != 0 && size != 0

	if !isPaginationParamsPresent {
		page = rawg.DEFAULT_PAGE
		size = rawg.DEFAULT_PAGE_SIZE
	}

	var res []*rawgSdk.Game
	var err error

	games, err := h.getGamesSearchResultFromCache(ctx, search, page, size)
	if err == nil {
		// successful read from cache
		responseData := utils.CreateListResponse(games)
		utils.GinSuccessResponse(ctx, responseData)
		return
	}

	h.logger.Printf("Unable to fetch from cache: %s\n", err)

	res, _, err = rawg.GetGamesFromKeywordSearchPaginated(h.r, search, size, page)

	if err != nil {
		if err, ok := err.(*rawgSdk.RawgError); ok {
			errorMsg := fmt.Errorf("Error from RAWG: %s", err.Body)
			ctx.AbortWithStatusJSON(err.HttpCode, gin.H{
				"message": errorMsg,
				"success": false,
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"message": fmt.Errorf("Error from RAWG: %s", err),
			"success": false,
		})
		return
	}

	defer func() {
		go h.saveGamesSearchResultToCache(ctx, res, search, page, size)
	}()

	utils.GinSuccessResponse(ctx, utils.CreateListResponse(res))
}
