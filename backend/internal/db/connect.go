package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func ConnectDB() {
	connStr := "postgres://postgres:your_password@localhost:5432/myapp"

	var err error
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	fmt.Println("✅ Подключение к PostgreSQL установлено")
}
