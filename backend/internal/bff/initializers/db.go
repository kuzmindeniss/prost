package initializers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DbConn *pgx.Conn

func ConnectToDb() {
	var err error
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		panic("DB_URL is not set")
	}

	DbConn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
}
