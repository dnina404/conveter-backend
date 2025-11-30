package main

import (
	"Conveter/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run("192.168.0.15:8080")
}
