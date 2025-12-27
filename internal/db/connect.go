package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	"Conveter/tables"
	"Conveter/tables/functions"
)

var Conn *pgx.Conn
var ctx = context.Background()

func ConnectDB() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("database url is empty")
	}
	var err error
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database %v", err)
	}

	if err := Conn.Ping(ctx); err != nil {
		log.Fatal("Unable to ping database")
	}

	fmt.Println("Connect to Postgresql succesfully")
}

func GetTable(table *functions.Currencies) error {
	sql := `SELECT id, code, full_name, sign, to_dollar 
	FROM currencies`

	rows, err := Conn.Query(ctx, sql)
	if err != nil {
		return fmt.Errorf("Unavalible to execute rows from database %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var currency tables.Currency
		err := rows.Scan(
			&currency.Id,
			&currency.Code,
			&currency.FullName,
			&currency.Sign,
			&currency.ToDollar,
		)
		if err != nil {
			return fmt.Errorf("error scaning row: %w", err)
		}

		table.Currcs = append(table.Currcs, currency)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating  rows %w", err)
	}
	return nil

}
