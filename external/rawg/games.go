package rawg

import (
	"context"

	"github.com/dimuska139/rawg-sdk-go/v3"
)

const SEARCH_LIMIT = 20

func GetGamesFromKeywordSearch(r *rawg.Client, keyword string) ([]*rawg.Game, int, error) {
	return GetGamesFromKeywordSearchPaginated(r, keyword, SEARCH_LIMIT, 1)
}

func GetGamesFromKeywordSearchPaginated(r *rawg.Client, keyword string, limit int, page int) ([]*rawg.Game, int, error) {
	rawgFilter := rawg.NewGamesFilter().SetPageSize(limit).SetPage(page).SetSearch(keyword)

	return r.GetGames(context.Background(), rawgFilter)
}
