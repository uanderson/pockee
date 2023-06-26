package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

// Pool for the database connection
var Pool *pgxpool.Pool

// Init the database connection pool
func Init() {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_NAME"))

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}

	log.Printf("connected to '%s' database", os.Getenv("DATABASE_NAME"))

	Pool = pool
}
