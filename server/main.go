package main

import (
	"github.com/gin-gonic/gin"
	"myapp/struct"
	"fmt"
)

func main() {
	server := gin.Default()
	a := 4
	fmt.Println(a)

	server.Run(":8080")
}