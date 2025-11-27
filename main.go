package main

import (
	"Conveter/routes"

	"Conveter/tables/functions"

	"github.com/gin-gonic/gin"
)

func main() {
	r := routes.SetupRouter()
	r.GET("/all", func(ctx *gin.Context) {
		ctx.JSON(functions.Coins.ShowAll())

	})
	r.Run(":8080")
}
