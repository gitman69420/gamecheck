package main

import (
	"gamecheck-backend/cacher"
	"gamecheck-backend/config"
	"gamecheck-backend/external/rawg"
	"gamecheck-backend/internal/handlers"
	"gamecheck-backend/internal/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := log.Default()

	// load env variables
	cfg, err := config.LoadEnvConfig()
	if err != nil {
		logger.Fatal(err)
	}

	vkClient, err := cacher.NewClient(cfg.VkAddress, cfg.VkPassword)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Print("Connection to Valkey: Success")

	q, err := config.InitSqlcQueries(&cfg.DbConfig)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Print("Connection to DB: Success")

	rawgClient := rawg.NewClient(cfg.ExternalAPISecret)

	h := handlers.NewHander(q, rawgClient, vkClient, logger)

	router := gin.Default()

	authRouter := router.Group("/auth")
	gamesRouter := router.Group("/games")
	searchRouter := router.Group("/search")
	userRouter := router.Group("/user")

	middlewares.InitiateMiddlewares(gamesRouter)
	middlewares.InitiateMiddlewares(searchRouter)
	middlewares.InitiateMiddlewares(userRouter)

	authRouter.POST("create-session", h.CreateAuthSessionHandler) // POST /auth/create-session

	gamesRouter.GET("", h.ListGamesHandler)      // GET /games
	gamesRouter.PUT(":gameId", h.PutGameHandler) // PUT /games/:gameId

	searchRouter.GET( // GET /search/games
		"games",
		middlewares.AcceptPaginationParamsMiddleware(middlewares.PaginationParams{}),
		h.SearchGamesHandler,
	)

	userRouter.GET("me", h.GetMyInfoHandler) // GET /user/me

	router.Run()
}
