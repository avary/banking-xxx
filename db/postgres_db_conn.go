package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/url"
)

const (
	scheme   = "postgres"
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "1234"
	dbname   = "bank"
)

// GetDbClient returns a connection pool to the database
// Using pgxpool as it's concurrency safe
func GetDbClient() *pgxpool.Pool {
	dsn := url.URL{
		Scheme: scheme,
		User:   url.UserPassword(user, password),
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   dbname,
	}

	q := dsn.Query()
	q.Set("sslmode", "disable")
	dsn.RawQuery = q.Encode()

	dbPool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	//defer dbPool.Close()  will do this in server.go

	if err = dbPool.Ping(context.Background()); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Printf("successfully connected to database %s", dsn.String())

	return dbPool
}
