package rawg

import (
	"context"

	"github.com/dimuska139/rawg-sdk-go/v3"
)

const DEFAULT_PAGE_SIZE = 20
const DEFAULT_PAGE = 1

func GetGamesFromKeywordSearch(r *rawg.Client, keyword string) ([]*rawg.Game, int, error) {
	return GetGamesFromKeywordSearchPaginated(r, keyword, DEFAULT_PAGE_SIZE, DEFAULT_PAGE)
}

func GetGamesFromKeywordSearchPaginated(r *rawg.Client, keyword string, limit uint, page uint) ([]*rawg.Game, int, error) {
	rawgFilter := rawg.NewGamesFilter().SetPageSize(int(limit)).SetPage(int(page)).SetSearch(keyword)

	return r.GetGames(context.Background(), rawgFilter)
}
