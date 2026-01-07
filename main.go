package main

import (
	"gamecheck-backend/external/rawg"
	"gamecheck-backend/internal/handlers"
	"gamecheck-backend/internal/middlewares"
	"gamecheck-backend/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := log.Default()

	// load env variables
	cfg, err := utils.LoadEnvConfig()
	if err != nil {
		logger.Fatal(err)
	}

	q, err := utils.InitSqlcQueries(&cfg.DbConfig)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Print("Connection to DB: Success")

	rawgClient := rawg.NewClient(cfg.ExternalAPISecret)

	h := handlers.NewHander(q, rawgClient)

	router := gin.Default()

	gamesRouter := router.Group("/games")
	searchRouter := router.Group("/search")

	middlewares.InitiateMiddlewares(gamesRouter)
	middlewares.InitiateMiddlewares(searchRouter)

	gamesRouter.GET("", h.ListGamesHandler)      // GET /games
	gamesRouter.PUT(":gameId", h.PutGameHandler) // PUT /games/:gameId

	searchRouter.GET( // GET /search/games
		"games",
		middlewares.AcceptPaginationParamsMiddleware(middlewares.PaginationParams{}),
		h.SearchGamesHandler,
	)

	router.Run()
}
