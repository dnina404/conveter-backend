package main

import (
	"Conveter/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
