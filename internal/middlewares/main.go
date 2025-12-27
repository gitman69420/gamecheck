package middlewares

import "github.com/gin-gonic/gin"

var generalMiddlewares []gin.HandlerFunc = []gin.HandlerFunc{
	Auth,
}

func InitiateMiddlewares(g *gin.RouterGroup) {
	for _, m := range generalMiddlewares {
		g.Use(m)
	}
}
