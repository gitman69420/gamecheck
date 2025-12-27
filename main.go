package main

import (
	"gamecheck-backend/internal/handlers"
	"gamecheck-backend/internal/middlewares"
	"gamecheck-backend/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	logger := log.Default()
	q, err := utils.InitSqlcQueries()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Print("Connection to DB: Success")

	h := handlers.NewHander(q)

	router := gin.Default()

	gamesRouter := router.Group("/games")

	middlewares.InitiateMiddlewares(gamesRouter)

	gamesRouter.GET("", h.ListGamesHandler)
	gamesRouter.PUT(":gameId", h.PutGameHandler)

	router.Run()
}
