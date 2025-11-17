package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello world", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "idi nahuy, we Russian persons!",
		})
	})
	r.GET("/main", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "html",
		})
	})
	if err := r.Run(":4333"); err != nil {
		log.Fatalf("failed to run server: %v", err)

	}
}
