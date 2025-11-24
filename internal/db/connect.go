package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func ConnectDB() {
	connStr := "DATABASE_URL"

	var err error
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	fmt.Println("Подключение к PostgreSQL установлено")
}
