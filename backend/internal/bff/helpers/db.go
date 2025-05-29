package helpers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

var DbConn *pgx.Conn

var Repo *repository.Queries

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

	Repo = repository.New(DbConn)
}
