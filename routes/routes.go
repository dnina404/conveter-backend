package routes

import (
	"Conveter/middleware"
	"Conveter/tables/functions"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger())
	main := r.Group("/main")
	{
		main.GET("/", functions.CurrObj.ShowAll)
		main.GET("/pong", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}
	return r
}
