package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectPostgres() error {
	godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return err
	}

	// test connection
	err = pool.Ping(context.Background())
	if err != nil {
		return err
	}

	DB = pool
	fmt.Println("PostgreSQL Connected!")
	return nil
}
