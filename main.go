package main

import (
	"Conveter/internal/db"
	"Conveter/routes"
	"Conveter/tables/functions"
	"log"
)

func main() {
	db.ConnectDB()
	err := db.GetTable(&functions.CurrObj)
	if err != nil {
		log.Fatal("you have error with db.GetTable", err)
	}
	r := routes.SetupRouter()
	r.Run("192.168.0.15:8080")
}
