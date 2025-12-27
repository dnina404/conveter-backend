package routes

import (
	"Conveter/middleware"
	"Conveter/tables/functions"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	log.Print("router starteds")
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger())
	main := r.Group("/main")
	{
		main.GET("/showall", functions.CurrObj.ShowAll)
		main.GET("/pong", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}
	return r
}
