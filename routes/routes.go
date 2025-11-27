package routes

import (
	"Conveter/internal/controllers"
	"Conveter/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger())
	main := r.Group("/main")
	{
		main.GET("/", controllers.GetAllCoins)

	}
	return r
}
