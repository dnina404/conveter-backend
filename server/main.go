package main

import (
	"github.com/gin-gonic/gin"
	"myapp/struct"
)

func main() {
	server := gin.Default()


	server.Run(":8080")
}