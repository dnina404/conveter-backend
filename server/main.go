package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	a := 4
	fmt.Println(a)

	server.Run(":8080")
}
