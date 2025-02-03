package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type Database struct {
	Pool *pgxpool.Pool
}

func New() *Database {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s",
		os.Getenv("POCKEE_DATABASE_USERNAME"),
		os.Getenv("POCKEE_DATABASE_PASSWORD"),
		os.Getenv("POCKEE_DATABASE_HOST"),
		os.Getenv("POCKEE_DATABASE_NAME"))

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}

	log.Printf("connected to '%s' database", os.Getenv("POCKEE_DATABASE_NAME"))

	return &Database{
		Pool: pool,
	}
}
