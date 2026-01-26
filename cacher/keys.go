package cacher

import "fmt"

// CacheKeyForSearch returns the cache key which we use for the sorted set of game cache keys
func CacheKeyForSearch(keyword string) string {
	return fmt.Sprintf("search:%s", keyword)
}

// CacheKeyForSearchMaxIndex returns the cache key which we use for the maximum search index
func CacheKeyForSearchMaxIndex(keyword string) string {
	return fmt.Sprintf("search:%s:max-index", keyword)
}

// CacheKeyForGame returns the cache key which we use for the specific game
func CacheKeyForGame(id int) string {
	return fmt.Sprintf("game:%d", id)
}
